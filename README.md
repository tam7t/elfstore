# elfstore

I recently read [An unlikely database migration](https://tailscale.com/blog/an-unlikely-database-migration/) where
Tailscale described their first database implementation - writing one big `json` object to a text file.

This got me thinking. There had to be a better way. Sure, `json` and a single text file is pretty simple, but what if we
could do more. Wouldn't it be even more simple, and easier to maintain, if there was no file at all?

If your organization isn't quite ready for [no-code](https://github.com/kelseyhightower/nocode), then maybe a stepping
stone would be no-data.

Enter `elfstore`.

`elfstore` allows you to persist your application's data without the need for a database, or even a file. `elfstore`
does this by modifying itself in-place.

Need to migrate you data to a new server? Easy! just copy your binary.

Thank you for coming to my Ted Talk.

## Example

```golang
// load data
data, err := elfstore.Load()

// save data
err := elfstore.Save(data)
```
