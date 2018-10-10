# Welcome to Buffalo!

Thank you for choosing Buffalo for your web development needs.

## Starting the Application

Buffalo ships with a command that will watch your application and automatically rebuild the Go binary and any assets for you. To do that run the "buffalo dev" command:

	$ buffalo dev

If you point your browser to [http://127.0.0.1:3000](http://127.0.0.1:3000) you should see a "Welcome to Buffalo!" page.

**Congratulations!** You now have your Buffalo application up and running.

## What Next?

We recommend you heading over to [http://gobuffalo.io](http://gobuffalo.io) and reviewing all of the great documentation there.


## Environment Variables

I recommend creating a `.env` file for your sensitive app configuration details (it is generated automatically when creating a Buffalo app, but doesnt get uploaded to Github/Gitlab)

Example `.env`

``` 
# This .env file was generated by buffalo, add here the env variables you need 
# buffalo to load into the ENV on application startup so your application works correctly.
# To add variables use KEY=VALUE format, you can later retrieve this in your application
# by using os.Getenv("KEY").
#
# Example:
# DATABASE_PASSWORD=XXXXXXXXX
# SESSION_SECRET=XXXXXXXXX
# SMTP_SERVER=XXXXXXXXX
PORT=8888
```

Good luck!

[Powered by Buffalo](http://gobuffalo.io)
