/*
Package gmail is a simple Go library for sending emails from a Gmail account.

NB: The attachment code was inspired by scorredoira's email (https://github.com/scorredoira/email)  and full credit goes
to him.

		package main

		import "github.com/SlyMarbo/gmail"

		func main() {
			email := gmail.Compose("Email subject", "Email body")
			email.From = "username@gmail.com"
			email.Password = "password"

			// Normally you'll only need one of these, but I thought I'd show both.
			email.AddRecipient("recipient@example.com")
			email.AddRecipients("another@example.com", "more@example.com")

			err := email.Send()
			if err != nil {
				// handle error.
			}
		}

Note:

If you have problems with authentication, be sure to check your password settings. While
developing the package, I had forgotten that I have application-specific passwords enabled
on my Google account, so my account password wasn't working; I had to sign into my
account and create an application-specific password for the package and use that.
*/
package gmail
