# Crayon
An interpreted programming language written in Go for idiots.

# What the Fuck?
There are only tags, values, constant values, commands (statements), variables and scope. Well, for now.

# What the Fuck is a Tag?
It works like a function, but you can't directly call it, only system can, at least for now.

## Kinds of Tag
1. MAIN
```
[MAIN]
```
It will execute when the engine ready.
2. FRAME
```
[FRAME XX]
```
It will execute after XX seconds.
3. LOOP
```
[LOOP]
```
It will execute every time until you exit the program.

# Values
1. String
```
"This is a string"
```

2. Number (Float 64)
```
10.10
```

3. Bool
```
#yes
#no
```

4. PI
```
#PI
```

5. Object (Untested)
```
{
    key1 = <Any>,
    key2 = <Any>
}
```

6. Array (Untested)
```
[<Any>, <Any>]
```


# Commands
1. Exit the program
```
K THX BYE
```

2. Console
```
write <Any> to console
```

3. Define
```
define variable $[variable]
```

4. Assign
```
assign value <Any> to $[variable]
```

5. Delete 
```
delete variable $[variable]
```

6. Math Operators
```
add <Number> by <Number>
sub <Number> by <Number>
divide <Number> by <Number>
multiply <Number> by <Number>
```

7. Logical Operators (All the results are <Bool>)
Protip: You can use commands within commands like
```
command1 @(command2)
```
```
is <Any> equal with <Any>
is <Any> not equal with <Any> 

not <Bool>

is <Number> less than <Number>
is <Number> less than or equal with <Number>

is <Number> more than <Number>
is <Number> more than or equal with <Number>

<Bool> and <Bool>
<Bool> or <Bool>
```

8. If Else
```
if <Bool> {
    this is scope
}

if <Bool> {
    this is scope
} else {
    this is scope
}
```