
/*
Propósito y Objetivo
Esta API REST ha sido desarrollada siguiendo el enfoque de arquitectura hexagonal (Ports and Adapters), con el objetivo de ofrecer una solución modular, desacoplada y fácilmente testeable. El diseño permite que la lógica de negocio se mantenga independiente de detalles de infraestructura, facilitando su evolución y mantenimiento a largo plazo.

El propósito principal de esta API CORE, para ser desplegada en la capa CORE de la arquitectura Banco, es exponer operaciones que permitan interactuar con las entidades del dominio, incluyendo:

Gestión de personas: A través de operaciones CRUD (crear, leer, actualizar y eliminar), se permite el manejo completo de la entidad Persona. Estas operaciones ilustran la interacción tradicional con bases de datos relacionales bajo un enfoque orientado a servicios.

Gestión de documentos: Se incorpora el uso de MongoDB Atlas como base de datos NoSQL para almacenar y consultar documentos estructurados de forma flexible. La entidad Documento representa un ejemplo concreto de cómo trabajar con documentos JSON, permitiendo operaciones que se adaptan a estructuras de datos dinámicas.

Esta API sirve tanto como base técnica para proyectos que requieran una separación clara entre lógica de negocio e infraestructura, como también como ejemplo educativo para ilustrar la implementación práctica de la arquitectura hexagonal en conjunto con tipo de almacenamiento relacional.
*/
   // String de conexión SqlServer: sqlserver://[SERVIDOR]?database=[BASE_DE_DATOS]&user id=[USUARIO]&password=[CLAVE]

      Ejemplo BD AZURE: sqlserver://svr-persona.database.windows.net?database=bd-arquitectura&user id=arq&password=City8990**2002

      Ejemplo BD LOCAL: sqlserver://localhost?database=MinimalAPIBD&user id=sa&password=santiago


  

