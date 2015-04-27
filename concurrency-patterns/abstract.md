Concurrency bugs are the worst. Go gives us awesome tools, yet still most of us have been up at 3am trying to figure out why 'go test -race' fails when “I just fixed that.”  

It’s almost too easy to write concurrency bugs in Go. Lots of great stuff has been written and said on how to do concurrency right, but we still lack good conventions on how to write concurrent code for a large codebase.

In this talk, we’ll discuss some conventions for writing concurrent code in large codebases without breaking the build, all with the standard library.
