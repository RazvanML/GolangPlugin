# GolangPlugin

The main idea of the plugin system is to allow behavioural modifications of a software without (much) changing the binary of the application.

The common scenario had in mind while writing the plugin is that we develop an enterprise application. At an early phase, the client has requests for customization, for example a new form, new fields are being added, as well as additional tables and fields in existing tables. The addition of the new features does not impact functionally the existing customer implementations, as they won't see the changes, as well as does not impact the stability of the product by introducing new bugs in the stock version. Previously this was done by having branches for each customer, which is a costly endeavor, because of the need of repeated merges between branches.

The idea of the plugin system is to define hooks inside the code, which hooks are nonoperational for the COTS system. For this reason, adding additional hooks does not significantly introduces bugs or changes the behaviour of the distribution. On the custom system, an additional dynamic linked module is being supplied, containing the customization desired by the specific client.
