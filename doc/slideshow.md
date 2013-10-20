title: Scramble your Email
author:
  name: "Jae Kwon"
  url: "hashed.im"
output: slideshow.html
controls: true

--

# Scramble
## NSA-proof Secure Webmail

--

### HTTPS is Insecure

* Certificate Authorities
* Perfect Forward Secrecy
* Government Order

--

### Certificate Authorities (CA)

<div style="font-size: 1.5em">
<pre>


$ ls -lt /etc/ssl/certs/ | wc -l
458
</pre>
</div>

--

### Certificate Authorities (CA)

* March 2011 Comodo<br/>
  `mail.google.com`,<br/>
  `login.yahoo.com`, ...
* August 2011 Diginotar<br />
  `*.google.com`
* FireFox & CNNIC

<i>Top down security isn't working!</i>

--

### Convergence.io

An alternative!

<i>"allows you to configure a dynamic set of <b>Notaries</b> which use network perspective to validate your communication"</i>

--

### Convergence.io

<img src="http://i.imgur.com/r51q1g5.png"/>

--

### Convergence.io

<img src="http://i.imgur.com/uKjWeKJ.png"/>

--

Except,

--

### Perfect Forward Secrecy

<b>Forward secrecy</b> requires that the private keys for a connection are not kept in <b>persistent storage</b>.

"Retroactive decryption"

<!-- Without perfect forward secrecy, if the server's private key is compromised, not only will all future TLS-encrypted sessions using that server certificate be compromised, but also any past sessions that used it as well -->

--

### Perfect Forward Secrecy

<b>Ephemeral Diffie-Hellman key exchange</b>

Lets two computers to establish a shared secret over an insecure communication channel.

(Magic!)

--

Except,

--

### Government Order

<img src="http://imgs.xkcd.com/comics/security.png"/>

--

### Government Order

* PRISM
* Lavabit
* Silent Circle (Zimmerman)
* TorMail (Marques)
* HushMail

--

<b>(How) can we secure email?</b>

--

### Secure email

* Best crypto primitives
* Webmail w/ untrusted servers
* Open source & federated

--

### Best crypto primitives

* RSA2048 for signing
* RSA2048 & AES128 for encryption

OpenPGP supports this!

--

### OpenPGP

Message:

<pre style="color:red">
To, Alice

Eve is sketchy

 - Bob
</pre>

--

### OpenPGP

Message:
<pre style="color:red">
To, Alice

Eve is sketchy

 - Bob
</pre>

Session key:
<pre style="color:red">
"SECRET_AES_SESSION_KEY_!@#K!J@#"
</pre>

--

### OpenPGP

Cipher Message:
<pre style="color:green">
+Q5QwQM+y8/y2OzrnZAt63RRmmm8AHmO6RrIWhvmb47BcWofaehMGAitrEEz
Oc4sn/Nn15P2/Ffch9X929gYqj2Bq9zFIbo9bTDBGvNgQ6mnPY7E/9nD9p6X
</pre>

Session key:
<pre style="color:red">
"SECRET_AES_SESSION_KEY_!@#K!J@#"
</pre>

--

### OpenPGP

Cipher Message:
<pre style="color:green">
+Q5QwQM+y8/y2OzrnZAt63RRmmm8AHmO6RrIWhvmb47BcWofaehMGAitrEEz
Oc4sn/Nn15P2/Ffch9X929gYqj2Bq9zFIbo9bTDBGvNgQ6mnPY7E/9nD9p6X
</pre>

Session key:
<pre style="color:red">
"SECRET_AES_SESSION_KEY_!@#K!J@#"
</pre>

Recipient public key:
<pre style="color:green">-----BEGIN PGP PUBLIC KEY BLOCK-----
xsBNBFI/k+wBCADO3eL0Beu5Hqeot4aRTO3ijSD1ddkCiEpTfnd1pCG72E72
wLxsqMt+lI3gVNxeje6eqFlc9K6PrP9hAScQKM0f6wp2NCqfdWmGk9NvTyVp
...
-----END PGP PUBLIC KEY BLOCK-----
</pre>


--

### OpenPGP

Cipher Message:
<pre style="color:green">
+Q5QwQM+y8/y2OzrnZAt63RRmmm8AHmO6RrIWhvmb47BcWofaehMGAitrEEz
Oc4sn/Nn15P2/Ffch9X929gYqj2Bq9zFIbo9bTDBGvNgQ6mnPY7E/9nD9p6X
</pre>

Encrypted Session key:
<pre style="color:green">
ilLIn+VuSZYR4Xym0BTX0r/gKKMVfNBgRPIxuHhtFBEyisCWwhDxrDsHP/vU
I9UqeqYQl67o+DcVWAcD0NUpCqtp2OH0s8TVn5Wac7g7n6MK7bPGCnschPfR
</pre>

Recipient Public Key:
<pre style="color:green">-----BEGIN PGP PUBLIC KEY BLOCK-----
xsBNBFI/k+wBCADO3eL0Beu5Hqeot4aRTO3ijSD1ddkCiEpTfnd1pCG72E72
wLxsqMt+lI3gVNxeje6eqFlc9K6PrP9hAScQKM0f6wp2NCqfdWmGk9NvTyVp
...<br/>-----END PGP PUBLIC KEY BLOCK-----
</pre>


--

### OpenPGP

Encrypted Message:
<pre style="color:green">-----BEGIN PGP MESSAGE BLOCK-----
xsBNBFJEx2kBCADzzfl2qJMPKHND7VFDJgDPc9U01kuzb1Mlo+wUm0YezfYQZbHn
7ZaAiYXLw3hLzSVXnve3mva2vhxmFQQW4GovR9dysQ/i9jcTOwmwMyBWrIz1ncbp
...<br/>-----END PGP MESSAGE BLOCK-----
</pre>

Recipient Public Key:
<pre style="color:green">-----BEGIN PGP PUBLIC KEY BLOCK-----
xsBNBFI/k+wBCADO3eL0Beu5Hqeot4aRTO3ijSD1ddkCiEpTfnd1pCG72E72
wLxsqMt+lI3gVNxeje6eqFlc9K6PrP9hAScQKM0f6wp2NCqfdWmGk9NvTyVp
...<br/>-----END PGP PUBLIC KEY BLOCK-----
</pre>


--

### Webmail w/ untrusted servers

--

### Webmail w/ untrusted servers

Encrypt all of the things!

* Mail subject & body
* Contacts
* Index blobs (future)

Also, store the encrypted private key on the server using `Scrypt(<PassPhrase>)`

--

### Webmail w/ untrusted servers

Can't trust the server,<br/>
Can't trust the javascript.

* Browser plugin
* HTML App Cache

--

### Name Resolution

--

### Name Resolution

Hi, I am:

```
-----BEGIN PGP PUBLIC KEY BLOCK-----
xsBNBFI/k+wBCADO3eL0Beu5Hqeot4aRTO3ijSD1ddkCiEpTfnd1pCG72E72
wLxsqMt+lI3gVNxeje6eqFlc9K6PrP9hAScQKM0f6wp2NCqfdWmGk9NvTyVp
3WiXQhDNucm+c79FgyvYiIUM8Xjt5AfOQNQ8dyqgoSiQR7lwGbbmen/C4aKS
ilLIn+VuSZYR4Xym0BTX0r/gKKMVfNBgRPIxuHhtFBEyisCWwhDxrDsHP/vU
I9UqeqYQl67o+DcVWAcD0NUpCqtp2OH0s8TVn5Wac7g7n6MK7bPGCnschPfR
mAadQ/Z453AseIlQEtUGV0cia6AP8hoT30/Lh++8TMTB0LMpVzJXqvVjABEB
AAHNAMLAXAQQAQIAEAUCUj+T8AkQS0aTOFk3BcEAAFiwCACjxQo/3bnNOjL4
+Q5QwQM+y8/y2OzrnZAt63RRmmm8AHmO6RrIWhvmb47BcWofaehMGAitrEEz
Oc4sn/Nn15P2/Ffch9X929gYqj2Bq9zFIbo9bTDBGvNgQ6mnPY7E/9nD9p6X
VZW5lSgoXMgkDuuWf1uG38pqJCu9m4YLHzYgqBeUmfMjy5ZWzfc1Z4DQEzNy
T2xaAnhKt8RRIahSl3vqti6Acy5ZHE+GEXzL89D6yQV8uDJCRbaHdMUNPAc+
KyO/vJuje0QPxkloHmy0JKmOIuB+dPkxYaIQUKubm0xoh6BWWg1I3obE7cCn
1IRp6TufNOyBKlkL+2tCczTHdNP5
=o8xa
-----END PGP PUBLIC KEY BLOCK-----
```

--

### Name Resolution

Hi, I am: `tntlt4ltqc74asv7@hashed.im`

Still hard to remember!

--

### Name Resolution

Hi, I am: `jaekwon@hashed.im`

* NameCoin?
* Public Key servers?

--

### Name Resolution

* 1B addresses require < 100Gb to store
* Distributed notary system

Like Convergence.io

--

### Name Resolution

* Client has a list of trusted notaries,<br/>
  `hashed.im` & `scramble.io`
* Client asks server,<br/>
    "Tell me what `hashed.im` and `scramble.io` think `jaekwon@hashed.im` resolves to"
* Server forwards query to `hashed.im` & `scramble.io`
* Server aggregates & responds
* Client makes a decision

--

### Open source & federated

* What if Lavabit had been open source?

--

### Open source & federated

* What if Lavabit had been open source?
* Federation == Shutdown Resistance

--

### Open source & federated

* What if Lavabit had been open source?
* Federation == Shutdown Resistance
* Federation == Notary Perspectives

--

### Introducing, Scramble

It already exists

http://github.com/dcposch/scramble

Signup at https://hashed.im

--

### Scramble + Future

* Normal email compatibility
* Enigmail compatibility
* Search w/ encrypted indices
* Chat

--

### Questions?
