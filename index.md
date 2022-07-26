---
theme: gaia
_class: lead
paginate: true
backgroundColor: #fff
backgroundImage: url('https://marp.app/assets/hero-background.svg')
marp: true
style: |
    section.movetov2 em {
      font-size: 0.7em;
    }
    section.questions p {
        text-align: center;
    }
    section.questions img {
       vertical-align: middle;
       width: 150px;
       height: 150px;
    }
---

![bg left:40% 80%](https://www.krakend.io/images/logo-krakend.svg)

# **Krakend v2**

A Quick Introduction

---

# What is up with v2?

-   Better documentation
-   Security fixes
-   Configuration file json schema support
-   Built-in logger access from plugins
-   Krakend `check-plugin` support.

<!--
Notes:
- The current version is confusing because it is not aligned with the latest documentation of krakend (v1.4.1).
- Security patches are always important and we shouldn't ignore them no matter how small they are.
- Less room for error, syntax highlighting, json property autocomplete and documentation.
- To integrate seamlessly with krakend logger, they've added this to inject logger in the plugins.
- Helps a lot in checking your custom plugin for incompatibilities in golang and module version.
-->

---

<!-- _class: movetov2 -->

# What is up with v2?

-   Request and Response modifier plugins
-   Rest to GraphQL Support
-   Async Agents for Pub/Sub
-   Shorter built-in plugin names
-   [Router options](https://www.krakend.io/docs/service-settings/router-options/)
-   [And more](https://github.com/krakendio/krakend-ce/releases)

<!--
Notes:
 - Before it was not possible to create custom request and response modifier plugin.
 -
 - Async agents means that when a message has been published for a queue, it will be triggered unlike before where a requests should have made first before it will consume message from the queue.
 - Helps in creating a neater krakend/lura config.
 - Router options such as disable_health, health_path, return_error_msg and more.
-->

---

<!-- _class: movetov2 -->

# Why move to v2 IMHO?

-   Easier to upgrade for future releases
-   Included a lot of security fixes
-   \*Build slimmer and cleaner api gateways
-   \*\*Migration should be easy with [Migration Guide](https://www.krakend.io/docs/configuration/migrating/)

<br/>
<br/>
<br/>

_\*Make sure to use the binary instead of compiling krakend_
_\*\*As long as you didn't do some crazy internal modification in your api gateway_

<!--
Notes:
- Just upgrade the docker container and we should be good to go.
-
- If we use the compiled krakend, we enforce ourselves to build slimmer and cleaner api gateways by only creating custom plugins and not modifying it's internals.
- They've provided a tool to easily migrate from the older version to the latest one.

-->

---

# Examples

1. [json-schema integration](https://www.krakend.io/docs/configuration/structure/)
2. [Sequential requests](https://www.krakend.io/docs/endpoints/sequential-proxy/)
3. [Conditional Request](https://www.krakend.io/docs/endpoints/common-expression-language-cel/)
4. [Server plugins](https://www.krakend.io/docs/extending/http-server-plugins/)
5. [Client plugins](https://www.krakend.io/docs/extending/http-client-plugins/)
6. [Req/Res modifier plugins](https://www.krakend.io/docs/extending/http-client-plugins/)

<!--
Note:
1. Found in ../krakend/krakend.json#L2
2. Found in ../krakend/krakend.json#L111
3. Found in ../krakend/krakend.json#L136
4. Found in ../krakend/krakend.json#L14
5. Found in ../krakend/krakend.json#L292
5. Found in ../krakend/krakend.json#L246
-->

---

# Common Problems

-   Wrong golang version
-   Unable to build plugin with golang alpine
-   I don't know if I should write a custom plugin?
-   My configuration isn't working
-   Why nothing works in v1.4?

<!--
 Notes:
 - Plugins should be built in the same golang version as krakend was built from.
 - Use `krakend check-plugin` command if necessary, it helps a lot in identifying incompatibilities of your plugin.
 - Make sure to include dev dependencies in your docker image.
 -
 - Check first if your requirement is already supported by the existing plugins.
    https://github.com/krakendio/krakend-martian
    https://www.krakend.io/docs/endpoints/common-expression-language-cel/
    https://github.com/krakendio/krakend-cel

 - Make sure that you've restarted krakend or the krakend container. If you want to automatically restart krakend on configuration changes, you can use devopsfaith/krakend:watch as your base image for development purposes.
-->

---

# References

-   [Krakend](https://www.krakend.io/docs/overview/)
-   [Krakend Custom Plugins](https://www.krakend.io/docs/extending/)
-   [cel-spec](https://github.com/google/cel-spec)
-   [Flexible configuration](https://www.krakend.io/docs/configuration/flexible-config/)
-   [Martian Plugins](https://www.krakend.io/docs/backends/martian/)
-   [Martian Repository](https://github.com/google/martian)

---

<!-- _class: questions -->

# Questions?

🤔

---

# _Thank you!_

[Source code](https://github.com/jbactad/krakend-v2-presentation)
