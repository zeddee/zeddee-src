---
title: "Don't Email Your Passwords and Other Adventures"
date: 2019-08-23T00:31:24+0800
slug: "dont-email-your-passwords-and-other-adventures"
tags: ["security","passwords"]
draft: false
---

So I signed up for an account with a local tech thing in the early hours of Tuesday morning,
and found my mailbox graced with this email:

> ![Please don't email me my password](/img/2019-posts/please-dont-email-me-my-password.jpg)

I reached out to them via the contact form at ~2AM (20 Aug 2019), and haven't heard back.

The first thing you should notice here is that I've put up my 
user name and password in all it's 36-character glory
in plain text, on my blog. Yes, I have since changed my password.  
No, I still do not know what a manteau is.

The second thing you should notice is that the
[AI Singapore Makerspace Team](https://makerspace.aisingapore.org/)
has emailed me my password.
I don't remember the last time an organization did this for me,
and the reason is probably that emailing your users their
passwords _as is_ is bad.

Let's break it down.

## Email is insecure by default

First, premise: all private data on the internet should be encrypted.

For reference, [Article 5 of the GDPR](https://gdpr.eu/article-5-how-to-process-personal-data/) states that "Personal data shall be: […] processed in a manner that ensures appropriate security of the personal data, including protection against unauthorised or unlawful processing and against accidental loss, destruction or damage, using appropriate technical or organisational measures (‘integrity and confidentiality’)." Here, "appropriate security" would usually be taken to mean "encrypt the damn thing". [^1]

The GDPR also [doesn't require emails in general](https://gdpr.eu/email-encryption/) to be encrypted, but it definitely requires you to keep private data secure.

Second, when we talk encryption, there are usually two layers to consider:

- **Encryption at-rest**, which refers to encryption of stored data (whether on a physical disk or on a cloud service, which ultimately is on a physical disk somewhere).
- **Encryption in-transit**, which refers to encryption of data that you're moving from place to place e.g. sending from your email account to another.

Emails are not guaranteed to be encrypted in-transit.
Gmail [sends mail over SSL](https://transparencyreport.google.com/safer-email/overview)
and with [S/MIME encryption](https://support.google.com/mail/answer/6330403?hl=en)
whenever possible, but for this to work,the
[receiving mailbox](https://www.blog.google/products/gmail/making-email-safer-for-you-posted-by/)
must also support these encryption protocols.

This is not always the case, especially if the route between
your mail server and the receipient's mail server is not
transparent to you. This detail is important because in any email exchange, your sent email may be hopping across a series of mailservers that receive, store, and relay that email to its intended recipient.

Email at-rest is not encrypted unless your email
provider explicitly tells you it is, like in [Proton Mail's](https://protonmail.com/security-details) and in 
[GSuite's documentation](https://support.google.com/googlecloud/answer/6056693?hl=en). This means that when your email lands anywhere allowed by aforementioned TLS/S/MIME encryption, it's stored on the mail server readily read by anyone who gains access to it.

> **Aside**: This is where I insert my standard **PSA to everyone** to **_never_** use your work email for personal communication. 
>
> **Case in point**: I used to work at a company that makes software that scans systems for sensitive data like credit card numbers and personal IDs. One day, one of the senior engineers was showing the team how a new feature he wrote works. He pulled up the web interface for our application, plugged in our work Gmail account (of which he had superadmin access to), and proceeded to 'scan' it. The team watched agape as **_all_** our emails started to fill the screen (the scan went really quickly across our _whole_ Gmail organization; that was the feature he was showing us). As part of the demonstration, he clicked on one of the emails the scan picked up, and showed the team that the application had picked up a test credit card number that had been sent over email. Of course, anyone who has administered an Exchange server or a GSuite organization _knows_ you can read anyone's email, but seeing it happen _en masse_ is another thing altogether.
>
> **Conclusion**: Don't use your work email to send shit.

So, when I received my password in plain text on my Gmail account, my first question was: so who else has it, and will I ever know?

## Secure Password Storage

As a developer, you **must never save passwords in a format you can read**. This is a [security vulnerability](https://www.owasp.org/index.php/Password_Plaintext_Storage),
and **an invasion of privacy** for these reasons:

- A password is a secret, both in the conventional sense and in a development sense. A secret should only be available to the party who generates it and the parties whom it is explicitly shared with.
- As a layperson, I may be reusing passwords. This means that by exposing their password as plain text, you may have compromised other accounts that they own. Less generously, one may assume that you have gained unauthorized access to those accounts.
- Anyone with read access to the user table/database will be able to pull passwords as is from _all_ users (as [OWASP demonstrates here](https://www.owasp.org/index.php/Password_Plaintext_Storage)).

As per [OWASP's Password Storage](https://cheatsheetseries.owasp.org/cheatsheets/Password_Storage_Cheat_Sheet.html) recommendation,
you should be **hashing** passwords. For reference, see how WordPress stores user passwords in a MySQL database:

> ![How WordPress stores user passwords](/img/2019-posts/user-pass.jpg)

Notice that it's _not_ an actual password, but an alphanumeric string representing a SHA256 (salted) hash of the original password.

Here's a small sample Go program that shows how password hashing checking works (You can run it on the [Go playground](https://play.golang.org/p/IzF9AuTX-7m)):

```go
// NOTE: This comes without important features like adding salt to further harden the hashed password
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	hasher := sha256.New()
	encryptMe := []byte("supersecretpassword")
	hasher.Write(encryptMe)
	endHash := hasher.Sum(nil)

	fmt.Println("The SHA256 hash of \"supersecretpassword\" is:", string(endHash[:]))
	
	
	fmt.Println("This password is correct:", checkPassword("supersecretpassword", endHash))
	fmt.Println("This password is wrong:", checkPassword("wrongpassword", endHash))
}

func checkPassword(enteredPassword string, correctPassword []byte) bool {
	hasher := sha256.New()
	hasher.Write([]byte(enteredPassword))
	checkThis := hasher.Sum(nil)
	return (string(correctPassword[:]) == string(checkThis[:]))
}
```

That's how a _rudimentary_ password hashing mechanism looks like.
There are other things that you have to do as a developer to keep your users' passwords secure,
like salting them to prevent rainbow table attacks, and [not sending them in plain text via email](#email-is-insecure-by-default).

## Conclusion

- Don't email me my password.
- You shouldn't know what my password is.
- If you're building an application or system that stores passwords, **only store the hashes**.

Useful resources:

- OWASP Cheatsheets: https://cheatsheetseries.owasp.org
- HaveIBeenPwned, online databse of compromised passwords: https://haveibeenpwned.com/About
