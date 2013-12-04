Ball Clocks
-----------

Based on the problem published here: http://www.chilton.com/~jimw/ballclk.html

I began this project coming up with a model of ball clocks that would allow me to
accurately predict, based on the queue size, how long it takes the queue to reorder
itself. I like to write it out:

Let's simulate

    2,3
    (1,2,3,4,5,6) () ()
    (1,2,3,4,5) (6) ()
    (6,1,2,3,4) () (5)
    (6,1,2,3) (4) (5)
    (4,6,1,2) () (5,3)
    (4,6,1) (2) (5,3)
    (1,3,5,2,4,6) () ()

I initially built a ball clock simulator around the type BallClock and the Tick() methods.
The ball clock accurately simulated the motion and ordering of the balls through the machine.
I've called this algorithm, unoriginally, "iterateminutes". It was enough to solve all the problems
from 27 to 127 in about 8 seconds on my i7 CPU using goroutines. Consider this my naive solution.
I am sorting the output because I'm using goroutines and waitgroups, so the output is non-deterministic.
Also my timings were taken from pre-compiled versions of the bench program, not using go run.

    go run bench/main.go -algo iterateminutes < test2.txt | sort -n


Here comes the google, you can stop reading here if you want
------------------------------------------------------------

It turns out that, as promised, ball clocks are a google'able thing. So after watching a few lego clocks
online (which really helped in step 1). I came across the original problem statement. I took away the key insight
that every 12 hours the same permutation is applied to the balls, and then I left. This gave rise to my second
strategy (permutedays), which solved the problem (27-127) in under .4 seconds. The strategy was pretty simple,
use iterateminutes until I've obtained the permutation vector, and then permute day by day until I get back to
the identity permutation (BallClock.IsPristine()).

    go run bench/main.go -algo permutedays < test2.txt | sort -n

That's kind of a boring middle step. I do remember from my old abstract algebra classes that permutations of finite numbers
can be written as a sequence of cycles. For example the permutation switch even and odd spots might look something like this:

    (2,1)(4,3)(6,5)...

It's clear that the period of this thing is 2 cycles. Now I didn't rigorously prove this, but it's a well known
fact that the period of a permutation is the least common multiple of the individual cycles, which is pretty
intuitive actually. That gave the final, simplest solution which completes almost instantly (well 14 milliseconds)

    go run bench/main.go -algo lcm < test2.txt | sort -n

Not bad, and a fun project. I'd like to talk about some of the pieces. The style was that I solved
each problem invidiually, hence the existance of clock2, and clock3. I also used golang method receivers to
do table driven testing of each algorithm against known data. Some commmon functionality around permutations
was put in it's own file. the golang package layout makes this sort of flexible structure pretty easy.

I also used the new go cover tool. We've got about [95% test coverage](http://shanemhansen.github.io/ballclock/coverage.html), I augment that with the actual deliverable (bench/main.go)
which reads the numbers from stdin and outputs the messages.
