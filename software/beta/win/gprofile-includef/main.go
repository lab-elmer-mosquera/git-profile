package main

import (
	"bufio"
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
	includePath := flag.String(
		"include",
		"",
		"Directorio al que se asociará el perfil",
	)

	profileName := flag.String(
		"profile",
		"",
		"Nombre del perfil",
	)

	flag.Parse()

	homeDir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("No se pudo obtener el directorio del usuario: %v", err)
	}

	// Sin argumentos: creación interactiva.
	if *includePath == "" && *profileName == "" {
		crearPerfil(homeDir)
		return
	}

	// Ambos argumentos son necesarios para vincular un perfil.
	if *includePath == "" || *profileName == "" {
		log.Fatal(
			"Uso: gprofile-includef.exe --include <directorio> --profile <perfil>",
		)
	}

	vincularPerfil(
		homeDir,
		*includePath,
		*profileName,
	)
}

func crearPerfil(homeDir string) {
	var nombre string
	var correo string
	var perfil string

	fmt.Println("=== Configuración de Nuevo Perfil Git/SSH ===")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa tu nombre para Git: ")
	nombre, _ = reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre)

	fmt.Print("Ingresa tu correo para Git: ")
	correo, _ = reader.ReadString('\n')
	correo = strings.TrimSpace(correo)

	fmt.Print("Ingresa el nombre del perfil (ej: udea): ")
	perfil, _ = reader.ReadString('\n')
	perfil = strings.TrimSpace(perfil)

	if nombre == "" || correo == "" || perfil == "" {
		log.Fatal("Todos los campos son obligatorios.")
	}

	sshDir := filepath.Join(homeDir, ".ssh")

	if err := os.MkdirAll(sshDir, 0700); err != nil {
		log.Fatalf("No se pudo crear el directorio .ssh: %v", err)
	}

	sshKeyPath := filepath.Join(
		sshDir,
		"id_ed25519_"+perfil,
	)

	gitConfigProfile := filepath.Join(
		homeDir,
		".gitconfig-"+perfil,
	)

	fmt.Printf(
		"\n[1/3] Generando llave SSH Ed25519 en: %s...\n",
		sshKeyPath,
	)

	generarLlaveSSH(sshKeyPath, correo)

	fmt.Printf(
		"[2/3] Creando archivo de configuración Git: %s...\n",
		gitConfigProfile,
	)

	crearConfiguracionGit(
		gitConfigProfile,
		nombre,
		correo,
		sshKeyPath,
	)

	fmt.Println("[3/3] ¡Perfil creado con éxito!")
	fmt.Println()

	fmt.Printf(
		"Para activar este perfil en una carpeta, corre:\n"+
			"gprofile-includef.exe --include . --profile %s\n",
		perfil,
	)
}

func generarLlaveSSH(sshKeyPath, correo string) {
	cmd := exec.Command(
		"ssh-keygen",
		"-t",
		"ed25519",
		"-C",
		correo,
		"-f",
		sshKeyPath,
		"-N",
		"",
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf(
			"Error creando la llave SSH:\n%s\n%s",
			err,
			string(output),
		)
	}
}

func crearConfiguracionGit(
	configPath string,
	nombre string,
	correo string,
	sshKeyPath string,
) {
	// Git utiliza cómodamente / incluso en Windows.
	gitSSHPath := filepath.ToSlash(sshKeyPath)

	contenido := fmt.Sprintf(
		"[user]\n"+
			"\tname = %s\n"+
			"\temail = %s\n"+
			"\n"+
			"[core]\n"+
			"\tsshCommand = ssh -i \"%s\" -F NUL\n",
		nombre,
		correo,
		gitSSHPath,
	)

	if err := os.WriteFile(
		configPath,
		[]byte(contenido),
		0644,
	); err != nil {
		log.Fatalf(
			"No se pudo crear la configuración Git: %v",
			err,
		)
	}
}

func vincularPerfil(
	homeDir string,
	includePath string,
	perfil string,
) {
	absPath, err := filepath.Abs(includePath)
	if err != nil {
		log.Fatalf("No se pudo resolver el directorio: %v", err)
	}

	absPath = filepath.Clean(absPath)

	gitConfigProfile := filepath.Join(
		homeDir,
		".gitconfig-"+perfil,
	)

	if _, err := os.Stat(gitConfigProfile); err != nil {
		if os.IsNotExist(err) {
			log.Fatalf(
				"El perfil '%s' no existe. "+
					"Créalo primero ejecutando gprofile-includef.exe.",
				perfil,
			)
		}

		log.Fatalf(
			"No se pudo comprobar el perfil: %v",
			err,
		)
	}

	// Convertimos C:\... a C:/...
	gitDir := filepath.ToSlash(absPath)

	// includeIf.gitdir necesita reconocer el directorio.
	if !strings.HasSuffix(gitDir, "/") {
		gitDir += "/"
	}

	includeKey := fmt.Sprintf(
		"includeIf.gitdir:%s.path",
		gitDir,
	)

	// Git acepta la ruta del archivo de configuración.
	configPath := filepath.ToSlash(gitConfigProfile)

	// Consultamos primero para evitar duplicar la misma entrada.
	existente := obtenerConfiguracion(includeKey)

	if existente == configPath {
		fmt.Printf(
			"Aviso: el directorio '%s' ya está vinculado al perfil '%s'.\n",
			absPath,
			perfil,
		)
		return
	}

	cmd := exec.Command(
		"git",
		"config",
		"--global",
		includeKey,
		configPath,
	)

	if output, err := cmd.CombinedOutput(); err != nil {
		log.Fatalf(
			"No se pudo vincular el perfil:\n%s\n%s",
			err,
			string(output),
		)
	}

	fmt.Printf(
		"Perfil '%s' vinculado con éxito a:\n%s\n",
		perfil,
		absPath,
	)
}

func obtenerConfiguracion(key string) string {
	cmd := exec.Command(
		"git",
		"config",
		"--global",
		"--get",
		key,
	)

	output, err := cmd.Output()
	if err != nil {
		return ""
	}

	return strings.TrimSpace(
		string(bytes.TrimSpace(output)),
	)
}
