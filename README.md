![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.001.jpeg)

**UNIVERSIDAD PERUANA DE CIENCIAS APLICADAS**


Programación Concurrente y Distribuida - CC57

TB4




**INTEGRANTES**

Benites Diaz, Juan Alejandro			u202010449

Tumbalobos Cubas, Kevin William		u201922518

Guevara Dominguez, Sebastian Aaron 		u20181h207


**SECCIÓN**

CC72



**DOCENTE**

Carlos Alberto Jara García




**CICLO 2023-02**


#
<a name="_rs2c1fha6y95"></a>[**Planteamiento del caso de uso	3**](#_1bgjlz3py0rm)**

[**Explicación del algoritmo usado	3**](#_9z6wk8533j62)

[**Resultados	5**](#_e6mimbttus87)

[**Conclusiones	7**](#_kp5dpkz3xr23)

[**Bibliografía	7**](#_t72z6mb3s54b)

[**Anexos	7**](#_ljpw8i4eqc4i)
#
#






























# <a name="_b2i44mbw61xi"></a><a name="_vzsi6oe79r1r"></a><a name="_1bgjlz3py0rm"></a>Planteamiento del caso de uso
**Problema Social:** En las redes sociales modernas, uno de los desafíos más significativos es la personalización de la experiencia del usuario. Sin una segmentación adecuada, las plataformas pueden mostrar contenido irrelevante, ofrecer recomendaciones inexactas y dirigir anuncios que no interesan a los usuarios. Esto no solo disminuye la satisfacción del usuario, sino que también puede reducir la efectividad de los anuncios y la retención de usuarios en la plataforma.

**Características del Problema:**

1. **Heterogeneidad de Usuarios:** Los usuarios de las redes sociales varían ampliamente en términos de edad, ubicación, intereses, comportamientos y preferencias.
1. **Volumen de Datos:** Las redes sociales manejan enormes cantidades de datos que deben procesarse de manera eficiente para extraer información útil.
1. **Necesidad de Personalización:** La capacidad de ofrecer contenido, recomendaciones y anuncios personalizados es crucial para mantener el interés y la participación de los usuarios.
1. **Eficiencia en el Procesamiento:** Procesar grandes volúmenes de datos en un tiempo razonable es un desafío técnico importante.

**Solución Propuesta:** Este proyecto utiliza el algoritmo de k-means para agrupar a los usuarios de una red social en diferentes segmentos basados en atributos específicos como edad, ubicación e intereses. Además, se implementa el procesamiento paralelo mediante goroutines para mejorar la eficiencia del procesamiento de grandes conjuntos de datos.

1. **Segmentación de Usuarios:** Aplicando k-means, se identifican grupos de usuarios con características similares, lo que permite a la red social:
   1. Personalizar el contenido mostrado a cada grupo de usuarios.
   1. Mejorar las recomendaciones de amigos y contenido.
   1. Dirigir anuncios más relevantes y efectivos.
1. **Procesamiento Paralelo:** Utilizando goroutines para dividir el conjunto de datos y procesar las partes en paralelo, se reduce significativamente el tiempo total de procesamiento, permitiendo una actualización más rápida y eficiente de los segmentos de usuarios.
# <a name="_9z6wk8533j62"></a>Explicación del algoritmo usado
A continuación se muestra la función que de manera concurrente realiza el algoritmo de k-means:
# ![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.002.png)
![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.003.png)

<a name="_agbsfd8yqj1e"></a>El algoritmo *kmeans* en el código inicializa *k* clústeres con centroides aleatorios y luego itera 1000 veces para asignar cada punto de datos al clúster más cercano, calculando la distancia entre los puntos y los centroides. Tras la asignación, los centroides se recalculan como el promedio de los puntos asignados a cada clúster. Si un clúster no tiene puntos, su centroide se reasigna aleatoriamente. Este proceso asegura la agrupación de puntos similares y la optimización iterativa de los centroides para mejorar la segmentación.


# <a name="_e6mimbttus87"></a>Resultados
Una vez ejecutado el código, se muestra los siguientes resultados:
# ![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.004.png)
<a name="_by6l3drg4vro"></a>Los resultados mostrados corresponden a la salida del servidor que utiliza el algoritmo de k-means para la segmentación de datos. Cada línea en la consola representa un conjunto de datos procesados, que incluye un identificador, las coordenadas de los puntos en el espacio de características, y el clúster al que pertenecen. Estos datos permiten observar cómo los usuarios han sido agrupados según sus similitudes en atributos específicos, facilitando así la personalización del contenido y la mejora en la eficiencia del procesamiento de grandes volúmenes de datos

Diferentes Simulaciones

![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.005.png)

![](Aspose.Words.c6fc557d-d73c-4504-ad02-671650eb3a93.006.png)
# <a name="_kp5dpkz3xr23"></a>Conclusiones
- La implementación de la segmentación de usuarios mediante el algoritmo de k-means ha demostrado ser altamente efectiva para mejorar la personalización del contenido en las redes sociales. Al identificar grupos de usuarios con características similares, la plataforma puede ofrecer experiencias más relevantes y atractivas, lo que resulta en una mayor satisfacción y retención de los usuarios. Los usuarios están más inclinados a interactuar con la plataforma cuando el contenido que reciben está alineado con sus intereses y necesidades específicas.
- Además, la utilización de goroutines para el procesamiento paralelo de datos ha permitido una significativa reducción en el tiempo necesario para manejar grandes volúmenes de datos. Esta eficiencia en el procesamiento no solo mejora la escalabilidad del sistema, sino que también asegura que las actualizaciones y personalizaciones se realicen de manera oportuna. En un entorno donde los datos de usuario crecen continuamente, esta capacidad para procesar información rápidamente es crucial para mantener la relevancia y efectividad de la segmentación.
- Por último, la optimización de anuncios a través de una mejor segmentación de usuarios permite a los anunciantes dirigir sus campañas de manera más precisa y efectiva. Esto no solo mejora el retorno de inversión en publicidad, sino que también enriquece la experiencia del usuario al reducir la exposición a anuncios irrelevantes. En conjunto, estos beneficios posicionan a la red social como una plataforma más eficiente y atractiva tanto para los usuarios como para los anunciantes, asegurando un crecimiento sostenido y una mayor competitividad en el mercado.

# <a name="_t72z6mb3s54b"></a>Bibliografía
*Go Concurrency Patterns: Pipelines and cancellation - The Go Programming Language*. (s. f.). <https://go.dev/blog/pipelines>

(S/f). Spinroot.com. Recuperado el 1 de mayo de 2024, de https://spinroot.com/spin/whatispin.html
# <a name="_ljpw8i4eqc4i"></a>Anexos
Link del repositorio de GitHub:

` `<https://github.com/MrPepePollo/TF_Concurrente> 

Link del video: 
https://www.youtube.com/watch?v=EcyrJlc9BV8








