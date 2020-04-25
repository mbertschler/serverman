`serverman` Server Setup Tooling
================================

Serverman (server manager) is a framework and tooling around setting up software on servers using Go code.

Why
---

The goal of `serverman` is to provide a simple way to install software on a single server and keep it up to date.
You declare a configuration of things you want to be installed using Go code. 
It chooses a statically typed programming language instead of config files to take advantage of the benefits
of type checking and compilation.
This also allows you to easily add custom logic which is harder to do in a config file based tool.
It also makes it possible to use Go modules for importing a specific version of third party config packages. 

Should I use it?
----------------

If you want to

- configure the software you want to install in Go
- manage your config in a Git repository
- automatically keep your installed packages up to date
- use a Debian/Ubuntu based Linux distribution

you should give `serverman` a try.

How does it work?
-----------------

For simplicity the tool is stateless and gets all required information from the system every time it runs.
It also prefers using CLIs over APIs for ease of debugging and manual testing. 

Planned Config Packages
-----------------------

- [ ] `apt` - install and update Debian packages
- [ ] `systemd` - manage services
- ...

---

To Do:
- How it plays together with other tools
- What it does during bootstrap and config updates
