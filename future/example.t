* TypeFast is a language developed to make language dev easier *

* Eventuall will have *
block Info -> {
	content: string,
	size:	usize,
	path: PATH
};

fnc init_Info(filename: string) -> Info
{
	block Info info = new block Info();

	info.content = read_data_from(filename);
	
	if len(info.content) == 0
	{
		raise new Error("File is invalid: No Information");
	}
	
	info.size = len(info.content);
	info.path = new PATH(filename);
	
	ret info;
}

* TypeFast supports Function Wrappers *
* A function wrapper has the following syntax: #[wrapper_name] *

* For example, the wrapper #[ignore] will ignore the code below it until it reaches } *
* #[ignore] is for functions only *
* All ignored functions act as test functions *
#[ignore]
fnc print() -> None
{
	print "Hey"
}
