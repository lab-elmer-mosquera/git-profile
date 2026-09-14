package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	includePath := flag.String("include", "", "Ruta de la carpeta del proyecto para enlazar (ej: . o /ruta/proyecto)")
	profileFlag := flag.String("profile", "", "Nombre del perfil a incluir (ej: udea)")
	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("No se pudo obtener el directorio Home: %v", err)
	}

	if *includePath != "" && *profileFlag != "" {
		ejecutarModoInclude(homeDir, *includePath, *profileFlag)
		return
	}

	ejecutarModoCreacion(homeDir)
}

func ejecutarModoCreacion(homeDir string) {
	var nombre, correo, perfil string

	fmt.Println("=== Configuración de Nuevo Perfil Git/SSH ===")
	fmt.Print("Ingresa tu nombre para Git: ")
	fmt.Scanln(&nombre)
	fmt.Print("Ingresa tu correo para Git: ")
	fmt.Scanln(&correo)
	fmt.Print("Ingresa el nombre del perfil (ej: udea): ")
	fmt.Scanln(&perfil)

	if nombre == "" || correo == "" || perfil == "" {
		log.Fatal("Todos los campos son obligatorios.")
	}

	sshKeyPath := filepath.Join(homeDir, ".ssh", "id_ed25519_"+perfil)
	customGitConfig := filepath.Join(homeDir, fmt.Sprintf(".gitconfig-%s", perfil))

	fmt.Printf("\n[1/3] Generando llave SSH Ed25519 en: %s...\n", sshKeyPath)
	cmdSSH := exec.Command("ssh-keygen", "-t", "ed25519", "-C", correo, "-f", sshKeyPath, "-N", "")
	if out, err := cmdSSH.CombinedOutput(); err != nil {
		log.Fatalf("Error creando SSH: %s\n%s", err, string(out))
	}

	fmt.Printf("[2/3] Creando archivo de configuración Git: %s...\n", customGitConfig)
	configContent := fmt.Sprintf("[user]\n\tname = %s\n\temail = %s\n[core]\n\tsshCommand = ssh -i %s -F /dev/null\n", nombre, correo, sshKeyPath)

	err := os.WriteFile(customGitConfig, []byte(configContent), 0644)
	if err != nil {
		log.Fatalf("Error al escribir archivo .gitconfig personalizado: %v", err)
	}

	fmt.Println("[3/3] ¡Perfil creado con éxito!")
	fmt.Printf("\nPara activar este perfil en una carpeta, corre:\ngprofile-includef --include . --profile %s\n", perfil)
}

func ejecutarModoInclude(homeDir, includePath, perfil string) {
	absPath, err := filepath.Abs(includePath)
	if err != nil {
		log.Fatalf("Error al resolver la ruta: %v", err)
	}

	if !strings.HasSuffix(absPath, string(filepath.Separator)) {
		absPath += string(filepath.Separator)
	}

	gitConfigOriginal := filepath.Join(homeDir, ".gitconfig")
	customGitConfig := filepath.Join(homeDir, fmt.Sprintf(".gitconfig-%s", perfil))

	if _, err := os.Stat(customGitConfig); os.IsNotExist(err) {
		log.Fatalf("El perfil '%s' no existe. Créalo primero ejecutando el programa sin banderas.", perfil)
	}

	gitPathFormat := filepath.ToSlash(absPath)

	// Creamos la cabecera de la sección para buscar si ya existe
	searchString := fmt.Sprintf("[includeIf \"gitdir:%s\"]", gitPathFormat)
	includeBlock := fmt.Sprintf("\n%s\n\tpath = %s\n", searchString, customGitConfig)

	// Leer el archivo original para comprobar duplicados
	content, err := os.ReadFile(gitConfigOriginal)
	if err != nil && !os.IsNotExist(err) {
		log.Fatalf("Error al leer el .gitconfig original: %v", err)
	}

	// EVITAR DUPLICADOS: Si ya contiene la sección, no hacemos nada
	if bytes.Contains(content, []byte(searchString)) {
		fmt.Printf("Aviso: La carpeta '%s' ya se encuentra vinculada en tu .gitconfig. No se realizaron cambios.\n", absPath)
		return
	}

	// Guardar el cambio si no está duplicado
	f, err := os.OpenFile(gitConfigOriginal, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("No se pudo abrir el .gitconfig original: %v", err)
	}
	defer f.Close()

	if _, err := f.WriteString(includeBlock); err != nil {
		log.Fatalf("Error al escribir el include en .gitconfig: %v", err)
	}

	fmt.Printf("¡Enlazado con éxito! A partir de ahora, lo que esté dentro de '%s' usará el perfil '%s'.\n", absPath, perfil)
}
