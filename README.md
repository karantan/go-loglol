# loglol
A simple repo where I can play with different go logging settings.

The easiest (best?) way to setup logging is just to simply print to the console and then
in the production wrap the program in a [systemd service](https://www.freedesktop.org/software/systemd/man/systemd.service.html).
With this everything printed to the console will be automatically sent to journald log.

Example of a one-shot systemd service:


`loglol.service` file:
```
[Unit]

[Service]
Environment="..."
ExecStart=/<PATH_TO_EXEC>/loglol-start
Type=oneshot

```

And in nix format:
```
systemd.services.loglol = {
    serviceConfig.Type = "oneshot";
    script = ''
        ./<PATH_TO_EXEC>/loglol
    '';
};
```

With this, we will be able to see CPU consumed, network IO, etc
(with `systemctl status loglol`). We can also limit the resources via `MemoryLimit`,
`CPUQuota`, etc.

See [systemd.resource-control](https://www.freedesktop.org/software/systemd/man/systemd.resource-control.html)
for more info.
