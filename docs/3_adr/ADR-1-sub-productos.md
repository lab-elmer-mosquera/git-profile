---
tipo: adr
id: ADR-001
status: Aceptada
fecha: 2026-09-18
vision_origen: VA-1
dominio: "Mecanismos de Perfil"
reemplaza:
superado_por:
requisitos_derivados:
  - REQ-001
  - REQ-002
rama: "adr/ADR-001-estrategia-subproductos-beta"
---

# [ADR-001] Estrategia de Validación Incremental mediante Subproductos Betas Derivados

## 🎯 1. Contexto y Justificación (El Método)
La Visión **VA-1** define a Git Profile como una entidad de primer nivel dentro del flujo de Git. Sin embargo, desarrollar directamente el producto final consolidado implica asumir un riesgo técnico, decisiones prematuras sobre el diseño del dominio y una barrera de entrada por el aprendizaje del lenguaje objetivo (Go).

Para tener responder a mis necesidades actuales, se decide adoptar una **estrategia de arquitectura evolutiva basada en subproductos betas independientes**. Esta estrategia permite construir herramientas ligeras orientadas a validar mecanismos específicos de perfil (como resolución de claves SSH o configuración `includeIf`) en entornos reales de trabajo antes de implementar la arquitectura final del producto.

## 🛠️ 2. Decisión Técnica
1. **Estrategia de Subproductos Betas:** El repositorio de Git Profile alojará subproductos derivados que implementen subconjuntos específicos de la visión **VA-1**.
2. **Ciclo de Vida Independiente:** Cada subproducto se distribuirá como un componente desacoplado (binarios o scripts específicos por plataforma como Windows/Linux).
3. **Carácter Transitorio:** Ningún subproducto beta constituye la arquitectura definitiva ni reemplaza el modelo conceptual completo de **VA-1**; actúan como vehículos de validación de hipótesis.

## 🔄 3. Alternativas Evaluadas

* **Opción A (Elegida) — Validación Incremental mediante Subproductos Betas:**
  * *Ventajas:* Permite obtener feedback temprano de usuarios en Windows y Linux, aísla la complejidad técnica y evita el acoplamiento prematuro a una tecnología o lenguaje.
  * *Desventajas:* Requiere mantener temporalmente la gobernanza de herramientas auxiliares en el repositorio.

* **Opción B — Desarrollo Monolítico Directo del Producto Final (Descartada):**
  * *Motivo de rechazo:* Incurre en un  riesgo de sobre-ingeniería y retrasos en la entrega de valor al intentar implementar toda la visión de **VA-1** en Go sin haber probado de primera mano la necesida de que exista todas features de la arquitectura visionada.