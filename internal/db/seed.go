package db

import (
	"context"
	"log"

	"github.com/Elevate-Techworks/social/internal/store"
)

var userNames = []string{
	"Alejandro", "Beatriz", "Carlos", "Daniela", "Eduardo", "Fernanda", "Gabriel", "Helena", "Ignacio", "Julia",
	"Kevin", "Lucia", "Manuel", "Natalia", "Oscar", "Patricia", "Ricardo", "Sofia", "Tomas", "Ursula",
	"Victor", "Wendy", "Xavier", "Yolanda", "Zacarias", "Adriana", "Bernardo", "Carmen", "David", "Elena",
	"Felipe", "Gloria", "Hugo", "Isabel", "Javier", "Karla", "Luis", "Marta", "Nicolas", "Olivia",
	"Pablo", "Raquel", "Santiago", "Teresa", "Ulises", "Valeria", "William", "Ximena", "Yosef", "Zoe",
}
var postTitles = []string{
	"El arte de la programación funcional",
	"Introducción a los microservicios",
	"Optimización de bases de datos SQL",
	"Patrones de diseño en Go",
	"Construyendo APIs RESTful escalables",
	"Guía completa de Docker para desarrolladores",
	"Entendiendo Kubernetes desde cero",
	"Testing automatizado en aplicaciones web",
	"Arquitectura limpia en proyectos reales",
	"Inteligencia artificial aplicada al desarrollo",
	"Mejores prácticas de seguridad web",
	"Ciencia de datos con Python",
	"Desarrollo mobile multiplataforma",
	"Integración continua y despliegue continuo",
	"Manejo de estado en aplicaciones frontend",
	"WebSockets y comunicación en tiempo real",
	"Caché y optimización de rendimiento",
	"GraphQL vs REST: cuándo usar cada uno",
	"Programación concurrente en Go",
	"Machine learning para principiantes",
	"Refactorización de código legado",
	"Sistemas distribuidos explicados",
	"Autorización y autenticación con JWT",
	"Monitoreo y observabilidad en producción",
	"Event-driven architecture patterns",
	"Serverless computing en la práctica",
	"Contenedores vs máquinas virtuales",
	"Bases de datos NoSQL: MongoDB y Redis",
	"Microfrontends: arquitectura moderna",
	"Automatización de infraestructura con Terraform",
	"Clean code y principios SOLID",
	"Debugging avanzado de aplicaciones",
	"Compiladores e intérpretes por dentro",
	"Blockchain y contratos inteligentes",
	"Desarrollo de videojuegos indie",
	"Realidad aumentada con WebAR",
	"Procesamiento de lenguaje natural",
	"Sistemas de recomendación personalizados",
	"Análisis de rendimiento de APIs",
	"DevOps para startups",
	"Git avanzado: flujos de trabajo profesionales",
	"Diseño de interfaces accesibles",
	"PWA: aplicaciones web progresivas",
	"Streaming de video en la web",
	"Criptografía aplicada a la web moderna",
	"Bases de datos vectoriales y búsqueda semántica",
	"Edge computing y sus casos de uso",
	"Observabilidad con OpenTelemetry",
	"Desarrollo ético de software",
	"El futuro del desarrollo web",
}
var postContents = []string{
	"La programación funcional nos permite escribir código más predecible y testeable. En este artículo exploramos los principios fundamentales como la inmutabilidad, las funciones puras y la composición de funciones aplicados a proyectos reales.",
	"Los microservicios han revolucionado la forma en que construimos aplicaciones. Aprenderemos a dividir un monolito en servicios independientes, cada uno con su propia base de datos y ciclo de vida.",
	"El rendimiento de las consultas SQL puede hacer o deshacer una aplicación. Exploramos índices, query plans y técnicas de optimización que todo desarrollador debería conocer.",
	"Los patrones de diseño son soluciones probadas a problemas comunes. Desde Singleton hasta Observer, veremos cómo implementarlos de manera elegante en Go.",
	"Las APIs RESTful son el estándar de la industria. Descubriremos cómo diseñar endpoints intuitivos, manejar versionado y documentar con OpenAPI.",
	"Docker ha cambiado la forma en que desplegamos software. Desde Dockerfile hasta docker-compose, esta guía cubre todo lo necesario para containerizar aplicaciones.",
	"Kubernetes puede parecer intimidante al principio, pero una vez entendidos los conceptos de pods, services y deployments, todo comienza a tener sentido.",
	"El testing automatizado es esencial para mantener la calidad del código. Unit tests, integration tests y end-to-end tests forman el pirámide de testing moderna.",
	"La arquitectura limpia propone separar el código en capas con responsabilidades bien definidas. El dominio permanece independiente de frameworks y bases de datos.",
	"La IA está transformando el desarrollo de software. Desde la generación de código hasta el testing automático, las herramientas de IA aumentan nuestra productividad.",
	"La seguridad web no es opcional. XSS, CSRF, SQL injection y otras vulnerabilidades deben ser entendidas y prevenidas desde el diseño de la aplicación.",
	"Python se ha convertido en el lenguaje dominante en ciencia de datos. Con bibliotecas como pandas, numpy y scikit-learn, el análisis de datos es más accesible que nunca.",
	"React Native y Flutter permiten crear aplicaciones nativas desde una sola base de código. Analizamos las ventajas y desventajas de cada enfoque.",
	"CI/CD automatiza el proceso de integración y despliegue. Con pipelines bien configurados, cada commit puede ser probado y desplegado automáticamente.",
	"El manejo de estado es uno de los mayores desafíos en el frontend. Redux, Zustand y Context API ofrecen diferentes soluciones para diferentes escalas.",
	"Los WebSockets permiten comunicación bidireccional en tiempo real. Perfectos para chats, notificaciones y actualizaciones en vivo sin polling constante.",
	"El caché puede mejorar el rendimiento órdenes de magnitud. Redis, Memcached y CDNs son herramientas esenciales en el arsenal de optimización.",
	"GraphQL ofrece flexibilidad en las consultas pero añade complejidad. Analizamos cuándo vale la pena el cambio y cuándo REST sigue siendo la mejor opción.",
	"Go brilla en programación concurrente con goroutines y channels. Patrones como worker pools y fan-in/fan-out permiten aprovechar al máximo los recursos.",
	"Machine learning no requiere un doctorado en matemáticas. Con frameworks como TensorFlow y scikit-learn, cualquier desarrollador puede crear modelos predictivos.",
	"Refactorizar código legado es un arte. Técnicas como el strangler pattern permiten modernizar sistemas sin interrumpir el servicio.",
	"Los sistemas distribuidos presentan desafíos únicos: consistencia, disponibilidad y tolerancia a particiones. El teorema CAP nos obliga a tomar decisiones difíciles.",
	"JWT se ha convertido en el estándar para autenticación stateless. Pero su implementación requiere cuidado con la expiración, el refresh y el almacenamiento seguro.",
	"Sin observabilidad, estamos operando a ciegas. Métricas, logs y traces forman los tres pilares que nos permiten entender el comportamiento de nuestros sistemas.",
	"La arquitectura basada en eventos desacopla los componentes del sistema. Kafka y RabbitMQ son las herramientas más populares para implementar este patrón.",
	"Serverless permite ejecutar código sin gestionar servidores. AWS Lambda, Cloud Functions y Azure Functions ofrecen escalabilidad automática y pago por uso.",
	"Los contenedores comparten el kernel del host mientras las VMs virtualizan hardware completo. Cada enfoque tiene su lugar en la arquitectura moderna.",
	"MongoDB y Redis representan dos enfoques diferentes de NoSQL. Documentos versus clave-valor, cada uno optimizado para diferentes patrones de acceso.",
	"Los microfrontends extienden los principios de microservicios al frontend. Equipos independientes pueden desarrollar y desplegar partes de la UI sin coordinarse.",
	"Terraform permite definir infraestructura como código. Con un solo archivo podemos crear redes, instancias y bases de datos de manera reproducible.",
	"El código limpio no es un lujo, es una necesidad. Los principios SOLID guían el diseño de software mantenible y extensible a largo plazo.",
	"Debugging efectivo requiere más que print statements. Delve, pprof y las herramientas de profiling nos ayudan a encontrar cuellos de botella y bugs sutiles.",
	"Entender cómo funcionan los compiladores nos hace mejores programadores. Lexer, parser, AST y code generation son las etapas fascinantes de la compilación.",
	"Blockchain va más allá de las criptomonedas. Los contratos inteligentes permiten crear aplicaciones descentralizadas con lógica ejecutable en la cadena.",
	"El desarrollo de videojuegos indie es más accesible que nunca. Godot y Unity ofrecen herramientas profesionales gratuitas para creadores independientes.",
	"WebAR trae realidad aumentada al navegador sin necesidad de apps nativas. Con bibliotecas como AR.js, experiencias inmersivas están a un click de distancia.",
	"El procesamiento de lenguaje natural permite a las máquinas entender texto humano. Desde análisis de sentimiento hasta traducción automática, las posibilidades son infinitas.",
	"Los sistemas de recomendación impulsan Netflix, Amazon y Spotify. Algoritmos de filtrado colaborativo y contenido crean experiencias personalizadas.",
	"Analizar el rendimiento de APIs revela cuellos de botella ocultos. Profiling, tracing y métricas customizadas nos dan visibilidad completa del sistema.",
	"DevOps en startups requiere pragmatismo. No se necesita toda la parafernalia enterprise; con las herramientas correctas se logra automatización efectiva.",
	"Git branch, rebase, cherry-pick y bisect son herramientas poderosas. Flujos como Git Flow y trunk-based development organizan el trabajo en equipo.",
	"La accesibilidad web no es solo ética, es legal. WCAG, ARIA labels y testing con screen readers aseguran que todos puedan usar nuestras aplicaciones.",
	"Las PWAs combinan lo mejor de web y nativo. Service workers, manifest y caché offline crean experiencias tipo app desde el navegador.",
	"Streaming de video requiere optimización de bitrate, codecs y CDN. Protocolos como HLS y DASH adaptan la calidad al ancho de banda disponible.",
	"La criptografía moderna protege nuestras comunicaciones. TLS, certificados y cifrado de extremo a extremo son fundamentales para la privacidad.",
	"Las bases de datos vectoriales como Pinecone y Weaviate permiten búsqueda semántica. Embeddings y similitud coseno revolucionan cómo encontramos información.",
	"Edge computing acerca el procesamiento al usuario. Cloudflare Workers y Deno Deploy ejecutan código en cientos de ubicaciones globales simultáneamente.",
	"OpenTelemetry estandariza la telemetría de aplicaciones. Traces distribuidos, métricas y logs se integran en un solo estándar vendor-neutral.",
	"El desarrollo ético considera el impacto social del software. Privacidad, sesgos algorítmicos y transparencia son responsabilidades del desarrollador moderno.",
	"El futuro del desarrollo web incluye WebAssembly, AI-assisted coding y edge computing. Las herramientas evolucionan pero los fundamentos permanecen.",
}
var postTags = [][]string{
	{"funcional", "paradigmas", "código limpio"},
	{"microservicios", "arquitectura", "backend"},
	{"sql", "optimización", "bases de datos"},
	{"go", "patrones de diseño", "arquitectura"},
	{"api", "rest", "backend"},
	{"docker", "contenedores", "devops"},
	{"kubernetes", "orquestación", "devops"},
	{"testing", "calidad", "automatización"},
	{"arquitectura limpia", "código limpio", "diseño"},
	{"ia", "machine learning", "productividad"},
	{"seguridad", "web", "vulnerabilidades"},
	{"python", "ciencia de datos", "análisis"},
	{"mobile", "react native", "flutter"},
	{"cicd", "devops", "automatización"},
	{"frontend", "estado", "react"},
	{"websockets", "tiempo real", "comunicación"},
	{"caché", "rendimiento", "redis"},
	{"graphql", "rest", "api"},
	{"go", "concurrencia", "goroutines"},
	{"machine learning", "tensorflow", "principiantes"},
	{"refactorización", "código legado", "modernización"},
	{"sistemas distribuidos", "cap", "arquitectura"},
	{"jwt", "autenticación", "seguridad"},
	{"monitoreo", "observabilidad", "producción"},
	{"eventos", "kafka", "arquitectura"},
	{"serverless", "aws lambda", "cloud"},
	{"contenedores", "vm", "infraestructura"},
	{"nosql", "mongodb", "redis"},
	{"microfrontends", "frontend", "arquitectura"},
	{"terraform", "iac", "infraestructura"},
	{"clean code", "solid", "buenas prácticas"},
	{"debugging", "profiling", "herramientas"},
	{"compiladores", "intérpretes", "teoría"},
	{"blockchain", "smart contracts", "web3"},
	{"videojuegos", "indie", "godot"},
	{"webar", "realidad aumentada", "navegador"},
	{"nlp", "machine learning", "texto"},
	{"recomendación", "algoritmos", "personalización"},
	{"api", "rendimiento", "tracing"},
	{"devops", "startups", "automatización"},
	{"git", "control de versiones", "colaboración"},
	{"accesibilidad", "wcag", "diseño"},
	{"pwa", "service workers", "mobile"},
	{"streaming", "video", "hls"},
	{"criptografía", "tls", "privacidad"},
	{"vectores", "búsqueda semántica", "embeddings"},
	{"edge computing", "cloudflare", "rendimiento"},
	{"opentelemetry", "observabilidad", "tracing"},
	{"ética", "privacidad", "responsabilidad"},
	{"futuro", "webassembly", "tendencias"},
}
var commentTexts = []string{
	"Excelente artículo, me aclaró muchas dudas.",
	"No estaba de acuerdo con esto al principio pero me convenciste.",
	"¿Podrías profundizar más en este tema?",
	"Lo implementé en mi proyecto y funcionó de maravilla.",
	"Gracias por compartir, llevaba tiempo buscando algo así.",
	"Me gustaría ver un ejemplo práctico de esto.",
	"Interesante punto de vista, nunca lo había pensado así.",
	"Esto debería ser parte de todo currículo de desarrollo.",
	"¿Hay alguna librería que recomiendes para esto?",
	"El rendimiento mejoró significativamente tras aplicar estos consejos.",
	"Me costó entenderlo al principio pero con los ejemplos quedó claro.",
	"¿Cómo manejarías esto en un entorno de producción?",
	"Gran aporte, lo comparto con mi equipo.",
	"Tenía un bug exactamente así, esto lo solucionó.",
	"Muy bien explicado, ojalá más contenido así.",
	"¿Sería aplicable esto a arquitecturas más grandes?",
	"Me salvaste horas de investigación, gracias.",
	"Solo una corrección menor: el enfoque ha evolucionado un poco.",
	"Esto cambió completamente mi forma de pensar sobre el tema.",
	"¿Podrías hacer una comparación con otras alternativas?",
	"Lo probé en Go y los resultados fueron impresionantes.",
	"El artículo tiene un nivel técnico perfecto.",
	"Me pregunto cómo escalaría esto con más tráfico.",
	"Gran contenido como siempre, sigue así.",
	"Esto debería ser obligatorio leerlo para juniors.",
}

func Seed(s store.Storage) error {
	ctx := context.Background()
	users := generateUsers(100)
	for _, u := range users {
		if err := s.Users.Create(ctx, u); err != nil {
			log.Println("failed to create user:", err)
			return err
		}
	}

	posts := generatePosts(50, users)
	for _, p := range posts {
		if err := s.Posts.Create(ctx, p); err != nil {
			log.Println("failed to create post:", err)
			return err
		}
	}

	comments := generateComments(100, posts, users)
	for _, c := range comments {
		if err := s.Comments.Create(ctx, c); err != nil {
			log.Println("failed to create comment:", err)
			return err
		}
	}

	log.Println("Seed completed successfully")
	return nil
}

func generateUsers(num int) []*store.User {
	users := make([]*store.User, num)

	for i := range num {
		users[i] = &store.User{
			Username: userNames[i%len(userNames)],
			Email:    userNames[i%len(userNames)] + string(rune('0'+i%10)) + "@example.com",
			Password: "123123",
		}
	}
	return users
}

func generatePosts(num int, users []*store.User) []*store.Post {
	posts := make([]*store.Post, num)

	for i := range num {
		posts[i] = &store.Post{
			UserId:  int64(users[i%len(users)].ID),
			Title:   postTitles[i],
			Content: postContents[i],
			Tags:    postTags[i],
		}
	}
	return posts
}

func generateComments(num int, posts []*store.Post, users []*store.User) []*store.Comment {
	comments := make([]*store.Comment, 0, num)

	for i := range num {
		comments = append(comments, &store.Comment{
			PostID:  posts[i%len(posts)].ID,
			UserID:  int64(users[(i+3)%len(users)].ID),
			Content: commentTexts[i%len(commentTexts)],
		})
	}
	return comments
}
