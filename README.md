# Ionix-challenge

Prueba técnica Backend

Desarrolla un servicio API REST para coordinar registros de vacunación
Requerimientos
● La base de datos a utilizar debe ser Postgres.
● La tecnología para crear el servicio debe ser Go.
Modelos
Los modelos del servicio deben ser 3: User, Drug, Vaccination.
User
Campo Tipo de datos Descripción
id integer identificador único.
name string Nombre del usuario, puede ser nulo.
email string Email del usuario, se utiliza para iniciar sesión.
password string Contraseña de usuario.

Drug
Campo Tipo de datos Descripción
id integer Identificador único.
name string Nombre de la droga.
approved boolean Indica si está aprobada para uso.
min_dose integer Dosis mínima.
max_dose integer Dosis máxima.
available_at datetime Desde esta fecha esta permitida utilizarse.

Vaccination
Campo Tipo de datos Descripción
id integer identificador único.
name string Nombre de la persona que se vacunara.
drug_id integer ID de la droga a vacunar.
dose integer Dosis por utilizar en la vacunación.
date datetime Fecha y hora de la vacunación.

API REST
Crear rutas para crear instancias de los modelos con las siguientes características:
Rutas de Auth:
● POST /signup -> Crear instancia de “user” entregando los campos de “name”,
“email’ y “password”.
● POST /login -> Capturar el “email” y “password” para iniciar sesion. En caso de que
las credenciales sean correctas devolver un JWT con vencimiento definido en variable
de entorno.
Rutas de Drugs:
Para ejecutar todas estas rutas deben validarse primero si el JWT es válido.
● POST /drugs -> Crear instancia de “drug” entregando los campos de “name”,
“approved”, “min_dose”, “max_dose”, “available_at”.
● PUT /drugs/:id -> Ruta para poder actualizar instancia de “drug”.
● GET /drugs -> Obtener todas las instancias de “drug”.
● DELETE /drugs/:id -> Eliminar instancia de “drug”.
Rutas de Vaccination:
Para ejecutar todas estas rutas deben validarse primero si el JWT es válido.
● POST /vaccination -> Crear instancia de “vaccination” entregando los campos de
“name”, “drug_id”, “dose”, “date”. Se debe validar que la dosis este dentro de lo
permitido y que la fecha de vacunación sea posterior a la fecha de uso permitida.
● PUT /vaccination /:id -> Actualizar instancia de “vaccination”.
● GET / vaccination -> Obtener todas las instancias de “vaccination”.
● DELETE / vaccination /:id -> Eliminar instancia de “vaccination”.
Opcional
Si tienes tiempo puedes incluir algunas de las siguientes funcionalidades:
● Crear tests.
● Crear un Dockerfile y/o Docker Compose para levantar el servicio y base de datos.
● Crear un pipeline con Github Actions, Gitlab CI o Jenkins. Debe ejecutar los tests, hacer el
build de la aplicación con el Dockerfile y publicar la imagen en algún registry como “Github
Packages”, “Gitlab Container Registry” o “Dockerhub”.

Comandos:

go test ./...
