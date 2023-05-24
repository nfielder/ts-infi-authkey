# ts-infi-authkey

This CLI program can be used for generating Tailscale authkeys. It is designed to be used as part of CI pipelines with OAuth credentials due to user generated auth keys having an maximum expiry time of 90 days.

It is heavily inspired by the [`get-authkey`](https://tailscale.com/kb/1215/oauth-clients/#get-authkey-utility) utility from Tailscale.

## Usage

The program relies on the presence of two environment variables to work:
- `TS_API_CLIENT_ID` - The OAuth credentials client ID.
- `TS_API_CLIENT_SECRET` - The OAuth credentials client secret.

```
Usage of ts-infi-authkey:
  -ephemeral
    allocate an ephemeral authkey (default true)
  -expiry duration
    time until expiry of the authkey. Accepts string similar to 5m or 1h (default 5m0s)
  -preauth
    set the authkey as pre-authorised (default true)
  -reusable
    allocate a reusable authkey
  -tags string
    comma-separated list of tags to apply to the authkey
```
