//vim: set filetype=cfg

package conf

const baseConfig = `;Sample Configuration File
[gottp]
listen="/tmp/shortener.sock"
EmailHost=""
EmailPort=""
EmailPassword=""
EmailUsername=""
EmailSender="Shortener"
EmailFrom=""
ErrorTo=""
EmailDummy=true

[shortener]
StoragePath="/tmp/shortener" #Override this config file`
