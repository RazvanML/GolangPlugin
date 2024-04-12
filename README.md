# GolangPlugin

The main idea of the plugin system is to allow behavioural modifications of a software without (much) changing the binary of the application.

The common scenario had in mind while writing this library is that we develop a product, sold as an enterprise or even cloud application. At an early phase of the sales process, the client requests customizations, for example a new form, new fields are being added, as well as additional tables and fields in existing tables. The addition of the new features does not propagate to the existing customers, as well as does not impact the stability of the product by introducing new bugs in the release version. Previously this was done by having branches for each customer, which is manpower expensive, because of the need of repeated merges between branches.

The central idea of the plugin system is to define hooks inside the code. The hooks are nonoperational for the COTS system. For this reason, adding additional hooks does not significantly introduce bugs or changes the behaviour of the distribution. On the custom system, we either supply an additional dynamic linked module, or we compile the code with specific tags.

Given the current practical state-of-the-art in Go, we provide the mechanism of statically linking the plugins. This means we can include plugins in a static fashion, even prohibit the library to look for plugins in te plugin directory. Such a feature provides security and stability. It also addresses the Windows operating system, where currenty the toolchain cannot build a DLL.