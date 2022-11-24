# Introdução
Esse projeto é uma prova de conceito para demonstrar um servidor de streaming para a matéria de Sistemas Distribuídos administrada pelo Prof. Rodrigo Rosa (2022/2)

## Membros

- André Plass
- César Hoffmann
- Vanessa Schenkel

## Gerando conteúdo fracionado para streaming

1. Instale o pacote ffmpeg para poder fracionar as mídias a serem disponibilizadas via streaming

```ssh
brew install ffmpeg
```

2. Fracione a mídia em questão usando o comando sugerido abaixo

```ssh
ffmpeg -i musica_a_ser_fracionada.mp3 -c:a libmp3lame -b:a 128k -map 0:0 -f segment -segment_time 5 -segment_list lista_de_fracoes.m3u8 -segment_format mpegts fracao%03d.ts

```

Os arquivos HLS gerados irão ser usados pelo servidor para servirmos via HTTP! :)

3. Adicione esses arquivos ao diretório `files` para que sejam consumidos pelo cliente

## Local development

1. Instale Golang na sua máquina local. Se estiver usando OSX é através do comando

```sh
brew install go
```

2. Inicie o servidor

```sh
go run main.go
```