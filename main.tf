#[allow(inter_access)] * allow us to work with the interpreter within the language *

* You can only require modules that the interpreter will understand *
* lang_info.json supports information about the language being created *
require "lang_info.json"

print inter.LangInfo.lang_name|

print inter.LangInfo.lang_version|
