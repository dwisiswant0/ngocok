# ngocok

ngrok Collaborator Link — yet another Burp Collaborator alternative for free with ngrok.

## Prerequisites

* **ngrok authtoken**: Authentication token from your ngrok account is required to establish a ngrok tunnel. See [Your Authtoken](https://dashboard.ngrok.com/get-started/your-authtoken).

> [!TIP]
> If `ngrok` is already configured on your machine with `ngrok config add-authtoken`, you can [install](#installation) and run it directly.
>
> To find the location of the configuration file, consult the [ngrok Agent Configuration File](https://ngrok.com/docs/agent/config/) documentation page.

That's it!

## Installation

> [!NOTE]
> Go version [**1.20+**](https://go.dev/doc/install) should be installed & configured.

```bash
$ go install github.com/dwisiswant0/ngocok@latest
```

### — or

Manually build the executable from the source:

```bash
$ git clone https://github.com/dwisiswant0/ngocok ngocok/
$ cd $_
$ go build . # or
$ go install .
```

## Usage

It's quite straightforward. Simply execute `ngocok`, and you're ready to capture the out-of-band requests!

Here are the supported options:

|      **Flag**     	|                  **Description**                  	| **Default** 	|
|-------------------	|---------------------------------------------------	|-------------	|
| `-e`/`--endpoint` 	| ngrok tunnel endpoint ("`http`" or "`tcp`")       	| "`http`"    	|
| `-t`/`--token`    	| ngrok authentication token                        	| ""          	|
| `--unstrip`       	| Unstrip `X-Forwarded-{For,Host,Proto}` headers    	| `false`     	|
| `-o`/`--output`   	| Log incoming requests to a file instead of stdout 	| ""          	|

> [!IMPORTANT]
> Using a `-t`/`--token` flag will takes precedence over the `NGROK_AUTHTOKEN` environment variable, and using `NGROK_AUTHTOKEN` environment variable will takes precedence over the ngrok config file<sup>[[?](#prerequisites)]</sup>.

### Examples

Start a tunnel with default endpoint (HTTP) and use the ngrok authentication token from the config file.

```bash
$ ngocok
```

Start a tunnel with a TCP endpoint.

```bash
$ ngocok --endpoint tcp
```

Start a tunnel with an HTTP endpoint and provide the ngrok authentication token using the `-t`/`--token` flag.


```bash
$ ngocok --token [authtoken]
```

Start a tunnel with an HTTP endpoint and use the ngrok authentication token from the `NGROK_AUTHTOKEN` environment variable.

```bash
$ NGROK_AUTHTOKEN="..." ngocok
```

Start a tunnel with an HTTP endpoint and log incoming requests to a file.

```bash
$ ngocok --output /path/to/requests.log
```

## Why is this name?

<details>
<summary><code>¯\_(ツ)_/¯</code></summary>

<blockquote>
	<!-- <img src="https://pbs.twimg.com/media/FD6QH3fVcAIMaCt?format=jpg&name=900x900" width="250" alt="judul Cara menghilangkan kebiasaan masturbasi pada anak kecil oleh Enny, dijawab oleh dr. Fadhilah Az Zahro, pertanyaannya Dok kenapa ya anak umur 2 tahun sudah melakukan masturbasi? Gimana cara menghilangkan kebiasaan itu, tiap dimarahi langsung nangis, sudah dibilangin baik-baik kalo itu...judul Cara menghilangkan kebiasaan masturbasi pada anak kecil oleh Enny, dijawab oleh dr. Fadhilah Az Zahro, pertanyaannya Dok kenapa ya anak umur 2 tahun sudah melakukan masturbasi? Gimana cara menghilangkan kebiasaan itu, tiap dimarahi langsung nangis, sudah dibilangin baik-baik kalo itu..." href="#"><br> -->
	<img src="https://i.pinimg.com/originals/fd/ae/d6/fdaed6ceb6fbc9e91759f0efe95c4c4d.jpg" width="150" alt="TUKAN NGOCOK" href="#">
</blockquote>
</details>

## Similar Projects

* [projectdiscovery/interactsh](https://github.com/projectdiscovery/interactsh): An OOB interaction gathering server and client library.

## License

**`ngocok`** is made with ♥ by [**@dwisiswant0**](https://github.com/dwisiswant0) under MIT license. See [LICENSE](/LICENSE). 