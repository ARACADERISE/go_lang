* Function with a chained function *
* A chained function is a function that works with the return value *
* of the function before it *

* Example of a chained function *
* A function that will have a chained function needs return types *
* int_return_type, possible return type
* Chained functions can change the return value of a function *
fnc get_age_total(ageOne: int, ageTwo: int) -> int, pos err
{
	ret ageOne + ageTwo
} -> {
	if ret_val < 21 -> ret Err("Not of age")
	* No need to re-return the value. If there is no ret found in the *
	* chained function that seems to be used, then the value is that *
	* of what was returned in the above function *
}

get_age_total(16, 5) * -> 21 *

* Reduced function are functions removed after they're called *
* Reduced functions declared with the #[reduced] function wrapper *
#[reduced]
fnc return_value(val vartype) -> vartype
{
	ret val
}() * this function automatically runs at the end *
    * So, after it runs, if you try running return_value it will error with *
    * "Undefined Function Error" *

* #[allow_print] allows variables(or blocks) to be printed *
* Example *
#[allow_print]
let ages = [1,2,3,4];
print ages|

#[allow_print]
fnc print_name(name string) -> None
{
	print name|
}

print print_name| * {type: function, params: name, return_type: None} *

#[allow_print]
block Info {
	name	str
	age	int
	names	[]str
}

print Info| * {type: Block, variables: [name,age, names], var_types: [str, int, []str], name: Info} *

* #[allow(inter_access)] allows access to the interpreter. Meaning you can change *
* types of variables, you can do specific things such as set new names to variables *
* add things to blocks, see information about required packages you included in *
* your project etc *
let a: int;
#[allow(inter_access)]
fnc change_type(var varname, new_type vartype) -> None
{
	inter.setNew(var, new_type) * .setNew(varname, vartype) *				                         * sets varname to vartype, accordingly *
}(a, str)


* You will normally see functions like this in the beginning: *
#[allow(inter_access)]
fnc all() -> None {
	inter.SetAll(settings, true)
}()
* The above function runs .SetAll, which takes in a block of information, and sets all values to the set_value *
* In the above example, settings is the function wrappers block that contains values, accordingly. In all(), we set all *
* those values to true..consistently *
