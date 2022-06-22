sfleet
=========

Simple utility to manage multiple ssh

## Install

```shell
go install github.com/j3ssie/sfleet@latest
```

## Usage

```shell
# simple usage
sfleet exec -t '1.2.3.4' --cmd 'ls -la /'
echo '1.2.3.4' | sfleet exec --cmd 'ls -la /'

# add worker to the config at ~/.sfleet/config
cat list-of-ips | sfleet worker add 

# run command on all the existing worker
sfleet exec --cmd 'ls -la /'
```

## Donation

[![paypal](https://www.paypalobjects.com/en_US/i/btn/btn_donateCC_LG.gif)](https://paypal.me/j3ssiejjj)

[!["Buy Me A Coffee"](https://www.buymeacoffee.com/assets/img/custom_images/orange_img.png)](https://www.buymeacoffee.com/j3ssie)
