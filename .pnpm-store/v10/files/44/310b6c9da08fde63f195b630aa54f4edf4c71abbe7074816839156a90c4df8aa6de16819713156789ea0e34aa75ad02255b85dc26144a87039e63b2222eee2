# A Course in Combinatorial Optimization

## Alexander Schrijver

```
CWI,
Kruislaan 413,
1098 SJ Amsterdam,
The Netherlands
```

```
and
```

```
Department of Mathematics,
University of Amsterdam,
Plantage Muidergracht 24,
1018 TV Amsterdam,
The Netherlands.
```

```
March 23, 2017
copyright ©c A. Schrijver
```

## 1. Shortest paths and trees

### 1.1. Shortest paths with nonnegative lengths

Let _D_ = (_V, A_) be a directed graph, and let _s, t_ ∈ _V_. A _walk_ is a sequence _P_ =
(_v_[^0]:_, a_[^1]:_, v_[^1]:_,... , am, vm_) where _ai_ is an arc from _vi_−[^1]: to _vi_ for _i_ = 1_,... , m_. If _v_[^0]:_,... , vm_
all are different, _P_ is called a _path_.
If _s_ = _v_[^0]: and _t_ = _vm_, the vertices _s_ and _t_ are the _starting_ and _end vertex_ of _P_,
respectively, and _P_ is called an _s_ − _t walk_, and, if _P_ is a path, an _s_ − _t path_. The
_length_ of _P_ is _m_. The _distance_ from _s_ to _t_ is the minimum length of any _s_ − _t_ path.
(If no _s_ − _t_ path exists, we set the distance from _s_ to _t_ equal to ∞.)
It is not difficult to determine the distance from _s_ to _t_: Let _Vi_ denote the set of
vertices of _D_ at distance _i_ from _s_. Note that for each _i_:

(1) _Vi_+1 is equal to the set of vertices _v_ ∈ _V_ \ (_V_[^0]: ∪ _V_[^1]: ∪···∪ _Vi_) for which
(_u, v_) ∈ _A_ for some _u_ ∈ _Vi_.

This gives us directly an algorithm for determining the sets _Vi_: we set _V_[^0]: := {_s_} and
next we determine with rule (1) the sets _V_[^1]:_, V_[^2]:_,..._ successively, until _Vi_+1 = ∅.
In fact, it gives a linear-time algorithm:

**Theorem 1.1.** _The algorithm has running time O_(|_A_|)_._

**Proof.** Directly from the description.

In fact the algorithm finds the distance from _s_ to all vertices reachable from _s_.
Moreover, it gives the shortest paths. These can be described by a rooted (directed)
tree _T_ = (_V_ ′_, A_′), with root _s_, such that _V_ ′ is the set of vertices reachable in _D_ from
_s_ and such that for each _u, v_ ∈ _V_ ′, each directed _u_ − _v_ path in _T_ is a shortest _u_ − _v_
path in _D_.[^1]
Indeed, when we reach a vertex _t_ in the algorithm, we store the arc by which _t_ is
reached. Then at the end of the algorithm, all stored arcs form a rooted tree with
this property.
There is also a trivial min-max relation characterizing the minimum length of an
_s_ − _t_ path. To this end, call a subset _A_′ of _A_ an _s_ − _t cut_ if _A_′ = _δ_out(_U_) for some
subset _U_ of _V_ satisfying _s_ ∈ _U_ and _t_ 6∈ _U_.[^2] Then the following was observed by
Robacker [1956]:

[^1]:A _rooted tree_, with _root s_, is a directed graph such that the underlying undirected graph is a
tree and such that each vertex _t_ [^6]:= _s_ has indegree 1. Thus each vertex _t_ is reachable from _s_ by a
unique directed _s_ − _t_ path.
[^2]:_δ_out(_U_) and _δ_in(_U_) denote the sets of arcs leaving and entering _U_, respectively.

**Theorem 1.2.** _The minimum length of an s_−_t path is equal to the maximum number_
_of pairwise disjoint s_ − _t cuts._

**Proof.** Trivially, the minimum is at least the maximum, since each _s_−_t_ path intersects
each _s_−_t_ cut in an arc. The fact that the minimum is equal to the maximum follows

by considering the _s_−_t_ cuts _δ_out(_Ui_) for _i_ = 0_,... , d_−1, where _d_ is the distance from
_s_ to _t_ and where _Ui_ is the set of vertices of distance at most _i_ from _s_.

This can be generalized to the case where arcs have a certain ‘length’. For any
‘length’ function _l_ : _A_ → Q+ and any walk _P_ = (_v_[^0]:_, a_[^1]:_, v_[^1]:_,... , am, vm_), let _l_(_P_) be
the length of _P_. That is:

(2) _l_(_P_) :=

```
∑m
```

```
i=1
```

```
l(ai).
```

Now the _distance_ from _s_ to _t_ (with respect to _l_) is equal to the minimum length of
any _s_ − _t_ path. If no _s_ − _t_ path exists, the distance is +∞.
Again there is an easy algorithm, due to Dijkstra [1959], to find a minimum-length
_s_−_t_ path for all _t_. Start with _U_ := _V_ and set _f_(_s_) := 0 and _f_(_v_) = ∞ if _v_ [^6]:= _s_. Next
apply the following iteratively:

(3) Find _u_ ∈ _U_ minimizing _f_(_u_) over _u_ ∈ _U_. For each _a_ = (_u, v_) ∈ _A_ for which
_f_(_v_) _> f_(_u_) + _l_(_a_), reset _f_(_v_) := _f_(_u_) + _l_(_a_). Reset _U_ := _U_ \{_u_}.

We stop if _U_ = ∅. Then:

**Theorem 1.3.** _The final function f gives the distances from s._

**Proof.** Let dist(_v_) denote the distance from _s_ to _v_, for any vertex _v_. Trivially,
_f_(_v_) ≥ dist(_v_) for all _v_, throughout the iterations. We prove that throughout the
iterations, _f_(_v_) = dist(_v_) for each _v_ ∈ _V_ \ _U_. At the start of the algorithm this is
trivial (as _U_ = _V_ ).
Consider any iteration (3). It suffices to show that _f_(_u_) = dist(_u_) for the chosen
_u_ ∈ _U_. Suppose _f_(_u_) _>_ dist(_u_). Let _s_ = _v_[^0]:_, v_[^1]:_,... , vk_ = _u_ be a shortest _s_ − _u_ path.
Let _i_ be the smallest index with _vi_ ∈ _U_.
Then _f_(_vi_) = dist(_vi_). Indeed, if _i_ = 0, then _f_(_vi_) = _f_(_s_) = 0 = dist(_s_) = dist(_vi_).
If _i >_ 0, then (as _vi_−[^1]: ∈ _V_ \ _U_):

(4) _f_(_vi_) ≤ _f_(_vi_−[^1]:) + _l_(_vi_−[^1]:_, vi_) = dist(_vi_−[^1]:) + _l_(_vi_−[^1]:_, vi_) = dist(_vi_)_._

This implies _f_(_vi_) ≤ dist(_vi_) ≤ dist(_u_) _< f_(_u_), contradicting the choice of _u_.

Clearly, the number of iterations is |_V_ |, while each iteration takes _O_(|_V_ |) time.
So the algorithm has a running time _O_(|_V_ |[^2]). In fact, by storing for each vertex _v_ the
last arc _a_ for which (3) applied we find a rooted tree _T_ = (_V_ ′_, A_′) with root _s_ such
that _V_ ′ is the set of vertices reachable from _s_ and such that if _u, v_ ∈ _V_ ′ are such that
_T_ contains a directed _u_ − _v_ path, then this path is a shortest _u_ − _v_ path in _D_.
Thus we have:

**Theorem 1.4.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V , and a length function_
_l_ : _A_ → Q+_, a shortest s_ − _t path can be found in time O_(|_V_ |[^2])_._

**Proof.** See above.

```
For an improvement, see Section 1.2.
A weighted version of Theorem 1.2 is as follows:
```

**Theorem 1.5.** _Let D_ = (_V, A_) _be a directed graph, s, t_ ∈ _V , and let l_ : _A_ → Z+_._
_Then the minimum length of an s_−_t path is equal to the maximum number k of s_−_t_
_cuts C_[^1]:_,... , Ck (repetition allowed) such that each arc a is in at most l_(_a_) _of the cuts_
_Ci._

**Proof.** Again, the minimum is not smaller than the maximum, since if _P_ is any _s_−_t_

path and _C_[^1]:_,... , Ck_ is any collection as described in the theorem:[^3]

(5) _l_(_P_) =

```
∑
```

```
a∈AP
```

```
l(a) ≥
```

```
∑
```

```
a∈AP
```

```
( number of i with a ∈ Ci)
```

```
=
```

```
∑k
```

```
i=1
```

```
|Ci ∩ AP|≥
```

```
∑k
```

```
i=1
```

```
1 = k.
```

To see equality, let _d_ be the distance from _s_ to _t_, and let _Ui_ be the set of vertices
at distance less than _i_ from _s_, for _i_ = 1_,... , d_. Taking _Ci_ := _δ_out(_Ui_), we obtain a
collection _C_[^1]:_,... , Cd_ as required.

**Application 1.1: Shortest path.** Obviously, finding a shortest route between cities is an
example of a shortest path problem. The length of a connection need not be the geographical
distance. It might represent the time or energy needed to make the connection. It might
cost more time or energy to go from _A_ to _B_ than from _B_ to _A_. This might be the case, for
instance, when we take differences of height into account (when routing trucks), or air and
ocean currents (when routing airplanes or ships).
Moreover, a route for an airplane flight between two airports so that a minimum amount
of fuel is used, taking weather, altitude, velocities, and air currents into account, can be

```
3AP denotes the set of arcs traversed by P
```

```
found by a shortest path algorithm (if the problem is appropriately discretized — otherwise
it is a problem of ‘calculus of variations’). A similar problem occurs when finding the
optimum route for boring say an underground railway tunnel.
```

```
Application 1.2: Dynamic programming. A company has to perform a job that will
take 5 months. For this job a varying number of extra employees is needed:
```

```
(6) month number of extra employees needed
1 b1=10
2 b2=7
3 b3=9
4 b4=8
5 b5=11
```

```
Recruiting and instruction costs EUR 800 per employee, while stopping engagement costs
EUR 1200 per employee. Moreover, the company has costs of EUR 1600 per month for
each employee that is engaged above the number of employees needed that month. The
company now wants to decide what is the number of employees to be engaged so that the
total costs will be as low as possible.
Clearly, in the example in any month i, the company should have at least bi and at most
11 extra employees for this job. To solve the problem, make a directed graph D = (V,A)
with
```

```
(7) V := {(i,x) | i = 1,...,5;bi ≤ x ≤ 11}∪{(0,0),(6,0)},
A := {((i,x),(i + 1,y)) ∈ V × V | i = 0,...,5}.
```

```
(Figure 1.1).
At the arc from (i,x) to (i + 1,y) we take as length the sum of
```

(8) (i) the cost of starting or stopping engagement when passing from _x_ to _y_ employees
(this is equal to 8(_y_ − _x_) if _y_ ≥ _x_ and to 12(_x_ − _y_) if _y < x_);
(ii) the cost of keeping the surplus of employees in month _i_+1 (that is, 16(_y_−_bi_+1))

```
(taking EUR 100 as unit).
Now the shortest path from (0,0) to (6,0) gives the number of employees for each month
so that the total cost will be minimized. Finding a shortest path is thus a special case of
dynamic programming.
```

```
Exercises
```

```
1.1. Solve the dynamic programming problem in Application 1.2 with Dijkstra’s method.
```

```
8
64
```

```
40
```

```
24
```

```
64
72
60
48
```

```
40
48
```

```
24
```

```
28
```

```
12
```

```
16
```

```
40
```

```
24
```

```
48
```

```
56 32 16
```

```
56 64
```

```
40
```

```
32
```

```
44
```

```
16
12 36
```

```
28
40
```

```
8
```

```
16
```

```
24
```

```
104
```

```
80
```

```
132
```

```
36
```

```
32 48 0
```

```
0
```

```
56
```

```
0 1 2 3 4 5 6
```

```
0
```

```
7
```

```
8
```

```
9
```

```
10
```

```
11
```

```
44
52
```

```
Figure 1.1
```

### 1.2. Speeding up Dijkstra’s algorithm with heaps

For dense graphs, a running time bound of _O_(|_V_ |[^2]) for a shortest path algorithm is
best possible, since one must inspect each arc. But if |_A_| is asymptotically smaller
than |_V_ |[^2], one may expect faster methods.
In Dijkstra’s algorithm, we spend _O_(|_A_|) time on updating the values _f_(_u_) and
_O_(|_V_ |[^2]) time on finding a _u_ ∈ _U_ minimizing _f_(_u_). As |_A_| ≤ |_V_ |[^2], a decrease in the
running time bound requires a speed-up in finding a _u_ minimizing _f_(_u_).
A way of doing this is based on storing the _u_ in some order so that a _u_ minimizing
_f_(_u_) can be found quickly and so that it does not take too much time to restore the
order if we delete a minimizing _u_ or if we decrease some _f_(_u_).
This can be done by using a ‘heap’, which is a rooted forest (_U, F_) on _U_, with the
property that if (_u, v_) ∈ _F_ then _f_(_u_) ≤ _f_(_v_).[^4] So at least one of the roots minimizes
_f_(_u_).

```
Let us first consider the 2-heap. This can be described by an ordering u1,... , un
```

[^4]:A _rooted forest_ is an acyclic directed graph _D_ = (_V,A_) such that each vertex has indegree at
most 1. The vertices of indegree 0 are called the _roots_ of _D_. If (_u,v_) ∈ _A_, then _u_ is called the _parent_
of _v_ and _v_ is called a _child_ of _u_.
If the rooted forest has only one root, it is a _rooted tree_.

of the elements of _U_ such that if _i_ = ⌊_j_
[^2]:⌋ then _f_(_ui_) ≤ _f_(_uj_). The underlying rooted
forest is in fact a rooted tree: its arcs are the pairs (_ui, uj_) with _i_ = ⌊_j_
[^2]:⌋.
In a 2-heap, one easily finds a _u_ minimizing _f_(_u_): it is the root _u_[^1]:. The following
theorem is basic for estimating the time needed for updating the 2-heap:

**Theorem 1.6.** _If u_[^1]: _is deleted or if some f_(_ui_) _is decreased, the_ [^2]:_-heap can be restored_
_in time O_(log _p_)_, where p is the number of vertices._

**Proof.** To remove _u_[^1]:, perform the following ‘sift-down’ operation. Reset _u_[^1]: := _un_
and _n_ := _n_ − 1. Let _i_ = 1. While there is a _j_ ≤ _n_ with 2_i_ + 1 ≤ _j_ ≤ [^2]:_i_ + 2 and
_f_(_uj_) _< f_(_ui_), choose one with smallest _f_(_uj_), swap _ui_ and _uj_, and reset _i_ := _j_.
If _f_(_ui_) has decreased perform the following ‘sift-up’ operation. While _i >_ 0 and
_f_(_uj_) _> f_(_ui_) for _j_ := ⌊_i_−[^1]
[^2]: ⌋, swap _ui_ and _uj_, and reset _i_ := _j_. The final 2-heap is as
required.
Clearly, these operations give 2-heaps as required, and can be performed in time
_O_(log |_U_|).

```
This gives the result of Johnson [1977]:
```

**Corollary 1.6a.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V and a length function_
_l_ : _A_ → Q+_, a shortest s_ − _t path can be found in time O_(|_A_|log |_V_ |)_._
**Proof.** Since the number of times a minimizing vertex _u_ is deleted and the number
of times a value _f_(_u_) is decreased is at most |_A_|, the theorem follows from Theorem
1.6.

Dijkstra’s algorithm has running time _O_(|_V_ |[^2]), while Johnson’s heap implemen-
tation gives a running time of _O_(|_A_|log |_V_ |). So one is not uniformly better than the
other.
If one inserts a ‘Fibonacci heap’ in Dijkstra’s algorithm, one gets a shortest path
algorithm with running time _O_(|_A_| + |_V_ |log |_V_ |), as was shown by Fredman and
Tarjan [1984].
A _Fibonacci forest_ is a rooted forest (_V, A_), so that for each _v_ ∈ _V_ the children of
_v_ can be ordered in such a way that the _i_th child has at least _i_ − 2 children. Then:[^5]

**Theorem 1.7.** _In a Fibonacci forest_ (_V, A_)_, each vertex has at most_ 1 + 2 log |_V_ |
_children._

**Proof.** For any _v_ ∈ _V_ , let _σ_(_v_) be the number of vertices reachable from _v_. We show
that _σ_(_v_) ≥ [^2]:(_d_out(_v_)−1)_/_[^2], which implies the theorem.[^6]

[^5]:_d_out(_v_) and _d_in(_v_) denote the outdegree and indegree of _v_.
[^6]:In fact, _σ_(_v_) ≥ _F_(_d_out(_v_)), where _F_(_k_) is the _k_th Fibonacci number, thus explaining the name
Fibonacci forest.

```
Let k := dout(v) and let vi be the ith child of v (for i = 1,... , k). By induction,
σ(vi) ≥ 2(dout(vi)−1)/[^2] ≥ 2(i−3)/[^2], as dout(vi) ≥ i − 2. Hence σ(v) = 1 +
```

```
∑k
i=1 σ(vi) ≥
1 +
∑k
i=1[^2](i−3)/[^2] = 2(k−1)/[^2] + 2(k−2)/[^2] +[^1]
2 −[^1]
2
```

```
√
2 ≥ 2(k−1)/[^2].
```

```
Now a Fibonacci heap consists of a Fibonacci forest (U, F), where for each v ∈ U
the children of v are ordered so that the ith child has at least i − 2 children, and a
subset T of U with the following properties:
```

(9) (i) if (_u, v_) ∈ _F_ then _f_(_u_) ≤ _f_(_v_);
(ii) if _v_ is the _i_th child of _u_ and _v_ 6∈ _T_ then _v_ has at least _i_ − 1 children;
(iii) if _u_ and _v_ are two distinct roots, then _d_out(_u_) [^6]:= _d_out(_v_).

```
So by Theorem 1.7, (9)(iii) implies that there exist at most 2 + 2 log |U| roots.
The Fibonacci heap will be described by the following data structure:
```

(10) (i) for each _u_ ∈ _U_, a doubly linked list _Cu_ of children of _u_ (in order);
(ii) a function _p_ : _U_ → _U_, where _p_(_u_) is the parent of _u_ if it has one, and
_p_(_u_) = _u_ otherwise;

```
(iii) the function dout : U → Z+;
(iv) a function b : {0,... , t}→ U (with t := 1+⌊2 log |V |⌋) such that b(dout(u)) =
u for each root u;
(v) a function l : U →{0,1} such that l(u) = 1 if and only if u ∈ T.
```

```
Theorem 1.8. When finding and deleting n times a u minimizing f(u) and decreas-
ing m times the value f(u), the structure can be updated in time O(m + p + nlog p),
where p is the number of vertices in the initial forest.
```

```
Proof. Indeed, a u minimizing f(u) can be identified in time O(log p), since we can
scan f(b(i)) for i = 0,... , t. It can be deleted as follows. Let v1,... , vk be the children
of u. First delete u and all arcs leaving u from the forest. In this way, v1,... , vk have
become roots, of a Fibonacci forest, and conditions (9)(i) and (ii) are maintained. To
repair condition (9)(iii), do for each r = v1,... , vk the following:
```

```
(11) repair(r):
if dout(s) = dout(r) for some root s 6= r, then:
{if f(s) ≤ f(r), add s as last child of r and repair(r);
otherwise, add r as last child of s and repair(s)}.
```

```
Note that conditions (9)(i) and (ii) are maintained, and that the existence of a root
s 6= r with dout(s) = dout(r) can be checked with the functions b, dout, and p. (During
the process we update the data structure.)
```

```
If we decrease the value f(u) for some u ∈ U we apply the following to u:
```

(12) _make root_(_u_):
if _u_ has a parent, _v_ say, then:
{delete arc (_v, u_) and repair(_u_);
if _v_ 6∈ _T_, add _v_ to _T_; otherwise, remove _v_ from _T_ and make root(_v_)}.

Now denote by incr(_.._) and decr(_.._) the number of times we increase and decrease
.. , respectively. Then:

(13) number of calls of make root = decr(_f_(_u_)) + decr(_T_)
≤ decr(_f_(_u_)) + incr(_T_) + _p_ ≤ 2decr(_f_(_u_)) + _p_ = 2_m_ + _p_,

since we increase _T_ at most once after we have decreased some _f_(_u_).
This also gives, where _R_ denotes the set of roots:

(14) number of calls of repair= decr(_F_) + decr(_R_)
≤ decr(_F_) + incr(_R_) + _p_ = 2decr(_F_) + _p_
≤ 2(_n_log _p_+number of calls of make root)+_p_ ≤ 2(_n_log _p_ + 2_m_ + _p_) + _p_.

Since deciding calling make root or repair takes time _O_(1) (by the data structure),
we have that the algorithm takes time _O_(_m_ + _p_ + _n_log _p_).

```
As a consequence one has:
```

**Corollary 1.8a.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V and a length function_
_l_ : _A_ → Q+_, a shortest s_ − _t path can be found in time O_(|_A_| + |_V_ |log |_V_ |)_._
**Proof.** Directly from the description of the algorithm.

### 1.3. Shortest paths with arbitrary lengths

If lengths of arcs may take negative values, it is not always the case that a shortest
walk exists. If the graph has a directed circuit of negative length, then we can obtain
_s_ − _t_ walks of arbitrarily small negative length (for appropriate _s_ and _t_).
However, it can be shown that if there are no directed circuits of negative length,
then for each choice of _s_ and _t_ there exists a shortest _s_ − _t_ walk (if there exists at
least one _s_ − _t_ path).

**Theorem 1.9.** _Let each directed circuit have nonnegative length. Then for each pair_
_s, t of vertices for which there exists at least one s_ − _t walk, there exists a shortest_

```
s − t walk, which is a path.
```

```
Proof. Clearly, if there exists an s − t walk, there exists an s − t path. Hence there
exists also a shortest path P, that is, an s − t path that has minimum length among
all s − t paths. This follows from the fact that there exist only finitely many paths.
We show that P is shortest among all s − t walks. Let P have length L. Suppose
that there exists an s − t walk Q of length less than L. Choose such a Q with a
minimum number of arcs. Since Q is not a path (as it has length less than L), Q
contains a directed circuit C. Let Q′ be the walk obtained from Q by removing C.
As l(C) ≥ 0, l(Q′) = l(Q) − l(C) ≤ l(Q) < L. So Q′ is another s − t walk of length
less than L, however with a smaller number of arcs than Q. This contradicts the
assumption that Q has a minimum number of arcs.
```

```
Also in this case there is an easy algorithm, the Bellman-Ford method (Bellman
[1958], Ford [1956]), determining a shortest s − t path.
Let n := |V |. The algorithm calculates functions f0, f1, f2,... , fn : V → R ∪{∞}
successively by the following rule:
```

(15) (i) Put _f_[^0]:(_s_) := 0 and _f_[^0]:(_v_) := ∞ for all _v_ ∈ _V_ \{_s_}.
(ii) For _k < n_, if _fk_ has been found, put

```
fk+1(v) := min{fk(v), min
(u,v)∈A(fk(u) + l(u, v))}
```

```
for all v ∈ V.
```

```
Then, assuming that there is no directed circuit of negative length, fn(v) is equal to
the length of a shortest s − v walk, for each v ∈ V. (If there is no s − v path at all,
fn(v) = ∞.)
This follows directly from the following theorem:
```

```
Theorem 1.10. For each k = 0,... , n and for each v ∈ V ,
```

```
(16) fk(v) = min{l(P) |P is an s − v walk traversing at most k arcs}.
```

```
Proof. By induction on k from (15).
```

```
So the above method gives us the length of a shortest s−t path. It is not difficult
to derive a method finding an explicit shortest s − t path. To this end, determine
parallel to the functions f0,... , fn, a function g : V → V by setting g(v) = u when
we set fk+1(v) := fk(u) + l(u, v) in (15)(ii). At termination, for any v, the sequence
v, g(v), g(g(v)),... , s gives the reverse of a shortest s − v path. Therefore:
```

**Corollary 1.10a.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V and a length function_
_l_ : _A_ → Q_, such that D has no negative-length directed circuit, a shortest s_ − _t path_
_can be found in time O_(|_V_ ||_A_|)_._

**Proof.** Directly from the description of the algorithm.

**Application 1.3: Knapsack problem.** Suppose we have a knapsack with a volume of
8 liter and a number of articles 1_,_[^2]:_,_[^3]:_,_[^4]:_,_5. Each of the articles has a certain volume and a
certain value:

(17) article volume value
1 5 4
2 3 7
3 2 3
4 2 5
5 1 4

So we cannot take all articles in the knapsack and we have to make a selection. We want
to do this so that the total value of articles taken into the knapsack is as large as possible.

```
We can describe this problem as one of finding x1,x2,x3,x4,x5 such that:
```

(18) _x_[^1]:_,x_[^2]:_,x_[^3]:_,x_[^4]:_,x_[^5]: ∈{[^0]:_,_[^1]:},
[^5]:_x_[^1]: + 3_x_[^2]: + 2_x_[^3]: + 2_x_[^4]: + _x_[^5]: ≤ 8,
[^4]:_x_[^1]: + 7_x_[^2]: + 3_x_[^3]: + 5_x_[^4]: + 4_x_[^5]: is as large as possible.

We can solve this problem with the shortest path method as follows. Make a directed graph
in the following way (cf. Figure 1.2):

There are vertices (_i,x_) for 0 ≤ _i_ ≤ 6 and 0 ≤ _x_ ≤ 8 and there is an arc from (_i_ − [^1]:_,x_)
to (_i,y_) if _y_ = _x_ or _y_ = _x_ + _ai_ (where _ai_ is the volume of article _i_) if _i_ ≤ 5 and there are
arcs from each (5_,x_) to (6_,_8). We have deleted in the picture all vertices and arcs that do
not belong to any directed path from (0_,_0).

The length of arc ((_i_−[^1]:_,x_)_,_(_i,y_)) is equal to 0 if _y_ = _x_ and to −_ci_ if _y_ = _x_+_ai_ (where
_ci_ denotes the value of _i_). Moreover, all arcs ending at (6_,_8) have length 0.
Now a shortest path from (0_,_0) to (6_,_8) gives us the optimal selection.

**Application 1.4: PERT-CPM.** For building a house certain activities have to be ex-
ecuted. Certain activities have to be done before other and every activity takes a certain
number of days:

```
-4
```

```
0 0 0
```

```
0 0 0
```

```
0 0
```

```
0 0 0 0 0
```

```
0 0
```

```
0
```

```
0
```

```
0
```

```
0
```

```
0
```

```
0
```

```
0
```

```
0
```

```
-7
```

```
-7
-3
```

```
-3
```

```
-3
-5
```

```
-5
```

```
-5
```

```
-5
```

```
-4
```

```
-4
```

```
-4
```

```
-4
```

```
-4
```

```
0 0 0 0
```

```
0 1 2 3 4 5 6
```

```
0
```

```
1
```

```
2
```

```
3
```

```
4
```

```
5
```

```
6
```

```
7
```

```
8
```

```
0
```

```
-4
```

```
Figure 1.2
```

(19) activity days needed to be done before
activity #

1. groundwork 2 2
2. foundation 4 3
3. building walls 10 4,6,7
4. exterior plumbing 4 5,9
5. interior plumbing 5 10
6. electricity 7 10
7. roof 6 8
8. finishing off outer walls 7 9
9. exterior painting 9 14
10. panelling 8 11,12
11. floors 4 13
12. interior painting 5 13
13. finishing off interior 6
14. finishing off exterior 2

We introduce two dummy activities 0 (start) and 15 (completion), each taking 0 days, where
activity 0 has to be performed before all other activities and 15 after all other activities.
The project can be represented by a directed graph _D_ with vertices 0_,_[^1]:_,...,_[^14]:_,_15,
where there is an arc from _i_ to _j_ if _i_ has to be performed before _j_. The length of arc (_i,j_)
will be the number _ti_ of days needed to perform activity _i_. This graph with length function
is called the _project network_.

```
2 4
```

```
10
10
10
```

```
4
4
```

```
5
```

```
7
```

```
7 9
```

```
8 4 6
```

```
2
```

```
1 2 3
```

```
4
```

```
6
```

```
7 8 9
```

```
5
```

```
10
```

```
14
```

```
11
```

```
12
```

```
13
```

```
15
```

```
8 5
```

```
6
```

```
0
0
```

```
Figure 1.3
```

Now a _longest_ path from 0 to 15 gives the minimum number of days needed to build the
house. Indeed, if _li_ denotes the length of a longest path from 0 to _i_, we can start activity _i_
on day _li_. If activity _j_ has been done after activity _i_, then _lj_ ≥ _li_+_ti_ by definition of longest
path. So there is sufficient time for completing activity _i_ and the schedule is practically
feasible. That is, there is the following min-max relation:

(20) the minimum number of days needed to finish the project is equal to the maxi-
mum length of a path in the project network.

A longest path can be found with the Bellman-Ford method, as it is equivalent to a
shortest path when we replace each length by its opposite. Note that _D_ should not have
any directed circuits since otherwise the whole project would be infeasible.
So the project network helps in planning the project and is the basis of the so-called
‘Program Evaluation and Review Technique’ (PERT). (Actually, one often represents ac-
tivities by arcs instead of vertices, giving a more complicated way of defining the graph.)
Any longest path from 0 to 15 gives the minimum number of days needed to complete
the project. Such a path is called a _critical path_ and gives us the bottlenecks in the project.
It tells us which activities should be controlled carefully in order to meet a deadline. At
least one of these activities should be sped up if we wish to complete the project faster.
This is the basis of the ‘Critical Path Method’ (CPM).

**Application 1.5: Price equilibrium.** A small example of an economical application is
as follows. Consider a number of remote villages, say _B,C,D,E_ and _F_. Certain pairs of
villages are connected by routes (like in Figure 1.4).
If villages _X_ and _Y_ are connected by a route, let _kX,Y_ be the cost of transporting one
liter of oil from _X_ to _Y_.

```
B
```

```
C D
```

```
E F
Figure 1.4
```

At a certain day, one detects an oil well in village _B_, and it makes oil freely available
in village _B_. Now one can follow how the oil price will develop, assuming that no other oil
than that from the well in _B_ is available and that only once a week there is contact between
adjacent villages.
It will turn out that the oil prices in the different villages will follow the iterations in
the Bellman-Ford algorithm. Indeed in week 0 (the week in which the well was detected)
the price in _B_ equals 0, while in all other villages the price is ∞, since there is simply no
oil available yet.
In week 1, the price in _B_ equals 0, the price in any village _Y_ adjacent to _B_ is equal to
_kB,Y_ per liter and in all other villages it is still ∞.
In week _i_ + 1 the liter price _pi_+1_,Y_ in any village _Y_ is equal to the minimum value of
_pi,Y_ and all _pi,X_ + _kX,Y_ for which there is a connection from _X_ to _Y_.
There will be price equilibrium if for each village _Y_ one has:

(21) it is not cheaper for the inhabitants of _Y_ to go to an adjacent village _X_ and to
transport the oil from _X_ to _Y_.

Moreover, one has the min-max relation for each village _Y_ :

(22) the maximum liter price in village _Y_ is equal to the the minimum length of a
path in the graph from _B_ to _Y_

taking _kX,Y_ as length function.
A comparable, but less spatial example is: the vertices of the graph represent oil prod-
ucts (instead of villages) and _kX,Y_ denotes the cost per unit of transforming oil product _X_
to oil product _Y_. If oil product _B_ is free, one can determine the costs of the other products
in the same way as above.

**Exercises**

```
1.2. Find with the Bellman-Ford method shortest paths from s to each of the other vertices
in the following graphs (where the numbers at the arcs give the length):
```

```
(i)
```

```
3
```

```
−2
```

```
1
```

```
4
```

```
−1 −3
```

```
3
```

```
7 2
```

```
1
```

```
2 −5
```

```
s
```

```
(ii)
```

```
2 1 −2 −4 7 −1
```

```
1 1
```

```
3 −3 4
```

```
7 4 3 −8 3 2
```

```
2 −4
```

```
s
```

1.3. Let be given the distance table:

```
to: A B C D E F G
from: A 0 1 ∞ ∞ ∞ 2 12
B ∞[^0] ∞ ∞ ∞ ∞ ∞
C ∞ −15 0 4 8 ∞ ∞
D ∞ ∞ 4 0 ∞ ∞ −2
E ∞ ∞ ∞ 4 0 ∞ ∞
F ∞ ∞ ∞ 9 3 0 12
G ∞ −12 2 3 −[^1] −4 0
```

```
A distance ∞ from X to Y should be interpreted as no direct route existing from X
to Y.
Determine with the Bellman-Ford method the distance from A to each of the other
cities.
```

1.4. Solve the knapsack problem of Application 1.3 with the Bellman-Ford method.

1.5. Describe an algorithm that tests if a given directed graph with length function con-
tains a directed circuit of negative length.

```
1.6. Let D = (V,A) be a directed graph and let s and t be vertices of D, such that t is
reachable from s. Show that the minimum number of arcs in an s − t path is equal
to the maximum value of φ(t) − φ(s), where φ ranges over all functions φ : V → Z
such that φ(w) − φ(v) ≤ 1 for each arc (v,w).
```

### 1.4. Minimum spanning trees

Let _G_ = (_V, E_) be a connected undirected graph and let _l_ : _E_ → R be a function,
called the _length_ function. For any subset _F_ of _E_, the _length l_(_F_) of _F_ is, by definition:

(23) _l_(_F_) :=

```
∑
```

```
e∈F
```

```
l(e).
```

In this section we consider the problem of finding a spanning tree in _G_ of minimum
length. There is an easy algorithm for finding a minimum-length spanning tree,
essentially due to Bor ̊uvka [1926]. There are a few variants. The first one we discuss
is sometimes called the _Dijkstra-Prim method_ (after Prim [1957] and Dijkstra [1959]).

Choose a vertex _v_[^1]: ∈ _V_ arbitrarily. Determine edges _e_[^1]:_, e_[^2]:_..._ successively as
follows. Let _U_[^1]: := {_v_[^1]:}. Suppose that, for some _k_ ≥ 0, edges _e_[^1]:_,... , ek_ have been
chosen, forming a spanning tree on the set _Uk_. Choose an edge _ek_+1 ∈ _δ_(_Uk_) that has
minimum length among all edges in _δ_(_Uk_).[^7] Let _Uk_+1 := _Uk_ ∪ _ek_+1.
By the connectedness of _G_ we know that we can continue choosing such an edge
until _Uk_ = _V_. In that case the selected edges form a spanning tree _T_ in _G_. This tree
has minimum length, which can be seen as follows.
Call a forest _F greedy_ if there exists a minimum-length spanning tree _T_ of _G_ that
contains _F_.

**Theorem 1.11.** _Let F be a greedy forest, let U be one of its components, and let_
_e_ ∈ _δ_(_U_)_. If e has minimum length among all edges in δ_(_U_)_, then F_ ∪{_e_} _is again a_
_greedy forest._

**Proof.** Let _T_ be a minimum-length spanning tree containing _F_. Let _P_ be the unique
path in _T_ between the end vertices of _e_. Then _P_ contains at least one edge _f_
that belongs to _δ_(_U_). So _T_′ := (_T_ \{_f_}) ∪{_e_} is a tree again. By assumption,
_l_(_e_) ≤ _l_(_f_) and hence _l_(_T_′) ≤ _l_(_T_). Therefore, _T_′ is a minimum-length spanning tree.
As _F_ ∪{_e_}⊆ _T_′, it follows that _F_ ∪{_e_} is greedy.

**Corollary 1.11a.** _The Dijkstra-Prim method yields a spanning tree of minimum_
_length._

```
7δ(U) is the set of edges e satisfying |e ∩ U| = 1.
```

**Proof.** It follows inductively with Theorem 1.11 that at each stage of the algorithm
we have a greedy forest. Hence the final tree is greedy — equivalently, it has minimum
length.

```
In fact one may show:
```

**Theorem 1.12.** _Implementing the Dijkstra-Prim method using Fibonacci heaps gives_
_a running time of O_(|_E_| + |_V_ |log |_V_ |)_._

**Proof.** The Dijkstra-Prim method is similar to Dijkstra’s method for finding a short-
est path. Throughout the algorithm, we store at each vertex _v_ ∈ _V_ \ _Uk_, the length
_f_(_v_) of a shortest edge {_u, v_} with _u_ ∈ _Uk_, organized as a Fibonacci heap. A vertex
_uk_+1 to be added to _Uk_ to form _Uk_+1 should be identified and removed from the Fi-
bonacci heap. Moreover, for each edge _e_ connecting _uk_+1 and some _v_ ∈ _V_ \ _Uk_+1, we
should update _f_(_v_) if the length of _uk_+1_v_ is smaller than _f_(_v_).
Thus we find and delete ≤|_V_ | times a _u_ minimizing _f_(_u_) and we decrease ≤|_E_|
times a value _f_(_v_). Hence by Theorem 1.8 the algorithm can be performed in time
_O_(|_E_| + |_V_ |log |_V_ |).

The Dijkstra-Prim method is an example of a so-called _greedy_ algorithm. We
construct a spanning tree by throughout choosing an edge that seems the best at the
moment. Finally we get a minimum-length spanning tree. Once an edge has been
chosen, we never have to replace it by another edge (no ‘back-tracking’).
There is a slightly different method of finding a minimum-length spanning tree,
_Kruskal’s_ method (Kruskal [1956]). It is again a greedy algorithm, and again itera-
tively edges _e_[^1]:_, e_[^2]:_,..._ are chosen, but by some different rule.
Suppose that, for some _k_ ≥ 0, edges _e_[^1]:_,... , ek_ have been chosen. Choose an edge
_ek_+1 such that {_e_[^1]:_,... , ek, ek_+1} forms a forest, with _l_(_ek_+1) as small as possible. By
the connectedness of _G_ we can (starting with _k_ = 0) iterate this until the selected
edges form a spanning tree of _G_.

**Corollary 1.12a.** _Kruskal’s method yields a spanning tree of minimum length._

**Proof.** Again directly from Theorem 1.11.

```
In a similar way one finds a maximum-length spanning tree.
```

**Application 1.6: Minimum connections.** There are several obvious practical situations
where finding a minimum-length spanning tree is important, for instance, when designing a
road system, electrical power lines, telephone lines, pipe lines, wire connections on a chip.
Also when clustering data say in taxonomy, archeology, or zoology, finding a minimum
spanning tree can be helpful.

**Application 1.7: The maximum reliability problem.** Often in designing a network
one is not primarily interested in minimizing length, but rather in maximizing ‘reliability’
(for instance when designing energy or communication networks). Certain cases of this
problem can be seen as finding a _maximum_ length spanning tree, as was observed by Hu
[1961]. We give a mathematical description.

Let _G_ = (_V,E_) be a graph and let _s_ : _E_ → R+ be a function. Let us call _s_(_e_) the
_strength_ of edge _e_. For any path _P_ in _G_, the _reliability_ of _P_ is, by definition, the minimum
strength of the edges occurring in _P_. The _reliability rG_(_u,v_) of two vertices _u_ and _v_ is equal
to the maximum reliability of _P_, where _P_ ranges over all paths from _u_ to _v_.
Let _T_ be a spanning tree of maximum strength, i.e., with

∑
_e_∈_ET s_(_e_) as large as possible.
(Here _ET_ is the set of edges of _T_.) So _T_ can be found with any maximum spanning tree
algorithm.
Now _T_ has the same reliability as _G_, for each pair of vertices _u,v_. That is:

(24) _rT_(_u,v_) = _rG_(_u,v_) for each _u,v_ ∈ _V_.

We leave the proof as an exercise (Exercise 1.11).

**Exercises**

```
1.7. Find, both with the Dijkstra-Prim algorithm and with Kruskal’s algorithm, a span-
ning tree of minimum length in the graph in Figure 1.5.
```

```
3 2
```

```
2 4 1
```

```
5 3
```

```
3 6 3
```

```
5 4 4 6 2 3
```

```
4 3 5 7 4 2
```

```
Figure 1.5
```

```
1.8. Find a spanning tree of minimum length between the cities given in the following
distance table:
```

Ame Ams Ape Arn Ass BoZ Bre Ein Ens s-G Gro Haa DH s-H Hil Lee Maa Mid Nij Roe Rot Utr Win Zut Zwo
Amersfoort 0 47 47 46 139 123 86 111 114 81 164 67 126 73 18 147 190 176 63 141 78 20 109 65 70
Amsterdam 47 0 89 92 162 134 100 125 156 57 184 20 79 87 30 132 207 175 109 168 77 40 151 107 103
Apeldoorn 47 89 0 25 108 167 130 103 71 128 133 109 154 88 65 129 176 222 42 127 125 67 66 22 41
Arnhem 46 92 25 0 132 145 108 78 85 116 157 112 171 63 64 154 151 200 17 102 113 59 64 31 66
Assen 139 162 108 132 0 262 225 210 110 214 25 182 149 195 156 68 283 315 149 234 217 159 143 108 69
Bergen op Zoom 123 134 167 145 262 0 37 94 230 83 287 124 197 82 119 265 183 59 128 144 57 103 209 176 193
Breda 86 100 130 108 225 37 0 57 193 75 250 111 179 45 82 228 147 96 91 107 49 66 172 139 156
Eindhoven 111 125 103 78 210 94 57 0 163 127 235 141 204 38 107 232 100 153 61 50 101 91 142 109 144
Enschede 114 156 71 85 110 230 193 163 0 195 135 176 215 148 132 155 236 285 102 187 192 134 40 54 71
’s-Gravenhage 81 57 128 116 214 83 75 127 195 0 236 41 114 104 72 182 162 124 133 177 26 61 180 146 151
Groningen 164 184 133 157 25 287 250 235 135 236 0 199 147 220 178 58 308 340 174 259 242 184 168 133 94
Haarlem 67 20 109 112 182 124 111 141 176 41 199 0 73 103 49 141 203 165 129 184 67 56 171 127 123
Den Helder 126 79 154 171 149 197 179 204 215 114 147 73 0 166 109 89 276 238 188 247 140 119 220 176 144
’s-Hertogenbosch 73 87 88 63 195 82 45 38 148 104 220 103 166 0 69 215 123 141 46 81 79 53 127 94 129
Hilversum 18 30 65 64 156 119 82 107 132 72 178 49 109 69 0 146 192 172 81 150 74 16 127 83 88
Leeuwarden 147 132 129 154 68 265 228 232 155 182 58 141 89 215 146 0 306 306 171 256 208 162 183 139 91
Maastricht 190 207 176 151 283 183 147 100 236 162 308 203 276 123 192 305 0 242 135 50 188 176 213 182 217
Middelburg 176 175 222 200 315 59 96 153 285 124 340 165 238 141 172 306 242 0 187 203 98 156 264 231 246
Nijmegen 63 109 42 17 149 128 91 61 102 133 174 129 188 46 81 171 135 187 0 85 111 76 81 48 83
Roermond 141 168 127 102 234 144 107 50 187 177 259 184 247 81 150 256 50 203 85 0 151 134 166 133 168
Rotterdam 78 77 125 113 217 57 49 101 192 26 242 67 140 79 74 208 188 98 111 151 0 58 177 143 148
Utrecht 20 40 67 59 159 103 66 91 134 61 184 56 119 53 16 162 176 156 76 134 58 0 123 85 90
Winterswijk 109 151 66 64 143 209 172 142 40 180 168 171 220 127 127 183 213 264 81 166 177 123 0 44 92
Zutphen 65 107 22 31 108 176 139 109 54 146 133 127 176 94 83 139 182 231 48 133 143 85 44 0 48
Zwolle 70 103 41 66 69 193 156 144 71 151 94 123 144 129 88 91 217 246 83 168 148 90 92 48 0

```
1.9. Let G = (V,E) be a graph and let l : E → R be a ‘length’ function. Call a forest F
good if l(F′) ≥ l(F) for each forest F′ satisfying |F′| = |F|.
Let F be a good forest and e be an edge not in F such that F ∪{e} is a forest and
such that (among all such e) l(e) is as small as possible. Show that F ∪{e} is good
again.
```

```
1.10. Let G = (V,E) be a complete graph and let l : E → R+ be a length function satisfying
l(uw) ≥ min{l(uv),l(vw)} for all distinct u,v,w ∈ V. Let T be a longest spanning
tree in G.
Show that for all u,w ∈ V , l(uw) is equal to the minimum length of the edges in the
unique u − w path in T.
```

```
1.11. Prove (24).
```

## 2. Polytopes, polyhedra, Farkas’

## lemma, and linear programming

### 2.1. Convex sets

```
A subset C of Rn is called convex if for all x, y in C and any 0 ≤ λ ≤ 1 also λx+(1−λ)y
belongs to C. So C is convex if with any two points in C, the whole line segment
connecting x and y belongs to C.
Clearly, the intersection of any number of convex sets is again a convex set. So,
for any subset X of Rn, the smallest convex set containing X exists. This set is called
the convex hull of X and is denoted by conv.hull(X). One easily proves:
```

```
(1) conv.hull(X) = {x |∃t ∈ N,∃x1,... , xt ∈ X,∃λ1,... , λt ≥ 0 :
x = λ1x1 + ··· + λtxt, λ1 + ··· + λt = 1}.
```

```
A basic property of closed convex sets is that any point not in C can be separated
from C by a ‘hyperplane’. Here a subset H of Rn is called a hyperplane (or an affine
hyperplane) if there exist a vector c ∈ Rn with c 6= 0 and a δ ∈ R such that:
```

```
(2) H = {x ∈ Rn | cTx = δ}.
```

```
We say that H separates z and C if z and C are in different components of Rn \ H.
```

```
Theorem 2.1. Let C be a closed convex set in Rn and let z 6∈ C. Then there exists
a hyperplane separating z and C.
```

```
Proof. Since the theorem is trivial if C = ∅, we assume C 6= ∅. Then there exists a
vector y in C that is nearest to z, i.e., that minimizes ‖z − y‖.
(The fact that such a y exists, can be seen as follows. Since C 6= ∅, there exists
an r > 0 such that B(z, r) ∩ C 6= ∅. Here B(z, r) denotes the closed ball with center
z and radius r. Then y minimizes the continuous function ‖z − y‖ over the compact
set B(z, r) ∩ C.)
Now define:
```

```
(3) c := z − y, δ :=
```

```
1
2
```

```
(‖z‖[^2] −‖y‖[^2]).
```

```
We show
```

(4) (i) _cTz > δ_,
(ii) _cTx < δ_ for each _x_ ∈ _C_.

Indeed, _cTz_ = (_z_ − _y_)_Tz >_ (_z_ − _y_)_Tz_ −[^1]
[^2]:‖_z_ − _y_‖[^2] = _δ_. This shows (4)(i).
If (4)(ii) would not hold, there exists an _x_ in _C_ such that _cTx_ ≥ _δ_. Since _cTy <_
_cTy_ +[^1]
[^2]:‖_c_‖[^2] = _δ_, we know _cT_(_x_ − _y_) _>_ 0. Hence there exists a _λ_ with 0 _< λ_ ≤ 1 and

(5) _λ <_

```
2cT(x − y)
‖x − y‖[^2].
```

Define

(6) _w_ := _λx_ + (1 − _λ_)_y._

So _w_ belongs to _C_. Moreover,

(7) ‖_w_ − _z_‖[^2] = ‖_λ_(_x_ − _y_) + (_y_ − _z_)‖[^2] = ‖_λ_(_x_ − _y_) − _c_‖[^2]
= _λ_[^2]‖_x_ − _y_‖[^2] − [^2]:_λcT_(_x_ − _y_) + ‖_c_‖[^2] _<_ ‖_c_‖[^2] = ‖_y_ − _z_‖[^2]_._

Here _<_ follows from (5).
However, (7) contradicts the fact that _y_ is a point in _C_ nearest to _z_.

Theorem 2.1 implies a characterization of closed convex sets – see Exercise 2.1.
Call a subset _H_ of R_n_ a _halfspace_ (or an _affine halfspace_) if there exist a vector _c_ ∈ R_n_
with _c_ [^6]:= 0 and a _δ_ ∈ R such that

(8) _H_ = {_x_ ∈ R_n_ | _cTx_ ≤ _δ_}_._

Clearly, each affine halfspace is a closed convex set.
Theorem 2.1 implies that if _C_ is a closed convex set and _z_ 6∈ _C_, then there exists
an affine halfspace _H_ so that _C_ ⊆ _H_ and _z_ 6∈ _H_.

**Exercises**

```
2.1. Let C ⊆ Rn. Then C is a closed convex set if and only if C =
```

```
⋂
F for some collection
F of affine halfspaces.
```

```
2.2. Let C ⊆ Rn be a convex set and let A be an m × n matrix. Show that the set
{Ax | x ∈ C} is again convex.
```

```
2.3. Let X be a finite set of vectors in Rn. Show that conv.hull(X) is compact.
(Hint: Show that conv.hull(X) is the image under a continuous function of a compact
set.)
```

```
2.4. Show that if z ∈ conv.hull(X), then there exist affinely independent vectors x1,...,xm
in X such that z ∈ conv.hull{x1,...,xm}. (This is the affine form of ‘Carath ́eodory’s
theorem’ (Carath ́eodory [1911]).)
```

```
(Vectors x1,...,xm are called affinely independent if there are no reals λ1,...,λm,
such that λ1x1 +···+λmxm = 0 and λ1 +···+λm = 0 and such that λ1,...,λm are
not all equal to 0.)
```

```
2.5. (i) Let C and D be two nonempty, bounded, closed, convex subsets of Rn such that
C ∩ D = ∅. Derive from Theorem 2.1 that there exists an affine hyperplane H
separating C and D. (This means that C and D are in different components of
Rn \ H.)
(Hint: Consider the set C − D := {x − y | x ∈ C,y ∈ D}.)
(ii) Show that in (i) we cannot delete the boundedness condition.
```

### 2.2. Polytopes and polyhedra

Special classes of closed convex sets are formed by the polytopes and the polyhedra.
In the previous section we saw that each closed convex set is the intersection of affine
halfspaces, possibly infinitely many. If it is the intersection of a _finite_ number of affine
halfspaces, the convex set is called a _polyhedron_.
So a subset _P_ of R_n_ is a polyhedron if and only if there exists an _m_ × _n_ matrix _A_
and a vector _b_ ∈ R_m_ such that

(9) _P_ = {_x_ ∈ R_n_ | _Ax_ ≤ _b_}_._

Here _Ax_ ≤ _b_ means:

(10) _a_[^1]:_x_ ≤ _b_[^1]:_,... , amx_ ≤ _bm,_

where _a_[^1]:_,... , am_ are the rows of _A_.
The matrix _A_ may have zero rows, i.e. _m_ = 0. In that case, _P_ = R_n_.
Related is the notion of ‘polytope’. A subset _P_ of R_n_ is called a _polytope_ if _P_ is
the convex hull of a finite number of vectors. That is, there exist vectors _x_[^1]:_,... , xt_
in R_n_ such that

(11) _P_ = conv.hull{_x_[^1]:_,... , xt_}_._

We will show that a subset _P_ of R_n_ is a polytope if and only if it is a bounded
polyhedron. This might be intuitively clear, but a strictly mathematical proof requires
some work.
We first give a definition. Let _P_ be a convex set. A point _z_ ∈ _P_ is called a
_vertex_ of _P_ if _z_ is _not_ a convex combination of two other points in _P_. That is, there
do not exist points _x, y_ in _P_ and a _λ_ with 0 _< λ <_ 1 such that _x_ [^6]:= _z, y_ [^6]:= _z_ and
_z_ = _λx_ + (1 − _λ_)_y_.

To characterize vertices we introduce the following notation. Let _P_ = {_x_ | _Ax_ ≤
_b_} be a polyhedron and let _z_ ∈ _P_. Then _Az_ is the submatrix of _A_ consisting of those
rows _ai_ of _A_ for which _aiz_ = _bi_.
Then we can show:

**Theorem 2.2.** _Let P_ = {_x_ | _Ax_ ≤ _b_} _be a polyhedron in_ R_n and let z_ ∈ _P. Then z_
_is a vertex of P if and only if_ rank(_Az_) = _n._

**Proof.** _Necessity._ Let _z_ be a vertex of _P_ and suppose rank(_Az_) _< n_. Then there
exists a vector _c_ [^6]:= 0 such that _Azc_ = 0. Since _aiz < bi_ for every _ai_ that does not
occur in _Az_, there exists a _δ >_ 0 such that:

(12) _ai_(_z_ + _δc_) ≤ _bi_ and _ai_(_z_ − _δc_) ≤ _bi_

for every row _ai_ of _A_ not occurring in _Az_. Since _Azc_ = 0 and _Az_ ≤ _b_ it follows that

(13) _A_(_z_ + _δc_) ≤ _b_ and _A_(_z_ − _δc_) ≤ _b._

So _z_+_δc_ and _z_−_δc_ belong to _P_. Since _z_ is a convex combination of these two vectors,
this contradicts the fact that _z_ is a vertex of _P_.

_Sufficiency._ Suppose rank(_Az_) = _n_ while _z_ is not a vertex of _P_. Then there exist
points _x_ and _y_ in _P_ such that _x_ [^6]:= _z_ [^6]:= _y_ and _z_ =[^1]
[^2]:(_x_ + _y_). Then for every row _ai_ of
_Az_:

(14) _aix_ ≤ _bi_ = _aiz_ =⇒ _ai_(_x_ − _z_) ≤ 0, and
_aiy_ ≤ _bi_ = _aiz_ =⇒ _ai_(_y_ − _z_) ≤ 0.

Since _y_ − _z_ = −(_x_ − _z_), this implies that _ai_(_x_ − _z_) = 0. Hence _Az_(_x_ − _z_) = 0. Since
_x_ − _z_ [^6]:= 0, this contradicts the fact that rank(_Az_) = _n_.

Theorem 2.2 implies that a polyhedron has only a finite number of vertices: For
each two different vertices _z_ and _z_′ one has _Az_ [^6]:= _Az_′, since _Azx_ = _bz_ has only one
solution, namely _x_ = _z_ (where _bz_ denotes the part of _b_ corresponding to _Az_). Since
there exist at most 2_m_ collections of subrows of _A_, _P_ has at most 2_m_ vertices.
From Theorem 2.2 we derive:

**Theorem 2.3.** _Let P be a bounded polyhedron, with vertices x_[^1]:_,... , xt. Then P_ =
conv.hull{_x_[^1]:_,... , xt_}_._

**Proof.** Clearly

(15) conv.hull{_x_[^1]:_,... , xt_}⊆ _P,_

```
since x1,... , xt belong to P and since P is convex.
The reverse inclusion amounts to:
```

```
(16) if z ∈ P then z ∈ conv.hull{x1,... , xt}.
```

```
We show (16) by induction on n − rank(Az).
If n − rank(Az) = 0, then rank(Az) = n, and hence, by Theorem 2.2, z itself is a
vertex of P. So z ∈ conv.hull{x1,... , xt}.
If n − rank(Az) > 0, then there exists a vector c 6= 0 such that Azc = 0. Define
```

```
(17) μ0 := max{μ | z + μc ∈ P},
ν0 := max{ν | z − νc ∈ P}.
```

```
These numbers exist since P is compact. Let x := z + μ0c and y := z − ν0c.
Now
```

```
(18) μ0 = min{
```

```
bi − aiz
aic
```

```
| ai is a row of A; aic > 0}.
```

```
This follows from the fact that μ0 is the largest μ such that ai(z + μc) ≤ bi for each
i = 1,... , m. That is, it is the largest μ such that
```

```
(19) μ ≤
bi − aiz
aic
```

```
for every i with aic > 0.
Let the minimum (18) be attained by i0. So for i0 we have equality in (18).
Therefore
```

(20) (i) _Azx_ = _Azz_ + _μ_[^0]:_Azc_ = _Azz,_

```
(ii) ai0x = ai0(z + μ0c) = bi0.
```

```
So Ax contains all rows in Az, and moreover it contains row ai0. Now Azc = 0
while ai0c 6= 0. This implies rank(Ax) > rank(Az). So by our induction hypothesis, x
belongs to conv.hull{x1,... , xt}. Similarly, y belongs to conv.hull{x1,... , xt}. There-
fore, as z is a convex combination of x and y, z belongs to conv.hull{x1,... , xt}.
```

```
As a direct consequence we have:
```

```
Corollary 2.3a. Each bounded polyhedron is a polytope.
```

```
Proof. Directly from Theorem 2.3.
```

```
Conversely:
```

**Theorem 2.4.** _Each polytope is a bounded polyhedron._

**Proof.** Let _P_ be a polytope in R_n_, say

(21) _P_ = conv.hull{_x_[^1]:_,... , xt_}_._

We may assume that _t_ ≥ 1. We prove the theorem by induction on _n_. Clearly, _P_ is
bounded.
If _P_ is contained in some affine hyperplane, the theorem follows from the induction
hypothesis.
So we may assume that _P_ is not contained in any affine hyperplane. It implies
that the vectors _x_[^2]: − _x_[^1]:_,... , xt_ − _x_[^1]: span R_n_. It follows that there exist a vector _x_0
in _P_ and a real _r >_ 0 such that the ball _B_(_x_[^0]:_, r_) is contained in _P_.
Without loss of generality, _x_[^0]: = 0. Define _P_∗ by

(22) _P_∗ := {_y_ ∈ R_n_ | _xTy_ ≤ 1 for each _x_ ∈ _P_}_._

```
Then P∗ is a polyhedron, as
```

(23) _P_∗ = {_y_ ∈ R_n_ | _xT_
_j y_ ≤ 1 for _j_ = 1_,... , t_}_._

This follows from the fact that if _y_ belongs to the right hand set in (23) and _x_ ∈ _P_
then _x_ = _λ_[^1]:_x_[^1]: + ··· + _λtxt_ for certain _λ_[^1]:_,... , λt_ ≥ 0 with _λ_[^1]: + ··· + _λt_ = 1, implying

(24) _xTy_ =

```
∑t
```

```
j=1
```

```
λjxT
j y ≤
```

```
∑t
```

```
j=1
```

```
λj = 1.
```

So _y_ belongs to _P_∗.
Moreover, _P_∗ is bounded, since for each _y_ [^6]:= 0 in _P_∗ one has that _x_ := _r_ ·‖_y_‖−[^1] ·_y_
belongs to _B_(0_, r_) and hence to _P_. Therefore, _xTy_ ≤ 1, and hence

(25) ‖_y_‖ = (_xTy_)_/r_ ≤ [^1]:_/r._

So _P_∗ ⊆ _B_(0_,_[^1]:_/r_).
This proves that _P_∗ is a bounded polyhedron. By Corollary 2.3a, _P_∗ is a polytope.

So there exist vectors _y_[^1]:_,... , ys_ in R_n_ such that

(26) _P_∗ = conv.hull{_y_[^1]:_,... , ys_}_._

We show:

(27) _P_ = {_x_ ∈ R_n_ | _yT_
_j x_ ≤ 1 for all _j_ = 1_,... , s_}_._

This implies that _P_ is a polyhedron.
To see the inclusion ⊆ in (27), it suffices to show that each of the vectors _xi_
belongs to the right hand side in (27). This follows directly from the fact that for
each _j_ = 1_,... , s_, _yT_
_j xi_ = _xT_
_i yj_ ≤ 1, since _yj_ belongs to _P_∗.
To see the inclusion ⊇ in (25), let _x_ ∈ R_n_ be such that _yT_
_j x_ ≤ 1 for all _j_ = 1_,... , s_.
Suppose _x_ 6∈ _P_. Then there exists a hyperplane separating _x_ and _P_. That is, there
exist a vector _c_ [^6]:= 0 in R_n_ and a _δ_ ∈ R such that _cTx_′ _< δ_ for each _x_′ ∈ _P_, while
_cTx > δ_. As 0 ∈ _P_, _δ >_ 0. So we may assume _δ_ = 1. Hence _c_ ∈ _P_∗. So there exist
_μ_[^1]:_,... , μs_ ≥ 0 such that _c_ = _μ_[^1]:_y_[^1]: + ··· _μsys_ and _μ_[^1]: + ··· + _μs_ = 1. This gives the
contradiction:

(28) [^1]: _< cTx_ =

```
∑s
```

```
j=1
```

```
μjyT
j x ≤
```

```
∑s
```

```
j=1
```

```
μj = 1.
```

**Convex cones**

Convex cones are special cases of convex sets. A subset _C_ of R_n_ is called a _convex_
_cone_ if for any _x, y_ ∈ _C_ and any _λ, μ_ ≥ 0 one has _λx_ + _μy_ ∈ _C_.
For any _X_ ⊆ R_n_, cone(_X_) is the smallest cone containing _X_. One easily checks:

(29) cone(_X_) = {_λ_[^1]:_x_[^1]: + ···_λtxt_ | _x_[^1]:_,... , xt_ ∈ _X_;_λ_[^1]:_,... , λt_ ≥ [^0]:}_._

```
A cone C is called finitely generated if C = cone(X) for some finite set X.
```

**Exercises**

```
2.6. Determine the vertices of the following polyhedra:
```

```
(i) P = {(x,y) | x ≥ 0,y ≥ 0,y − x ≤ 2,x + y ≤ 8,x + 2y ≤ 10,x ≤ 4}.
(ii) P = {(x,y,z) | x + y ≤ 2,y + z ≤ 4,x + z ≤ 3,−2x − y ≤ 3,−y − 2z ≤
3,−2x − z ≤ 2}.
(iii) P = {(x,y) | x + y ≤ 1,x − y ≤ 2}.
(iv) P = {(x,y) | x + y = 1,x ≥ 3}.
(v) P = {(x,y,z) | x ≥ 0,y ≥ 0,x + y ≤ 1}.
(vi) P = {(x,y,z) | x + y ≥ 1,x + z ≥ 1,y − z ≥ 0}.
(vii) P = {(x,y) | 3x + 2y ≤ 18,x − y ≥−6,5x + 2y ≤ 20,x ≥ 0,y ≥ 0}.
```

```
2.7. Let C ⊆ Rn. Then C is a closed convex cone if and only if C =
```

```
⋂
F for some
collection F of linear halfspaces.
```

```
(A subset H of Rn is called a linear halfspace if H = {x ∈ Rn | cTx ≤ 0} for some
nonzero vector c.)
```

```
2.8. Show that if z ∈ cone(X), then there exist linearly independent vectors x1,...,xm
in X such that z ∈ cone{x1,...,xm}. (This is the linear form of ‘Carath ́eodory’s
theorem’.)
```

```
2.9. Let A be an m × n matrix of rank m and let b ∈ Rm. Derive from Exercise 2.8 that
the system Ax = b has a nonnegative solution x if and only if it has a nonnegative
basic solution.
(A submatrix B of A is called a basis of A if B is a nonsingular m × m submatrix of
A. A solution x of Ax = b is a basic solution if A has a basis B so that x is 0 in those
coordinates not corresponding to columns in B.)
```

2.10. Prove that every finitely generated convex cone is a closed set. (This can be derived
from Exercise 2.3 and Corollary 2.3a.)

2.11. Prove that a convex cone is finitely generated if and only if it is the intersection of
finitely many linear halfspaces.
(_Hint:_ Use Corollary 2.3a and Theorem 2.4.)

2.12. Let _P_ be a subset of R_n_. Show that _P_ is a polyhedron if and only if _P_ = _Q_ + _C_ for
some polytope _Q_ and some finitely generated convex cone _C_.

```
(Hint: Apply Exercise 2.11 to cone(X) in Rn+1, where X is the set of vectors
```

```
(
1
x
```

```
)
```

```
in Rn+1 with x ∈ P.)
```

2.13. For any subset _X_ of R_n_, define

```
(30) X∗ := {y ∈ Rn | xTy ≤ 1 for each x ∈ X}.
```

```
(i) Show that for each convex cone C, C∗ is a closed convex cone.
(ii) Show that for each closed convex cone C, (C∗)∗ = C.
```

2.14. Let _P_ be a polyhedron.

```
(i) Show that P∗ is again a polyhedron.
(Hint: Use previous exercises.)
(ii) Show that P contains the origin if and only if (P∗)∗ = P.
(iii) Show that the origin is an internal point of P if and only if P∗ is bounded.
```

### 2.3. Farkas’ lemma

Let _A_ be an _m_ × _n_ matrix and let _b_ ∈ R_m_. With the Gaussian elimination method
one can prove that

```
(31) Ax = b
```

```
has a solution x if and only if there is no solution y for the following system of linear
equations:
```

```
(32) yTA = 0, yTb = −1.
```

```
Farkas’ lemma (Farkas [1894,1896,1898]) gives an analogous characterization for
the existence of a nonnegative solution x for (31).
```

```
Theorem 2.5 (Farkas’ lemma). The system Ax = b has a nonnegative solution if
and only if there is no vector y satisfying yTA ≥ 0 and yTb < 0.
```

```
Proof. Necessity. Suppose Ax = b has a solution x0 ≥ 0 and suppose there exists a
vector y0 satisfying yT
0 A ≥ 0 and yT
0 b < 0. Then we obtain the contradiction
```

```
(33) 0 > yT
0 b = yT
0 (Ax[^0]) = (yT
0 A)x[^0] ≥[^0].
```

```
Sufficiency. Suppose Ax = b has no solution x ≥ 0. Let a1,... , an be the columns
of A. So
```

```
(34) b 6∈ C := cone{a1,... , an}.
```

```
So by Exercise 2.7 there exists a linear halfspace H containing C and not containing
b. That is, there exists a vector c such that cTb < 0 while cTx ≥ 0 for each x in C.
In particular, cTaj ≥ 0 for j = 1,... , n. So y := c satisfies yTA ≥ 0 and yTb < 0.
```

```
So Farkas’ lemma states that exactly one of the following two assertions is true:
```

(35) (i) ∃_x_ ≥ 0 : _Ax_ = _b_,
(ii) ∃_y_ : _yTA_ ≥ 0 and _yTb <_ 0.

```
There exist several variants of Farkas’ lemma, that can be easily derived from
Theorem 2.5.
```

```
Corollary 2.5a. The system Ax ≤ b has a solution x if and only if there is no vector
y satisfying y ≥ 0, yTA = 0 and yTb < 0.
```

```
Proof. Let A′ be the matrix
```

```
(36) A′ := [A − A I],
```

```
where I denotes the m × m identity matrix.
```

Then _Ax_ ≤ _b_ has a solution _x_ if and only if the system _A_′_x_′ = _b_ has a nonnegative
solution _x_′. Applying Theorem 2.5 to _A_′_x_′ = _b_ gives the corollary.

```
Another consequence is:
```

**Corollary 2.5b.** _Suppose that the system Ax_ ≤ _b has at least one solution. Then for_
_every solution x of Ax_ ≤ _b one has cTx_ ≤ _δ if and only if there exists a vector y_ ≥ 0
_such that yTA_ = _cT and yTb_ ≤ _δ._

**Proof.** _Sufficiency._ If such a vector _y_ exists, then for every vector _x_ one has

(37) _Ax_ ≤ _b_ =⇒ _yTAx_ ≤ _yTb_ =⇒ _cTx_ ≤ _yTb_ =⇒ _cTx_ ≤ _δ._

_Necessity._ Suppose that such a vector _y_ does not exist. It means that the following
system of linear inequalities in the variables _y_ and _λ_ has no solution (_yT λ_) ≥ (0 0):

(38) (_yT λ_)

```
(
A b
0 1
```

```
)
= (cT δ).
```

According to Farkas’ lemma this implies that there exists a vector

```
(
z
μ
```

```
)
so that
```

(39)

```
(
A b
0 1
```

```
)(
z
μ
```

```
)
≥
```

```
(
0
0
```

```
)
and (cT δ)
```

```
(
z
μ
```

```
)
< 0.
```

We distinguish two cases.

**Case 1:** _μ_ = 0. Then _Az_ ≥ 0 and _cTz <_ 0. However, by assumption, _Ax_ ≤ _b_ has
a solution _x_[^0]:. Then, for _τ_ large enough:

(40) _A_(_x_[^0]: − _τ z_) ≤ _b_ and _cT_(_x_[^0]: − _τ z_) _> δ._

This contradicts the fact that _Ax_ ≤ _b_ implies _cTx_ ≤ _δ_.

**Case 2:** _μ >_ 0. As (39) is homogeneous, we may assume that _μ_ = 1. Then for
_x_ := −_z_ one has:

(41) _Ax_ ≤ _b_ and _cTx > δ._

Again this contradicts the fact that _Ax_ ≤ _b_ implies _cTx_ ≤ _δ_.

**Exercises**

2.15. Prove that there exists a vector _x_ ≥ 0 such that _Ax_ ≤ _b_ if and only if for each _y_ ≥ 0
satisfying _yTA_ ≥ 0 one has _yTb_ ≥ 0.

2.16. Prove that there exists a vector _x >_ 0 such that _Ax_ = 0 if and only if for each _y_
satisfying _yTA_ ≥ 0 one has _yTA_ = 0. (Stiemke’s theorem (Stiemke [1915]).)

2.17. Prove that there exists a vector _x_ [^6]:= 0 satisfying _x_ ≥ 0 and _Ax_ = 0 if and only if
there is no vector _y_ satisfying _yTA >_ 0. (Gordan’s theorem (Gordan [1873]).)

2.18. Prove that there exists a vector _x_ satisfying _Ax < b_ if and only if _y_ = 0 is the only
solution for _y_ ≥ [^0]:_,yTA_ = 0_,yTb_ ≤ 0.

2.19. Prove that there exists a vector _x_ satisfying _Ax < b_ and _A_′_x_ ≤ _b_′ if and only if for all
vectors _y,y_′ ≥ 0 one has:

```
(i) if yTA + y′TA′ = 0 then yTb + y′Tb′ ≥ 0, and
(ii) if yTA + y′TA′ = 0 and y 6= 0 then yTb + y′Tb′ > 0.
```

```
(Motzkin’s theorem (Motzkin [1936]).)
```

2.20. Let _A_ be an _m_ × _n_ matrix and let _b_ ∈ R_m_, with _m_ ≥ _n_ + 1. Suppose that _Ax_ ≤ _b_
has no solution _x_. Prove that there exist indices _i_[^0]:_,...,in_ so that the system _ai_[^0]:_x_ ≤
_bi_[^0]:_,...,ainx_ ≤ _bin_ has no solution _x_. Here _ai_ is the _i_th row of _A_ and _bi_ is the _i_th
component of _b_.
(_Hint:_ Combine Farkas’ lemma with Carath ́eodory’s theorem.)

### 2.4. Linear programming

One of the standard forms of a linear programming (LP) problem is:

(42) maximize _cTx_,
subject to _Ax_ ≤ _b_.

So linear programming can be considered as maximizing a ‘linear function’ _cTx_ over
a polyhedron _P_ = {_x_ | _Ax_ ≤ _b_}. Geometrically, this can be seen as shifting a
hyperplane to its ‘highest’ level, under the condition that it intersects _P_.
Problem (42) corresponds to determining the following maximum:

(43) max{_cTx_ | _Ax_ ≤ _b_}_._

This is the form in which we will denote an LP-problem.
If _P_ = {_x_ | _Ax_ ≤ _b_} is a nonempty polytope, then it is clear that max{_cTx_ | _Ax_ ≤
_b_} is attained by a _vertex_ of _P_ (cf. Exercise 2.21).
Clearly, also any _minimization_ problem can be transformed to form (43), since

(44) min{_cTx_ | _Ax_ ≤ _b_} = −max{−_cTx_ | _Ax_ ≤ _b_}_._

One says that _x_ is a _feasible solution_ of (43) if _x_ satisfies _Ax_ ≤ _b_. If _x_ moreover
attains the maximum, _x_ is called an _optimum solution_.
The famous method to solve linear programming problems is the _simplex method_,
designed by Dantzig [1951b]. The first polynomial-time method for LP-problems is
due to Khachiyan [1979,1980], based on the _ellipsoid method_. In 1984, Karmarkar
[1984] published another polynomial-time method for linear programming, the _inte-_
_rior point method_, which turns out to be competitive in practice with the simplex
method.

The Duality theorem of linear programming, due to von Neumann [1947], states
that if the maximum (43) is finite, then the maximum value is equal to the minimum
value of another, so-called _dual_ LP-problem:

(45) min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}_._

In order to show this, we first prove:

**Lemma 2.1.** _Let P be a polyhedron in_ R_n and let c_ ∈ R_n. If_ sup{_cTx_ | _x_ ∈ _P_} _is_
_finite, then_ max{_cTx_ | _x_ ∈ _P_} _is attained._

**Proof.** Let _δ_ := sup{_cTx_ | _x_ ∈ _P_}. Choose matrix _A_ and vector _b_ so that _P_ = {_x_ |
_Ax_ ≤ _b_}. We must show that there exists an _x_ ∈ R_n_ such that _Ax_ ≤ _b_ and _cTx_ ≥ _δ_.
Suppose that such an _x_ does not exist. Then by Farkas’ lemma, in the form of
Corollary 2.5a, there exists a vector _y_ ≥ 0 and a real number _λ_ ≥ 0 such that:

(46) _yTA_ − _λcT_ = 0, _yTb_ − _λδ <_ [^0]:_._

This gives

(47) _λδ_ = _λ_sup{_cTx_ | _Ax_ ≤ _b_} = sup{_λcTx_ | _Ax_ ≤ _b_} = sup{_yTAx_ | _Ax_ ≤ _b_}≤
_yTb < λδ,_

a contradiction.

```
From this we derive:
```

**Theorem 2.6** (Duality theorem of linear programming)**.** _Let A be an m_ × _n matrix,_
_b_ ∈ R_m, c_ ∈ R_n. Then_

(48) max{_cTx_ | _Ax_ ≤ _b_} = min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}_,_

_provided that both sets are nonempty._

**Proof.** First note that

(49) sup{_cTx_ | _Ax_ ≤ _b_}≤ inf{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}_,_

because if _Ax_ ≤ _b, y_ ≥ [^0]:_, yTA_ = _cT_, then

(50) _cTx_ = (_yTA_)_x_ = _yT_(_Ax_) ≤ _yTb._

As both sets are nonempty,the supremum and the infimum are finite. By Lemma 2.1
it suffices to show that we have equality in (49).
Let _δ_ := sup{_cTx_ | _Ax_ ≤ _b_}. Hence:

(51) if _Ax_ ≤ _b_ then _cTx_ ≤ _δ._

So by Corollary 2.5b there exists a vector _y_ such that

(52) _y_ ≥ [^0]:_, yTA_ = _cT, yTb_ ≤ _δ._

This implies that the infimum in (49) is at most _δ_.

```
The Duality theorem can be interpreted geometrically as follows. Let
```

(53) max{_cTx_ | _Ax_ ≤ _b_} =: _δ_

be attained at a point _x_∗. Without loss of generality we may assume that the first _k_
rows of _A_ belong to the matrix _Ax_∗. So _a_[^1]:_x_ ≤ _b_[^1]:_,... , akx_ ≤ _bk_ are those inequalities
in _Ax_ ≤ _b_ for which _aix_∗ = _bi_ holds. Elementary geometric insight (cf. Figure
2.1) gives that _cTx_ = _δ_ must be a nonnegative linear combination of the equations
_a_[^1]:_x_ = _b_[^1]:_,... , akx_ = _bk_.
That is, there exist _λ_[^1]:_,... , λk_ ≥ 0 such that:

(54) _λ_[^1]:_a_[^1]: + ··· + _λkak_ = _cT,_
_λ_[^1]:_b_[^1]: + ··· + _λkbk_ = _δ._

Define

(55) _y_∗ := (_λ_[^1]:_,... , λk,_[^0]:_,... ,_0)_T._

Then _y_∗ is a feasible solution for the dual problem min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}.
Therefore,

```
c
```

```
a x=b
2
2
```

```
c x=T
δ
```

```
a2
a1
```

```
P
```

```
x*
a x=b
1
1
```

```
Figure 2.1
```

(56) max{_cTx_ | _Ax_ ≤ _b_} = _δ_ = _λ_[^1]:_b_[^1]: + ··· + _λkbk_ ≥ min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}_._

Since trivially the converse inequality holds:

(57) max{_cTx_ | _Ax_ ≤ _b_}≤ min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}

(cf. (50)), _y_∗ is an optimum solution of the dual problem.

```
There exist several variants of the Duality theorem.
```

**Corollary 2.6a.** _Let A be an m_ × _n matrix, b_ ∈ R_m, c_ ∈ R_n. Then_

(58) max{_cTx_ | _x_ ≥ 0;_Ax_ = _b_} = min{_yTb_ | _yTA_ ≥ _cT_}_,_

_provided that both sets are nonempty._

**Proof.** Define

(59) _A_ ̃ :=

```


```

```
A
−A
−I
```

```

```

```
, ̃b :=
```

```


```

```
b
−b
0
```

```

```

```
.
```

Then

(60) max{_cTx_ | _x_ ≥ 0;_Ax_ = _b_} = max{_cTx_ | _Ax_ ̃ ≤ ̃_b_} =
min{_zT_ ̃_b_ | _z_ ≥ 0;_zTA_ ̃ = _cT_} =
min{_uTb_ − _vTb_ + _wT_[^0]: | _u, v, w_ ≥ 0;_uTA_ − _vTA_ − _wT_ = _cT_} =
min{_yTb_ | _yTA_ ≥ _cT_}_._

The last equality follows by taking _y_ := _u_ − _v_.

**Exercises**

2.21. Let _P_ = {_x_ | _Ax_ ≤ _b_} be a nonempty polytope. Prove that max{_cTx_ | _Ax_ ≤ _b_} is
attained by a vertex of _P_.

2.22. Let _P_ = {_x_ | _Ax_ ≤ _b_} be a (not necessarily bounded) polyhedron, such that _P_ has at
least one vertex. Prove that if max{_cTx_ | _Ax_ ≤ _b_} is finite, it is attained by a vertex
of _P_.

2.23. Prove the following variant of the Duality theorem:

```
(61) max{cTx | x ≥ 0;Ax ≤ b} = min{yTb | y ≥ 0;yTA ≥ cT}
```

```
(assuming both sets are nonempty).
```

2.24. Prove the following variant of the Duality theorem:

```
(62) max{cTx | Ax ≥ b} = min{yTb | y ≤ 0;yTA = cT}
```

```
(assuming both sets are nonempty).
```

2.25. Let a matrix, a column vector, and a row vector be given:

```
(63)
```

```


```

```
A B C
D E F
G H K
```

```

```

```
,
```

```


```

```
a
b
c
```

```

```

```
,(d e f),
```

```
where A,B,C,D,E,F,G,H,K are matrices, a,b,c are column vectors, and d,e,f are
row vectors (of appropriate dimensions). Then
```

```
(64) max{dx + ey + fz | x ≥ 0;z ≤ 0;
Ax + By + Cz ≤ a;
Dx + Ey + Fz = b;
Gx + Hy + Kz ≥ c}
= min{ua + vb + wc | u ≥ 0;w ≤ 0;
uA + vD + wG ≥ d;
uB + vE + wH = e;
uC + vF + wK ≤ f},
```

```
assuming that both sets are nonempty.
```

2.26. Give an example of a matrix _A_ and vectors _b_ and _c_ for which both {_x_ | _Ax_ ≤ _b_} and
{_y_ | _y_ ≥ 0;_yTA_ = _cT_} are empty.

2.27. Let ̃_x_ be a feasible solution of max{_cTx_ | _Ax_ ≤ _b_} and let ̃_y_ be a feasible solution
of min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}. Prove that ̃_x_ and ̃_y_ are optimum solutions of the
maximum and minimum, respectively if and only if for each _i_ = 1_,...,m_ one has:
_y_ ̃_i_ = 0 or _aix_ ̃ = _bi_.
(Here _A_ has _m_ rows and _ai_ denotes the _i_th row of _A_.)

2.28. Let _A_ be an _m_ × _n_ matrix and let _b_ ∈ R_m_. Let {_x_ | _Ax_ ≤ _b_} be nonempty and let
_C_ be the convex cone {_x_ | _Ax_ ≤ [^0]:}. Prove that the set of all vectors _c_ for which
max{_cTx_ | _Ax_ ≤ _b_} is finite, is equal to _C_∗.

## 3. Matchings and covers in

## bipartite graphs

### 3.1. Matchings, covers, and Gallai’s theorem

Let _G_ = (_V, E_) be a graph. A _stable set_ is a subset _C_ of _V_ such that _e_ 6⊆ _C_ for each
edge _e_ of _G_. A _vertex cover_ is a subset _W_ of _V_ such that _e_ ∩ _W_ [^6]:= ∅ for each edge _e_
of _G_. It is not difficult to show that for each _U_ ⊆ _V_ :

(1) _U_ is a stable set ⇐⇒ _V_ \ _U_ is a vertex cover.

A _matching_ is a subset _M_ of _E_ such that _e_ ∩ _e_′ = ∅ for all _e, e_′ ∈ _M_ with _e_ [^6]:= _e_′.
A matching is called _perfect_ if it covers all vertices (that is, has size[^1]
[^2]:|_V_ |). An _edge_
_cover_ is a subset _F_ of _E_ such that for each vertex _v_ there exists _e_ ∈ _F_ satisfying
_v_ ∈ _e_. Note that an edge cover can exist only if _G_ has no isolated vertices.
Define:

(2) _α_(_G_) := max{|_C_|| _C_ is a stable set},
_τ_(_G_) := min{|_W_|| _W_ is a vertex cover},
_ν_(_G_) := max{|_M_|| _M_ is a matching},
_ρ_(_G_) := min{|_F_|| _F_ is an edge cover}.

These numbers are called the _stable set number_, the _vertex cover number_, the _matching_
_number_, and the _edge cover number_ of _G_, respectively.
It is not difficult to show that:

(3) _α_(_G_) ≤ _ρ_(_G_) and _ν_(_G_) ≤ _τ_(_G_).

The triangle _K_[^3]: shows that strict inequalities are possible. In fact, equality in one of
the relations (3) implies equality in the other, as Gallai [1958,1959] proved:

**Theorem 3.1** (Gallai’s theorem)**.** _For any graph G_ = (_V, E_) _without isolated vertices_
_one has_

(4) _α_(_G_) + _τ_(_G_) = |_V_ | = _ν_(_G_) + _ρ_(_G_)_._

**Proof.** The first equality follows directly from (1).
To see the second equality, first let _M_ be a matching of size _ν_(_G_). For each of the
|_V_ |−[^2]:|_M_| vertices _v_ missed by _M_, add to _M_ an edge covering _v_. We obtain an edge
cover of size |_M_| + (|_V_ |− [^2]:|_M_|) = |_V_ |−|_M_|. Hence _ρ_(_G_) ≤|_V_ |− _ν_(_G_).

```
Second, let F be an edge cover of size ρ(G). For each v ∈ V delete from F, dF(v)−1
edges incident with v. We obtain a matching of size at least |F|−
```

```
∑
v∈V (dF(v)−1) =
|F|− (2|F|−|V |) = |V |−|F|. Hence ν(G) ≥|V |− ρ(G).
```

```
This proof also shows that if we have a matching of maximum cardinality in any
graph G, then we can derive from it a minimum cardinality edge cover, and conversely.
```

```
Exercises
```

```
3.1. Let G = (V,E) be a graph without isolated vertices. Define:
```

```
(5) α2(G) := the maximum number of vertices such that no edge
contains more than two of these vertices;
ρ2(G) := the minimum number of edges such that each vertex
is contained in at least two of these edges;
τ2(G) := the minimum number of vertices such that each edge
contains at least two of these vertices
ν2(G) := the maximum number of edges such that no vertex is
contained in more than two of these edges;
```

```
possibly taking vertices (edges, respectively) more than once.
```

```
(i) Show that α2(G) ≤ ρ2(G) and that ν2(G) ≤ τ2(G).
(ii) Show that α2(G) + τ2(G) = 2|V |.
(iii) Show that ν2(G) + ρ2(G) = 2|V |.
```

### 3.2. _M_-augmenting paths

```
Basic in matching theory are M-augmenting paths, which are defined as follows. Let
M be a matching in a graph G = (V, E). A path P = (v0, v1,... , vt) in G is called
M-augmenting if
```

(6) (i) _t_ is odd,

```
(ii) v1v2, v3v4,... , vt−2vt−1 ∈ M,
(iii) v0, vt 6∈
```

```
⋃
M.
```

```
Note that this implies that v0v1, v2v3,... , vt−1vt do not belong to M.
```

```
Clearly, if P = (v0, v1,... , vt) is an M-augmenting path, then
```

```
(7) M′ := M△EP
```

```
edge in M
edge not in M vertex not covered by M
```

```
vertex covered by M
```

```
Figure 3.1
```

is a matching satisfying |_M_′| = |_M_| + 1.[^8]
In fact, it is not difficult to show that:

**Theorem 3.2.** _Let G_ = (_V, E_) _be a graph and let M be a matching in G. Then_
_either M is a matching of maximum cardinality, or there exists an M-augmenting_
_path._

**Proof.** If _M_ is a maximum-cardinality matching, there cannot exist an _M_-augmenting
path _P_, since otherwise _M_△_EP_ would be a larger matching.
If _M_′ is a matching larger than _M_, consider the components of the graph _G_′ :=
(_V, M_ ∪ _M_′). As _G_′ has maximum valency two, each component of _G_′ is either a
path (possibly of length 0) or a circuit. Since |_M_′| _>_ |_M_|, at least one of these
components should contain more edges of _M_′ than of _M_. Such a component forms
an _M_-augmenting path.

### 3.3. K ̋onig’s theorems

A classical min-max relation due to K ̋onig [1931] (extending a result of Frobenius
[1917]) characterizes the maximum size of a matching in a bipartite graph (we follow
the proof of De Caen [1988]):

**Theorem 3.3** (K ̋onig’s matching theorem)**.** _For any bipartite graph G_ = (_V, E_) _one_
_has_

(8) _ν_(_G_) = _τ_(_G_)_._

_That is, the maximum cardinality of a matching in a bipartite graph is equal to the_
_minimum cardinality of a vertex cover._

**Proof.** By (3) it suffices to show that _ν_(_G_) ≥ _τ_(_G_). We may assume that _G_ has at
least one edge. Then:

(9) _G_ has a vertex _u_ covered by each maximum-size matching.

```
8EP denotes the set of edges in P. △ denotes symmetric difference.
```

To see this, let _e_ = _uv_ be any edge of _G_, and suppose that there are maximum-size
matchings _M_ and _N_ missing _u_ and _v_ respectively[^9]. Let _P_ be the component of _M_ ∪_N_
containing _u_. So _P_ is a path with end vertex _u_. Since _P_ is not _M_-augmenting (as _M_
has maximum size), _P_ has even length, and hence does not traverse _v_ (otherwise, _P_
ends at _v_, contradicting the bipartiteness of _G_). So _P_∪_e_ would form an _N_-augmenting
path, a contradiction (as _N_ has maximum size). This proves (9).
Now (9) implies that for the graph _G_′ := _G_ − _u_ one has _ν_(_G_′) = _ν_(_G_) − 1.
Moreover, by induction, _G_′ has a vertex cover _C_ of size _ν_(_G_′). Then _C_ ∪{_u_} is a
vertex cover of _G_ of size _ν_(_G_′) + 1 = _ν_(_G_).

```
Combination of Theorems 3.1 and 3.3 yields the following result of K ̋onig [1932].
```

**Corollary 3.3a** (K ̋onig’s edge cover theorem)**.** _For any bipartite graph G_ = (_V, E_)_,_
_without isolated vertices, one has_

(10) _α_(_G_) = _ρ_(_G_)_._

_That is, the maximum cardinality of a stable set in a bipartite graph is equal to the_
_minimum cardinality of an edge cover._

**Proof.** Directly from Theorems 3.1 and 3.3, as _α_(_G_) = |_V_ |− _τ_(_G_) = |_V_ |− _ν_(_G_) =
_ρ_(_G_).

**Exercises**

```
3.2. (i) Prove that a k-regular bipartite graph has a perfect matching (if k ≥ 1).
(ii) Derive that a k-regular bipartite graph has k disjoint perfect matchings.
(iii) Give for each k > 1 an example of a k-regular graph not having a perfect
matching.
```

```
3.3. Prove that in a matrix, the maximum number of nonzero entries with no two in the
same line (=row or column), is equal to the minimum number of lines that include
all nonzero entries.
```

```
3.4. Let A = (A1,...,An) be a family of subsets of some finite set X. A subset Y of X is
called a transversal or a system of distinct representatives (SDR) of A if there exists
a bijection π : {1,...,n}→ Y such that π(i) ∈ Ai for each i = 1,...,n.
Decide if the following collections have an SDR:
```

```
(i) {3,4,5},{2,5,6},{1,2,5},{1,2,3},{1,3,6},
(ii) {1,2,3,4,5,6},{1,3,4},{1,4,7},{2,3,5,6},{3,4,7},{1,3,4,7},{1,3,7}.
9M misses a vertex u if u 6∈
```

```
⋃
M. Here
```

⋃
_M_ denotes the union of the edges in _M_; that is, the
set of vertices covered by the edges in _M_.

```
3.5. Let A = (A1,...,An) be a family of subsets of some finite set X. Prove that A has
an SDR if and only if
```

```
(11)
∣∣
```

```
⋃
```

```
i∈I
```

```
Ai
∣∣ ≥|I|
```

```
for each subset I of {1,...,n}.
[Hall’s ‘marriage’ theorem (Hall [1935]).]
```

```
3.6. Let A = (A1,...,An) be subsets of the finite set X. A subset Y of X is called a
partial transversal or a partial system of distinct representatives (partial SDR) if it is
a transversal of some subcollection (Ai1,...,Aik) of (A1,...,An).
Show that the maximum cardinality of a partial SDR of A is equal to the minimum
value of
```

```
(12) |X \ Z| + |{i | Ai ∩ Z 6= ∅}|,
where Z ranges over all subsets of X.
```

```
3.7. Let A = (A1,...,An) be a family of finite sets and let k be a natural number. Show
that A has k pairwise disjoint SDR’s of A if and only if
```

```
(13)
∣∣
```

```
⋃
```

```
i∈I
```

```
Ai
∣∣ ≥ k|I|
```

```
for each subset I of {1,...,n}.
```

```
3.8. Let A = (A1,...,An) be a family of subsets of a finite set X and let k be a natural
number. Show that X can be partitioned into k partial SDR’s if and only if
```

```
(14) k ·|{i | Ai ∩ Y 6= ∅}|≥|Y |
```

```
for each subset Y of X.
(Hint: Replace each Ai by k copies of Ai and use Exercise 3.6 above.)
```

```
3.9. Let (A1,...,An) and (B1,...,Bn) be two partitions of the finite set X.
```

```
(i) Show that (A1,...,An) and (B1,...,Bn) have a common SDR if and only if for
each subset I of {1,...,n}, the set
```

```
⋃
i∈I Ai intersects at least |I| sets among
B1,...,Bn.
(ii) Suppose that |A1| = ··· = |An| = |B1| = ··· = |Bn|. Show that the two
partitions have a common SDR.
```

3.10. Let (_A_[^1]:_,...,An_) and (_B_[^1]:_,...,Bn_) be two partitions of the finite set _X_. Show that the
minimum cardinality of a subset of _X_ intersecting each set among _A_[^1]:_,...,An,B_[^1]:_,...,_
_Bn_ is equal to the maximum number of pairwise disjoint sets in _A_[^1]:_,...,An,B_[^1]:_,...,Bn_.

3.11. A matrix is called _doubly stochastic_ if it is nonnegative and each row sum and each
column sum is equal to 1. A matrix is called a _permutation matrix_ if each entry is 0
or 1 and each row and each column contains exactly one 1.

```
(i) Show that for each doubly stochastic matrix A = (ai,j)n
i,j=1 there exists a per-
mutation π ∈ Sn such that ai,π(i) 6= 0 for all i = 1,...,n.
(ii) Derive that each doubly stochastic matrix is a convex linear combination of
permutation matrices.
```

```
[Birkhoff-von Neumann theorem (Birkhoff [1946], von Neumann [1953]).]
```

3.12. Let _G_ = (_V,E_) be a bipartite graph with colour classes _U_ and _W_. Let _b_ : _V_ → Z+
be so that

```
∑
v∈U b(v) =
```

```
∑
v∈W b(v) =: t.
A b-matching is a function c : E → Z+ so that for each vertex v of G:
```

```
(15)
```

```
∑
```

```
e∈E,v∈e
```

```
c(e) = b(v)
```

```
Show that there exists a b-matching if and only if
```

```
(16)
```

```
∑
```

```
v∈X
```

```
b(v) ≥ t
```

```
for each vertex cover X.
```

3.13. Let _G_ = (_V,E_) be a bipartite graph and let _b_ : _V_ → Z+. Show that _G_ has a subgraph
_G_′ = (_V,E_′) such that deg_G_′(_v_) = _b_(_v_) for each _v_ ∈ _V_ if and only if each _X_ ⊆ _V_
contains at least

```
(17)
```

```
1
2
```

```
(
```

```
∑
```

```
v∈X
```

```
b(v) −
```

```
∑
```

```
v∈V\X
```

```
b(v))
```

```
edges.
```

3.14. Let _G_ = (_V,E_) be a bipartite graph and let _b_ : _V_ → Z+. Show that the maximum
number of edges in a subset _F_ of _E_ so that each vertex _v_ of _G_ is incident with at
most _b_(_v_) of the edges in _F_, is equal to

```
(18) min
X⊆V
```

```
∑
```

```
v∈X
```

```
b(v) + |E(V \ X)|.
```

3.15. Let _G_ = (_V,E_) be a bipartite graph and let _k_ ∈ N. Prove that _G_ has _k_ disjoint
perfect matchings if and only if each _X_ ⊆ _V_ contains at least _k_(|_X_|−[^1]
[^2]:|_V_ |) edges.
3.16. Show that each 2_k_-regular graph contains a set _F_ of edges so that each vertex is
incident with exactly two edges in _F_.

### 3.4. Cardinality bipartite matching algorithm

We now focus on the problem of finding a maximum-sized matching in a bipartite
graph algorithmically.
In any graph, if we have an algorithm finding an _M_-augmenting path for any
matching _M_ (if it exists), then we can find a maximum cardinality matching: we
iteratively find matchings _M_[^0]:_, M_[^1]:_,..._, with |_Mi_| = _i_, until we have a matching _Mk_
such that there does not exist any _Mk_-augmenting path.
We now describe how to find an _M_-augmenting path in a bipartite graph.

**Matching augmenting algorithm for bipartite graphs**

**input:** a bipartite graph _G_ = (_V, E_) and a matching _M_,
**output:** a matching _M_′ satisfying |_M_′| _>_ |_M_| (if there is one).
**description of the algorithm:** Let _G_ have colour classes _U_ and _W_. Orient each
edge _e_ = {_u, w_} of _G_ (with _u_ ∈ _U, w_ ∈ _W_) as follows:

(19) if _e_ ∈ _M_ then orient _e_ from _w_ to _u_,
if _e_ 6∈ _M_ then orient _e_ from _u_ to _w_.

Let _D_ be the directed graph thus arising. Consider the sets

(20) _U_′ := _U_ \

```
⋃
M and W′ := W \
```

```
⋃
M.
```

Now an _M_-augmenting path (if it exists) can be found by finding a directed path
in _D_ from any vertex in _U_′ to any vertex in _W_′. Hence in this way we can find a
matching larger than _M_.

```
This implies:
```

**Theorem 3.4.** _A maximum-size matching in a bipartite graph can be found in time_
_O_(|_V_ ||_E_|)_._

**Proof.** The correctness of the algorithm is immediate. Since a directed path can
be found in time _O_(|_E_|), we can find an augmenting path in time _O_(|_E_|). Hence a
maximum cardinality matching in a bipartite graph can be found in time _O_(|_V_ ||_E_|)
(as we do at most |_V_ | iterations).

```
Hopcroft and Karp [1973] gave an O(|V |[^1]/[^2]|E|) algorithm.
```

**Application 3.1: Assignment problem.** Suppose we have _k_ machines at our disposal:
_m_[^1]:_,...,mk_. On a certain day we have to carry out _n_ jobs: _j_[^1]:_,...,jn_. Each machines
is capable of performing some jobs, but can do only one job a day. E.g., we could have

five machines _m_[^1]:_,...,m_[^5]: and five jobs _j_[^1]:_,...,j_[^5]: and the capabilities of the machines are
indicated by crosses in the following table:

```
j1 j2 j3 j4 j5
m1 X X X
m2 X X X X
m3 X X
m4 X
m5 X
```

We want to assign the machines to the jobs in such a way that every machine performs
at most one job and that a largest number of jobs is carried out.
In order to solve this problem we represent the machines and jobs by vertices _m_[^1]:_,...,mk_
and _j_[^1]:_,...,jn_ of a bipartite graph _G_ = (_V,E_), and we make an edge from _mi_ to _jj_ if job _j_
can be performed by machine _i_. Thus the example gives Figure 3.2. Then a maximum-size
matching in _G_ corresponds to a maximum assignment of jobs.

```
3
```

```
4
```

```
2
```

```
m1
```

```
m
```

```
m3
```

```
m4
```

```
m5 j5
```

```
j
```

```
j
```

```
j2
```

```
j1
```

```
Figure 3.2
```

**Exercises**

3.17. Find a maximum-size matching and a minimum vertex cover in the bipartite graph
in Figure 3.3.

3.18. Solve the assignment problem given in Application 3.1.

3.19. Derive K ̋onig’s matching theorem from the cardinality matching algorithm for bipar-
tite graphs.

3.20. Show that a minimum-size vertex cover in a bipartite graph can be found in polyno-
mial time.

3.21. Show that, given a family of sets, a system of distinct representatives can be found
in polynomial time (if it exists).

```
1
```

```
a b c d e f
```

```
2 3 4 5
```

```
g h i j
```

```
6 7 8 9 10
Figure 3.3
```

### 3.5. Weighted bipartite matching

We now consider the problem of finding a matching of maximum weight for which
we describe the so-called _Hungarian method_ developed by Kuhn [1955], using work of
Egerv ́ary [1931] (see Corollary 3.7b below).
Let _G_ = (_V, E_) be a graph and let _w_ : _E_ → R be a ‘weight’ function. For any
subset _M_ of _E_ define the _weight w_(_M_) of _M_ by

(21) _w_(_M_) :=

```
∑
```

```
e∈M
```

```
w(e).
```

The maximum-weight matching problem consists of finding a matching of maximum
weight.
Again, augmenting paths are of help at this problem. Call a matching _M extreme_
if it has maximum weight among all matchings of cardinality |_M_|.
Let _M_ be an extreme matching. Define a ‘length’ function _l_ : _E_ → R as follows:

(22) _l_(_e_) :=

```
{
w(e) if e ∈ M,
−w(e) if e 6∈ M.
```

Then the following holds:

**Proposition 1.** _Let P be an M-augmenting path of minimum length. If M is_
_extreme, then M_′ := _M_△_EP is extreme again._

**Proof.** Let _N_ be any extreme matching of size |_M_| + 1. As |_N_| _>_ |_M_|, _M_ ∪ _N_ has
a component _Q_ that is an _M_-augmenting path. As _P_ is a shortest _M_-augmenting
path, we know _l_(_Q_) ≥ _l_(_P_). Moreover, as _N_△_EQ_ is a matching of size |_M_|, and as
_M_ is extreme, we know _w_(_N_△_EQ_) ≤ _w_(_M_). Hence

(23) _w_(_N_) = _w_(_N_△_EQ_) − _l_(_Q_) ≤ _w_(_M_) − _l_(_P_) = _w_(_M_′)_._

Hence _M_′ is extreme.

This implies that if we are able to find a minimum-length _M_-augmenting path in
polynomial time, we can find a maximum-weight matching in polynomial time: find
iteratively extreme matchings _M_[^0]:_, M_[^1]:_,..._ such that |_Mk_| = _k_ for each _k_. Then the
matching among _M_[^0]:_, M_[^1]:_,..._ of maximum weight is a maximum-weight matching.
If _G_ is bipartite, we can find a minimum-length _M_-augmenting path as follows. Let
_G_ have colour classes _U_ and _W_. Orient the edges of _G_ as in (19), making the directed
graph _D_, and let _U_′ and _W_′ as in (20). Then a minimum-length _M_-augmenting path
can be found by finding a minimum-length path in _D_ from any vertex in _U_′ to any
vertex in _W_′. This can be done in polynomial time, since:

**Theorem 3.5.** _Let M be an extreme matching. Then D has no directed circuit of_
_negative length._

**Proof.** Suppose _C_ is a directed circuit in _D_ with length _l_(_C_) _<_ 0. We may assume
_C_ = (_u_[^0]:_, w_[^1]:_, u_[^1]:_,... , wt, ut_) with _u_[^0]: = _ut_ and _u_[^1]:_,... , ut_ ∈ _U_ and _w_[^1]:_,... , wt_ ∈ _W_.
Then the edges _w_[^1]:_u_[^1]:_,... , wtut_ belong to _M_ and the edges _u_[^0]:_w_[^1]:_, u_[^1]:_w_[^2]:_,... , ut_−[^1]:_wt_ do
not belong to _M_. Then _M_′′ := _M_△_EC_ is a matching of cardinality _k_ of weight
_w_(_M_′′) = _w_(_M_) − _l_(_C_) _> w_(_M_), contradicting the fact that _M_ is extreme.

This gives a polynomial-time algorithm to find a maximum-weight matching in a
bipartite graph. The description above yields:

**Theorem 3.6.** _A maximum-weight matching in a bipartite graph G_ = (_V, E_) _can be_
_found in O_(|_V_ |[^2]|_E_|) _time._

**Proof.** We do _O_(|_V_ |) iterations, each consisting of finding a shortest path (in a graph
without negative-length directed circuits), which can be done in _O_(|_V_ ||_E_|) time (with
the Bellman-Ford algorithm — see Corollary 1.10a).

In fact, a sharpening of this method (by transmitting a ‘potential’ _p_ : _V_ → Q
throughout the matching augmenting iterations, making the length function _l_ non-
negative, so that Dijkstra’s method can be used) gives an _O_(|_V_ |(|_E_| + |_V_ |log |_V_ |))
algorithm.

**Application 3.2: Optimal assignment.** Suppose that we have _n_ jobs and _m_ machines
and that each job can be done on each machine. Moreover, let a cost function (or cost
matrix) _ki,j_ be given, specifying the cost of performing job _j_ by machine _i_. We want to
perform the jobs with a minimum of total costs.
This can be solved with the maximum-weight bipartite matching algorithm. To this
end, we make a complete bipartite graph _G_ with colour classes of cardinality _m_ and _n_. Let
_K_ be the maximum of _ki,j_ over all _i,j_. Define the weight of the edge connecting machine _i_
and job _j_ to be equal to _K_ − _ki,j_. Then a maximum-weight matching in _G_ corresponds to

an optimum assignment of machines to jobs.
So the algorithm for solving the assignment problem counters the remarks made by
Thorndike [1950] in an Address delivered on September 9, 1949 at a meeting of the American
Psychological Association at Denver, Colorado:

```
There are, as has been indicated, a finite number of permutations in the assign-
ment of men to jobs. When the classification problem as formulated above was
presented to a mathematician, he pointed to this fact and said that from the
point of view of the mathematician there was no problem. Since the number of
permutations was finite, one had only to try them all and choose the best. He
dismissed the problem at that point. This is rather cold comfort to the psy-
chologist, however, when one considers that only ten men and ten jobs mean
over three and a half million permutations. Trying out all the permutations
may be a mathematical solution to the problem, it is not a practical solution.
```

**Application 3.3: Transporting earth.** Monge [1784] was one of the first to consider
the assignment problem, in the role of the problem of transporting earth from one area to
another, which he considered as the discontinuous, combinatorial problem of transporting
molecules:

```
Lorsqu’on doit transporter des terres d’un lieu dans un autre, on a coutime de
donner le nom de D ́eblai au volume des terres que l’on doit transporter, & le
nom de Remblai `a l’espace qu’elles doivent occuper apr`es le transport.
Le prix du transport d’une mol ́ecule ́etant, toutes choses d’ailleurs ́egales, pro-
portionnel `a son poids & `a l’espace qu’on lui fait parcourir, & par cons ́equent le
prix du transport total devant ˆetre proportionnel `a la somme des produits des
mol ́ecules multipli ́ees chacune par l’espace parcouru, il s’ensuit que le d ́eblai &
le remblai ́etant donn ́e de figure & de position, il n’est pas indiff ́erent que telle
mol ́ecule du d ́eblai soit transport ́ee dans tel ou tel autre endroit du remblai,
mais qu’il y a une certaine distribution `a faire des mol ́ecules du premier dans
le second, dapr`es laquelle la somme de ces produits sera la moindre possible, &
le prix du transport total sera minimum.[^10]
```

Monge describes an interesting geometric method to solve the assignment problem in this
case: let _l_ be a line touching the two areas from one side; then transport the earth molecule

[^10]:When one must transport earth from one place to another, one usually gives the name of _D ́eblai_
to the volume of earth that one must transport, & the name of _Remblai_ to the space that they
should occupy after the transport.
The price of the transport of one molecule being, if all the rest is equal, proportional to its weight
& to the distance that one makes it covering, & hence the price of the total transport having to be
proportional to the sum of the products of the molecules each multiplied by the distance covered,
it follows that, the d ́eblai & the remblai being given by figure and position, it makes difference if a
certain molecule of the d ́eblai is transported to one or to another place of the remblai, but that there
is a certain distribution to make of the molcules from the first to the second, after which the sum of
these products will be as little as possible, & the price of the total transport will be a _minimum_.

touched in one area to the position touched in the other area. Then repeat, until all
molecules are transported.

**Exercises**

3.22. Five mechanics, stationed in the cities _A,B,C,D,E_, have to perform jobs in the cities
_F,G,H,I,J_. The jobs must be assigned in such a way to the mechanics that everyone
gets one job and that the total distance traveled by them is as small as possible. The
distances are given in the tables below. Solve these assignment problems with the
weighted matching algorithm.

```
(i)
```

```
F G H I J
A 6 17 10 1 3
B 9 23 21 4 5
C 2 8 5 0 1
D 19 31 19 20 9
E 21 25 22 3 9
```

```
(ii)
```

```
F G H I J
A 11 5 21 7 18
B 17 4 20 9 25
C 4 1 3 2 4
D 6 2 19 3 9
E 19 7 23 18 26
```

3.23. Derive from the weighted matching algorithm for bipartite graphs an algorithm for
finding a minimum-weight perfect matching in a bipartite graph _G_ = (_V,E_). (A
matching _M_ is _perfect_ if

```
⋃
M = V .)
```

3.24. Let _A_[^1]:_,...,An_ be subsets of the finite set _X_ and let _w_ : _X_ → R+ be a ‘weight’
function. Derive from the weighted matching algorithm a polynomial-time algorithm
to find a minimum-weight SDR.

### 3.6. The matching polytope

The weighted matching problem is related to the ‘matching polytope’. Let _G_ = (_V, E_)
be a graph. For each matching _M_ let the _incidence vector χM_ : _E_ → R of _M_ be
defined by:

(24) _χM_(_e_) := 1 if _e_ ∈ _M_,
_χM_(_e_) := 0 if _e_ 6∈ _M_,

for _e_ ∈ _E_.
It is important to realize that the set of functions _f_ : _E_ → R can be considered
as a vector space and each such function as a vector. Thus we can denote _f_(_e_) by _fe_.
The function _χM_ can be considered alternatively as a vector in the vector space R_E_.
Similarly for functions _g_ : _V_ → R.

```
The matching polytope of G is defined as:
```

(25) _P_matching(_G_) :=conv.hull{_χM_ | _M_ is a matching in _G_}.

So _P_matching(_G_) is a polytope in R_E_.
The matching polytope is a polyhedron, and hence can be described by linear
inequalities. For bipartite graphs, these inequalities are quite simple. To this end
it is convenient first to consider⋃ _perfect_ matchings. (A matching _M_ is _perfect_ if
_M_ = _V_ .) Now the _perfect matching polytope_ of _G_ is defined by:

(26) _P_perfect matching(_G_) :=conv.hull{_χM_ | _M_ is a perfect matching in _G_}.

Again, _P_perfect matching(_G_) is a polytope in R_E_. Now the following can be derived quite
directly from Exercise 3.11:

**Theorem 3.7.** _Let G_ = (_V, E_) _be a bipartite graph. Then the perfect matching_
_polytope Pperfect matching_(_G_) _is equal to the set of vectors x_ ∈ R_E satisfying:_

(27) ∑_xe_ ≥ 0 for each _e_ ∈ _E;_

```
e∋v
```

```
xe = 1 for each v ∈ V.
```

**Proof.** Left to the reader (Exercise 3.25).

Clearly, each vector _x_ in _P_perfect matching(_G_) should satisfy (27), since each vector
_χM_ satisfies (27). The essence of the theorem is that the inequalities (27) are enough
to define the polytope _P_perfect matching(_G_).
(An alternative way of proving Theorem 3.7 is using the ‘total unimodularity’ of
the incidence matrix of a bipartite graph, together with the Hoffman-Kruskal theorem
on integer solutions to linear programming problems with integer data and totally
unimodular constraint matrix — see Section 8.3.)
From Theorem 3.7 one can derive the linear inequalities describing the matching
polytope of a bipartite graph:

**Corollary 3.7a.** _Let G_ = (_V, E_) _be a bipartite graph. Then the matching polytope_
_Pmatching_(_G_) _is equal to the set of vectors x_ ∈ R_E satisfying:_

(28) ∑_xe_ ≥ 0 for each _e_ ∈ _E;_

```
e∋v
```

```
xe ≤ 1 for each v ∈ V.
```

**Proof.** Left to the reader (Exercise 3.26).

Clearly, one cannot delete the bipartiteness condition: if _G_ is the triangle _K_[^3]: then
the function _x_ defined by _xe_ := 1_/_2 for each edge _e_ satisfies (28), but does not belong
to the matching polytope.
Corollary 3.7a asserts that the weighted matching problem can be formulated as
a linear programming problem:

(29) maximize _wTx_,
subject to ∑_xe_ ≥ 0 for each _e_ ∈ _E_;

```
e∋v
```

```
xe ≤ 1 for each v ∈ V.
```

With linear programming duality one can derive from this a ‘weighted’ extension
of K ̋onig’s matching theorem, due to Egerv ́ary [1931]:

**Corollary 3.7b.** _Let G_ = (_V, E_) _be a bipartite graph and let w_ : _E_ → R _be a ‘weight’_
_function. Then the maximum weight of a matching is equal to th_∑ _e minimum value of_
_v_∈_V y_(_v_)_, where y ranges over all functions y_ : _V_ → R+ _satisfying y_(_u_)+_y_(_v_) ≥ _w_(_e_)
_for each edge e_ = _uv of G._

**Proof.** The maximum weight of a matching in _G_ is equal to

(30) max{_wTχM_ | _M_ is a matching in _G_}.

Since _P_matching(_G_) is the convex hull of the _χM_, (30) is equal to

(31) max{_wTx_ | _x_ ∈ _P_matching(_G_)}.

By Corollary 3.7a this is equal to

(32) max{_wTx_ | ∑ _xe_ ≥ 0 for each _e_ ∈ _E_;
_e_∋_v xe_ ≤ 1 for each _v_ ∈ _V_ }.

By linear programming duality this is equal to

(33) min{

```
∑
v∈V yv | yv ≥ 0 for each v ∈ V ;
yu + yv ≥ we for each edge e = uv}.
```

This is exactly the minimum described in the Corollary.

An extension of this corollary gives a further extension of K ̋onig’s matching the-
orem (Theorem 3.3):

**Theorem 3.8.** _In Corollary_ 3.7b_, if w is integer-valued, then we can take also y_
_integer-valued._

**Proof.** Let _y_ ∈ R_V_
+ attain the minimum, and assume that we have chosen _y_ so that
the number of vertices _v_ with _yv_ 6∈ Z is as small as possible. Let _U_ and _W_ be the two
colour classes of _G_ and let _X_ be the set of vertices _v_ of _G_ with _yv_ 6∈ Z. If _X_ = ∅ we
are done, so assume that _X_ [^6]:= ∅. Without loss of generality, |_X_ ∩ _U_|≥|_X_ ∩ _W_|. Let
_u_ be a vertex in _X_ ∩ _U_ with _yu_ −⌊_yu_⌋ as small as possible. Let _ε_ := _yu_ −⌊_yu_⌋. Reset

(34) _y_ ̃_v_ :=

```


```

```

```

```
yv − ε if v ∈ X ∩ U,
yv + ε if v ∈ X ∩ W,
yv if v 6∈ X.
```

One easily checks that again ̃_yv_ + ̃_yv_′ ≥ _w_(_e_) for each edge _e_ = _vv_′ of _G_ (using the fact
that _w_ is integer-valued). Moreover,

```
∑
v∈V y ̃v =
```

∑
∑ _v_∈_V yv_ − _ε_|_X_ ∩ _U_| + _ε_|_X_ ∩ _W_| ≤
_v_∈_V yv_. So ̃_y_ also attains the minimum. However, ̃_y_ has fewer noninteger-valued
components than _y_ (as ̃_yu_ ∈ Z), contradicting our assumption.

**Exercises**

3.25. Derive Theorem 3.7 from Exercise 3.11.

3.26. Derive Corollary 3.7a from Theorem 3.7.

## 4. Menger’s theorem, flows, and

## circulations

### 4.1. Menger’s theorem

In this section we study the maximum number _k_ of pairwise disjoint paths in a graph
connecting two vertices, or two sets of vertices. Here ‘disjoint’ can mean: _vertex-_
_disjoint_ (= having no vertex in common), or _internally vertex-disjoint_ (= having no
vertex in common except for their end vertices), or _arc-disjoint_ (= having no arc in
common). By the way, we often say ‘disjoint’ to mean ‘pairwise disjoint’.
Let _D_ = (_V, A_) be a directed graph and let _S_ and _T_ be subsets of _V_. A path is
called an _S_ − _T path_ if it runs from a vertex in _S_ to a vertex in _T_.
Menger [1927] gave a min-max theorem for the maximum number of disjoint _S_−_T_
paths. We follow the proof given by G ̈oring [2000].
A set _C_ of vertices is called _S_ − _T disconnecting_ if _C_ intersects each _S_ − _T_ path
(_C_ may intersect _S_ ∪ _T_).

**Theorem 4.1** (Menger’s theorem (directed vertex-disjoint version))**.** _Let D_ = (_V, A_)
_be a digraph and let S, T_ ⊆ _V. Then the maximum number of vertex-disjoint S_ − _T_
_paths is equal to the minimum size of an S_ − _T disconnecting vertex set._

**Proof.** Obviously, the maximum does not exceed the minimum. Equality is shown
by induction on |_A_|, the case _A_ = ∅ being trivial.
Let _k_ be the minimum size of an _S_ − _T_ disconnecting vertex set. Choose _a_ =
(_u, v_) ∈ _A_. If each _S_ − _T_ disconnecting vertex set in _D_ − _a_ has size at least _k_, then
inductively there exist _k_ vertex-disjoint _S_ − _T_ paths in _D_ − _a_, hence in _D_.
So we can assume that _D_ − _a_ has an _S_ − _T_ disconnecting vertex set _C_ of size
≤ _k_ − 1. Then _C_ ∪{_u_} and _C_ ∪{_v_} are _S_ − _T_ disconnecting vertex sets of _D_ of size
_k_. So |_C_| = _k_ − 1.
Now each _S_ −(_C_ ∪{_u_}) disconnecting vertex set _B_ of _D_ −_a_ has size at least _k_, as
it is _S_ − _T_ disconnecting in _D_. Indeed, each _S_ − _T_ path _P_ in _D_ intersects _C_ ∪{_u_},
and hence _P_ contains an _S_ − (_C_ ∪{_u_}) path in _D_ − _a_. So _P_ intersects _B_.
So by induction, _D_ −_a_ contains _k_ disjoint _S_ −(_C_ ∪{_u_}) paths. As |_C_ ∪{_u_}| = _k_,
this means that for each _c_ ∈ _C_ ∪{_u_} there exists an _S_ − _c_ path _Pc_ such that for
distinct _c, c_′ ∈ _C_ ∪{_u_}, _Pc_ and _Pc_′ are disjoint. So each _Pc_ intersects _C_ ∪{_u_} only in
_c_.
Similarly, there exists for each _c_ ∈ _U_ ∪{_v_} a _c_ − _T_ path such that for distinct
_c, c_′ ∈ _C_ ∪{_v_}, _Qc_ and _Qc_′ are disjoint. Again, each _Qc_ intersects _C_ ∪{_v_} only in _c_.
Note that for _c_ ∈ _C_ ∪{_u_} and _c_′ ∈ _C_ ∪{_v_}, if _Pc_ and _Qc_′ intersect in some vertex
_w_, then _c_ = _w_ = _c_′. For suppose not. By symmetry we can assume that _w_ [^6]:= _c_. Then

_w_ 6∈ _C_ ∪{_u_}, as otherwise _Pc_ and _Pw_ would intersect. Then the concatenation of the

_S_ − _w_ part of _Pc_ and the _w_ − _T_ part of _Qc_′ is a walk from _S_ to _T_ in _D_ − _a_ avoiding
_C_ (also in case _w_ = _v_). This contradicts the fact that _C_ is _S_ − _T_ disconnecting for
_D_ − _a_.
For each _c_ ∈ _C_, let _Rc_ be the _S_ − _T_ path obtained by concatenating _Pc_ and _Qc_.
Moreover, let _R_ be the _S_ − _T_ path obtained by concatenating _Pu_, arc _a_ = (_u, v_), and
_Qv_. By the above, the paths _Rc_ (_c_ ∈ _C_) and _R_ form _k_ pairwise disjoint _S_ − _T_ paths.

A consequence of this theorem is a variant on _internally vertex-disjoint s_−_t_ paths,
that is, _s_ − _t_ paths having no vertex in common except for _s_ and _t_. Recall that a set
_U_ of vertices is called an _s_ − _t vertex-cut_ if _s, t_ 6∈ _U_ and each _s_ − _t_ path intersects _U_.

**Corollary 4.1a** (Menger’s theorem (directed internally vertex-disjoint version))**.** _Let_
_D_ = (_V, A_) _be a digraph and let s and t be two distinct and nonadjacent vertices of_
_D. Then the maximum number of internally vertex-disjoint s_−_t paths is equal to the_
_minimum size of an s_ − _t vertex-cut._

**Proof.** Let _D_′ := _D_ − _s_ − _t_ and let _S_ and _T_ be the sets of outneighbours of _s_ and
of inneighbours of _t_, respectively. Then Theorem 4.1 applied to _D_′_, S, T_ gives the
corollary.

In turn, Theorem 4.1 follows from Corollary 4.1a by adding two new vertices _s_
and _t_ and arcs (_s, v_) for all _v_ ∈ _S_ and (_v, t_) for all _v_ ∈ _T_.
Also an arc-disjoint version can be derived (where paths are _arc-disjoint_ if they
have no arc in common).
Recall that a set _C_ of arcs is an _s_ − _t cut_ if _C_ = _δ_out(_U_) for some subset _U_ of _V_
with _s_ ∈ _U_ and _t_ 6∈ _U_.

**Corollary 4.1b** (Menger’s theorem (directed arc-disjoint version))**.** _Let D_ = (_V, A_)
_be a digraph and s, t_ ∈ _V with s_ [^6]:= _t. Then the maximum number of arc-disjoint s_−_t_
_paths is equal to the minimum size of an s_ − _t cut._

**Proof.** Let _L_(_D_) be the line digraph of _D_ and let _S_ := _δ_out(_s_) and _T_ := _δ_in(_t_).[^11]
Then Theorem 4.1 for _L_(_D_)_, S, T_ implies the corollary. Note that a minimum-size set
of arcs intersecting each _s_ − _t_ path necessarily is an _s_ − _t_ cut. Moreover, an _S_ − _T_
path in _L_(_D_) yields an _s_ − _t_ walk in _D_, which contains an _s_ − _t_ path in _D_. Hence
disjoint paths in _L_(_D_) yield arc-disjoint paths in _D_.

The internally vertex-disjoint version of Menger’s theorem can be derived in turn
from the arc-disjoint version: make a digraph _D_′ as follows from _D_: replace any

```
11For any vertex v, δout(v) = δout({v}) and δin(v) = δin({v}).
```

vertex _v_ by two vertices _v_′_, v_′′ and make an arc (_v_′_, v_′′); moreover, replace each arc
(_u, v_) by (_u_′′_, v_′). Then Corollary 4.1b for _D_′_, s_′′_, t_′ gives Corollary 4.1a for _D, s, t_.
Note that an _s_′′ − _t_′ cut _C_ in _D_′ yields a set _X_ of vertices and arcs in _D_ intersecting
each _s_−_t_ path, with |_X_| = |_C_|. Replacing each arc occurring in _X_ by one of its ends
different from _s, t_, yields an _s_ − _t_ vertex cut in _D_ of size at most |_C_|.
Similar theorems hold for _undirected_ graphs. They can be derived from the di-
rected case by replacing each undirected edge _uw_ by two opposite arcs (_u, w_) and
(_w, u_).

**Application 4.1: Routing airplanes.** An airline company carries out a certain number
of flights according to some fixed timetable, in a weekly cycle. The timetable is basically
given by a flight number _i_ (for instance 562), a departure city _dci_ (for instance Vancouver),
a departure time _dti_ (for instance Monday 23.15h), an arrival city _aci_ (for instance Tokyo),
and an arrival time _ati_ (for instance Tuesday 7.20h). All times include boarding and disem-
barking and preparing the plane for a next flight. Thus a plane with arrival time Tuesday
7.20h at city _c_, can be used for any flight from _c_ with departure time from Tuesday 7.20h
on.
The flights are carried out by _n_ airplanes of one type, denoted by _a_[^1]:_,...,an_. At each
weekday there should be an airplane for maintenance at the home basis, from 6.00h till
18.00h. Legal rules prescribe which of the airplanes _a_[^1]:_,...,an_ should be at the home basis
during one day the coming week, but it is not prescribed which airplane should be at the
home basis at which day (see Application 9.4 for an extension where this _is_ prescribed).
The timetable is made in such a way that at each city the number of incoming flights is
equal to the number of outgoing flights. Here ‘maintenance’ is also considered as a flight.
However, there is flexibility in assigning the airplanes to the flights: if at a certain moment
at a certain city two or more airplanes are available for a flight, in principle any of them
can be used for that flight. Which of the available airplanes will be used, is decided by the
main office of the company. This decision is made at 18.00h on the Saturday before. At
that time the company makes the exact routing of the planes for the coming week.
At that moment, certain planes are performing certain flights, while other planes are
grounded at certain cities. Routing the airplanes is easy as the timetable is set up in such
a way that at each moment and each city enough airplanes are available.
Indeed, one can make a directed graph _D_ (Figure 4.1) with vertex set all pairs (_dci,dti_)
and (_aci,ati_) for all flight numbers _i_. For each flight _i_ that is not in the air at Saturday
18.00h, one makes an arc from (_dci,dti_) to (_aci,ati_). We also do this for the “flights”
representing maintenance.
Moreover, for each city _c_ and each two consecutive times _t,t_′ at which any flight departs
or arrives at _c_, one makes _m_ parallel arcs from (_c,t_) to (_c,t_′), where _m_ is the number of
airplanes that will be in city _c_ during the period _t_–_t_′.
In this way we obtain a directed graph such that at each vertex the indegree is equal
to the outdegree, except at any (_c,tc_) where _tc_ is the earliest time after Saturday 18.00h
at which any flight arrives at or leaves city _c_. Hence we can find in _D_ arc-disjoint paths
_P_[^1]:_,...,Pn_ (where _n_ is the number of airplanes) in _D_ such that each arc is in exactly one of
the _Pi_. This would give a routing for the airplanes.
However, the restriction that some prescribed airplanes must undergo maintenance the

```
Sat Sun Mon Tue Wed Thu Fri Sat
```

```
maintenance maintenance maintenance maintenance maintenance
```

```
B
C
D
E
F
G
```

```
J
K
L
```

```
N
```

```
H
I
```

M

```
A
```

```
Figure 4.1
```

coming week gives some complications. It means for instance that if a certain airplane _ai_
(say) is now on the ground at city _c_ and should be home for maintenance the coming week,
then the path _Pi_ should start at (_c,tc_) and should traverse one of the arcs representing
maintenance. If _ai_ is now in the air, then path _Pi_ should start at (_c,t_) where _t_ is the
first-coming arrival time of _ai_ and should traverse a maintenance arc. So the company first
finds arc-disjoint paths _Pi_[^1]:_,...,Pik_, where _ai_[^1]:_,...,aik_ are the airplanes that should undergo
maintenance the coming week. These paths can be extended to paths _P_[^1]:_,...,Pn_ such that
each arc is traversed exactly once.
So the problem can be solved by finding arc-disjoint paths starting in a given set of
vertices and ending in a given set of vertices (by slightly extending the graph _D_).

**Exercises**

```
4.1. Let D = (V,A) be a directed graph and let s,t1,...,tk be vertices of D. Prove
that there exist pairwise arc-disjoint paths P1,...,Pk such that Pi is an s − ti path
(i = 1,...,k) if and only if for each U ⊆ V with s ∈ U one has
```

```
(1) |δout(U)|≥|{i | ti 6∈ U}|.
```

```
4.2. Let A = (A1,...,An) and B = (B1,...,Bn) be families of subsets of a finite set.
Show that A and B have a common SDR if and only if for all I,J ⊆ {1,...,n} one
has
```

```
(2)
∣∣
i∈I
```

```
Ai ∩
j∈J
```

```
Bj
∣∣ ≥|I| + |J|− n.
```

```
4.3. Let G = (V,E) be a bipartite graph, with colour classes V1 and V2, such that |V1| =
|V2|. Show that G has k pairwise disjoint perfect matchings if and only if for each
subset U of V1:
```

```
(3)
```

```
∑
```

```
v∈V2
```

```
min{k,|E(v) ∩ U|}≥ k|U|,
```

```
where E(v) denotes the set of vertices adjacent to v.
4.4. Let D = (V,A) be a simple directed graph and let s,t ∈ V. Let α be the minimum
length of an s−t path. Show that the maximum number of pairwise arc-disjont s−t
paths is at most (|V |/α)[^2].
(Hint: Let Uk denote the set of vertices at distance at most k from s. Show that
|δout(Uk)|≤ (|V |/α)[^2] for some k < α.)
```

### 4.2. Flows in networks

Other consequences of Menger’s theorem concern ‘flows in networks’. Let _D_ = (_V, A_)
be a directed graph and let _s, t_ ∈ _V_. Throughout we assume _s_ [^6]:= _t_. A function
_f_ : _A_ → R is called an _s_ − _t flow_ if:[^12]

(4) (i) _f_(_a_) ≥ [^0]: for each _a_ ∈ _A_;
(ii)

```
∑
```

```
a∈δin(v)
```

```
f(a) =
```

```
∑
```

```
a∈δout(v)
```

```
f(a) for each v ∈ V \{s, t}.
```

Condition (4)(ii) is called the _flow conservation law_: the amount of flow entering a
vertex _v_ [^6]:= _s, t_ should be equal to the amount of flow leaving _v_.
The _value_ of an _s_ − _t_ flow _f_ is, by definition:

(5) value(_f_) :=

```
∑
```

```
a∈δout(s)
```

```
f(a) −
```

```
∑
```

```
a∈δin(s)
```

```
f(a).
```

So the value is the net amount of flow leaving _s_. It can be shown that it is equal to
the net amount of flow entering _t_.
Let _c_ : _A_ → R+, called a _capacity function_. We say that a flow _f_ is _under c_ (or
_subject to c_) if

(6) _f_(_a_) ≤ _c_(_a_) for each _a_ ∈ _A._

The _maximum flow problem_ now is to find an _s_ − _t_ flow under _c_, of maximum value.
To formulate a min-max relation, define the _capacity_ of a cut _δ_out(_U_) by:
[^12]:_δ_out(_v_) and _δ_in(_v_) denote the sets of arcs leaving _v_ and entering _v_, respectively.

(7) _c_(_δ_out(_U_)) :=
_a_∈_δ_out(_U_)

```
c(a).
```

Then:

**Proposition 2.** _For every s_ − _t flow f under c and every s_ − _t cut δ_out(_U_) _one has:_

(8) value(_f_) ≤ _c_(_δ_out(_U_))_._

**Proof.**

(9) value(_f_) =

```
∑
```

```
a∈δout(s)
```

```
f(a) −
```

```
∑
```

```
a∈δin(s)
```

```
f(a)
```

```
=
```

```
∑
```

```
a∈δout(s)
```

```
f(a) −
```

```
∑
```

```
a∈δin(s)
```

```
f(a) +
```

```
∑
```

```
v∈U\{s}
```

```
(
```

```
∑
```

```
a∈δout(v)
```

```
f(a) −
```

```
∑
```

```
a∈δin(v)
```

```
f(a))
```

```
=
```

```
∑
```

```
v∈U
```

```
(
```

```
∑
```

```
a∈δout(v)
```

```
f(a) −
```

```
∑
```

```
a∈δin(v)
```

```
f(a)) =
```

```
∑
```

```
a∈δout(U)
```

```
f(a) −
```

```
∑
```

```
a∈δin(U)
```

```
f(a)
```

```
⋆
≤
```

```
∑
```

```
a∈δout(U)
```

```
f(a)
```

```
⋆⋆
≤
```

```
∑
```

```
a∈δout(U)
```

```
c(a) = c(δout(U)).
```

```
It is convenient to note the following:
```

(10) equality holds in (8) ⇐⇒ ∀_a_ ∈ _δ_in(_U_) : _f_(_a_) = 0 and
∀_a_ ∈ _δ_out(_U_) : _f_(_a_) = _c_(_a_)_._

This follows directly from the inequalities _⋆_ and _⋆⋆_ in (9).
Now from Menger’s theorem one can derive that equality can be attained in (8),
which is a theorem of Ford and Fulkerson [1956]:

**Theorem 4.2** (max-flow min-cut theorem)**.** _For any directed graph D_ = (_V, A_)_,_

_s, t_ ∈ _V with s_ [^6]:= _t, and c_ : _A_ → R+_, the maximum value of an s_ − _t flow under c is_
_equal to the minimum capacity of an s_ − _t cut. In formula:_

(11) max
_f s_-_t_ flowvalue(_f_) = min
_δ_out(_U_) _s_-_t_ cut

```
c(δout(U)).
```

**Proof.** If _c_ is integer-valued, the corollary follows from Menger’s theorem (directed
arc-disjoint version — Corollary 4.1b) by replacing each arc _a_ by _c_(_a_) parallel arcs.
If _c_ is rational-valued, there exists a natural number _N_ such that _N c_(_a_) is integer
for each _a_ ∈ _A_. Changing each _c_(_a_) to _N c_(_a_) multiplies both the maximum and the
minimum by _N_. So the equality follows from the case where _c_ is integer-valued.

If _c_ is real-valued, we can derive the corollary from the case where _c_ is rational-
valued, by continuity and compactness arguments.

```
Moreover, one has (Dantzig [1951a]):
```

**Corollary 4.2a** (Integrity theorem)**.** _If c is integer-valued, there exists an integer-_
_valued maximum flow._

**Proof.** Directly from Menger’s theorem.

**Exercises**

```
4.5. Let D = (V,A) be a directed graph and let s,t ∈ V. Let f : A → R+ be an s−t flow
of value β. Show that there exists an s − t flow f′ : A → Z+ of value ⌈β⌉ such that
⌊f(a)⌋≤ f′(a) ≤⌈f(a)⌉ for each arc a. (Integer flow theorem (Dantzig [1951a]).)
```

```
4.6. Let G = (V,E) be a graph and let c : E → R+ be a ‘capacity’ function. Let K be
the complete graph on V. For each edge st of K, let w(st) be the minimum capacity
of any s − t cut in G. [An s − t cut is any subset δ(W) with s ∈ W,t 6∈ W.]
Let T be a spanning tree in K of maximum total weight with respect to the function
w. Prove that for all s,t ∈ V , w(st) is equal to the minimum weight of the edges of
T in the unique s − t path in T.
(Hint: Use Exercise 1.10.)
```

### 4.3. Finding a maximum flow

Let _D_ = (_V, A_) be a directed graph, let _s, t_ ∈ _V_ , and let _c_ : _A_ → Q+ be a ‘capacity’
function. We now describe the algorithm of Ford and Fulkerson [1956] to find an _s_−_t_
flow of maximum value under _c_.
In this section, by _flow_ we will mean an _s_ − _t_ flow under _c_, and by _cut_ an _s_ − _t_
cut. A _maximum flow_ is a flow of maximum value.
We now describe the algorithm of Ford and Fulkerson [1957] to determine a max-
imum flow. We assume that _c_(_a_) _>_ 0 for each arc _a_. Moreover, for the sake of
exposition we assume that for no arc (_u, v_), also (_v, u_) is an arc. (We can achieve
this by, if necessary, ‘inserting a new vertex on an arc’.) First we give an important
subroutine:

**Flow augmenting algorithm**

**input:** a flow _f_.
**output:** either (i) a flow _f_′ with value(_f_′) _>_ value(_f_),
or (ii) a cut _δ_out(_W_) with _c_(_δ_out(_W_)) = value(_f_).

```
description of the algorithm: For any pair a = (v, w) define a−[^1] := (w, v). Make
an auxiliary graph Df = (V, Af) by the following rule: for any arc a ∈ A,
```

```
(12) if f(a) < c(a) then a ∈ Af,
if f(a) > 0 then a−[^1] ∈ Af.
```

```
So if 0 < f(a) < c(a) then both a and a−[^1] are arcs of Af.
Now there are two possibilities:
```

(13) **Case 1:** _There exists an s_ − _t path in Df,_
**Case 2:** _There is no s_ − _t path in Df._

```
Case 1: There exists an s − t path P = (v0, a1, v1,... , ak, vk) in Df = (V, Af).
So v0 = s and vk = t. We may assume that P is a path. As a1,... , ak belong to Af,
we know by (12) that for each i = 1,... , k:
```

```
(14) either (i) ai ∈ A and σi := c(ai) − f(ai) > 0
or (ii) a−[^1]
i ∈ A and σi := f(a−[^1]
i ) >[^0].
```

```
Define α := min{σ1,... , σk}. So α > 0. Let f′ : A → R+ be defined by, for a ∈ A:
```

```
(15) f′(a) := f(a) + α, if a = ai for some i = 1,... , k;
:= f(a) − α, if a = a−[^1]
i for some i = 1,... , k;
:= f(a), for all other a.
```

```
Then f′ again is an s − t flow under c. The inequalities 0 ≤ f′(a) ≤ c(a) hold
because of our choice of α. It is easy to check that also the flow conservation law
(4)(ii) is maintained.
Moreover,
```

```
(16) value(f′) = value(f) + α,
```

```
since either (v0, v1) ∈ A, in which case the outgoing flow in s is increased by α, or
(v1, v0) ∈ A, in which case the ingoing flow in s is decreased by α.
Path P is called a flow augmenting path.
```

```
Case 2: There is no path in Df = (V, Af) from s to t.
Now define:
```

```
(17) U := {u ∈ V | there exists a path in Df from s to u}.
```

```
Then s ∈ U while t 6∈ U, and so δout(U) is an s − t cut.
```

By definition of _U_, if _u_ ∈ _U_ and _v_ 6∈ _U_, then (_u, v_) 6∈ _Af_ (as otherwise also _v_
would belong to _U_). Therefore:

(18) if (_u, v_) ∈ _δ_out(_U_), then (_u, v_) 6∈ _Af_, and so (by (12)): _f_(_u, v_) = _c_(_u, v_),
if (_u, v_) ∈ _δ_in(_U_), then (_v, u_) 6∈ _Af_, and so (by (12)): _f_(_u, v_) = 0.

Then (10) gives:

(19) _c_(_δ_out(_U_)) = value(_f_)_._

This finishes the description of the flow augmenting algorithm. The description
of the _(Ford-Fulkerson) maximum flow algorithm_ is now simple:

**Maximum flow algorithm**

**input:** directed graph _D_ = (_V, A_)_, s, t_ ∈ _V, c_ : _A_ → R+.
**output:** a maximum flow _f_ and a cut _δ_out(_U_) of minimum capacity, with value(_f_) =
_c_(_δ_out(_U_)).
**description of the algorithm:** Let _f_[^0]: be the ‘null flow’ (that is, _f_[^0]:(_a_) = 0 for
each arc _a_). Determine with the flow augmenting algorithm flows _f_[^1]:_, f_[^2]:_,... , fN_ such
that _fi_+1 = _f_′
_i_, until, in the (_N_ + 1)-th iteration, say, we obtain output (ii) of the
flow augmenting algorithm. Then we have flow _fN_ and a cut _δ_out(_U_) with the given
properties.

```
We show that the algorithm terminates, provided that all capacities are rational.
```

**Theorem 4.3.** _If all capacities c_(_a_) _are rational, the algorithm terminates._

**Proof.** If all capacities are rational, there exists a natural number _K_ so that _Kc_(_a_)
is an integer for each _a_ ∈ _A_. (We can take for _K_ the l.c.m. of the denominators of
the _c_(_a_).)
Then in the flow augmenting iterations, every flow _fi_(_a_) and every _α_ is a multiple
of 1_/K_. So at each iteration, the flow value increases by at least 1_/K_. Since the flow
value cannot exceed _c_(_δ_out(_s_)), we can have only finitely many iterations.

We should note here that this theorem is not true if we allow general real-valued
capacities.
In Section 4.4 we shall see that if we choose always a shortest path as flow aug-
menting path, then the algorithm has polynomially bounded running time.
Note that the algorithm also implies the max-flow min-cut theorem (Theorem
4.2). Note moreover that in the maximum flow algorithm, if all capacities are integer,
then the maximum flow will also be integer-valued. So it also implies the integrity
theorem (Corollary 4.2a).

**Application 4.2: Transportation problem.** Suppose that there are _m_ factories, that
all produce the same product, and _n_ customers that use the product. Each month, factory
_i_ can produce _si_ tons of the product. Customer _j_ needs every month _dj_ tons of the product.
From factory _i_ to customer _j_ we can transport every month at most _ci,j_ tons of the product.
The problem is: can the needs of the customers be fulfilled?

```
s t
```

```
1 b
```

```
f2
```

```
1
```

```
b4
```

```
f3
```

```
b3
```

```
b5
```

```
b2
```

```
f
```

```
Figure 4.2
```

In order to solve the problem with the maximum-flow algorithm, we make the graph as
in Figure 4.2 (for _m_ = 3, _n_ = 5). We define a capacity function _c_ on the arcs as follows:

(20) _c_(_s,fi_) := _si_ for _i_ = 1_,...,m_,
_c_(_fi,bj_) := _ci,j_ for _i_ = 1_,...,m_;_j_ = 1_,...,n_,
_c_(_bj,t_) := _dj_ for _j_ = 1_,...,n_.

Now we have:

(21) the needs of the customers can be fulfilled ⇐⇒ there is an _s_−_t_ flow under _c_ with
value _d_[^1]: + ··· + _dn_.

Since there cannot exist an _s_ − _t_ flow under _c_ of value larger than _d_[^1]: + ··· + _dn_ (since
_c_(_δ_in(_t_)) = _d_[^1]: + ··· + _dn_), the problem can be solved with the maximum-flow algorithm.
If there exists a flow of value _d_[^1]: +···+_dn_, then the flow on arc (_fi,bj_) gives the amount
that should be transported each month from factory _i_ to customer _j_. The flow on arc (_s,fi_)
gives the amount to be produced each month by factory _fi_.

**Exercises**

```
4.7. Determine with the maximum flow algorithm an s − t flow of maximum value and
an s − t cut of minimum capacity in the following graphs (where the numbers at the
arcs give the capacities):
```

```
(i)
```

```
1 11
```

```
5
```

```
2
```

```
5
```

```
1
```

```
2
```

```
4
```

```
2
```

```
7
```

```
4
```

```
10 2 1
```

```
s[^2][^2] t
```

```
2
```

```
(ii)
```

```
12
```

```
7
3
```

```
4
1 1
```

```
3
```

```
1
```

```
9
```

```
1 11
2
```

```
1
```

```
3 2
```

```
2 5 3
```

```
2
7
5
```

```
s t
```

(iii) s t

```
3 3
```

```
6 9
```

```
8
```

```
4
```

```
5
```

```
4
```

```
4
```

```
1
```

```
2 5
```

```
7
```

```
1
```

```
2
6
```

(iv) s t

```
2 2
```

```
4
```

```
4
```

```
5
```

```
2
```

```
2
```

```
10 4
```

```
12
```

```
1
```

```
5 7
```

```
6
```

```
2
```

```
3
```

```
3
```

```
4.8. Solve the transportation problem with the maximum-flow algorithm for the following
data: m = n = 3,s1 = 13,s2 = 9,s3 = 4,d1 = 3,d2 = 7,d3 = 12,
```

```
ci,j j = 1 j = 2 j = 3
i = 1 2 0 8
i = 2[^3][^8][^3]
i = 3 0 1 3
```

```
4.9. Describe the problem of finding a maximum-size matching in a bipartite graph as a
maximum flow problem.
```

4.10. Determine with the maximum-flow algorithm if there exists a 3 × 3 matrix _A_ = (_ai,j_)
satisfying:[^13]

```
ai,j ≥ 0 for all i,j = 1,2,3;
```

```
A1 ≤
```

```


```

```
13
9
4
```

```

```

```
;
```

```
1TA = (3,7,12);
```

```
A ≤
```

```


```

```
2 0 8
3 8 3
0 1 3
```

```

```

```
.
```

4.11. Give an example of a directed graph with irrational capacities, in which, at a bad
choice of flow augmenting paths, the maximum flow algorithm does not terminate.

4.12. Let _D_ = (_V,A_) be a directed graph, let _s,t_ ∈ _V_ and let _f_ : _A_ → Q+ be an _s_ − _t_ flow
of value _b_. Show that for each _U_ ⊆ _V_ with _s_ ∈ _U,t_ 6∈ _U_ one has:

```
(22)
```

```
∑
```

```
a∈δout(U)
```

```
f(a) −
```

```
∑
```

```
a∈δin(U)
```

```
f(a) = b.
```

### 4.4. Speeding up the maximum flow algorithm

We saw that the number of iterations in the maximum flow algorithm is finite, if all
capacities are rational. If we choose as our flow augmenting path _P_ in the auxiliary

graph _Df_ an _arbitrary s_ − _t_ path, the number of iterations yet can get quite large.
For instance, in the graph in Figure 4.3 the number of iterations, at a bad choice of
paths, can become 2 · [^10]:_k_.

```
131 denotes the vector (1,1,1)T.
```

```
s 1 t
```

```
10 10
```

```
10 10
```

```
k k
```

```
k k
```

```
Figure 4.3
```

However, if we choose always a _shortest s_ − _t_ path in _Df_ as our flow augmenting
path _P_ (that is, with a minimum number of arcs), then the number of iterations is
at most |_V_ |·|_A_|. This was shown by Dinits [1970] and Edmonds and Karp [1972].
Again, for any directed graph _D_ = (_V, A_) and _s, t_ ∈ _V_ , let _μ_(_D_) denote the
minimum length of an _s_−_t_ path. Moreover, let _α_(_D_) denote the set of arcs contained
in at least one shortest _s_ − _t_ path. Then one has:

**Proposition 3.** _Let D_ = (_V, A_) _and s, t_ ∈ _V. Let D_′ := (_V, A_ ∪ _α_(_D_)−[^1])_. Then_
_μ_(_D_′) = _μ_(_D_) _and α_(_D_′) = _α_(_D_)_._

**Proof.** It suffices to show that _μ_(_D_) and _α_(_D_) are invariant if we add _a_−[^1] to _D_ for
one arc _a_ ∈ _α_(_D_). Suppose not. Then there is an _s_ − _t_ path _P_ traversing _a_−[^1], of
length at most _μ_(_D_). As _a_ ∈ _α_(_D_), there is an _s_ − _t_ path _Q_ traversing _a_, of length
_μ_(_D_). Hence _AP_ ∪ _AQ_ \{_a, a_−[^1]} contains an _s_ − _t_ path of length less than _μ_(_D_), a
contradiction.

```
This implies the result of Dinits [1970] and Edmonds and Karp [1972]:
```

**Theorem 4.4.** _If we choose in each iteration a shortest s_−_t path as flow augmenting_
_path, the number of iterations is at most_ |_V_ ||_A_|_._

**Proof.** If we augment flow _f_ along a shortest path _P_, obtaining flow _f_′, then _Df_′
is a subgraph of _D_′ := (_V, Af_ ∪ _α_(_Df_)−[^1]). Hence _μ_(_Df_′) ≥ _μ_(_D_′) = _μ_(_Df_) (by
Proposition 3). Moreover, if _μ_(_Df_′) = _μ_(_Df_), then _α_(_Df_′) ⊆ _α_(_D_′) = _α_(_Df_) (again
by Proposition 3). As at least one arc in _P_ belongs to _Df_ but not to _Df_′, we have a
strict inclusion.
So throughout the iterations, (_μ_(_Df_)_,_−|_α_(_Df_)|) strictly increases lexicographi-
cally. Since _μ_(_Df_) _<_ |_V_ | and |_α_(_Df_)|≤|_A_| for each _f_, the number of iterations is at
most |_V_ ||_A_|.

```
Since a shortest path can be found in time O(|A|), this gives:
```

**Corollary 4.4a.** _The maximum flow problem can be solved in time O_(|_V_ ||_A_|[^2])_._

**Proof.** Directly from Theorem 4.4.

This algorithm can be improved, as was shown by Karzanov [1974]. In each
iteration we find a shortest path in _O_(|_A_|) time. But as long as the distance from _s_
to _t_ does not increase, we could use the data-structure of the previous shortest path
search so as to find the next shortest path.
This can be described as follows. Call an _s_ − _t_ flow _f blocking_ if for each _s_ − _t_
flow _g_ ≥ _f_ one has _g_ = _f_. Now Karzanov [1974] showed the following (we give the
short proof of Malhotra, Kumar, and Maheshwari [1978]; see also Tarjan [1984]):

**Theorem 4.5.** _Given an acyclic directed graph D_ = (_V, A_)_, s, t_ ∈ _V , and a capacity_
_function c_ : _A_ → Q+_, a blocking flow can be found in time O_(|_V_ |[^2])_._

**Proof.** First order the vertices reachable from _s_ as _s_ = _v_[^1]:_, v_[^2]:_,... , vn_−[^1]:_, vn topologi-_
_cally_; that is, if (_vi, vj_) ∈ _A_ then _i < j_. This can be done in time _O_(|_A_|).[^14]
We describe the procedure recursively. Consider the minimum of the values
_c_(_δ_in(_v_)) for all _v_ ∈ _V_ \{_s_} and _c_(_δ_out(_v_)) for all _v_ ∈ _V_ \{_t_}. Let the minimum

be attained by _vi_ and _c_(_δ_out(_vi_)) (without loss of generality). Define _f_(_a_) := _c_(_a_) for
each _a_ ∈ _δ_out(_vi_) and _f_(_a_) := 0 for all other _a_.
Next for _j_ = _i_+1_,... , n_−1, redefine _f_(_a_) for each _a_ ∈ _δ_out(_vj_) so that _f_(_a_) ≤ _c_(_a_)
and so that _f_(_δ_out(_vj_)) = _f_(_δ_in(_vj_)). By the minimality of _vi_ and _c_(_δ_in(_v_)), we can
always do this, as initially _f_(_δ_in(_vj_)) ≤ _c_(_δ_out(_vi_)) ≤ _c_(_δ_in(_vj_)). We do this in such a
way that finally _f_(_a_) ∈{[^0]:_, c_(_a_)} for all but at most one _a_ in _δ_out(_vj_).
After that, for _j_ = _i, i_ − [^1]:_,... ,_2, redefine similarly _f_(_a_) for _a_ ∈ _δ_in(_vj_) so that
_f_(_a_) ≤ _c_(_a_) and so that _f_(_δ_in(_vj_)) = _f_(_δ_out(_vj_)).
If _vi_ ∈{_s, t_} we stop, and _f_ is a blocking flow.
If _vi_ 6∈ {_s, t_}, set _c_′(_a_) := _c_(_a_) − _f_(_a_) for each _a_ ∈ _A_, and delete all arcs _a_ with
_c_′(_a_) = 0 and delete _vi_ and all arcs incident with _vi_, thus obtaining the directed graph
_D_′ = (_V_ ′_, A_′). Obtain (recursively) a blocking flow _f_′ in _D_′ subject to the capacity
function _c_′. Define _f_′′(_a_) := _f_(_a_) + _f_′(_a_) for _a_ ∈ _A_′ and _f_′′(_a_) = _f_(_a_) for _a_ ∈ _A_ \ _A_′.
Then _f_′′ is a blocking flow in _D_.
This describes the algorithm. The correctness can be seen as follows. If _vi_ ∈{_s, t_}
the correctness is immediate. If _vi_ 6∈ {_s, t_}, suppose _f_′′ is not a blocking flow in _D_,
and let _P_ be an _s_ − _t_ path in _D_ such that _f_′′(_a_) _< c_(_a_) for each arc _a_ in _P_. Then
each arc of _P_ belongs to _A_′, since _f_′′(_a_) = _f_(_a_) = _c_(_a_) for each _a_ ∈ _A_ \ (_A_′ ∪ _δ_in(_vi_)).
So for each arc _a_ of _P_ one has _c_′(_a_) = _c_(_a_) − _f_(_a_) _> f_′′(_a_) − _f_(_a_) = _f_′(_a_). This
contradicts the fact that _f_′ is a blocking flow in _D_′.
The running time of the algorithm is _O_(|_V_ |[^2]), since the running time of the iter-

[^14]:This can be done recursively as follows (cf. Knuth [1968], Tarjan [1974]). If _δ_out(_s_) = ∅, then
the ordering is trivial. If _δ_out(_s_) [^6]:= ∅, choose (_s,v_) ∈ _δ_out(_s_), and order the vertices reachable from _v_
topologically, as _w_[^1]:_,...,wm_, delete them from _D_, and order the remaining vertices reachable from
_s_ topologically as _v_[^1]:_,...,vk_; then _v_[^1]:_,...,vk,w_[^1]:_,...,wm_ gives a required topological ordering.

ation is _O_(|_V_ | + |_A_ \ _A_′|), and since there are at most |_V_ | iterations. (Note that we
determine the topological ordering only once, at the preprocessing.)

```
Theorem 4.5 has the following consequence:
```

**Corollary 4.5a.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V , and a capacity function_
_c_ : _A_ → Q_, a flow f satisfying μ_(_Df_) _> μ_(_D_) _can be found in time O_(|_V_ |[^2])_._

**Proof.** Let _D_ ̃ be the subgraph of _D_ consisting of all arcs that are contained in a
shortest _s_ − _t_ path in _D_. Find a blocking flow in _D_ ̃. Then _μ_(_Df_) _> μ_(_D_) (by
Proposition 3).

```
Hence we have:
```

**Corollary 4.5b.** _Given a directed graph D_ = (_V, A_)_, s, t_ ∈ _V , and a capacity function_
_c_ : _A_ → Q_, a maximum s_ − _t flow can be found in time O_(|_V_ |[^3])_._

**Proof.** Directly from the foregoing.

Goldberg and Tarjan [1990] gave an _O_(|_A_|log(|_V_ |[^2]_/_|_A_|)) algorithm for finding
a blocking flow in an acyclic directed graph, implying an _O_(|_V_ ||_A_|log(|_V_ |[^2]_/_|_A_|))
algorithm for finding a maximum flow in any directed graph. An alternative approach
finding a maximum flow in time _O_(|_V_ ||_A_|log(|_V_ |[^2]_/_|_A_|)) was described in Goldberg
and Tarjan [1988].
For surveys on maximum flow algorithms, see Goldberg, Tardos, and Tarjan [1990]
and Ahuja, Magnanti, and Orlin [1993].

### 4.5. Circulations

A theorem related to the max-flow min-cut theorem is due to Hoffman [1960] and
concerns circulations. Let _D_ = (_V, A_) be a directed graph. A function _f_ : _A_ → R is
called a _circulation_ if for each vertex _v_ ∈ _V_ one has:

(23)

```
∑
```

```
a∈δin(v)
```

```
f(a) =
```

```
∑
```

```
a∈δout(v)
```

```
f(a).
```

So now the flow conservation law holds at _each_ vertex _v_.
Hoffman [1960] proved the following theorem (which can also be derived from the
max-flow min-cut theorem, but a direct proof seems shorter). For any directed graph
_D_ = (_V, A_), and any _d, c, f_ : _A_ → R with _d_(_a_) ≤ _f_(_a_) ≤ _c_(_a_) for each _a_ ∈ _A_, we
define

(24) _Af_ := {_a_ | _f_(_a_) _< c_(_a_)}∪{_a_−[^1] | _d_(_a_) _< f_(_a_)}_,_

and _Df_ := (_V, Af_).

**Theorem 4.6** (Hoffman’s circulation theorem)**.** _Let D_ = (_V, A_) _be a directed graph_
_and let d, c_ : _A_ → R _be such that d_(_a_) ≤ _c_(_a_) _for each arc a. Then there exists a_
_circulation f such that_

(25) _d_(_a_) ≤ _f_(_a_) ≤ _c_(_a_)

_for each arc a if and only if_

(26)

```
∑
```

```
a∈δin(U)
```

```
d(a) ≤
```

```
∑
```

```
a∈δout(U)
```

```
c(a)
```

_for each subset U of V._

**Proof.** To see necessity of (26), suppose that a circulation _f_ satisfying (25) exists.
Then[^15]

(27) _d_(_δ_in(_U_)) ≤ _f_(_δ_in(_U_)) = _f_(_δ_out(_U_)) ≤ _c_(_δ_out(_U_))_._

```
To see sufficiency, define for any f : A → R and any v ∈ V ,
```

(28) loss_f_(_v_) := _f_(_δ_out(_v_)) − _f_(_δ_in(_v_))_._

So loss_f_ can be seen as a vector in R_V_.
Choose a function _f_ satisfying _d_ ≤ _f_ ≤ _c_ and minimizing ‖loss_f_‖[^1]:. Let

(29) _S_ := {_v_ ∈ _V_ | loss_f_(_v_) _<_ [^0]:} and _T_ := {_v_ ∈ _V_ | loss_f_(_v_) _>_ [^0]:}.

If _S_ = ∅, then _f_ is a circulation, and we are done. So assume _S_ [^6]:= ∅. If _Df_ contains
an _S_ − _T_ path, we can modify _f_ so as to reduce ‖loss_f_‖[^1]:. So _Df_ does not contain
any _S_ − _T_ path. Let _U_ be the set of vertices reachable in _Df_ from _S_. Then for each
_a_ ∈ _δ_out(_U_) we have _a_ 6∈ _Af_ and hence _f_(_a_) = _c_(_a_). Similarly, for each _a_ ∈ _δ_in(_U_) we
have _a_−[^1] 6∈ _Af_ and hence _f_(_a_) = _d_(_a_). Therefore

(30) _c_(_δ_out(_U_)) − _d_(_δ_in(_U_)) = _f_(_δ_out(_U_)) − _f_(_δ_in(_U_)) = loss_f_(_U_) = loss_f_(_S_) _<_ [^0]:_,_

contradicting (26).

```
15For any function g : A → R and subset B of A, we write g(B) =
```

```
∑
a∈B g(a).
```

```
One has moreover:
```

**Theorem 4.7.** _In Theorem_ 4.6_, if c and d are integer and there exists a circulation_
_f satisfying d_ ≤ _f_ ≤ _c, then there exists an integer-valued circulation f_′ _satisfying_
_d_ ≤ _f_′ ≤ _c._

**Proof.** Directly from the proof above.

**Exercises**

4.13. Let _D_ = (_V,A_) be a directed graph and let _f_ : _A_ → R be a circulation. Show that
there exists a circulation _f_′ such that _f_′ is integer-valued and such that ⌊_f_(_a_)⌋ ≤
_f_′(_a_) ≤⌈_f_(_a_)⌉ for each arc _a_.

4.14. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be partitions of a finite set _X_ and let _k_
be a natural number. Prove that _X_ can be covered by _k_ common SDR’s of A and B
if and only if

```
(31)
∣∣(
```

```
⋃
```

```
i∈I
```

```
Ai ∪
```

```
⋃
```

```
j∈J
```

```
Bj)
∣∣ ≥|X| + k(|I| + |J|− n)
```

```
for all I,J ⊆{1,...,n} with
```

```
⋃
i∈I Ai ∩
```

```
⋃
j∈J Bj = ∅.
```

4.15. Let _D_ = (_V,A_) be a directed graph, and let _f_ : _A_ → R+. Let C be the collection
of directed circuits in _D_. For each directed circuit _C_ in _D_ let _χC_ be the incidence
vector of _C_. That is, _χC_ : _A_ →{[^0]:_,_[^1]:}, with _χC_(_a_) = 1 if _C_ traverses _a_ and _χC_(_a_) = 0
otherwise.
Show that _f_ is a nonnegative circulation if and only if there exists a function _λ_ : C →
R+ such that

```
(32) f =
```

```
∑
```

```
C∈C
```

```
λ(C)χC.
```

```
That is, the nonnegative circulations form the cone generated by {χC | C ∈C}.
```

### 4.6. Minimum-cost flows

In the previous sections we were searching for flows of maximum value. In this section
we consider the problem of finding a flow of maximum value with the additional
property that it has ‘minimum cost’.
Let be given again a directed graph _D_ = (_V, A_), distinct vertices _s_ and _t_ of _D_,
and a capacity function _c_ : _A_ → R+. Let moreover be given a function _k_ : _A_ → R+,

```
called the cost function. Again we assume that c(a) > 0 for each a, and that for no
u, v ∈ V , both (u, v) and (v, u) occur as arc of D.
Define for any function f : A → R+ the cost of f as:
```

```
(33) cost(f) :=
```

```
∑
```

```
a∈A
```

```
k(a)f(a).
```

```
The following is the minimum-cost flow problem (or min-cost flow problem):
```

(34) given: a directed graph _D_ = (_V, A_), _s, t_ ∈ _V_ , a capacity function _c_ : _A_ → R+
and a cost function _k_ : _A_ → R+;
find: an _s_ − _t_ flow subject to _c_ of maximum value, such that _f_ has minimum
cost among all _s_ − _t_ flows subject to _c_ of maximum value.

```
This problem can be solved with an adaptation of the algorithm described in
Section 4.3. Let us define an s−t flow f ≤ c to be an extreme flow if f has minimum
cost among all s − t flows g ≤ c with value(g) = value(f).
So an extreme flow does not need to have maximum value. An extreme flow is a
flow f that has minimum cost among all flows with the same value as f.
Define a length function l : A ∪ A−[^1] → R by:
```

```
(35) l(a) :=
```

```
{
k(a) if a ∈ A,
−k(a−[^1]) if a−[^1] ∈ A
```

```
for each a ∈ A ∪ A−[^1]. For any directed path or circuit P in A ∪ A−[^1], let l(P) be the
sum of l(a) over arcs a in P.
For any flow f, let again Df = (V, Af) be the auxiliary graph corresponding to f
(in the sense of the flow augmenting algorithm).
```

```
Proposition 4. f is an extreme flow if and only if Df has no directed circuit C with
l(C) < 0.
```

```
Proof. Necessity. Suppose that C = (a1,... , ak) is a directed circuit in Df of negative
length; that is,
```

```
(36) l(C) = l(a1) + l(a2) + ··· + l(ak) < 0.
```

```
So a1,... , ak are arcs in Df. Define for i = 1,... , k:
```

```
(37) σi :=
```

```
{
c(ai) − f(ai) if ai ∈ A,
f(a−[^1]
i ) if a−[^1]
i ∈ A.
```

Note that by definition of _Df_, _σi >_ 0 for each _i_ = 1_,... , k_. Let _α_ := min{_σ_[^1]:_,... , σk_}
and define for each _a_ ∈ _A_:

(38) _g_(_a_) :=

```


```

```

```

```
f(a) + α if a ∈ C,
f(a) − α if a−[^1] ∈ C,
f(a) otherwise.
```

Then _g_ is again an _s_ − _t_ flow subject to _c_, with value(_g_) = value(_f_). Moreover one
has

(39) cost(_g_) = cost(_f_) + _α_ · _l_(_C_) _<_ cost(_f_)_._

So _f_ is not an extreme flow.

_Sufficiency._ Let _g_ be any flow with value(_g_) =value(_f_). Define _h_ : _Af_ → R+ by:

(40) _h_(_a_) := _g_(_a_) − _f_(_a_) if _g_(_a_) _> f_(_a_), and
_h_(_a_−[^1]) := _f_(_a_) − _g_(_a_) if _g_(_a_) _< f_(_a_),

for _a_ ∈ _A_, while _h_(_a_) = 0 for all other arcs _a_ of _Af_. Then _h_ is a circulation in _Df_.
∑ Hence, by Exercise 4.15, there exists a function _λ_ : C → R+ such that _h_ =
_C_∈C _λ_(_C_)_χC_. Hence cost(_g_) − cost(_f_) =

∑
_C_∈C _λ_(_C_)_l_(_C_). Assuming _Df_ has no
directed circuits of negative length, it follows that cost(_g_) ≥ cost(_f_). So _f_ is an
extreme flow.

```
With this we can show:
```

**Proposition 5.** _Let f be an extreme flow. Let f_′ _arise by choosing in the flow_
_augmenting algorithm a path P in Df minimizing l_(_P_)_. Then f_′ _is an extreme flow_
_again._

**Proof.** If _Df_′ has a directed circuit _C_ with _l_(_C_) _<_ 0, we could have chosen a shorter
_s_ − _t_ path in _Df_. To see this, let _A_(_P_) and _A_(_C_) be the sets of arcs in _P_ and _C_,
respectively. Then _A_(_C_) ⊆ _Af_′ ⊆ _Af_ ∪ _A_(_P_)−[^1]. Hence for _B_ := _A_(_C_) ∩ _A_(_P_)−[^1], the
set (_A_(_P_) \ _B_−[^1]) ∪ (_A_(_C_) \ _B_) is a subset of _Af_. Moreover, it can be decomposed
into a directed _s_ − _t_ path _Q_ and a number of directed circuits _C_[^1]:_,... , Ck_, counting
multiplicities: if an arc occurs both in _A_(_P_) and _A_(_C_), then it occurs twice in this
decomposition. (This decomposition exists since adding arc (_t, s_) makes (_V,_(_A_(_P_) \
_B_−[^1]) ∪ (_A_(_C_) \ _B_)) Eulerian.)
Since _l_(_Ci_) ≥ 0 for each _i_ (as _Ci_ is a directed circuit in _Af_ and _f_ is extreme), and
since _l_(_B_−[^1]) = −_l_(_B_), we have

(41) _l_(_P_) _> l_(_P_)+_l_(_C_) = _l_(_P_)−_l_(_B_−[^1])+_l_(_C_)−_l_(_B_) = _l_(_Q_)+

```
∑k
```

```
i=1
```

```
l(Ci) ≥ l(Q).
```

This contradicts the fact that _P_ is shortest in _Df_.

This implies that the min-cost flow problem can be solved by choosing in the flow
augmenting algorithm a shortest path in the auxiliary graph throughout. The first
flow, the all-zero flow _f_[^0]:, is trivially an extreme flow. Hence also all further flows
_f_[^1]:_, f_[^2]:_, f_[^3]:_,..._ will be extreme flows by Proposition 5. Therefore, also the last flow,
which is of maximum value, is an extreme flow. So we have a solution to the min-cost
flow problem. (Here we assume that all capacities are rational.)

```
In this process, we should be able to find a shortest s − t path in the auxiliary
```

graphs _Df_. This is indeed possible with the Bellman-Ford method, since _Df_ does not
have directed circuits of negative length as we saw in Proposition 4.

The algorithm can be modified so that all lengths are nonnegative throughout
the iterations, and this yields a running time of _O_(_M_ · (_m_ + _n_log _n_)), where _M_ is
the value of a maximum flow (assuming all capacities to be integer). This is not
polynomial-time. Tardos [1985] gave the first polynomial-time algorithm to find a
minimum-cost flow of maximum value. At the moment of writing, the asymptotically
fastest method was given by Orlin [1988,1993] and runs in _O_(_m_log _n_(_m_ + _n_log _n_))
time.

In a similar way one can describe a _minimum-cost circulation_ algorithm. For more
about network flows we refer to the books of Ford and Fulkerson [1962] and Ahuja,
Magnanti, and Orlin [1993].

**Application 4.3: Minimum-cost transportation problem.** Beside the data in Appli-
cation 4.2 one may also have a cost function _ki,j_, giving the cost of transporting 1 ton from
factory _i_ to costumer _j_. Moreover, there is given a cost _ki_ of producing 1 ton by factory
_i_ (for each _i_). We want to make a production and transportation plan that minimizes the
total cost.
This problem can be solved by assigning also costs to the arcs in Application 4.2. We
can take the costs on the arcs from _bj_ to _t_ equal to 0.

**Application 4.4: Routing empty freighters.** Historically, in his paper “Optimum
utilization of the transportation system”, Koopmans [1948] was one of the first studying
the minimum-cost transportation problem, in particular with application to the problem of
routing empty freighters. Koopmans considered the surplus and need of register ton of ship
capacity at harbours all over the world, as given by the following table (data are aggregated
to main harbours):

```
Net receipt of dry cargo in overseas trade, 1925
Unit: Millions of metric tons per annum
Harbour Received Dispatched Net receipts
New York 23.5 32.7 −9.2
San Francisco 7.2 9.7 −2.5
St. Thomas 10.3 11.5 −1.2
Buenos Aires 7.0 9.6 −2.6
Antofagasta 1.4 4.6 −3.2
Rotterdam 126.4 130.5 − 4.1
Lisbon 37.5 17.0 20.5
Athens 28.3 14.4 13.9
Odessa 0.5 4.7 −4.2
Lagos 2.0 2.4 −0.4
Durban 2.1 4.3 −2.2
Bombay 5.0 8.9 −3.9
Singapore 3.6 6.8 −3.2
Yokohama 9.2 3.0 6.2
Sydney 2.8 6.7 −3.9
Total 266.8 266.8 0.0
```

Given is moreover a distance table between these harbours. Koopmans wondered how
ships should be routed between harbours so as to minimize the total amount of ton kilome-
ters made by empty ships.
This problem is a special case of the min-cost flow problem. Make a graph with vertex
set all harbours, together with two dummy harbours _s_ and _t_. From any harbour _u_ with
a surplus (positive net receipt) to any harbour _w_ with a need (negative net receipt) make
an arc with cost equal to the distance between _u_ and _w_, and with capacity ∞. Moreover,
from _s_ to any harbour _u_ with a surplus _σ_, make an arc with cost 0 and capacity equal to
_σ_. Similarly, from any harbour _w_ with a need _ν_, make an arc to _t_, with cost 0 and capacity
equal to _ν_.
Now a maximum flow of minimum cost corresponds to an optimum routing of ships
between harbours.
A similar model applies to the problem of routing empty box cars in a railway network
(Feeney [1957], cf. Norman and Dowling [1968], White and Bomberault [1969]).

**Application 4.5: Routing of railway stock.** NS (Nederlandse Spoorwegen = Dutch
Railways) performs a daily schedule on its line Amsterdam–Vlissingen, with the following
(weekday) timetable:

```
ride number 2123 2127 2131 2135 2139 2143 2147 2151 2155 2159 2163 2167 2171 2175 2179 2183 2187 2191
Amsterdam d 6.48 7.55 8.56 9.56 10.56 11.56 12.56 13.56 14.56 15.56 16.56 17.56 18.56 19.56 20.56 21.56 22.56
Rotterdam a 7.55 8.58 9.58 10.58 11.58 12.58 13.58 14.58 15.58 16.58 17.58 18.58 19.58 20.58 21.58 22.58 23.58
Rotterdam d 7.00 8.01 9.02 10.03 11.02 12.03 13.02 14.02 15.02 16.00 17.01 18.01 19.02 20.02 21.02 22.02 23.02
Roosendaal a 7.40 8.41 9.41 10.43 11.41 12.41 13.41 14.41 15.41 16.43 17.43 18.42 19.41 20.41 21.41 22.41 23.54
Roosendaal d 7.43 8.43 9.43 10.45 11.43 12.43 13.43 14.43 15.43 16.45 17.45 18.44 19.43 20.43 21.43
Vlissingen a 8.38 9.38 10.38 11.38 12.38 13.38 14.38 15.38 16.38 17.40 18.40 19.39 20.38 21.38 22.38
ride number 2108 2112 2116 2120 2124 2128 2132 2136 2140 2144 2148 2152 2156 2160 2164 2168 2172 2176
Vlissingen d 5.30 6.54 7.56 8.56 9.56 10.56 11.56 12.56 13.56 14.56 15.56 16.56 17.56 18.56 19.55
Roosendaal a 6.35 7.48 8.50 9.50 10.50 11.50 12.50 13.50 14.50 15.50 16.50 17.50 18.50 19.50 20.49
Roosendaal d 5.29 6.43 7.52 8.53 9.53 10.53 11.53 12.53 13.53 14.53 15.53 16.53 17.53 18.53 19.53 20.52 21.53
Rotterdam a 6.28 7.26 8.32 9.32 10.32 11.32 12.32 13.32 14.32 15.32 16.32 17.33 18.32 19.32 20.32 21.30 22.32
Rotterdam d 5.31 6.29 7.32 8.35 9.34 10.34 11.34 12.34 13.35 14.35 15.34 16.34 17.35 18.34 19.34 20.35 21.32 22.34
Amsterdam a 6.39 7.38 8.38 9.40 10.38 11.38 12.38 13.38 14.38 15.38 16.40 17.38 18.38 19.38 20.38 21.38 22.38 23.38
```

The rides are carried out by one type of stock, that consists of two-way units that can
be coupled with each other. The length of the trains can be changed at the end stations
and at two intermediate stations: Rotterdam and Roosendaal. So in this example, each
train ride consists of three ride ‘segments’.
Based on the expected number of passengers, NS determines for each ride segment a
minimum number of units that should be deployed for that segment:

```
ride number 2123 2127 2131 2135 2139 2143 2147 2151 2155 2159 2163 2167 2171 2175 2179 2183 2187 2191
Amsterdam-Rotterdam 3 5 4 3 3 3 3 3 3 4 5 5 3 2 2 2 1
Rotterdam-Roosendaal 2 3 4 4 2 3 3 3 3 4 5 5 4 2 2 2 1
Roosendaal-Vlissingen 3 2 2 2 2 3 2 3 3 3 4 4 3 2 1
ride number 2108 2112 2116 2120 2124 2128 2132 2136 2140 2144 2148 2152 2156 2160 2164 2168 2172 2176
Vlissingen-Roosendaal 2 4 4 4 2 2 3 2 2 2 3 3 2 2 2
Roosendaal-Rotterdam 2 4 5 4 5 3 3 3 2 3 3 4 3 2 2 2 2
Rotterdam-Amsterdam 1 3 5 4 4 5 3 3 3 3 3 4 5 3 2 2 2 2
```

```
17
```

```
19
```

```
20
```

```
22
```

```
24
```

```
1
```

```
10
```

```
6
```

```
7
```

```
8
```

```
9
```

```
11
12
13
14
```

```
15
```

```
16
```

```
18
```

```
21
```

```
23
```

```
2
```

```
3
```

```
4
```

```
5
```

```
Roosendaal
```

```
Amsterdam
```

```
Vlissingen
```

```
Rotterdam
```

```
Figure 4.4
```

A unit uncoupled from a train at a station can be coupled at any other later train, in
the same direction or the other. Moreover, for each segment there is a maximum number

of units given that can be used for that segment (depending for instance on the length of
station platforms).
The company now wishes to find the minimum number of units that should be used to
run the schedule (excluding maintenance).
As was observed by Bartlett [1957] (cf. van Rees [1965]) this problem can be considered
as a minimum-cost circulation problem (cf. Figure 4.4). Make a directed graph _D_ with
vertex set all pairs (_s,t_) where _s_ is any station where the train composition can be changed
(in our example: the end stations and the two intermediate stations) and _t_ is any time at
which there is a train arriving at or leaving _s_. For each ride segment make an arc from (_s,t_)
to (_s_′_,t_′) if the segment leaves _s_ at time _t_ and arrives at _s_′ at time _t_′.

Moreover, for each station _s_ and each two consecutive times _t,t_′ at which segments
arrive or leave, one makes an arc from (_s,t_) to (_s,t_′). One also does this overnight.
Now for each arc _a_ coming from a segment assign a lower bound _d_(_a_) equal to the
number given in the table above for the segment. Moreover, define an upper bound _c_(_a_)
equal to the maximum number of units that can be used for that segment. For each arc _a_
from (_s,t_) to (_s,t_′) let _d_(_a_) := 0 and _c_(_a_) := ∞.
For each arc _a_ define a cost _k_(_a_) := 0, except if _a_ corresponds to an overnight stay at
one of cities, when _k_(_a_) := 1. Then a minimum-cost circulation corresponds to a routing of
the stock using a minimum number of units.
There are several variations possible. Instead of an upper bound _c_(_a_) = ∞ for the arcs _a_
from (_c,t_) to (_s,t_′) one can take _c_(_a_) equal to the capacity of the storage area at _s_. Instead
of a cost _k_(_a_) = 0 at each segment one can take _k_(_a_) equal to the cost of riding one unit of
stock over that segment. This can be weighed against the cost of buying extra units.
A similar model for routing airplanes was considered by Ferguson and Dantzig [1955].

**Exercises**

4.16. Determine in the following networks a maximum _s_ − _t_ flow of minimum-cost (_cost_ in
_italics_, **capacity** in **bold**):

```
(i) s t
```

```
5
```

```
3
```

```
1
```

```
2
```

```
3
```

```
2
```

```
5
```

```
6
```

```
7
```

```
8
```

```
3
```

```
5
```

```
22
```

```
6
```

```
5
```

```
10
```

```
18
```

```
4
```

```
4
```

```
7
```

```
(ii) s t
```

```
3
```

```
8
```

```
9
```

```
7 3 1
```

```
7
```

```
2
```

```
4
5
```

```
3
```

```
1
```

```
4 1
```

```
2
```

```
2
```

```
20
```

```
6
```

```
2
```

```
3 2
```

```
3
```

```
3
```

```
6
```

```
40
```

```
2
```

```
1 1
```

```
30
```

```
5
```

```
(iii) s t
```

```
2
```

```
4
```

```
1
```

```
1
```

```
3
```

```
4 1
```

```
3
```

```
8
```

```
1
```

```
5
```

```
3 2 4
```

```
7 8
```

```
3
```

```
8
```

```
6
```

```
1 1
```

```
5
```

```
2
```

```
7
```

```
6
```

```
8
```

```
1
```

```
1
```

```
1
```

```
9
```

```
2
```

```
8
```

```
2
```

```
7
```

4.17. Solve the minimum-cost transportation problem for the following data sets:

```
(i) m = n = 3,s1 = 9,s2 = 15,s3 = 7,d1 = 5,d2 = 13,d3 = 7,k1 = 2,k2 = 3,k3 =
2,
```

```
ci,j j = 1 j = 2 j = 3
i = 1 6 4 0
i = 2 3 9 4
i = 3 0 2 6
```

```
ki,j j = 1 j = 2 j = 3
i = 1 8 3 5
i = 2 2 7 1
i = 3 2 5 9
```

```
(ii) m = n = 3,s1 = 11,s2 = 7,s3 = 6,d1 = 9,d2 = 7,d3 = 5,k1 = 4,k2 = 3,k3 = 3,
```

```
ci,j j = 1 j = 2 j = 3
i = 1 7 4 0
i = 2 3 3 2
i = 3 0 2 4
```

```
ki,j j = 1 j = 2 j = 3
i = 1 3 2 4
i = 2 2 8 4
i = 3 1 3 2
```

4.18. Describe the problem of finding a maximum-weight matching in a bipartite graph as
a minimum-cost flow problem.

4.19. Reduce the problem of finding an extreme flow of given value, to the min-cost flow
problem as described above.

## 5. Nonbipartite matching

### 5.1. Tutte’s 1-factor theorem and the Tutte-Berge

### formula

A basic result on matchings in arbitrary (not necessarily bipartite) graphs was found
by Tutte [1947]. It characterizes graphs that have a perfect matching. A _perfect_
_matching_ (or a 1−_factor_) is a matching _M_ that covers all vertices of the graph. (So
_M_ partitions the vertex set of _G_.)
Berge [1958] observed that Tutte’s theorem implies a min-max formula for the
maximum size of a matching in a graph, the Tutte-Berge formula, which we prove
first.
Call a component of a graph _odd_ if it has an odd number of vertices. For any
graph _G_, define

(1) _o_(_G_) := number of odd components of _G_.

Let _ν_(_G_) denotes the maximum size of a matching. For any graph _G_ = (_V, E_) and
_U_ ⊆ _V_ , the graph obtained by deleting all vertices in _U_ and all edges incident with
_U_, is denoted by _G_ − _U_.
Then:

**Theorem 5.1** (Tutte-Berge formula)**.** _For each graph G_ = (_V, E_)_,_

(2) _ν_(_G_) = min
_U_⊆_V_

```
1
2(|V | + |U|− o(G − U)).
```

**Proof.** To see ≤, we have for each _U_ ⊆ _V_ :

(3) _ν_(_G_) ≤|_U_|+_ν_(_G_−_U_) ≤|_U_|+[^1]
[^2]:(|_V_ \_U_|−_o_(_G_−_U_)) =[^1]
[^2]:(|_V_ |+|_U_|−_o_(_G_−_U_))_._

We prove the reverse inequality by induction on |_V_ |, the case _V_ = ∅ being trivial.
We can assume that _G_ is connected, since otherwise we can apply induction to the
components of _G_.
First assume that there exists a vertex _v_ covered by all maximum-size matchings.
Then _ν_(_G_ − _v_) = _ν_(_G_) − 1, and by induction there exists a subset _U_′ of _V_ \{_v_} with

(4) _ν_(_G_ − _v_) =[^1]
[^2]:(|_V_ \{_v_}| + |_U_′|− _o_(_G_ − _v_ − _U_′)).

Then _U_ := _U_′ ∪{_v_} gives equality in (2), since

(5) _ν_(_G_) = _ν_(_G_ − _v_) + 1 =[^1]
[^2]:(|_V_ \{_v_}| + |_U_′|− _o_(_G_ − _v_ − _U_′)) + 1
=[^1]
[^2]:(|_V_ | + |_U_|− _o_(_G_ − _U_)).

So we can assume that there is no such _v_. In particular, _ν_(_G_) _<_[^1]
[^2]:|_V_ |. We show
that there exists a matching of size[^1]
[^2]:(|_V_ |− 1), which implies the theorem (taking
_U_ := ∅).
Indeed, suppose to the contrary that each maximum-size matching _M_ misses at
least two distinct vertices _u_ and _v_. Among all such _M, u, v_, choose them such that
the distance dist(_u, v_) of _u_ and _v_ in _G_ is as small as possible.
If dist(_u, v_) = 1, then _u_ and _v_ are adjacent, and hence we can augment _M_ by the
edge _uv_, contradicting the maximality of |_M_|. So dist(_u, v_) ≥ 2, and hence we can
choose an intermediate vertex _t_ on a shortest _u_−_v_ path. By assumption, there exists
a maximum-size matching _N_ missing _t_. Choose such an _N_ with |_M_ ∩ _N_| maximal.
By the minimality of dist(_u, v_), _N_ covers both _u_ and _v_. Hence, as _M_ and _N_ cover
the same number of vertices, there exists a vertex _x_ [^6]:= _t_ covered by _M_ but not by _N_.
Let _x_ ∈ _e_ = _xy_ ∈ _M_. Then _y_ is covered by some edge _f_ ∈ _N_, since otherwise _N_ ∪{_e_}
would be a matching larger than _N_. Replacing _N_ by (_N_ \{_f_}) ∪{_e_} would increase
its intersection with _M_, contradicting the choice of _N_.

(This proof is based on the proof of Lov ́asz [1979] of Edmonds’ matching polytope
theorem.)
The Tutte-Berge formula immediately implies Tutte’s 1-factor theorem.

**Corollary 5.1a** (Tutte’s 1-factor theorem)**.** _A graph G_ = (_V, E_) _has a perfect match-_
_ing if and only if G_ − _U has at most_ |_U_| _odd components, for each U_ ⊆ _V._

**Proof.** Directly from the Tutte-Berge formula (Theorem 5.1), since _G_ has a perfect
matching if and only if _ν_(_G_) ≥[^1]
[^2]:|_V_ |.

In the following sections we will show how to find a maximum-size matching
algorithmically.
With Gallai’s theorem, the Tutte-Berge formula implies a formula for the edge
cover number _ρ_(_G_), where _o_(_U_) denotes the number of odd components of the sub-
graph _G_[_U_] of _G_ induced by _U_ (meaning that _G_[_U_] = (_U,_{_e_ ∈ _E_ | _e_ ⊆ _U_})):

**Corollary 5.1b.** _Let G_ = (_V, E_) _be a graph without isolated vertices. Then_

(6) _ρ_(_G_) = max
_U_⊆_V_

```
|U| + o(U)
2
```

_._

**Proof.** By Gallai’s theorem (Theorem 3.1) and the Tutte-Berge formula (Theorem
5.1),

(7) _ρ_(_G_) = |_V_ |− _ν_(_G_) = |_V_ |− min
_W_⊆_V_

```
|V | + |W|− o(V \ W)
2
```

```
= max
U⊆V
```

```
|U| + o(U)
2
```

_._

**Exercises**

```
5.1. (i) Show that a tree has at most one perfect matching.
(ii) Show (not using Tutte’s 1-factor theorem) that a tree G = (V,E) has a perfect
matching if and only if the subgraph G −v has exactly one odd component, for
each v ∈ V.
```

```
5.2. Let G be a 3-regular graph without any bridge. Show that G has a perfect matching.
(A bridge is an edge e not contained in any circuit; equivalently, deleting e increases
the number of components; equivalently, {e} is a cut.)
```

```
5.3. Let A1,...,An be a collection of nonempty subsets of the finite set X so that each
element in X is in exactly two sets among A1,...,An. Show that there exists a set
Y intersecting all sets A1,...,An, and satisfying |Y |≤ t if and only if for each subset
I of {1,...,n} the number of components of (Ai | i ∈ I) containing an odd number
of sets in (Ai | i ∈ I) is at most 2t −|I|.
(Here a subset Y of X is called a component of (Ai | i ∈ I) if it is a minimal nonempty
subset of X with the property that for each i ∈ I: Ai ∩ Y = ∅ or Ai ⊆ Y .)
5.4. Let G = (V,E) be a graph and let T be a subset of V. Then G has a matching
covering T if and only if the number of odd components of G − W contained in T is
at most |W|, for each W ⊆ V.
```

```
5.5. Let G = (V,E) be a graph and let b : V → Z+. Show that there exists a function
f : E → Z+ so that for each v ∈ V :
```

```
(8)
```

```
∑
```

```
e∈E,v∈e
```

```
f(e) = b(v)
```

```
if and only if for each subset W of V the number β(W) is at most b(V \ W).
(Here for any subset W of V , b(W) :=
```

```
∑
v∈W b(v). Moreover, β(W) denotes the
following. Let U be the set of isolated vertices in the graph G|W induced by W and
let t denote the number of components C of the graph G|W \U with b(C) odd. Then
β(W) := b(U) + t.)
```

```
5.6. Let G = (V,E) be a graph and let b : V → Z+. Show that G has a subgraph
G′ = (V,E′) such that degG′(v) = b(v) for each v ∈ V if and only if for each two
disjoint subsets U and W of V one has
```

```
(9)
```

```
∑
```

```
v∈U
```

```
b(v) ≥ q(U,W) +
```

```
∑
```

```
v∈W
```

```
(b(v) − dG−U(v)).
```

```
Here q(U,W) denotes the number of components K of G − (U ∪ W) for which b(K)
plus the number of edges connecting K and W, is odd. Moreover, dG−U(v) is the
degree of v in the subgraph induced by V \ U.
```

### 5.2. Cardinality matching algorithm

We now investigate the problem of finding a maximum-cardinality matching algorith-
mically. Like in the bipartite case, the key is to find an augmenting path. However,
the idea for bipartite graphs to orient the edges using the two colour classes, does not
apply to nonbipartite graphs.
Yet one could try to find an _M_-augmenting path by finding a so-called _M_-
alternating walk, but such a path can run into a loop that cannot immediately be
deleted. It was J. Edmonds who found the trick to resolve this problem, namely
by ‘shrinking’ the loop (which he called a ‘blossom’). Then applying recursion to a
smaller graph solves the problem.
We first describe the operation of _shrinking_. Let _X_ and _Y_ be sets. Then we define
_X/Y_ as follows:

(10) _X/Y_ := _X_ if _X_ ∩ _Y_ = ∅,
_X/Y_ := (_X_ \ _Y_ ) ∪{_Y_ } if _X_ ∩ _Y_ [^6]:= ∅.

So if _G_ = (_V, E_) is a graph and _C_ ⊆ _V_ , then _V /C_ arises from _V_ by deleting all
vertices in _C_, and adding one new vertex called _C_. For any edge _e_ of _G_, _e/C_ = _e_ if
_e_ is disjoint from _C_, while _e/C_ = _uC_ if _e_ = _uv_ with _u_ 6∈ _C_, _v_ ∈ _C_. (If _e_ = _uv_ with
_u, v_ ∈ _C_, then _e/C_ is a loop _CC_; they can be neglected in the context of matchings.)
Then for any _F_ ⊆ _E_:

(11) _F/C_ := {_f /C_ | _f_ ∈ _F_}_._

So _G/C_ := (_V /C, E/C_) is again a graph. We say that _G/C_ arises from _G_ by _shrinking_
_C_.
Let _G_ = (_V, E_) be a graph and let _M_ be a matching in _G_, and let _W_ be the set of
vertices missed by _M_. A walk _P_ = (_v_[^0]:_, v_[^1]:_,... , vt_) is called _M-alternating_ if for each
_i_ = 1_,... , t_ − 1 exactly one of _vi_−[^1]:_vi_ and _vivi_+1 belongs to _M_. Note that one can find
a shortest _M_-alternating _W_ −_W_ walk of positive length, by considering the auxiliary
directed graph _D_ = (_V, A_) with

(12) _A_ := {(_w, w_′) |∃_x_ ∈ _V_ : {_w, x_}∈ _E,_{_x, w_′}∈ _M_}_._

Then _M_-alternating _W_ − _W_ walks correspond to directed walks in _D_ from a vertex
in _W_ to a vertex that is adjacent to at least one vertex in _W_.

```
So an M-augmenting path is an M-alternating W − W walk of positive length,
in which all vertices are distinct. By Theorem 3.2, a matching M has maximum size
if and only if there is no M-augmenting path. We call an M-alternating walk P an
M-blossom if v0,... , vt−1 are distinct, v0 is missed by M, and vt = v0.
The core of the algorithm is the following observation.
```

```
Theorem 5.2. Let C be an M-blossom in G. Then M has maximum size in G if
and only if M/C has maximum size in G/C.
```

```
Proof. Let C = (v0, v1,... , vt), G′ := G/C and M′ := M/C.
First let P be an M-augmenting path in G. We may assume that P does not start
at v0 (otherwise we can reverse P). If P does not traverse any vertex in C, then P
is also M′-augmenting in G′. If P does traverse a vertex in C, we can decompose P
as P = QR, where Q ends at a vertex in C, and no other vertex on Q belongs to C.
Then by replacing the last vertex of Q by C makes Q to an M′-augmenting path in
G′.
Conversely, let P′ be an M′-augmenting path in G′. If P′ does not traverse vertex
C of G′, then P′ is also an M-augmenting path in G. If P′ traverses vertex C of G′,
we may assume it ends at C (as C is missed by M′). So we can replace C in P′ by
some vertex vi ∈ C to obtain a path Q in G ending at vi. If i is odd, extending Q
by vi+1,... , vt−1, vt gives an M-augmenting path in G. If i is even, extending Q by
vi−1,... , v1, v0 gives an M-augmenting path in G.
```

```
Another useful observation is (where a W − v walk is a walk starting at a vertex
in W and ending at v):
```

```
Theorem 5.3. Let P = (v0, v1,... , vt) be a shortest even-length M-alternating W −v
walk. Then either P is simple or there exist i < j such that vi = vj, i is even, j is
odd, and v0,... , vj−1 are all distinct.
```

```
Proof. Assume P is not simple. Choose i < j such that vj = vi and such that j is as
small as possible. If j − i is even, we can delete vi+1,... , vj from P so as to obtain
a shorter M-alternating W − v walk. So j − i is odd. If j is even and i is odd, then
vi+1 = vj−1 (as it is the vertex matched to vi = vj), contradicting the minimality of
j.
```

```
We now describe an algorithm for the following problem:
```

(13) given: a matching _M_;
find: a matching _N_ with |_N_| = |_M_| + 1 or conclude that _M_ is a maximum-size
matching.

```
Let W be the set of vertices missed by M.
```

(14) **Case 1.** _There is no M-alternating W_ − _W walk._ Then _M_ has maximum size
(as there is no _M_-augmenting path).

```
Case 2. There is an M-alternating W − W walk. Let P = (v0, v1,... , vt) be a
shortest such walk.
Case 2a. P is path. Hence P is an M-augmenting path. Then output N :=
M△EP.
Case 2b. P is not a path. That is, not all vertices in P are different. Choose
i < j such that vi = vj with j as small as possible. Reset M :=
M△{v0v1, v1v2,... , vi−1vi}. Then C := (vi, vi+1,... , vj) is an M-blossom.
Apply the algorithm (recursively) to G′ = G/C and M′ := M/C.
```

• If it gives an _M_′-augmenting path _P_′ in _G_′, transform _P_′ to an
_M_-augmenting path in _G_ (as in the proof of Theorem 5.2).
• If it concludes that _M_′ has maximum size in _G_′, then _M_ has max-
imum size in _G_ (by Theorem 5.2).

```
This gives a polynomial-time algorithm to find a maximum-size matching, which
is a basic result of Edmonds [1965c].
```

```
Theorem 5.4. Given an undirected graph, a maximum-size matching can be found
in time O(|V |[^2]|E|).
```

```
Proof. The algorithm directly follows from algorithm (14), since one can iteratively
apply it, starting with M = ∅, until a maximum-size matching is attained.
By using (12), a shortest M-alternating W −W walk can be found in time O(|E|).
Moreover, the graph G/C can be constructed in time O(|E|). Since the recursion has
depth at most |V |, each application of algorithm (14) takes O(|V ||E|) time. Since the
number of applications is at most |V |, we have the time bound given in the theorem.
```

```
In fact, the method can be sharpened to O(|V |[^3]) (Balinski [1969]), O(|V |[^5]/[^2]) (Even
and Kariv [1975]) and even to O(|V |[^1]/[^2]|E|) (Micali and Vazirani [1980]). For surveys,
see Schrijver [2003].
```

```
Application 5.1: Pairing. If a certain group of people has to be split into pairs, where
certain pairs fit and other pairs do not fit (for instance, when assigning hotel rooms or bus
seats to a touring group), we have an example of a (perfect) matching problem.
```

```
Application 5.2: Two-processor scheduling. Suppose we have to carry out certain
jobs, where some of the jobs have to be done before other. We can represent this by a
partially ordered set (X,≤) where X is the set of jobs and x < y indicates that job x has
to be done before job y. Each job takes one time-unit, say one hour.
Suppose now that there are two workers, each of which can do one job at a time.
Alternatively, suppose that you have one machine, that can do at each moment two jobs
```

simultaneously (such a machine is called a _two-processor_).
We wish to do all jobs within a minimum total time span. This problem can be solved
with the matching algorithm as follows. Make a graph _G_ = (_X,E_), with vertex set _X_ (the
set of jobs) and with edge set

(15) _E_ := {{_u,v_}| _u_ 6≤ _v_ and _v_ 6≤ _u_}.

(So (_X,E_) is the complementary graph of the ‘comparability graph’ associated with (_X,_≤).)
Consider now a possible schedule of the jobs. That is, we have a sequence _p_[^1]:_,...,pt_,
where each _pi_ is either a singleton vertex or an edge of _G_ so that _p_[^1]:_,...,pt_ partition _X_ and
so that if _x_ ∈ _pi_ and _y_ ∈ _pj_ and _x < y_ then _i < j_.[^16]
Now the pairs in this list should form a matching _M_ in _G_. Hence _t_ = |_X_|−|_M_|. In
particular, _t_ cannot be smaller than |_X_|− _ν_(_G_), where _ν_(_G_) is the matching number of _G_.
Now it can be shown that in fact one can always make a schedule with _t_ = |_X_|− _ν_(_G_).
To this end, let _Q_ be a minimum partition of _V_ into vertices and edges of _G_, and let _Y_ be
the set of minimal elements of _X_. If _q_ ⊆ _Y_ for some _q_ ∈ _Q_, we can replace _X_ by _X_ \ _q_ and
_Q_ by _Q_ \{_q_}, and apply induction.
So we may assume that each _y_ ∈ _Y_ is contained in an edge _yz_ ∈ _Q_ with _z_ 6∈ _Y_. Choose
an edge _yz_ ∈ _Q_ such that _y_ ∈ _Y_ and such that the height of _z_ is as small as possible. (The
_height_ of an element _z_ is the maximum size of a chain in (_X,_≤) with maximum element _z_.)
As _z_ 6∈ _Y_ there exists an _y_′_z_′ ∈ _Q_ with _y_′ ∈ _Y_ and _y_′ _< z_.
Now clearly _yy_′ is an edge of _G_, as _y_ and _y_′ are minimal elements. Moreover, _zz_′ is an
edge of _G_. For if _z < z_′ then _y_′ _< z < z_′, contradicting the fact that _y_′_z_′ ∈ _EG_; and if
_z_′ _< z_ than _z_′ would have smaller height than _z_.
So replacing _yz_ and _y_′_z_′ in _Q_ by _yy_′ and _zz_′, we have _yy_′ ⊆ _Y_ , and we can apply
induction as before.

**Exercises**

```
5.7. Apply the matching augmenting algorithm to the matchings in the following graphs:
```

```
(i)
```

```
(ii)
```

```
16Here we identify a vertex v with the set {v}.
```

```
(iii)
```

### 5.3. Weighted matching algorithm

```
Edmonds [1965a] proved that also the maximum-weight matching problem can be
solved in polynomial time. Equivalently, the minimum-weight perfect matching prob-
lem can be solved in polynomial time. It is the problem:
```

(16) given: a graph _G_ = (_V, E_) and a ‘weight’ function _w_ : _E_ → Q;
find: a perfect matching _M_ minimizing

```
∑
e∈M w(e).
```

```
We describe the algorithm, assuming without loss of generality that G has at least
one perfect matching and that w(e) ≥ 0 for each edge e (we can add a constant to
all edge weights without changing the problem).
Like the cardinality matching algorithm, the weighted matching algorithm is based
on shrinking sets of vertices. Unlike the cardinality matching algorithm however, for
weighted matchings one has to ‘deshrink’ sets of vertices (the reverse operation of
shrinking). Thus we have to keep track of the shrinking history throughout the
iterations.
The algorithm is ‘primal-dual’. The ‘vehicle’ carrying us to a minimum-weight
perfect matching is a pair of a nested[^17] collection Ω of odd-size subsets of V , and a
function π : Ω → Q satisfying:
```

```
(17) (i) π(U) ≥ 0 if U ∈ Ω with |U|≥ 3,
(ii)
```

```
∑
```

```
U∈Ω
e∈δ(U)
```

```
π(U) ≤ w(e) for each e ∈ E.
```

```
This implies that for each perfect matching N in G one has w(N) ≥
```

```
∑
```

```
U∈Ω
```

```
π(U), since
```

```
(18) w(N) =
```

```
∑
```

```
e∈N
```

```
w(e) ≥
```

```
∑
```

```
e∈N
```

```
∑
```

```
U∈Ω
e∈δ(U)
```

```
π(U) =
```

```
∑
```

```
U∈Ω
```

```
π(U)|N ∩ δ(U)|≥
```

```
∑
```

```
U∈Ω
```

```
π(U).
```

```
Notation and assumptions. Let be given Ω and π : Ω → Q. Define
17A collection Ω of subsets of a set V is called nested if U ∩ W = ∅ or U ⊆ W or W ⊆ U for any
U,W ∈ Ω.
```

(19) _wπ_(_e_) := _w_(_e_) −
_U_∈Ω
_e_∈_δ_(_U_)

```
π(U)
```

for any edge _e_ ∈ _E_. (So (17)(ii) implies _wπ_(_e_) ≥ 0.)
_G/_Ω denotes the graph obtained from _G_ by shrinking all sets in Ωmax, the set of
inclusionwise maximal sets in Ω. We will assume throughout that {_v_} ∈ Ω for each
_v_ ∈ _V_. Hence, as Ω is nested and covers _V_ , Ωmax is a partition of _V_.
When shrinking a set _U_ ∈ Ω, we denote the new vertex representing the shrunk
set _U_ just by _U_. So _G/_Ω has vertices the sets in Ωmax, with two distinct elements
_U, U_′ ∈ Ωmax adjacent if and only if _G_ has an edge connecting _U_ and _U_′. We denote
any edge of _G/_Ω by the original edge in _G_.
Throughout we restrict ourselves to Ω and _π_ satisfying:

(20) for each _U_ ∈ Ω with |_U_|≥ 3, the graph obtained from _G_|_U_ by shrinking all
inclusionwise maximal proper subsets of _U_ that are in Ω, has a Hamiltonian
circuit _CU_ of edges _e_ with _wπ_(_e_) = 0.

```
X
```

```
edge in M
edge not in M
```

```
vertex covered by M
vertex not covered by M
Figure 5.1. An M-alternating forest
```

_M_**-alternating forests.** An important role in the algorithm is played by a so-called
‘_M_-alternating forest’ relative to a matching _M_ (cf. Figure 5.1).
Let _M_ be a matching in a graph _G_ = (_V, E_) and let _W_ be the set of vertices
missed by _M_. Then a subset _F_ of _E_ is an _M-alternating forest_ in _G_ if _F_ is a forest
containing _M_ such that each component of (_V, F_) consists either of an edge in _M_

or contains exactly one vertex in _W_ and such that each path in _F_ starting in _W_ is
_M_-alternating.
The set of vertices _v_ ∈ _V_ for which there exists an even-length (odd-length,
respectively) _W_ − _v_ path in _F_ is denoted by even(_F_) (odd(_F_), respectively).
**The algorithm.** We iterate with Ω and _π_ : Ω → Q satisfying (17) and (20), a

matching _M_ in _G/_Ω and an _M_-alternating forest _F_ in _G/_Ω with _wπ_(_F_) = 0.
Initially, we set _M_ := ∅, _F_ := ∅, Ω := {{_v_} | _v_ ∈ _V_ }, and _π_({_v_}) := 0 for
each _v_ ∈ _V_. Then, as long as _M_ is not a perfect matching in _G/_Ω, we perform the
following iteratively:

(21) Reset _π_(_U_) := _π_(_U_) − _α_ for _U_ ∈ odd(_F_) and _π_(_U_) := _π_(_U_) + _α_ for _U_ ∈
even(_F_), where _α_ is the largest value such that (17) is maintained. After
that
(i) there exists an edge _e_ of _G/_Ω with _wπ_(_e_) = 0 such that _e_
intersects even(_F_) but not odd(_F_),
or (ii) there exists a _U_ ∈ odd(_F_) with |_U_|≥ 3 and _π_(_U_) = 0.
First assume (i) holds. If only one end of _e_ belongs to even(_F_), extend _F_
by _e_. If both ends of _e_ belong to even(_F_) and _F_ ∪{_e_} contains a circuit_C_,
let _U_ := _V C_ and _CU_ := _C_, add _U_ to Ω (defining _π_(_U_) := 0), and replace _F_
by _F/U_ and _M_ by _M/U_. If both ends of _e_ belong to even(_F_) and _F_ ∪{_e_}
contains an _M_-augmenting path, augment _M_ and reset _F_ := _M_.
Next assume (ii) holds. Delete _U_ from Ω, replace _F_ by _F_ ∪ _P_ ∪ _N_ and
_M_ by _M_ ∪ _N_, where _P_ is the even-length path in _CU_ connecting the two
edges of _F_ incident with _U_ and where _N_ is the matching in _CU_ covering all
vertices in _U_ that are not covered by _M_.

(Note that in this iteration _α_ is bounded, since

∑
_U_∈Ω _π_(_U_) is bounded (by (18), as
there is at least one perfect matching), and since |even(_F_)| _>_ |odd(_F_)| (as _M_ is not
perfect).)
If _M_ is a perfect matching in _G/_Ω, we are done: by (20) we can expand _M_ to a
perfect matching _N_ in _G_ with _wπ_(_N_) = 0 and |_N_ ∩_δ_(_U_)| = 1 for each _U_ ∈ Ω; then _N_
has equality throughout in (18), and hence it is a minimum-weight perfect matching.

**Theorem 5.5.** _There are at most_ |_V_ |[^2] _iterations._

**Proof.** In any iteration where we augment _M_, the value of |_V_ (_G/_Ω)|−[^2]:|_M_| decreases
by 2. If there is no matching augmentation, this value remains invariant. So there
are at most[^1]
[^2]:|_V_ | matching augmentations.
Let _V_even be the set of vertices _v_ ∈ _V_ that are shrunk to a vertex in even(_F_).
Let Ω[^0]: be the set of vertices of _G/_Ω that do not belong to even(_F_). Then in any
iteration with no matching augmentation, 2|_V_even| + |Ω[^0]:| increases. Since this value
cannot exceed 2|_V_ |, between any two matching augmentations there are at most 2|_V_ |

iterations.

```
This gives the theorem of Edmonds [1965a]:
```

**Corollary 5.5a.** _A minimum-weight perfect matching can be found in polynomial_
_time._

**Proof.** The nestedness of Ω implies that |Ω|≤ [^2]:|_V_ | (which is an easy exercise — see
Exercise 5.10). Hence each iteration can be performed in polynomial time. With any

_U_ ∈ Ω with |_U_|≥ 3 we should keep the Hamiltonian circuit _CU_ of (20) — which we
had obtained earlier when shrinking _U_.

```
As a consequence one can derive:
```

**Corollary 5.5b.** _In any graph with weight function on the edges, a maximum-weight_
_matching can be found in polynomial time._

**Proof.** Left to the reader. (Exercise 5.9.)

The above algorithm can be implemented in time _O_(|_V_ |[^3]), which is a result of
Gabow [1973] and Lawler [1976]. Faster algorithms were given by Galil, Micali, and
Gabow [1986] (_O_(|_E_||_V_ |log |_V_ |)) and Gabow [1990] (_O_(|_V_ ||_E_| + |_V_ |[^2] log |_V_ |)).
For more about matchings we refer to the book of Lov ́asz and Plummer [1986].

**Application 5.3: Optimal pairing.** In several practical situations one has to find an
‘optimal pairing’, for example, when scheduling crews for airplanes. Also if one has to
assign bus seats optimally to the participants of an organized tour, or to accommodate the
participants most satisfactorily in two-bed hotel rooms, one has to solve a maximum-weight
perfect matching problem.

**Application 5.4: Airline timetabling.** A European airline company has for its European
flights a number of airplanes available. Each plane can make on any day two return flights to
European destinations (not necessarily the same destinations). The profit one makes on any
flight depends on the departure and arrival times of the flight (also due to intercontinental
connections). The company wants to make a timetable so that it can be performed by
the available fleet and so that the total profit is maximized. Assume that the number of
destinations to be reached is equal to twice the number of airplanes available.
To solve this problem, consider the complete graph with vertex set all possible destina-
tions. For each edge of this graph, connecting destinations _B_ and _C_ say, one calculates the
profit that will be made if one and the same air plane will make its flights to _B_ and _C_ (in
one order or the other). So one determines the optimum schedule for the flights to _B_ and _C_
so that the two return flights can be done by the same airplane and so that the total profit
on the two flights is maximized.
Now a timetable yielding maximum profit is found by determining a maximum-weight

```
perfect matching in this graph.
```

```
Application 5.5: Chinese postman problem. The Chinese postman problem, first
studied by Guan [1960], consists of the following. Given a connected graph G = (V,E) and
a length function l : E → Q+, find a minimum-length tour T that traverses each edge at
least once.
It is not difficult to see that if each vertex of G has an even degree, then the optimal
tour traverses each edge exactly once. But if the graph has vertices of odd degree, certain
edges have to be traversed more than once. To find such edges we can proceed as follows.
First determine the set U of vertices of odd degree. Note that |U| is even. For each pair
u,u′ of vertices in U determine the distance d(u,u′) between u and u′ in the graph G (taking
l as length). Consider the complete graph H = (U,E′) on U. Determine a minimum-weight
perfect matching M in H, taking d as weight function. For each edge uu′ in M we can
determine a path Pu,u′ in G of length d(u,u′). It can be shown that any two different such
paths do not have any edge in common (assuming that each edge has positive length) —
see Exercise 5.13. Let E ̃ be the set of edges occurring in the Pu,u′ (uu′ ∈ M). Then there
exists a tour T so that each edge e ∈ E \E ̃ is traversed exactly once and each edge e ∈ E ̃ is
traversed exactly twice. This tour T is a shortest ‘Chinese postman tour’ (Exercise 5.14).
```

```
Application 5.6: Christofides’ approximative algorithm for the traveling sales-
man problem. Christofides [1976] designed the following algorithm to find a short travel-
ing salesman tour in a graph (generally not the shortest however). The traveling salesman
problem is the problem, given a finite set V and a ‘length’ function l : V ×V → Q+, to find
a shortest traveling salesman tour. A traveling salesman tour (or Hamiltonian circuit) is a
circuit in the complete graph on V traversing each vertex exactly once.
Suppose that the length function is symmetric (that is, l(u,v) = l(v,u) for all u,v ∈ V )
and satisfies the triangle inequality:
```

```
(22) l(u,w) ≤ l(u,v) + l(v,w)
```

```
for all u,v,w ∈ V. Then a reasonably short traveling salesman tour can be found as follows.
First determine a shortest spanning tree S (with the greedy algorithm). Next, let U be
the set of vertices that have odd degree in S. Find a shortest perfect matching M on U,
taking l as weight function. Now ES ∪ M forms a set of edges such that each vertex has
even degree. (If an edge occurs both in ES and in M, we take it as two parallel edges.) So
we can make a cycle T such that each edge in ES ∪ M is traversed exactly once. Then T
traverses each vertex at least once. By inserting shortcuts we obtain a traveling salesman
tour T′ with length(T′) ≤length(T).
How far away is the length of T′ from the length of a shortest traveling salesman tour?
Let ρ be the length of a shortest traveling salesman tour. It is not difficult to show that:
```

(23) (i) length(_S_) ≤ _ρ_;
(ii) length(_M_) ≤[^1]
[^2]:_ρ_.

```
(Exercise 5.18.) Hence
```

(24) length(_T_′) ≤length(_T_) =length(_S_)+length(_M_) ≤[^3]
[^2]:_ρ_.

So the tour obtained with Christofides’ algorithm is not longer than[^3]
[^2]: times the optimal
tour.
The factor[^3]
[^2]: seems quite large, but it is the smallest factor for which a polynomial-time
method is known. Don’t forget moreover that it is a _worst-case_ bound, and that in practice
(or on average) the algorithm might have a much better performance.

**Exercises**

```
5.8. Find with the weighted matching algorithm a minimum-weight perfect matching in
the following weighted graphs:
```

```
(i)
```

```
1
```

```
1
```

```
1 3
```

```
4
```

```
2
```

```
7
```

```
5
```

```
6
```

```
(ii)
```

```
0
```

```
0
```

```
0
```

```
0
```

```
1
```

```
6
```

```
0 5
```

```
0
```

```
1
```

```
2
```

```
4
```

```
1
```

```
8
```

```
5
```

```
6
```

```
7
```

```
3
```

```
5.9. Derive Corollary 5.5b from Corollary 5.5a.
```

5.10. A collection Ω of subsets of a finite set _V_ is called _cross-free_ if:

```
(25) if X,Y ∈ Ω, then X ⊆ Y , or Y ⊆ X, or X ∩ Y = ∅, or X ∪ Y = V.
```

```
Show that if Ω is cross-free, then |Ω|≤ 4|V |.
```

5.11. Find a shortest Chinese postman route in the graph in Figure 5.2.

5.12. Find a shortest Chinese postman route in the map of Figure 5.3.

5.13. Show that the paths found in the algorithm for the Chinese postman problem pairwise
do not have any edge in common (if each edge has positive length).

```
2
```

```
6
```

```
5
```

```
9
```

```
4 3 1
3
2 1
```

```
8
3
3
```

```
3
6
```

```
2 3
```

```
3
```

```
3
4
5
3
4
3
```

```
4 5
```

```
3
```

```
5
4
1
2
```

```
2
```

```
Figure 5.2
```

5.14. Show that the tour found in Application 5.5 is indeed a shortest Chinese postman
tour.

5.15. Apply Christofides’ algorithm to the table in Exercise 1.8.

5.16. Let _G_ = (_V,E_) be a graph and let _T_ ⊆ _V_ with |_T_| even. Call a subset _F_ of _E_ a
_T-join_ if _T_ is equal to the set of vertices of odd degree in the graph (_V,F_).
Derive from Corollary 5.5a that a minimum-weight _T_-join can be found in polynomial
time.

5.17. Let _G_ = (_V,E_) be a graph and let _l_ : _E_ → Q be a length function such that each
circuit has nonnegative length. Let _s,t_ ∈ _V_.
Derive from the minimum-weight perfect matching algorithm an algorithm to find a
minimum-length _s_ − _t_ path in _G_.

5.18. Show (23).

### 5.4. The matching polytope

The weighted matching algorithm of Edmonds [1965a] gives as a side result a charac-
terization of the perfect matching polytope _P_perfect matching(_G_) of any graph _G_. This
is Edmonds’ matching polytope theorem.
The _perfect matching polytope_ of a graph _G_ = (_V, E_), denoted by _P_perfect matching(_G_),
is the convex hull of the incidence vectors of the perfect matchings in _G_.[^18] That is,
[^18]:For any finite set _X_ and any subset _Y_ of _X_, the _incidence vector_ or _incidence function_ of a
subset _Y_ of _X_ is the vector _χY_ ∈ R_X_ defined by: _χY_
_x_ := 1 if _x_ ∈ _Y_ and _χY_
_x_ := 0 otherwise.

```
Figure 5.3. Part of the Xuhui district of Shanghai
```

(26) _P_perfect matching(_G_) =conv.hull{_χM_ | _M_ perfect matching in _G_}.

So _P_perfect matching(_G_) is a polytope in R_E_.
In Section 3.6 we saw that for a bipartite graph _G_ = (_V, E_), the perfect matching
polytope is fully determined by the following set of inequalities:

(27) (i) _xe_ ≥ 0 for each _e_ ∈ _E_;
(ii)

```
∑
e∋v xe = 1 for each v ∈ V.
```

These inequalities are not enough for, say, _K_[^3]:: taking _x_(_e_) :=[^1]
[^2]: for each edge _e_ of _K_[^3]
gives a vector _x_ satisfying (27) but not belonging to the perfect matching polytope
of _K_[^3]:.
Edmonds [1965a] showed that it is enough to add the following set of inequalities:

(28)

```
∑
```

```
e∈δ(U)
```

```
xe ≥ 1 for each odd subset U of V.
```

It is clear that for any perfect matching _M_ in _G_ the incidence vector _χM_ satisfies
(28). So clearly, _P_perfect matching(_G_) is contained in the polyhedron _Q_ defined by (27)
and (28). The essence of Edmonds’ theorem is that one does not need more.
In order to show Edmonds’ theorem, we derive from Edmonds’ algorithm the
following theorem, where Podd(_V_ ) denotes the collection of odd subsets of _V_ :

**Theorem 5.6.** _Let G_ = (_V, E_) _be a graph and let w_ : _E_ → Q _be a ‘weight’ function._

_Then the minimum weight of a perfect matching is equal to the m_∑ _aximum value of_
_X_∈Podd(_V_) _π_(_X_) _where π ranges over all functions π_ : Podd(_V_ ) → Q _satisfying_ (17)_._

**Proof.** We may assume that _w_ is nonnegative: if _μ_ is the minimum value of _w_(_e_) over
all edges _e_, decreasing each _w_(_e_) by _μ_ decreases both the maximum and the minimum
by[^1]
[^2]:|_V_ |_μ_.
The fact that the minimum is not smaller than the maximum follows from (18).
Equality follows from the fact that in the algorithm the final perfect matching and
the final function _π_ have equality throughout in (18).

```
This implies:
```

**Corollary 5.6a** (Edmonds’ perfect matching polytope theorem)**.** _The perfect match-_
_ing polytope of any graph G_ = (_V, E_) _is determined by_ (27) _and_ (28)_._

**Proof.** By Theorem 5.6 and LP-duality, for any weight function _w_ ∈ Q_E_, the min-
imum weight of a perfect matching is equal to the minimum of _wTx_ taken over the
polytope determined by (27) and (28). Hence the two polytopes coincide, by Theorem
2.1.

From this one can derive Edmonds’ matching polytope theorem, characterizing
the _matching polytope_ of a graph _G_ = (_V, E_), denoted by _P_matching(_G_), which is the
convex hull of the incidence vectors of the matchings in _G_. That is,

(29) _P_matching(_G_) =conv.hull{_χM_ | _M_ matching in _G_}.

Again, _P_matching(_G_) is a polytope in R_E_.

**Corollary 5.6b** (Edmonds’ matching polytope theorem)**.** _For any graph G_ = (_V, E_)
_the matching polytope is determined by:_

(30) (i) _xe_ ≥ [^0]: for each _e_ ∈ _E;_
(ii)

```
∑
e∋v xe ≤[^1] for each v ∈ V ;
(iii)
```

```
∑
e⊂U xe ≤⌊[^1]
2|U|⌋ for each U ⊆ V with |U| odd.
```

**Proof.** Left to the reader (Exercise 5.21).

```
This in turn has the following consequence:
```

**Corollary 5.6c.** _Let G_ = (_V, E_) _be a graph and let w_ : _E_ → Q+_. Then the maximum_
_weight of a matching is equal to the minimum value of_

(31)

```
∑
```

```
v∈V
```

```
yv +
```

```
∑
```

```
U⊆V
```

```
zU⌊
```

```
1
2
```

```
|U|⌋,
```

_where y_ ∈ Q_V_
+ _and z_ ∈ Q

```
Podd(V)
+ satisfy
```

```
∑
v∈e yv +
```

∑
_U_∈Podd(_V_)_,e_⊆_U zU_ ≥ _w_(_e_) _for each_
_edge e._

**Proof.** Directly with LP-duality from Corollary 5.6b.

In fact, Cunningham and Marsh’ theorem shows that if _w_ is integer-valued, we
can restrict _y_ and _z_ to integer vectors — see Section 5.5.

**Exercises**

5.19. Show that for any graph _G_ = (_V,E_), if the inequalities (30)(i)(ii) fully determine the
matching polytope, then _G_ is bipartite.

5.20. Show that the perfect matching polytope of a graph _G_ = (_V,E_) is also determined
by the following inequalities:

```
(32) ∑ xe ≥ 0 for each e ∈ E;
```

```
e∈δ(U)
```

```
xe ≥ 1 for each odd subset U of V ;
∑
```

```
e∈E
```

```
xe =[^1]
2|V |.
```

5.21. Derive Edmonds’ matching polytope theorem from Edmonds’ perfect matching poly-
tope theorem.

5.22. Derive from Edmonds matching polytope theorem the linear inequalities determining
the convex hull of all _symmetric_ permutation matrices.

5.23. Let _G_ = (_V,E_) be a graph. Show that the convex hull of the incidence vectors of
matchings of size _k_ is equal to the intersection of the matching polytope of _G_ with
the hyperplane {_x_ | [^1]:_Tx_ = _k_}.

5.24. Let _G_ = (_V,E_) be a graph. Show that the convex hull of the incidence vectors of
matchings of size at least _k_ and at most _l_ is equal to the intersection of the matching
polytope of _G_ with the set {_x_ | _k_ ≤ [^1]:_Tx_ ≤ _l_}.

### 5.5. The Cunningham-Marsh formula

Cunningham and Marsh [1978] showed a more general result, which generalizes both
Edmonds’ matching polytope theorem and the Tutte-Berge formula. We give a direct
proof.

**Theorem 5.7** (Cunningham-Marsh formula)**.** _In Corollary_ 5.6c_, if w is integer, we_
_can take y and z integer._

**Proof.** We must give a matching _M_ and integer values _yv, zU_ as required with _w_(_M_)
equal to (31).
Let _T_ be equal to the maximum weight of a matching and let M be the set of
matchings _M_ of weight _T_. We prove the theorem by induction on _T_. We may assume
that _G_ is the complete graph on _V_. Let _G, w_ be a counterexample to the theorem
with (fixing _V_ and _T_)

∑
_e_∈_E w_(_e_) as large as possible.
First assume that there exists a vertex _u_ of _G_ covered by every matching _M_ ∈M.
Let _w_′ be obtained from _w_ by decreasing _w_(_e_) by 1 for each edge _e_ incident with _u_
with _w_(_e_) ≥ 1. Then the maximum of _w_′(_M_) over all matchings _M_ is equal to _T_ −1,
since each _M_ ∈ M contains an edge _e_ incident with _u_ with _w_(_e_) ≥ 1. Hence, by
induction, there exist _y_′
_v, z_′
_U_ as required for _w_′. Now increasing _y_′
_u_ by 1 and leaving
all other values of _y_′
_v, z_′
_U_ invariant, gives _yv, zU_ as required for _w_.
So we may assume that for each vertex _v_ there exists a matching _M_ ∈ M not
covering _v_. We show that for each three distinct vertices _a, b, c_ ∈ _V_ one has

(33) _w_(_ac_) ≥ min{_w_(_ab_)_, w_(_bc_)}_._

Indeed, by the maximality of

∑
_e_∈_E w_(_e_) there exists a matching _M_ ∈M containing
_ac_. (Otherwise we could increase the weight of _ac_ without increasing _T_, contradicting
the maximality of

∑
_e_∈_E w_(_e_).) Moreover, there exists a matching _M_′ ∈ M not
covering _b_. Let _P_ be the component of _M_∪_M_′ containing _ac_. At least one component,
_Q_ say, of _P_ \{_ac_} does not contain _b_. By symmetry of _a_ and _c_ we may assume that
_Q_ contains _a_. Then _M_△(_Q_ ∪{_ac_}) and _M_′△(_Q_ ∪{_ab_}) are matchings again. Now
_w_(_M_△(_Q_ ∪{_ac_})) ≤ _T_ = _w_(_M_), and so _w_(_Q_ ∩ _M_′) ≤ _w_(_Q_ ∩ _M_) + _w_(_ac_). Moreover,
_w_(_M_′△(_Q_ ∪{_ab_})) ≤ _T_ = _w_(_M_′), and so _w_(_Q_ ∩ _M_) + _w_(_ab_) ≤ _w_(_Q_ ∩ _M_′). Hence
_w_(_ab_) ≤ _w_(_ac_), proving (33).
For each natural number _n_ ≥ 1 let _Gn_ be the graph on _V_ with as edges all _e_ ∈ _E_
with _w_(_e_) ≥ _n_, and let K_n_ be the set of components of _Gn_. Consider some _n_ and
some _U_ ∈K_n_.
By (33), _G_|_U_ is a complete graph. We show that each _M_ ∈ M contains exactly
⌊[^1]
[^2]:|_U_|⌋ edges that are in _EU_ (= set of edges contained in _U_).
Suppose to the contrary that _U_ contains two vertices _a_ and _b_ such that _a_ and _b_ are
not covered by any edge in _M_ ∩_EU_. If _a_ or _b_ is not covered by _M_ we could replace the
edge in _M_ incident with _a_ or _b_ (if any) by the edge _ab_, thereby increasing the weight
— a contradiction. So we may assume that _ac, bd_ ∈ _M_ for some _c, d_ 6∈ _U_. By (33),
_w_(_cd_) ≥ min{_w_(_ac_)_, w_(_ad_)}≥ min{_w_(_ac_)_, w_(_ab_)_, w_(_bd_)} = min{_w_(_ac_)_, w_(_bd_)}. Since
_w_(_ab_) _>_ max{_w_(_ac_)_, w_(_bd_)} this implies _w_(_ab_) + _w_(_cd_) _> w_(_ac_) + _w_(_bd_). Therefore,
replacing _ac_ and _bd_ in _M_ by _ab_ and _cd_ would increase the weight — a contradiction.
So |_M_ ∩ _EU_| = ⌊[^1]
[^2]:|_U_|⌋.
For each _U_ ⊆ _V_ with |_U_| _>_ 1, define _zU_ as the number of natural numbers _n_ ≥ 1
for which _U_ ∈ K_n_. Then

∑
_U_⊇_e zU_ ≥ _w_(_e_) for each edge _e_ (since _e_ is in _w_(_e_) graphs
_Gn_). Moreover, choose _M_ ∈M arbitrarily. Then

(34)

```
∑
```

```
U⊆V
```

```
zU⌊
```

```
1
2
```

```
|U|⌋ =
```

```
∑∞
```

```
n=1
```

```
∑
```

```
U∈Kn
```

```
⌊
```

```
1
2
```

```
|U|⌋ =
```

```
∑∞
```

```
n=1
```

```
∑
```

```
U∈Kn
```

```
|M ∩ EU|
```

```
=
```

```
∑
```

```
e∈M
```

```
(number of n, U with e ⊆ U ∈Kn) =
```

```
∑
```

```
e∈M
```

```
w(e).
```

**Exercises**

5.25. Derive the Tutte-Berge formula from the Cunningham-Marsh formula (Theorem 5.7).

5.26. Derive Edmonds’ matching polytope theorem from the Cunningham-Marsh formula
(Theorem 5.7).

## 6. Problems, algorithms, and

## running time

### 6.1. Introduction

Probably most of the readers will have some intuitive idea about what is a problem
and what is an algorithm, and what is meant by the running time of an algorithm. Al-
though for the greater part of this course this intuition will be sufficient to understand
the substance of the matter, in some cases it is important to formalize this intuition.
This is particularly the case when we deal with concepts like NP and NP-complete.
The class of problems solvable in polynomial time is usually denoted by P. The
class NP, that will be described more precisely below, is a class of problems that
might be larger (and many people believe it _is_ larger). It includes most combinatorial
optimization problems, including all problems that are in P. That is: P⊆NP. In
particular, NP does **not** mean: “non-polynomial time”. The letters NP stand for
“nondeterministic polynomial-time”. The class NP consists, roughly speaking, of all
those questions with the property that for any input that has a positive answer, there
is a ‘certificate’ from which the correctness of this answer can be derived in polynomial
time.
For instance, the question:

(1) ‘Given a graph _G_, is _G_ Hamiltonian?’

belongs to NP. If the answer is ‘yes’, we can convince anyone that this answer is
correct by just giving a Hamiltonian circuit in _G_ as a certificate. With this certificate,
the answer ‘yes’ can be checked in polynomial time — in fact: trivially. Here it is
not required that we are able to _find_ the certificate in polynomial time. The only
requirement is that there _exists_ a certificate which can be checked in polynomial
time.
Checking the certificate in polynomial time means: checking it in time bounded
by a polynomial in the original input. In particular, it implies that the certificate
itself has size bounded by a polynomial in the original input.
To elucidate the meaning of NP, it is not known if for any graph _G_ for which
question (1) has a _negative_ answer, there is a certificate from which the correctness of
this answer can be derived in polynomial time. So there is an easy way of convincing
‘your boss’ that a certain graph is Hamiltonian (just by exhibiting a Hamiltonian
circuit), but no easy way is known for convincing this person that a certain graph is
non-Hamiltonian.
Within the class NP there are the “NP-complete” problems. These are by defi-
nition the hardest problems in the class NP: a problem Π in NP is NP-_complete_ if

every problem in NP can be reduced to Π, in polynomial time. It implies that if one
NP-complete problem can be proved to be solvable in polynomial time, then _each_
problem in NP can be solved in polynomial time. In other words: then P=NP would
follow.
Surprisingly, there are several prominent combinatorial optimization problems
that are NP-complete, like the traveling salesman problem and the problem of finding
a maximum clique in a graph. This pioneering eye-opener was given by Cook [1971]
and Karp [1972].
Since that time one generally sets the polynomially solvable problems against the
NP-complete problems, although there is no proof that these two concepts really are
distinct. For almost every combinatorial optimization problem one has been able
either to prove that it is solvable in polynomial time, or that it is NP-complete. But
theoretically it is still a possibility that these two concepts are just the same! Thus
it is unknown which of the two diagrams in Figure 6.1 applies.

```
NP
```

```
NP-c
```

```
P
```

```
NP-c P=NP
```

```
Figure 6.1
```

Below we make some of the notions more precise. We will not elaborate all tech-
nical details fully, but hope that the reader will be able to see the details with not
too much effort. For precise discussions we refer to the books by Aho, Hopcroft, and
Ullman [1974], Garey and Johnson [1979], and Papadimitriou [1994].

### 6.2. Words

If we use the computer to solve a certain graph problem, we usually do not put a
picture of the graph in the computer. (We are not working with analog computers,
but with digital computers.) Rather we put some appropriate encoding of the problem
in the computer, by describing it by a sequence of symbols taken from some fixed
finite ‘alphabet’ Σ. We can take for Σ for instance the ASCII set of symbols or the
set {[^0]:_,_[^1]:}. It is convenient to have symbols like ( , ) , { , } and the comma in Σ, and
moreover some symbol like meaning: ‘blank’. Let us fix one alphabet Σ.

We call any ordered finite sequence of elements from Σ a _word_. The set of all
words is denoted by Σ∗.

```
e
```

```
d
```

```
b
```

```
c
```

```
a
```

```
Figure 6.2
```

It is not difficult to encode objects like rational numbers, vectors, matrices, graphs,
and so on, as words. For instance, the graph given in Figure 6.2 can be encoded, as
usual, by the word:

(2) ({_a, b, c, d, e_}_,_{{_a, b_}_,_{_a, c_}_,_{_b, c_}_,_{_c, d_}_,_{_d, e_}_,_{_e, a_}})_._

A function _f_ defined on a finite set _X_ can be encoded by giving the set of pairs
(_x, f_(_x_)) with _x_ ∈ _X_. For instance, the following describes a function defined on the
edges of the graph above:

(3)
{({_a, b_}_,_32)_,_({_a, c_}_,_−17)_,_({_b, c_}_,_[^5]:_/_7)_,_({_c, d_}_,_6)_,_({_d, e_}_,_−1)_,_({_e, a_}_,_−9)}_._

A pair of a graph and a function can be described by the word (_w, v_), where _w_ is the
encoding of the graph and _v_ is the encoding of the function.
The _size_ of a word _w_ is the number of symbols used in _w_, counting multiplicities.
(So the word _abaa_[^32]:_bc_ has size 8.) The size is important when we make estimates on
the running time of algorithms.
Note that in encoding numbers (integers or rational numbers), the size depends
on the number of symbols necessary to encode these numbers. Thus if we encounter
a problem on a graph with numbers defined on the edges, then the size of the input
is the total number of bits necessary to represent this structure. It might be much
larger than just the number of nodes and edges of the graph, and much smaller than
the sum of all numbers occurring in the input.
Although there are several ways of choosing an alphabet and encoding objects by
words over this alphabet, any way chosen is quite arbitrary. We will be dealing with
solvability in polynomial time in this chapter, and for that purpose most encodings
are equivalent. Below we will sometimes exploit this flexibility.

### 6.3. Problems

What is a problem? Informally, it is a question or a task, for instance, “Does this given
graph have a perfect matching?” or “Find a shortest traveling salesman tour in this
graph!”. In fact there are two types of problems: problems that can be answered by
‘yes’ or ‘no’ and those that ask you to find an object with certain prescribed properties.
We here restrict ourselves to the first type of problems. From a complexity point of
view this is not that much of a restriction. For instance, the problem of finding a
shortest traveling salesman tour in a graph can be studied by the related problem:
Given a graph, a length function on the edges, and a rational number _r_, does there
exist a traveling salesman tour of length at most _r_? If we can answer this question
in polynomial time, we can find the length of a shortest tour in polynomial time, for
instance, by binary search.
So we study problems of the form: Given a certain object (or sequence of objects),
does it have a certain property? For instance, given a graph _G_, does it have a perfect
matching?
As we encode objects by words, a problem is nothing but: given a word _w_, does
it have a certain property? Thus the problem is fully described by describing the
“certain property”. This, in turn, is fully described by just the set of all words
that have the property. Therefore we have the following mathematical definition: a
_problem_ is any subset Π of Σ∗.
If we consider any problem Π ⊆ Σ∗, the corresponding ‘informal’ problem is:

(4) Given word _w_, does _w_ belong to Π?

In this context, the word _w_ is called an _instance_ or the _input_.

### 6.4. Algorithms and running time

An algorithm is a list of instructions to solve a problem. The classical mathematical
formalization of an algorithm is the _Turing machine_. In this section we will describe
a slightly different concept of an algorithm (the ‘Thue system’) that is useful for our
purposes (explaining NP-completeness). In Section 6.10 below we will show that it is
equivalent to the notion of a Turing machine.
A basic step in an algorithm is: replace subword _u_ by _u_′. It means that if word
_w_ is equal to _tuv_, where _t_ and _v_ are words, we replace _w_ by the word _tu_′_v_. Now
by definition, an _algorithm_ is a finite list of instructions of this type. It thus is fully
described by a sequence

(5) ((_u_[^1]:_, u_′
[^1]:)_,... ,_(_un, u_′
_n_))_,_

where _u_[^1]:_, u_′
[^1]:_,... , un, u_′
_n_ are words. We say that word _w_′ _follows from_ word _w_ if there

exists a _j_ ∈{[^1]:_,... , n_} such that _w_ = _tujv_ and _w_′ = _tu_′
_jv_ for certain words _t_ and _v_, in
such a way that _j_ is the smallest index for which this is possible and the size of _t_ is as
small as possible. The algorithm _stops at_ word _w_ if _w_ has no subword equal to one of
_u_[^1]:_,... , un_. So for any word _w_, either there is a unique word _w_′ that follows from _w_,
or the algorithm stops at _w_. A (finite or infinite) sequence of words _w_[^0]:_, w_[^1]:_, w_[^2]:_,..._ is
called _allowed_ if each _wi_+1 follows from _wi_ and, if the sequence is finite, the algorithm
stops at the last word of the sequence. So for each word _w_ there is a unique allowed
sequence starting with _w_. We say that _A accepts w_ if this sequence is finite.
For reasons of consistency it is important to have the ‘empty space’ at both sides
of a word as part of the word. Thus instead of starting with a word _w_, we start with
_w_ , where is a symbol indicating space.
Let _A_ be an algorithm and let Π ⊆ Σ∗ be a problem. We say that _A solves_ Π if
Π equals the set of words accepted by _A_. Moreover, _A_ solves Π _in polynomial-time_ if
there exists a polynomial _p_(_x_) such that for any word _w_ ∈ Σ∗: if _A_ accepts _w_, then
the allowed sequence starting with _w_ contains at most _p_(size(_w_)) words.
This definition enables us indeed to decide in polynomial time if a given word _w_
belongs to Π. We just take _w_[^0]: := _w_, and next, for _i_ = 0_,_[^1]:_,_[^2]:_,..._, we choose ‘the first’
subword _uj_ in _wi_ and replace it by _u_′
_j_ (for some _j_ ∈{[^1]_,... , n_}) thus obtaining _wi_+1.
If within _p_(size(_w_)) iterations we stop, we know that _w_ belongs to Π, and otherwise
we know that _w_ does not belong to Π.
Then P denotes the set of all problems that can be solved by a polynomial-time
algorithm.

### 6.5. The class NP

We mentioned above that NP denotes the class of problems for which a positive
answer has a ‘certificate’ from which the correctness of the positive answer can be
derived in polynomial time. We will now make this more precise.
The class NP consists of those problems Π ⊆ Σ∗ for which there exist a problem
Π′ ∈P and a polynomial _p_(_x_) such that for any _w_ ∈ Σ∗:

(6) _w_ ∈ Π if and only if there exists a word _v_ such that (_w, v_) ∈ Π′ and such
that size(_v_) ≤ _p_(size(_w_)).

So the word _v_ acts as a certificate showing that _w_ belongs to Π. With the polynomial-
time algorithm solving Π′, the certificate proves in polynomial time that _w_ belongs
to Π.
As examples, the problems

(7) Π[^1]: := {_G_ | _G_ is a graph having a perfect matching} and
Π[^2]: := {_G_ | _G_ is a Hamiltonian graph}

(encoding _G_ as above) belong to NP, since the problems

(8) Π′
[^1]: := {(_G, M_) | _G_ is a graph and _M_ is a perfect matching in _G_}
and
Π′
[^2]: := {(_G, H_) | _G_ is a graph and _H_ is a Hamiltonian circuit in
_G_}

belong to P.
Similarly, the problem

(9) TSP := {(_G, l, r_) | _G_ is a graph, _l_ is a ‘length’ function on the
edges of _G_ and _r_ is a rational number such that _G_ has a
Hamiltonian tour of length at most _r_}

(‘the traveling salesman problem’) belongs to NP, since the problem

(10) TSP′ := {(_G, l, r, H_) | _G_ is a graph, _l_ is a ‘length’ function on the
edges of _G_, _r_ is a rational number, and _H_ is a Hamiltonian
tour in _G_ of length at most _r_}

belongs to P.
Clearly, P⊆NP, since if Π belongs to P, then we can just take the empty string
as certificate for any word _w_ to show that it belongs to Π. That is, we can take
Π′ := {(_w,_) | _w_ ∈ Π}. As Π ∈P, also Π′ ∈P.
The class NP is apparently much larger than the class P, and there might be not
much reason to believe that the two classes are the same. But, as yet, nobody has
been able to show that they really are different! This is an intriguing mathematical
question, but besides, answering the question might also have practical significance.
If P=NP can be shown, the proof might contain a revolutionary new algorithm,
or alternatively, it might imply that the concept of ‘polynomial-time’ is completely
useless. If P[^6]:=NP can be shown, the proof might give us more insight in the reasons
why certain problems are more difficult than other, and might guide us to detect and
attack the kernel of the difficulties.

### 6.6. The class co-NP

By definition, a problem Π ⊆ Σ∗ belongs to the class co-NP if the ‘complementary’
problem Π := Σ∗ \ Π belongs to NP.
For instance, the problem Π[^1]: defined in (7) belongs to co-NP, since the problem

(11) Π′′
[^1]: := {(_G, W_) | _G_ is a graph and _W_ is a subset of the vertex set
of _G_ such that the graph _G_ − _W_ has more than |_W_| odd
components}

belongs to P. This follows from Tutte’s ‘1-factor theorem’ (Corollary 5.1a): a graph _G_
has no perfect matching if and only if there is a subset _W_ of the vertex set of _G_ with
the properties described in (11). (Here, strictly speaking, the complementary problem

Π[^1]: of Π[^1]: consists of all words _w_ that either do not represent a graph, or represent
a graph having no perfect matching. We assume however that there is an easy way
of deciding if a given word represents a graph. Therefore, we might assume that the
complementary problem is just {_G_ | _G_ is a graph having no perfect matching}.)
It is not known if the problems Π[^2]: and TSP belong to co-NP.
Since for any problem Π in P also the complementary problem Π belongs to P,
we know that P⊆co-NP. So P⊆NP∩co-NP. The problems in NP∩co-NP are those for
which there exist certificates both in case the answer is positive and in case the answer
is negative. As we saw above, the perfect matching problem Π[^1]: is such a problem.
Tutte’s theorem gives us the certificates. Therefore, Tutte’s theorem is called a _good_
_characterization_.
In fact, there are very few problems known that are proved to belong to NP∩co-NP,
but that are not known to belong to P. Most problems having a good characterization,
have been proved to be solvable in polynomial time. The notable exception for which
this is not yet proved is _primality testing_ (testing if a given natural number is a prime
number).

### 6.7. NP-completeness

The NP-complete problems are by definition the hardest problems in NP. To be more
precise, we first define the concept of a polynomial-time reduction. Let Π and Π′
be two problems and let _A_ be an algorithm. We say that _A_ is a _polynomial-time_
_reduction_ of Π′ to Π if _A_ is a polynomial-time algorithm (‘solving’ Σ∗), so that for
any allowed sequence starting with _w_ and ending with _v_ one has: _w_ ∈ Π′ if and only
if _v_ ∈ Π. A problem Π is called NP-_complete_, if Π ∈NP and for each problem Π′ in
NP there exists a polynomial-time reduction of Π′ to Π.
It is not difficult to see that if Π belongs to P and there exists a polynomial-time
reduction of Π′ to Π, then also Π′ belongs to P. It implies that if one NP-complete
problem can be solved in polynomial time, then each problem in NP can be solved in
polynomial time. Moreover, if Π belongs to NP, Π′ is NP-complete and there exists
a polynomial-time reduction of Π′ to Π, then also Π is NP-complete.

### 6.8. NP-completeness of the satisfiability problem

We now first show that in fact there exist NP-complete problems. In fact we show
that the so-called _satisfiability problem_, denoted by SAT, is NP-complete.
To define SAT, we need the notion of a _boolean expression_. Examples are:

(12) ((_x_[^2]: ∧ _x_[^3]:) ∨¬(_x_[^3]: ∨ _x_[^5]:) ∧ _x_[^2]:), ((¬_x_[^47]: ∧ _x_[^2]:) ∧ _x_[^47]:), ¬(_x_[^7]: ∧¬_x_[^7]:).

Boolean expressions can be defined inductively. First, for each natural number _n_,
the ‘word’ _xn_ is a boolean expression (using some appropriate encoding of natural
numbers and of subscripts). Next, if _v_ and _w_ are boolean expressions, then also
(_v_ ∧ _w_), (_v_ ∨ _w_) and ¬_v_ are boolean expressions. These rules give us all boolean
expressions. (If necessary, we may use other subscripts than the natural numbers.)
Now SAT is a subcollection of all boolean expressions, namely it consists of those
boolean expressions that are satisfiable. A boolean expression _f_(_x_[^1]:_, x_[^2]:_, x_[^3]:_,..._) is
called _satisfiable_ if there exist _α_[^1]:_, α_[^2]:_, α_[^3]:_,..._ ∈ {[^0]:_,_[^1]:} such that _f_(_α_[^1]:_, α_[^2]:_, α_[^3]:_,..._) = 1,
using the well-known identities

(13) [^0]: ∧ 0 = 0 ∧ 1 = 1 ∧ 0 = 0_,_[^1]: ∧ 1 = 1_,_
[^0]: ∨ 0 = 0_,_[^0]: ∨ 1 = 1 ∨ 0 = 1 ∨ 1 = 1_,_
¬0 = 1_,_¬1 = 0_,_(0) = 0_,_(1) = 1_._

**Exercise.** Let _n_ ≥ 1 be a natural number and let _W_ be a collection of words in
{[^0]:_,_[^1]:}∗ all of length _n_. Prove that there exists a boolean expression _f_(_x_[^1]:_,... , xn_) in
the variables _x_[^1]:_,... , xn_ such that for each word _w_ = _α_[^1]:_... αn_ in the symbols 0 and 1
one has: _w_ ∈ _W_ if and only if _f_(_α_[^1]:_,... , αn_) = 1.

The satisfiability problem SAT trivially belongs to NP: we can take as certificate
for a certain _f_(_x_[^1]:_, x_[^2]:_, x_[^3]:_,..._) to belong to SAT, the equations _xi_ = _αi_ that give _f_ the
value 1. (We only give those equations for which _xi_ occurs in _f_.)
To show that SAT is NP-complete, it is convenient to assume that Σ = {[^0]:_,_[^1]:}.
This is not that much a restriction: we can fix some order of the symbols in Σ, and
encode the first symbol by 10, the second one by 100, the third one by 1000, and so
on. There is an easy (certainly polynomial-time) way of obtaining one encoding from
the other.
The following result is basic for the further proofs:

**Theorem 6.1.** _Let_ Π ⊆ {[^0]:_,_[^1]:}∗ _be in_ P_. Then there exist a polynomial p_(_x_) _and_
_an algorithm that finds for each natural number n in time p_(_n_) _a boolean expression_
_f_(_x_[^1]:_, x_[^2]:_, x_[^3]:_,..._) _with the property:_

(14) _any word α_[^1]:_α_[^2]:_... αn in_ {[^0]:_,_[^1]:}∗ _belongs to_ Π _if and only if the boolean ex-_
_pression f_(_α_[^1]:_,... , αn, xn_+1_, xn_+2_,..._) _is satisfiable._

**Proof.** Since Π belongs to P, there exists a polynomial-time algorithm _A_ solving Π.
So there exists a polynomial _p_(_x_) such that a word _w_ belongs to Π if and only if the
allowed sequence for _w_ contains at most _p_(size(_w_)) words. It implies that there exists

```
a polynomial q(x) such that any word in the allowed sequence for w has size less than
q(size(w)).
We describe the algorithm meant in the theorem. Choose a natural number n.
Introduce variables xi,j and yi,j for i = 0,1,... , p(n), j = 1,... , q(n). Now there exists
(cf. the Exercise above) a boolean expression f in these variables with the following
properties. Any assignment xi,j := αi,j ∈{0,1} and yi,j := βi,j ∈{0,1} makes f equal
to 1 if and only if the allowed sequence starting with the word w0 := α0,1α0,2... α0,n
is a finite sequence w0,... , wk, so that:
```

(15) (i) _αi,j_ is equal to the _j_th symbol in the word _wi_, for each _i_ ≤ _k_ and each
_j_ ≤ size(_wi_);
(ii) _βi,j_ = 1 if and only if _i > k_ or _j_ ≤ size(_wi_).

```
The important point is that f can be found in time bounded by a polynomial in
n. To see this, we can encode the fact that word wi+1 should follow from word wi
by a boolean expression in the ‘variables’ xi,j and xi+1,j, representing the different
positions in wi and wi+1. (The extra variables yi,j and yi+1,j are introduced to indicate
the sizes of wi and wi+1.) Moreover, the fact that the algorithm stops at a word w
also can be encoded by a boolean expression. Taking the ‘conjunction’ of all these
boolean expressions, will give us the boolean expression f.
```

```
As a direct consequence we have:
```

```
Corollary 6.1a. Theorem 6.1 also holds if we replace P by NP in the first sentence.
```

```
Proof. Let Π ⊆ {0,1}∗ belong to NP. Then, by definition of NP, there exists a
problem Π′ in P and a polynomial r(x) such that any word w belongs to Π if and
only if (w, v) belongs to Π′ for some word v with size(v) ≤ r(size(w)). By properly
re-encoding, we may assume that for each n ∈ N, any word w ∈{0,1}∗ belongs to Π
if and only if wv belongs to Π′ for some word v of size r(size(w)). Applying Theorem
6.1 to Π′ gives the corollary.
```

```
Now the main result of Cook [1971] follows:
```

```
Corollary 6.1b (Cook’s theorem). The satisfiability problem SAT is NP-complete.
```

```
Proof. Let Π belong to NP. We describe a polynomial-time reduction of Π to SAT.
Let w = α1... αn ∈ {0,1}∗. By Corollary 6.1a we can find in time bounded by
a polynomial in n a boolean expression f such that w belongs to Π if and only if
f(α1,... , αn, xn+1,...) is satisfiable. This is the required reduction to SAT.
```

### 6.9. NP-completeness of some other problems

We next derive from Cook’s theorem some of the results of Karp [1972]. First we
show that the 3-_satisfiability problem_ 3-SAT is NP-complete. Let _B_[^1]: be the set of
all words _x_[^1]:_,_¬_x_[^1]:_, x_[^2]:_,_¬_x_[^2]:_,..._. Let _B_[^2]: be the set of all words (_w_[^1]: ∨···∨ _wk_), where
_w_[^1]:_,_··· _, wk_ are words in _B_[^1]: and 1 ≤ _k_ ≤ 3. Let _B_[^3]: be the set of all words _w_[^1]:∧_..._∧_wk_,
where _w_[^1]:_,... , wk_ are words in _B_[^2]:. Again, we say that a word _f_(_x_[^1]:_, x_[^2]:_,..._) ∈ _B_[^3]: is
_satisfiable_ if there exists an assignment _xi_ := _αi_ ∈ {[^0]:_,_[^1]:} (_i_ = 1_,_[^2]:_,..._) such that
_f_(_α_[^1]:_, α_[^2]:_,..._) = 1 (using the identities (13)).
Now the 3-satisfiability problem 3-SAT is: Given a word _f_ ∈ _B_[^3]:, decide if it is
satisfiable.

**Corollary 6.1c.** _The 3-satisfiability problem_ 3-SAT _is_ NP_-complete._

**Proof.** We give a polynomial-time reduction of SAT to 3-SAT. Let _f_(_x_[^1]:_, x_[^2]:_,..._) be a
boolean expression. Introduce a variable _yg_ for each subword _g_ of _f_ that is a boolean
expression.
Now _f_ is satisfiable if and only if the following system is satisfiable:

(16) _yg_ = _yg_′ ∨ _yg_′′ (if _g_ = _g_′ ∨ _g_′′),
_yg_ = _yg_′ ∧ _yg_′′ (if _g_ = _g_′ ∧ _g_′′),
_yg_ = ¬_yg_′ (if _g_ = ¬_g_′),
_yf_ = 1.

Now _yg_ = _yg_′ ∨ _yg_′′ can be equivalently expressed by: _yg_ ∨¬_yg_′ = 1_, yg_ ∨¬_yg_′′ =
[^1]:_,_¬_yg_ ∨ _yg_′ ∨ _yg_′′ = 1. Similarly, _yg_ = _yg_′ ∧ _yg_′′ can be equivalently expressed by:
¬_yg_ ∨ _yg_′ = 1_,_¬_yg_ ∨ _yg_′′ = 1_, yg_ ∨ ¬_yg_′ ∨ ¬_yg_′′ = 1. The expression _yg_ = ¬_yg_′ is
equivalent to: _yg_ ∨ _yg_′ = 1_,_¬_yg_ ∨¬_yg_′ = 1.
By renaming variables, we thus obtain words _w_[^1]:_,... , wk_ in _B_[^2]:, so that _f_ is satis-
fiable if and only if the word _w_[^1]: ∧_..._ ∧ _wk_ is satisfiable.

We next derive that the _partition problem_ PARTITION is NP-complete. This is
the problem: Given a collection C of subsets of a finite set _X_, is there a subcollection
of C that forms a partition of _X_?

**Corollary 6.1d.** _The partition problem_ PARTITION _is_ NP_-complete._

**Proof.** We give a polynomial-time reduction of 3-SAT to PARTITION. Let _f_ =
_w_[^1]:∧_..._∧_wk_ be a word in _B_[^3]:, where _w_[^1]:_,... , wk_ are words in _B_[^2]:. Let _x_[^1]:_,... , xm_ be the
variables occurring in _f_. Make a bipartite graph _G_ with colour classes {_w_[^1]:_,... , wk_}
and {_x_[^1]:_,... , xm_}, by joining _wi_ and _xj_ by an edge if and only if _xj_ or ¬_xj_ occurs in
_wi_. Let _X_ be the set of all vertices and edges of _G_.
Let C′ be the collection of all sets {_wi_}∪_E_′, where _E_′ is a nonempty subset of the

edge set incident with _wi_. Let C′′ be the collection of all sets {_xj_}∪_E_′
_j_ and {_xj_}∪_E_′′
_j_ ,
where _E_′
_j_ is the set of all edges {_wi, xj_} so that _xj_ occurs in _wi_ and where _E_′′
_j_ is the
set of all edges {_wi, xj_} so that ¬_xj_ occurs in _wi_.
Now _f_ is satisfiable if and only if the collection C′ ∪C′′ contains a subcollection
that partitions _X_. Thus we have a reduction of 3-SAT to PARTITION.

We derive the NP-completeness of the _directed Hamiltonian cycle problem_ DI-
RECTED HAMILTONIAN CYCLE: Given a directed graph, does it have a directed
Hamiltonian cycle?

**Corollary 6.1e.** DIRECTED HAMILTONIAN CYCLE _is_ NP_-complete._

**Proof.** We give a polynomial-time reduction of PARTITION to DIRECTED HAMIL-
TONIAN CYCLE. Let C = {_C_[^1]:_,... , Cm_} be a collection of subsets of the set _X_ =
{_x_[^1]:_,... , xk_}. Introduce ‘vertices’ _r_[^0]:_, r_[^1]:_,... , rm, s_[^0]:_, s_[^1]:_,... , sk_.
For each _i_ = 1_,... , m_ we do the following. Let _Ci_ = {_xj_[^1]:_,... , xjt_}. We construct a
directed graph on the vertices _ri_−[^1]:_, ri_, _sjh_−[^1]:_, sjh_ (for _h_ = 1_,... , t_) and 3_t_ new vertices,
as in Figure 6.3. Moreover, we make arcs from _rm_ to _s_[^0]: and from _sk_ to _r_[^0]:.
s

```
-1
```

```
j jt
-1 s
```

```
ri
```

```
t -1
```

```
r
```

```
j2
```

```
i
```

```
s s jt-1 jt-1-1 sj j 1-1
2
```

```
s s
1 sj
```

```
Figure 6.3
```

Let _D_ be the directed graph arising. Then it is not difficult to check that there
exists a subcollection C′ of C that partitions _X_ if and only if _D_ has a directed Hamil-
tonian cycle _C_. (Take: (_ri_−[^1]:_, ri_) ∈ _C_ ⇐⇒ _Ci_ ∈C′.)

From this we derive the NP-completeness of the _undirected Hamiltonian cycle_
_problem_ UNDIRECTED HAMILTONIAN CYCLE: Given a graph, does it have a
Hamiltonian cycle?

**Corollary 6.1f.** UNDIRECTED HAMILTONIAN CYCLE _is_ NP_-complete._

**Proof.** We give a polynomial-time reduction of DIRECTED HAMILTONIAN CY-
CLE to UNDIRECTED HAMILTONIAN CYCLE. Let _D_ be a directed graph. Re-
place each vertex _v_ by three vertices _v_′_, v_′′_, v_′′′, and make edges {_v_′_, v_′′} and {_v_′′_, v_′′′}.

Moreover, for each arc (_v_[^1]:_, v_[^2]:) of _D_, make an edge {_v_′
[^1]:_, v_′′′
[^2]: }. This makes the undi-
rected graph _G_. One easily checks that _D_ has a directed Hamiltonian cycle if and

only if _G_ has an (undirected) Hamiltonian cycle.

This trivially implies the NP-completeness of the _traveling salesman problem_ TSP:
Given a complete graph _G_ = (_V, E_), a ‘length’ function _l_ on _E_, and a rational _r_, does
there exist a Hamiltonian cycle of length at most _r_?

**Corollary 6.1g.** _The traveling salesman problem_ TSP _is_ NP_-complete._

**Proof.** We give a polynomial-time reduction of UNDIRECTED HAMILTONIAN
CYCLE to TSP. Let _G_ be a graph. Let _G_′ be the complete graph on _V_. Let _l_(_e_) := 0
for each edge _e_ of _G_ and let _l_(_e_) := 1 for each edge of _G_′ that is not an edge of _G_.
Then _G_ has a Hamiltonian cycle if and only if _G_′ has a Hamiltonian cycle of length
at most 0.

### 6.10. Turing machines

In Section 6.4 we gave a definition of ‘algorithm’. How adequate is this definition?
Can any computer program be modelled after that definition?
To study this question, we need to know what we understand by a ‘computer’.
Turing [1937] gave the following computer model, now called a _Turing machine_ or a
_one-tape Turing machine_.
A Turing machine consists of a ‘processor’ that can be in a finite number of ‘states’
and of a ‘tape’, of infinite length (in two ways). Moreover, there is a ‘read-write head’,
that can read symbols on the tape (one at a time). Depending on the state of the
processor and the symbol read, the processor passes to another (or the same) state,
the symbol on the tape is changed (or not) and the tape is moved one position ‘to
the right’ or ‘to the left’.
The whole system can be described by just giving the dependence mentioned in
the previous sentence. So, mathematically, a _Turing machine_ is just a function

(17) _T_ : _M_ × Σ → _M_ × Σ ×{+1_,_−[^1]:}.

Here _M_ and Σ are finite sets: _M_ is interpreted as the set of states of the processor,
while Σ is the set of symbols that can be written on the tape. The function _T_
describes an ‘iteration’: _T_(_m, σ_) = (_m_′_, σ_′_,_+1) should mean that if the processor is
in state _m_ and the symbol read on the tape is _σ_, then the next state will be _m_′, the
symbol _σ_ is changed to the symbol _σ_′ and the tape is moved one position to the right.
_T_(_m, σ_) = (_m_′_, σ_′_,_−1) has a similar meaning — now the tape is moved one position
to the left.

Thus if the processor is in state _m_ and has the word _w_′_α_′_σα_′′_w_′′ on the tape,
where the symbol indicated by _σ_ is read, and if _T_(_m, σ_) = (_m_′_, σ_′_,_+1), then next the

processor will be in state _m_′ and has the word _w_′_α_′_σ_′_α_′′_w_′′ on the tape, where the
symbol indicated by _α_′′ is read. Similarly if _T_(_m, σ_) = (_m_′_, σ_′_,_−1).
We assume that _M_ contains a certain ‘start state’ 0 and a certain ‘halting state’
∞. Moreover, Σ is assumed to contain a symbol meaning ‘blank’. (This is necessary
to identify the beginning and the end of a word on the tape.)
We say that the Turing machine _T accepts_ a word _w_ ∈ (Σ\{ })∗ if, when starting
in state 0 and with word _w_ on the tape (all other symbols being blank), so that
the read-write head is reading the first symbol of _w_, then after a finite number of
iterations, the processor is in the halting state ∞. (If _w_ is the empty word, the
symbol read initially is the blank symbol .)
Let Π be the set of words accepted by _T_. So Π is a problem. We say that _T solves_
Π. Moreover, we say that _T solves_ Π _in polynomial time_ if there exists a polynomial
_p_(_x_) such that if _T_ accepts a word _w_, it accepts _w_ in at most _p_(size(_w_)) iterations.
It is not difficult to see that the concept of algorithm defined in Section 6.4 above
is at least as powerful as that of a Turing machine. We can encode any state of the
computer model (processor+tape+read-write head) by a word (_w_′_, m, w_′′). Here _m_ is
the state of the processor and _w_′_w_′′ is the word on the tape, while the first symbol of
_w_′′ is read. We define an algorithm _A_ by:

(18) replace subword _, m, σ_ by _σ_′_, m_′, whenever _T_(_m, σ_) = (_m_′_, σ_′_,_+1) and _m_ [^6]:=
∞;
replace subword _α, m, σ_ by _m_′_, ασ_′, whenever _T_(_m, σ_) = (_m_′_, σ_′_,_−1) and
_m_ [^6]:= ∞.

To be precise, we should assume here that the symbols indicating the states in _M_
do not belong to Σ. Moreover, we assume that the symbols ( and ) are not in Σ.
Furthermore, to give the algorithm a start, it contains the tasks of replacing subword
_α_ by the word (_,_[^0]:_, α_ , and subword _α_ by _α_) (for any _α_ in Σ \{ }). Then, when
starting with a word _w_, the first two iterations transform it to the word (_,_[^0]:_, w_). After
that, the rules (18) simulate the Turing machine iterations. The iterations stop as
soon as we arrive at state ∞.
So _T_ accepts a word _w_ if and only if _A_ accepts _w_ — in (about) the same number
of iterations. That is, _T_ solves a problem Π (in polynomial time) if and only if _A_
solves Π (in polynomial time).
This shows that the concept of ‘algorithm’ defined in Section 6.4 is at least as
powerful as that of a Turing machine. Conversely, it is not hard (although technically
somewhat complicated) to simulate an algorithm by a Turing machine. But how
powerful is a Turing machine?
One could think of several objections against a Turing machine. It uses only one
tape, that should serve both as an input tape, and as a memory, and as an output
tape. We have only limited access to the information on the tape (we can shift only
one position at a time). Moreover, the computer program seems to be implemented in

the ‘hardware’ of the computer model; the Turing machine solves only one problem.
To counter these objections, several other computer models have been proposed
that model a computer more realistically: multi-tape Turing machines, random access
machines (RAM’s), the universal Turing machine. However, from a polynomial-time
algorithmic point of view, these models all turn out to be equivalent. Any problem
that can be solved in polynomial time by any of these computer models, can also
be solved in polynomial time by some one-tape Turing machine, and hence by an
algorithm in the sense of Section 6.4. We refer to Aho, Hopcroft, and Ullman [1974]
and Papadimitriou [1994] for an extensive discussion.

## 7. Cliques, stable sets, and

## colourings

### 7.1. Introduction

```
We have seen in Chapter 5 that in any graph G = (V, E), a matching of maximum
cardinality can be found in polynomial time. Similarly, an edge-cover of minimum
cardinality can be found in polynomial time.
On the other hand, it is NP-complete to find a maximum-cardinality stable set in
a graph. That is, determining α(G) is NP-complete. To be more precise, the problem
COCLIQUE is:
```

(1) given: a graph _G_ and a natural number _k_,
decide: if _α_(_G_) ≥ _k_.

```
Then:
```

```
Theorem 7.1. The problem COCLIQUE is NP-complete.
```

```
Proof. We reduce SAT to COCLIQUE. Let C1 ∧ ··· ∧ Ck be a boolean expres-
sion in the variables x1,... , xn, where each expression is a disjunction of the literals
x1,¬x1,... , xn,¬xn. Consider the graph G = (V, E) with V := {(σ, i) | σ is a literal
in Ci} and E := {{(σ, i),(τ, j)}| i = j or σ = ¬τ}. Then the expression is satisfiable
if and only if G has a stable set of size k.
```

```
Since by Gallai’s theorem Theorem 3.1, α(G) = |V |− τ(G), also determining the
vertex-cover number τ(G) is NP-complete.
A clique in a graph G = (V, E) is a subset C of V such that u and w are adjacent
for any two distinct u, w in C. The clique number of G, denoted by ω(G), is the
maximum cardinality of any clique in G.
Observe that a subset C of V is a clique in G if and only if C is a stable set in the
complementary graph G. So finding a maximum-cardinality clique in G is equivalent
to finding a maximum-cardinality stable set in G, and ω(G) = α(G). As determining
α(G) is NP-complete, also determining ω(G) is NP-complete.
A (vertex-)colouring of a graph G = (V, E) is a partition of V into stable sets
C1,... , Ck. The sets C1,... , Ck are called the colours of the colouring. The (vertex-
) colouring number, or (vertex-)chromatic number, of G, denoted by χ(G), is the
minimum number of colours in any vertex-colouring of G. A graph G is called k-
colourable if χ(G) ≤ k. Again, it is NP-complete to decide if a graph is k-colourable.
That is, let VERTEX-COLOURING be the problem:
```

(2) given: a graph _G_ and a natural number _k_,
decide: if _χ_(_G_) ≤ _k_.

```
Theorem 7.2. The problem VERTEX-COLOURING is NP-complete.
```

```
Proof. We show that COCLIQUE can be reduced to VERTEX-COLOURING. Let
G = (V, E) be an undirected graph and let k ∈ Z+. We want to decide if α(G) ≥ k.
To this end, let V ′ be a copy of V and let C be a set of size k, where V , V ′, and C
are disjoint. Make a graph H with vertex set V ∪V ′ ∪C as follows. A pair of vertices
in V is adjacent in H if and only if it is adjacent in G. The sets V ′ and C are cliques
in H. Each vertex in V is adjacent to each vertex in V ′ ∪ C, except to its copy in V ′.
No vertex in V ′ is adjacent to any vertex in C.
This defines the graph H. Then α(G) ≥ k if and only if χ(H) ≤|V | + 1.
```

```
Well-known is the four-colour conjecture (4CC), stating that χ(G) ≤ 4 for each
planar graph G. This conjecture was proved by Appel and Haken [1977] and Appel,
Haken, and Koch [1977], and is now called the four-colour theorem (4CT).
In fact, it is NP-complete to decide if a planar graph is 3-colourable. [Note that
one can decide in polynomial time if a graph G is 2-colourable, as bipartiteness can
be checked in polynomial time.]
These NP-completeness results imply that if NP6=co-NP, then one may not ex-
pect a min-max relation characterizing the stable set number α(G), the vertex-cover
number τ(G), the clique number ω(G), or the colouring number χ(G) of a graph G.
There is a trivial upper bound on the colouring number:
```

```
(3) χ(G) ≤ ∆(G) + 1,
```

```
where ∆(G) denotes the maximum valency of G. Brooks [1941] sharpened this in-
equality as follows:
```

```
Theorem 7.3 (Brooks’ theorem). For any connected graph G one has χ(G) ≤ ∆(G),
except if G = Kn or G = C2n+1 for some n ≥ 1.[^19]
```

```
Another inequality relates the clique number and the colouring number:
```

```
(4) ω(G) ≤ χ(G).
```

```
This is easy, since in any clique all vertices should have different colours.
But there are several graphs which have strict inequality in (4). We mention
the odd circuits C2k+1, with 2k + 1 ≥ 5: then ω(C2k+1) = 2 and χ(C2k+1) = 3.
```

```
19Here Ck denotes the circuit with k vertices.
```

Moreover, for the complement _C_[^2]:_k_+1 of any such graph we have: _ω_(_C_[^2]:_k_+1) = _k_ and
_χ_(_C_[^2]:_k_+1) = _k_ + 1.
It was a conjecture of Berge [1963] that these graphs are crucial, which was proved
in 2002 by Chudnovsky, Robertson, Seymour, and Thomas:[^20]

**Strong perfect graph theorem**: Let _G_ be a graph. If _ω_(_G_) _< χ_(_G_) then _G_
contains _Cn_ or _Cn_, for some odd _n_ ≥ 5, as an _induced_ subgraph.
Another conjecture is due to Hadwiger [1943]. Since there exist graphs with
_ω_(_G_) _< χ_(_G_), it is not true that if _χ_(_G_) ≥ _n_ then _G_ contains the complete graph

_Kn_ on _n_ vertices as a subgraph. However, Hadwiger conjectured the following, where
a graph _H_ is called a _minor_ of a graph _G_ if _H_ arises from some subgraph of _G_ by
contracting some (possible none) edges.

**Hadwiger’s conjecture**: If _χ_(_G_) ≥ _n_ then _G_ contains _Kn_ as a minor.

In other words, for each _n_, the graph _Kn_ is the only graph _G_ with the property that
_G_ is not (_n_ − 1)-colourable and each proper minor of _G_ is (_n_ − 1)-colourable.
Hadwiger’s conjecture is trivial for _n_ = 1_,_[^2]:_,_3, and was shown by Hadwiger for
_n_ = 4 (see Exercise 7.8). As planar graphs do not contain _K_[^5]: as a minor, Hadwiger’s
conjecture for _n_ = 5 implies the four-colour theorem. In fact, Wagner [1937] showed
that Hadwiger’s conjecture for _n_ = 5 is equivalent to the four-colour conjecture.
Recently, Robertson, Seymour, and Thomas [1993] showed that Hadwiger’s conjecture
is true also for _n_ = 6, by showing that in that case it is equivalent to the four-colour
theorem. For _n_ ≥ 7 Hadwiger’s conjecture is unsettled.

**Application 7.1: Map colouring.** A well-known application of colouring the vertices of
a graph is that of colouring the countries in a map in such a way that adjacent countries
obtain different colours. So the four-colour theorem implies that if each country is connected,
then the map can be coloured using not more than four colours. (One should not consider
countries as ‘adjacent’ if they have a common boundary of measure 0 only.)
There are several other cases where colouring a map amounts to finding a minimum
vertex-colouring in a graph. For instance, consider a map of the Paris M ́etro network
(Figure 7.1).
Suppose now that you want to print a coloured map of the network, indicating each of
the 13 lines by a colour, in such a way that lines that cross each other or meet each other
in a station, are indicated by different colours and in such a way that a minimum number
of colours is used. This easily reduces to a graph colouring problem.

**Application 7.2: Storage of goods, etc.** Suppose you are the director of a circus and
wish to transport your animals in a number of carriages, in such a way that no two of the
animals put into one carriage eat each other, and in such a way that you use a minimum

[^20]:Let _G_ = (_V,E_) be a graph and let _V_ ′ ⊆ _V_. Then the subgraph of _G induced by V_ ′, denoted by
_G_|_V_ ′ is the graph (_V_ ′_,E_′), where _E_′ equals the set of all edges in _E_ contained in _V_ ′. The graph
_G_|_V_ ′ is called an _induced_ subgraph of _G_.

```
1
```

```
2
```

```
3
```

```
3
```

```
4
```

```
4
```

```
5
```

```
5
```

```
6
```

```
2
```

```
7
```

```
7
```

```
13
12
```

```
11
```

```
9
```

```
10
```

```
12
13
```

```
8
7
```

```
13
```

```
10
9 1
8
```

```
6
```

```
11
```

```
1
```

```
10
12
```

```
11
13
```

```
8
9
```

```
2
3
4
5
6
7
Figure 7.1. The Paris M ́etro lines
```

number of carriages.
This trivially reduces to a graph colouring problem. A similar problem is obtained if
you have to store a number of chemicals in a minimum number of rooms of a storehouse,
in such a way that no two of the chemicals stored in one room react upon each other in an
unwanted way.
This problem may also occur when assigning multiple-bed rooms to school boys on a
school trip.

**Application 7.3: Assigning frequencies to radio stations, car phones, etc.** Suppose
one has to assign frequencies to radio stations in a certain area. Certain pairs of radio
stations that are too close to each other cannot be assigned the same frequency as it would
cause mutual interference. Such pairs of radio stations form the edge set of a graph _G_, with
vertex set the set of radio stations. The chromatic number of _G_ is equal to the minimum
number of different frequencies that one needs in order to assign a frequency to each of the
stations.
The problem occurs also when assigning frequencies to car phones, where often in a very
short time new frequencies should be determined.

**Exercises**

```
7.1. Determine ω(G) and χ(G) for the graph G obtained from the Paris M ́etro map given
in Application 7.1.
7.2. Colour the map of Figure 7.2 (from the April 1975 issue of Scientific American).
```

```
Figure 7.2
```

```
7.3. Show that if G is a bipartite graph, then ω(G) = χ(G).
7.4. Derive from K ̋onig’s edge cover theorem (Corollary 3.3a) that if G is the complement
of a bipartite graph, then ω(G) = χ(G).
7.5. Derive K ̋onig’s edge cover theorem (Corollary 3.3a) from the strong perfect graph
theorem.
7.6. Let H be a bipartite graph and let G be the complement of the line-graph of H.
Derive from K ̋onig’s matching theorem (Theorem 3.3) that ω(G) = χ(G).
7.7. Derive K ̋onig’s matching theorem (Theorem 3.3) from the strong perfect graph the-
orem.
```

```
7.8. Let G = (V,E) be a simple graph such that no minor of G is isomorphic to K4. Show
that χ(G) ≤ 3.
[Hint: One may assume that G is not a forest or a circuit. Then G has a circuit not
covering all vertices of G. As G has no K4-minor, G is not 3-connected, that is, G
has a vertex cut set of size less than 3; then χ(G) ≤ 3 follows by induction.]
```

### 7.2. Edge-colourings of bipartite graphs

For any graph _G_ = (_V, E_), an _edge-colouring_ is a partition Π = {_M_[^1]:_,... , Mp_} of the
edge set _E_, where each _Mi_ is a matching. Each of these matchings is called a _colour_.

Define the _edge-colouring number_ or _edge-chromatic number χ_′(_G_) by

(5) _χ_′(_G_) := min{|Π|| Π is an edge-colouring of _G_}.

So _χ_′(_G_) = _χ_(_L_(_G_)), where _L_(_G_) is the line graph of _G_.
Let ∆(_G_) denote the maximum degree of (the vertices of) _G_. Clearly,

(6) _χ_′(_G_) ≥ ∆(_G_)_,_

since at each vertex _v_, the edges incident with _v_ should have different colours. Again
the triangle _K_[^3]: has strict inequality. K ̋onig [1916] showed that for bipartite graphs
the two numbers are equal.

**Theorem 7.4** (K ̋onig’s edge-colouring theorem)**.** _For any bipartite graph G_ = (_V, E_)
_one has_

(7) _χ_′(_G_) = ∆(_G_)_._

_That is, the edge-colouring number of a bipartite graph is equal to its maximum de-_
_gree._

**Proof.** First notice that the theorem is easy if ∆(_G_) ≤ 2. In that case, _G_ consists of
a number of vertex-disjoint paths and even circuits.
In the general case, colour as many edges of _G_ as possible with ∆(_G_) colours,
without giving the same colour to two intersecting edges. If all edges are coloured we
are done, so suppose some edge _e_ = {_u, w_} is not coloured. At least one colour, say
_red_, does not occur among the colours given to the edges incident with _u_. Similarly,
there is a colour, say _blue_, not occurring at _w_. (Clearly, _red_[^6]:=_blue_, since otherwise we
could give edge _e_ the colour _red_.)
Let _H_ be the subgraph of _G_ having as edges all _red_ and _blue_ edges of _G_, together
with the edge _e_. Now ∆(_H_) = 2, and hence _χ_′(_H_) = ∆(_H_) = 2. So all edges
occurring in _H_ can be (re)coloured with _red_ and _blue_. In this way we colour more
edges of _G_ than before. This contradicts the maximality assumption.

This proof also gives a polynomial-time algorithm to find an edge-colouring with
∆(_G_) colours.
We remark here that Vizing [1964] proved that for general simple graphs _G_ one
has

(8) ∆(_G_) ≤ _χ_′(_G_) ≤ ∆(_G_) + 1_._

Here ‘simple’ cannot be deleted, as is shown by the graph _G_ with three vertices, where

any two vertices are connected by two parallel edges: then ∆(_G_) = 4 while _χ_′(_G_) = 6.

A theorem ‘dual’ to K ̋onig’s edge-colouring theorem was also shown by K ̋onig.
Note that the edge-colouring number _χ_′(_G_) of a graph _G_ is the minimum number of
matchings needed to cover the edges of a bipartite graph. Dually, one can define:

(9) _ξ_(_G_) := the maximum number of pairwise disjoint edge covers in _G_.

So, in terms of colours, _ξ_(_G_) is the maximum number of colours that can be used in
colouring the edges of _G_ in such a way that at each vertex all colours occur. Hence,
if _δ_(_G_) denotes the minimum degree of _G_, then

(10) _ξ_(_G_) ≤ _δ_(_G_)_._

The triangle _K_[^3]: again is an example having strict inequality. For bipartite graphs
however:

**Corollary 7.4a.** _For any bipartite graph G_ = (_V, E_) _one has_

(11) _ξ_(_G_) = _δ_(_G_)_._

_That is, the maximum number of pairwise disjoint edge covers is equal to the minimum_
_degree._

**Proof.** One may derive from _G_ a bipartite graph _H_, each vertex of which has degree
_δ_(_G_) or 1, by repeated application of the following procedure:

(12) for any vertex _v_ of degree larger than _δ_(_G_), add a new vertex _u_, and replace
one of the edges incident with _v_, {_v, w_} say, by {_u, w_}.

So there is a one-to-one correspondence between the edges of the final graph _H_ and
the edges of _G_. Since _H_ has maximum degree _δ_(_G_), by Theorem 7.4 the edges of _H_
can be coloured with _δ_(_G_) colours such that no two edges of the same colour intersect.
So at any vertex of _H_ of degree _δ_(_G_) all colours occur. This gives a colouring of the
edges of _G_ with _δ_(_G_) colours such that at any vertex of _G_ all colours occur.

**Application 7.4: Scheduling classes.** Suppose we have _n_ classes and _m_ teachers. In the
following scheme it is indicated by an X which classes should be taught by which teachers
(one lesson of one hour a day):

```
class: 1 2 3 4 5 6
teacher: a X X X
b X X X X
c X X X
d X X
e X X X X
f X X X X
g X X X X
```

The question is: What is the minimum timespan in which all lessons can be scheduled?
Theorem 7.4 tells us that all lessons can be scheduled within a timespan of 4 hours.
Indeed, make a bipartite graph _G_ with colour classes _T_ := set of teachers and _C_ := set of
classes, where _t_ ∈ _T_ and _c_ ∈ _C_ are connected if and only if teacher _t_ should teach class _c_;
that is, if there is an X in position (_t,c_) in the scheme.
In the above example _G_ will have maximum degree ∆(_G_) equal to 4. Hence according to
Theorem 7.4, the edge-colouring number _χ_′(_G_) of _G_ is also equal to 4. So we can colour the
edges of _G_ by 4 colours so that no two edges of the same colour have a vertex in common.
That is, we can colour the X’s in the scheme by 4 colours so that there are no two crosses
of the same colour in any row or column. If every colour represent one hour, we obtain a
schedule spanning 4 hours.
This application can be extended to the case where teachers can give more than one
lesson a day to a class. In that case we obtain a bipartite graph with multiple edges.
For any _k_-edge-colouring of a graph _G_ = (_V,E_), we can assume that any two colours
differ by at most 1 in size (if they differ more, one can exchange the two colours on one of the
path components of the union of the two colours, to bring their cardinalities closer together).
That is, each colour has size ⌊|_E_|_/k_⌋ or ⌈|_E_|_/k_⌉. It implies that there is a schedule in which
no more than ⌈|_E_|_/k_⌉ lessons are scheduled simultaneously. So the number of classrooms
needed is ⌈|_E_|_/k_⌉, which is clearly best possible if we want to schedule |_E_| lessons within _k_
hours.

**Exercises**

```
7.9. Determine a schedule for the following scheduling problems:
```

```
(i)
```

```
X X X X
X X X X
X X X X
X X X X
X X X X
```

```
(ii)
```

```
X X X X
X X X X
X X X X
X X X X
X X X X
X X X X
X X X X
```

(iii)

```
1 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18
```

```
J
```

```
L
```

```
R
S
T
```

```
W
X
```

```
Z
```

```
V
```

```
U
```

```
Q
```

```
P
```

```
O
```

```
M
N
```

```
K
```

```
I
```

```
G
H
```

```
F
```

```
E
```

```
C
D
```

```
B
```

```
A
```

```
Y
```

```
(Here the slots to be scheduled are indicated by open cells.)
```

7.10. Let _G_ be the line-graph of some bipartite graph _H_. Derive from K ̋onig’s edge-
colouring theorem (Theorem 7.4) that _ω_(_G_) = _χ_(_G_).
7.11. Derive K ̋onig’s edge-colouring theorem (Theorem 7.4) from the strong perfect graph
theorem.

7.12. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be partitions of a finite set _X_ such that
|_A_[^1]:| = ··· = |_An_| = |_B_[^1]:| = ··· = |_Bn_| = _k_. Show that A and B have _k_ pairwise
disjoint common transversals.

7.13. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be families of subsets of a finite set _X_.
(i) Let _k_ ∈ N. Suppose that _X_ can be partitioned into _k_ partial SDR’s of A, and
that _X_ also can be partitioned into _k_ partial SDR’s of B. Derive that _X_ can be
partitioned into _k common_ partial SDR’s for A and B.
(ii) Show that the minimum number of common partial SDR’s of A and B needed
to cover _X_ is equal to

```
(13) ⌈max
Y⊆X max{
```

```
|Y |
|{i|Ai ∩ Y 6= ∅}|
```

```
,
```

```
|Y |
|{i|Bi ∩ Y 6= ∅}|
```

```
}⌉.
```

```
(Hint: Use Exercise 3.8.)
```

7.14. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be families of subsets of a finite set _X_
and let _k_ ∈ N. Suppose that _X_ has a partition (_Y_[^1]:_,...,Yk_) such that each _Yi_ is an
SDR of A. Suppose moreover that _X_ has a partition (_Z_[^1]:_,...,Zk_) such that each _Zi_
is an SDR of B. Derive that _X_ has a partition (_X_[^1]:_,...,Xk_) such that each _Xi_ is an
SDR both of A and of B.

7.15. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be families of subsets of a finite set _X_
and let _k_ ∈ N. Suppose that _X_ has a partition (_Y_[^1]:_,...,Yn_) such that |_Yi_| = _k_ and
_Yi_ ⊆ _Ai_ for _i_ = 1_,...,n_. Suppose moreover that _X_ has a partition (_Z_[^1]:_,...,Zn_) such
that |_Zi_| = _k_ and _Zi_ ⊆ _Bi_ for _i_ = 1_,...,n_. Derive that _X_ has a partition (_X_[^1]:_,...,Xk_)
such that each _Xi_ is an SDR both of A and of B.
7.16. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bm_) be families of subsets of a finite set and
let _k_ be a natural number. Prove that A and B have _k_ pairwise disjoint common
SDR’s if and only if for all _I,J_ ⊆{[^1]:_,...,n_}:

```
(14)
∣∣
```

```
⋃
```

```
i∈I
```

```
Ai ∩
```

```
⋃
```

```
j∈J
```

```
Bj
∣∣ ≥ k(|I| + |J|− n).
```

```
(Hint: Use Exercise 7.15.)
```

7.17. Let A = (_A_[^1]:_,...,An_) and B = (_B_[^1]:_,...,Bn_) be families of subsets of a finite set _X_.

```
(i) Let k ∈ N. Suppose that A has k pairwise disjoint SDR’s and that also B
has k pairwise disjoint SDR’s. Derive that X can be partitioned into k subsets
X1,...,Xk such that each Xi contains an SDR of A and contains an SDR of B.
```

```
(ii) Show that the maximum number k for which there exists a partition as in (i) is
equal to
```

```
(15) ⌊ min
∅6=I⊆{1,...,n}
```

```
min{
```

```
∣∣⋃
i∈I Ai
∣∣
|I|
```

```
,
```

```
∣∣⋃
i∈I Bi
∣∣
|I|
```

```
}⌋.
```

```
(Hint: Use Exercise 3.7.)
```

### 7.3. Partially ordered sets

```
A partially ordered set is a pair (X,≤) where X is a set and where ≤ is a relation on
X satisfying (for all x, y, z ∈ X):
```

(16) (i) _x_ ≤ _x_;
(ii) if _x_ ≤ _y_ and _y_ ≤ _x_ then _x_ = _y_;
(iii) if _x_ ≤ _y_ and _y_ ≤ _z_ then _x_ ≤ _z_.

```
A subset C of X is called a chain if for all x, y ∈ C one has x ≤ y or y ≤ x. A subset
A of X is called an antichain if for all x, y ∈ A with x 6= y one has x 6≤ y and y 6≤ x.
Note that if C is a chain and A is an antichain then
```

```
(17) |C ∩ A|≤ 1.
```

```
First we observe the following easy min-max relation:
```

```
Theorem 7.5. Let (X,≤) be a partially ordered set, with X finite. Then the mini-
mum number of antichains needed to cover X is equal to the maximum cardinality of
any chain.
```

```
Proof. The fact that the maximum cannot be larger than the minimum follows easily
from (17). To see that the two numbers are equal, define for any element x ∈ X the
height of x as the maximum cardinality of any chain in X with maximum x. For any
i ∈ N, let Ai denote the set of all elements of height i.
Let k be the maximum height of the elements of X. Then A1,... , Ak are antichains
covering X, and moreover there exists a chain of size k.
```

```
Dilworth [1950] proved that the same theorem also holds when we interchange the
words ‘chain’ and ‘antichain’:
```

```
Theorem 7.6 (Dilworth’s decomposition theorem). Let (X,≤) be a partially ordered
set, with X finite. Then the minimum number of chains needed to cover X is equal
to the maximum cardinality of any antichain.
```

**Proof.** We apply induction on |_X_|. The fact that the maximum cannot be larger than
the minimum follows easily from (17). To see that the two numbers are equal, let _α_
be the maximum cardinality of any antichain and let _A_ be an antichain of cardinality
_α_. Define

(18) _A_↓ := {_x_ ∈ _X_ |∃_y_ ∈ _A_ : _x_ ≤ _y_}_,_
_A_↑ := {_x_ ∈ _X_ |∃_y_ ∈ _A_ : _x_ ≥ _y_}_._

Then _A_↓ ∪ _A_↑ = _X_ (since _A_ is a maximum antichain) and _A_↓ ∩ _A_↑ = _A_.
First assume _A_↓ [^6]:= _X_ and _A_↑ [^6]:= _X_. Then by induction _A_↓ can be covered with _α_
chains. Since _A_ ⊆ _A_↓, each of these chains contains exactly one element in _A_. For
each _x_ ∈ _A_, let _Cx_ denote the chain containing _x_. Similarly, there exist _α_ chains _C_′
_x_
(for _x_ ∈ _A_) covering _A_↑, where _C_′
_x_ contains _x_. Then for each _x_ ∈ _A_, _Cx_ ∪ _C_′
_x_ forms a
chain in _X_, and moreover these chains cover _X_.
So we may assume that for each antichain _A_ of cardinality _α_ one has _A_↓ = _X_ or
_A_↑ = _X_. It means that each antichain _A_ of cardinality _α_ is either the set of minimal
elements of _X_ or the set of maximal elements of _X_. Now choose a minimal element
_x_ and a maximal element _y_ of _X_ such that _x_ ≤ _y_. Then the maximum cardinality of
an antichain in _X_ \{_x, y_} is equal to _α_−1 (since each antichain in _X_ of cardinality _α_
contains _x_ or _y_). By induction, _X_ \{_x, y_} can be covered with _α_ − 1 chains. Adding
the chain {_x, y_} yields a covering of _X_ with _α_ chains.

**Application 7.5: Project scheduling.** Suppose you have to perform a project consisting
of several jobs. Each job takes one time-unit, say one hour. Certain jobs have to be done
before other jobs; this relation is given by a partial order on the jobs. Assuming that you
have sufficient workers, the time required to finish the project is equal to the size _γ_ of the
longest chain. Indeed, by Theorem 7.5, the jobs can be split into _γ_ antichains _A_[^1]:_,...,Aγ_;
in fact, these antichains can be chosen such that if _x_ ∈ _Ai_ and _y_ ∈ _Aj_ and _x < y_ then _i < j_.
As in each of these antichains, the jobs can be done simultaneously, we obtain a feasible
schedule.
This is an application quite similar to PERT-CPM (Application 1.4).

**Application 7.6: Bungalow assignment.** Suppose you are the manager of a bungalow
park, with bungalows that can be rented out during the holiday season. There have been
made a number of reservations, each for a connected period of some weeks, like in Figure
7.3. If the number of reservations during any of the weeks in the holiday season is not larger
than the total number of bungalows available, then there exists an allocation of customers to
bungalows, in such a way that no renter has to switch bungalows during his/her stay. This
rule well-known to bungalow park managers, is a special case of Dilworth’s decomposition
theorem.
Indeed, one can make a partial order as follows. Let _X_ be the set of reservations made,
and for any _x,y_ ∈ _X_ let _x < y_ if the last day for reservation _x_ is earlier than or equal to
the first day of reservation _y_.
Then the maximum size of any antichain of (_X,_≤) is equal to the maximum number _n_

```
Figure 7.3
```

of reservations made for any week in the season. By Dilworth’s decomposition theorem, _X_
can be split into _n_ chains. Each chain now gives a series of reservations that can be assigned
to one and the same bungalow.
A similar problem occurs when assigning hotel rooms to hotel guests.

**Application 7.7: Terminal and platform assignment.** A similar problem as in Appli-
cation 7.6 occurs when one has to assign airplanes to terminals at an airport, or trains or
buses to platforms in a train or bus station. The model has to be adapted however, if one
requires a _periodic_ assignment; this occurs for instance if the trains or buses run a periodic
timetable, say with period one hour.

**Exercises**

7.18. Let (_X,_≤) be a partially ordered set. Call a chain _maximal_ if it is not contained
in any other chain. Prove that the maximum number of pairwise disjoint maximal
chains is equal to the minimum cardinality of a set intersecting all maximal chains.

7.19. Derive K ̋onig’s edge cover theorem from Dilworth’s decomposition theorem.

7.20. Let _G_ = (_V,E_) be a bipartite graph, with colour classes _V_[^1]: and _V_[^2]:, with |_V_[^1]:| = |_V_[^2]:| =
_n_. Let _k_ be a natural number. Derive from Dilworth’s decomposition theorem that
the edges of _G_ can be covered by _k_ perfect matchings if and only if for each vertex
cover _W_ ⊆ _V_ the number of edges contained in _W_ is at most _k_(|_W_|− _n_).

7.21. Let I = (_I_[^1]:_,...,In_) be a family of intervals on R, in such a way that each _x_ ∈ R
is contained in at most _k_ of these intervals. Show that I can be partitioned into _k_
classes I[^1]:_,...,_I_k_ so that each I_j_ consists of pairwise disjoint intervals.
7.22. Let _D_ = (_V,A_) be an acyclic directed graph and let _s_ and _t_ be vertices of _D_ such that
each arc of _D_ occurs in at least one _s_−_t_ path. Derive from Dilworth’s decomposition
theorem that the minimum number of _s_ − _t_ paths needed to cover all arcs is equal to
the maximum cardinality of _δ_out(_U_), where _U_ ranges over all subsets of _V_ satisfying
_s_ ∈ _U,t_ 6∈ _U_ and _δ_in(_U_) = ∅.

7.23. A graph _G_ = (_V,E_) is called a _comparability graph_ if there exists a partial order ≤
on _V_ such that for all _u,w_ in _V_ with _u_ [^6]:= _w_ one has:

```
(19) {u,w}∈ E ⇔ u ≤ w or w ≤ u.
```

```
(i) Show that if G is a comparability graph, then ω(G) = χ(G).
(ii) Show that if G is the complement of a comparability graph, then ω(G) = χ(G).
(Hint: Use Dilworth’s decomposition theorem (Theorem 7.6).)
```

7.24. Let (_X,_≤) be a partially ordered set, with _X_ finite. Let C and A denote the collections
of chains and antichains in (_X,_≤), respectively. Let _w_ : _X_ → Z+ be a ‘weight’
function.

```
(i) Show that the maximum weight w(C) of any chain is equal to the minimum value
of
```

```
∑
A∈A λ(A), where the λ(A) range over all nonnegative integers satisfying
```

```
(20)
```

```
∑
```

```
A∈A,x∈A
```

```
λ(A) = w(x)
```

```
for each x ∈ X.
(ii) Show that the maximum weight w(A) of any antichain is equal to the mini-
mum value of
```

```
∑
C∈C λ(C), where the λ(C) range over all nonnegative integers
satisfying
```

```
(21)
```

```
∑
```

```
C∈C,x∈C
```

```
λ(C) = w(x)
```

```
for each x ∈ X.
(iii) Derive that the convex hull of the incidence vectors of antichains (as vectors in
RX) is equal to the set of all vectors f ∈ RX
+ satisfying f(C) ≤ 1 for each chain
C.
[For any finite set X and any subset Y of X, define the incidence vector χY ∈ RX
of Y as:
```

```
(22)
χY
x := 1 if x ∈ Y ;
:= 0 if x 6∈ Y .]
```

```
(iv) Derive also that the convex hull of the incidence vectors of chains (as vectors
in RX) is equal to the set of all vectors f ∈ RX
+ satisfying f(A) ≤ 1 for each
antichain A.
```

7.25. Derive Dilworth’s decomposition theorem (Theorem 7.6) from the strong perfect
graph theorem.

### 7.4. Perfect graphs

We now consider a general class of graphs, the ‘perfect’ graphs, that turn out to unify
several results in combinatorial optimization, in particular, min-max relations and
polyhedral characterizations.
As we saw before, the clique number _ω_(_G_) and the colouring number _χ_(_G_) of a
graph _G_ = (_V, E_) are related by the inequality:

(23) _ω_(_G_) ≤ _χ_(_G_)_._

There are graphs that have strict inequality; for instance, the circuit _C_[^5]: on five
vertices.
Having equality in (23) does not say that much about the internal structure of a
graph: any graph _G_ = (_V, E_) can be extended to a graph _G_′ = (_V_ ′_, E_′) satisfying
_ω_(_G_′) = _χ_(_G_′), simply by adding to _G_ a clique of size _χ_(_G_), disjoint from _V_.
However, if we require that equality in (23) holds for each induced subgraph of
_G_, we obtain a much more powerful condition. The idea for this was formulated by
Berge [1963]. He defined a graph _G_ = (_V, E_) te be _perfect_ if _ω_(_G_′) = _χ_(_G_′) holds for
each induced subgraph _G_′ of _G_.
Several classes of graphs could be shown to be perfect, and Berge [1961,1963]
observed the important phenomenon that for several classes of graphs that were shown
to be perfect, also the class of complementary graphs is perfect. (The _complement_
or the _complementary graph G_ of a graph _G_ = (_V, E_) is the graph with vertex set
_V_ , where any two distinct vertices in _V_ are adjacent in _G_ if and only if they are
nonadjacent in _G_.)
Berge therefore conjectured that the complement of any perfect graph is perfect
again. This conjecture was proved by Lov ́asz [1972b], and his _perfect graph theorem_
forms the kernel of perfect graph theory. It has several other theorems in graph theory
as consequence. Lov ́asz [1972a] gave the following stronger form of the conjecture,
which we show with the elegant linear-algebraic proof found by Gasparian [1996].

**Theorem 7.7.** _A graph G is perfect if and only if ω_(_G_′)_α_(_G_′) ≥ |_V_ (_G_′)| _for each_
_induced subgraph G_′ _of G._

**Proof.** Necessity is easy, since if _G_ is perfect, then _ω_(_G_′) = _χ_(_G_′) for each induced
subgraph _G_′ of _G_, and since _χ_(_G_′)_α_(_G_′) ≥|_V_ (_G_′)| for any graph _G_′.
To see sufficiency, suppose to the contrary that there exists an imperfect graph _G_
satisfying the condition, and choose such a graph with |_V_ (_G_)| minimal. So _χ_(_G_) _>_
_ω_(_G_), while _χ_(_G_′) = _ω_(_G_′) for each induced subgraph _G_′ [^6]:= _G_ of _G_.
Let _ω_ := _ω_(_G_) and _α_ := _α_(_G_). We can assume that _V_ (_G_) = {[^1]:_,... , n_}.
We first construct

(24) stable sets _C_[^0]:_,... , Cαω_ such that each vertex is covered by exactly _α_ of the

```
Ci.
```

Let _C_[^0]: be any stable set in _G_ of size _α_. By the minimality of _G_, we know that for
each _v_ ∈ _C_[^0]:, the subgraph of _G_ induced by _V_ (_G_) \{_v_} is perfect, and that hence its
colouring number is at most _ω_ (as its clique number is at most _ω_); therefore _V_ (_G_)\{_v_}

can be partitioned into _ω_ stable sets. Doing this for each _v_ ∈ _C_[^0]:, we obtain stable
sets as in (24).
Now for each _i_ = 0_,... , αω_, there exists a clique _Ki_ of size _ω_ with _Ki_ ∩ _Ci_ = ∅.
Otherwise, the subgraph _G_′ of _G_ induced by _V_ (_G_) \ _Ci_ would have _ω_(_G_′) _< ω_, and
hence it has colouring number at most _ω_ − 1. Adding _Ci_ as a colour would give an
_ω_-vertex colouring of _G_, contradicting the assumption that _χ_(_G_) _> ω_(_G_).
Then, if _i_ [^6]:= _j_ with 0 ≤ _i, j_ ≤ _αω_, we have |_Kj_ ∩ _Ci_| = 1. This follows from the
fact that _Kj_ has size _ω_ and intersects each _Ci_ in at most one vertex, and hence, by
(24), it intersects _αω_ of the _Ci_. As _Kj_ ∩ _Cj_ = ∅, we have that |_Kj_ ∩ _Ci_| = 1 if _i_ [^6]:= _j_.
Now consider the (_αω_ + 1) × _n_ incidence matrices _M_ = (_mi,j_) and _N_ = (_ni,j_)
of _C_[^0]:_,... , Cαω_ and _K_[^0]:_,... , Kαω_ respectively. So _M_ and _N_ are 0_,_1 matrices, with
_mi,j_ = 1 ⇔ _j_ ∈ _Ci_, and _ni,j_ = 1 ⇔ _j_ ∈ _Ki_, for _i_ = 0_,... , αω_ and _j_ = 1_,... , n_. By
the above, _M NT_ = _J_ − _I_, where _J_ is the (_αω_ + 1) × (_αω_ + 1) all-1 matrix, and _I_ the
(_αω_ + 1) × (_αω_ + 1) identity matrix. As _J_ − _I_ has rank _αω_ + 1, we have _n_ ≥ _αω_ + 1.
This contradicts the condition given in the theorem.

```
This implies:
```

**Corollary 7.7a** ((Lov ́asz’s) perfect graph theorem)**.** _The complement of a perfect_
_graph is perfect again._

**Proof.** Directly from Theorem 7.7, as the condition given in it is maintained under
taking the complementary graph.

In fact, Berge [1963] also made an even stronger conjecture, which was proved
in 2002 by Chudnovsky, Robertson, Seymour, and Thomas (we mentioned this in
Section 7.1 in a different but equivalent form):
**Strong perfect graph theorem.** A graph _G_ is perfect if and only if _G_ does not
contain any odd circuit _C_[^2]:_k_+1 with _k_ ≥ 2 or its complement as an induced subgraph.
We now show how several theorems we have seen before follow as consequences
from the perfect graph theorem. First observe that trivially, any bipartite graph _G_ is
perfect. This implies K ̋onig’s edge cover theorem (Theorem 3.3a):

**Corollary 7.7b** (K ̋onig’s edge cover theorem)**.** _The complement of a bipartite graph_
_is perfect. Equivalently, the edge cover number of any bipartite graph (without isolated_
_vertices) is equal to its stable set number._

**Proof.** Directly from the perfect graph theorem. Note that if _G_ is a bipartite graph,
then its cliques have size at most 2; hence _χ_(_G_) is equal to the edge cover number of
_G_ if _G_ has no isolated vertices.
Note moreover that the class of complements of bipartite graphs is closed under
taking induced subgraphs. Hence the second statement in the Corollary indeed is
equivalent to the first.

We saw in Section 3.3 that by Gallai’s theorem (Theorem 3.1), K ̋onig’s edge cover
theorem directly implies K ̋onig’s matching theorem (Theorem 3.3), saying that the
matching number of a bipartite graph _G_ is equal to its vertex cover number. That is,
the stable set number of the line graph _L_(_G_) of _G_ is equal to the minimum number
of cliques of _L_(_G_) that cover all vertices of _L_(_G_). As this is true for any induced
subgraph of _L_(_G_) we know that the complement _L_(_G_) of the line graph _L_(_G_) of any
bipartite graph _G_ is perfect.
Hence with the perfect graph theorem we obtain K ̋onig’s edge-colouring theorem
(Theorem 7.4):

**Corollary 7.7c** (K ̋onig’s edge-colouring theorem)**.** _The line graph of a bipartite graph_
_is perfect. Equivalently, the edge-colouring number of any bipartite graph is equal to_
_its maximum degree._

**Proof.** Again directly from K ̋onig’s matching theorem and the perfect graph theorem.

We can also derive Dilworth’s decomposition theorem (Theorem 7.6) easily from
the perfect graph theorem. Let (_V,_≤) be a partially ordered set. Let _G_ = (_V, E_) be
the graph with:

(25) _uv_ ∈ _E_ if and only if _u < v_ or _v < u_.

Any graph _G_ obtained in this way is called a _comparability graph_.
As Theorem 7.5 we saw the following easy ‘dual’ form of Dilworth’s decomposition
theorem:

**Theorem 7.8.** _In any partially ordered set_ (_V,_≤)_, the maximum size of any chain_
_is equal to the minimum number of antichains needed to cover V._

**Proof.** For any _v_ ∈ _V_ define the _height_ of _v_ as the maximum size of any chain in _V_
with maximum element _v_. Let _k_ be the maximum height of any element _v_ ∈ _V_. For
_i_ = 1_,... , k_ let _Ai_ be the set of elements of height _i_. Then _A_[^1]:_,... , Ak_ are antichains
covering _V_ , and moreover, there is a chain of size _k_, since there is an element of height
_k_.

Equivalently, we have _ω_(_G_) = _χ_(_G_) for any comparability graph. As the class of
comparability graphs is closed under taking induced subgraphs we have:

**Corollary 7.8a.** _Any comparability graph is perfect._

**Proof.** Directly from Theorem 7.8.

```
So by the perfect graph theorem:
```

**Corollary 7.8b.** _The complement of any comparability graph is perfect._

**Proof.** Directly from Corollary 7.8a and the perfect graph theorem (Corollary 7.7a).

```
That is:
```

**Corollary 7.8c** (Dilworth’s decomposition theorem)**.** _In any partially ordered set_
(_V,_≤)_, the maximum size of any antichain is equal to the minimum number of chains_
_needed to cover V._

**Proof.** Directly from Corollary 7.8b.

A further application of the perfect graph theorem is to ‘chordal graphs’, which
we describe in the next section.
We note here that it was shown with the help of the ‘ellipsoid method’ that
there exists a polynomial-time algorithm for finding a maximum-size clique and a
minimum vertex-colouring in any perfect graph (Gr ̈otschel, Lov ́asz, and Schrijver
[1981]). However no _combinatorial_ polynomial-time algorithm is known for these
problems.

**Exercises**

7.26. Show that the graph obtained from the Paris M ́etro network (see Application 7.1) is
perfect.

7.27. Show that Theorem 7.7 is implied by the strong perfect graph theorem.

### 7.5. Chordal graphs

We finally consider a further class of perfect graphs, the ‘chordal graphs’ (or ‘rigid
circuit graphs’ or ‘triangulated graphs’). A graph _G_ is called _chordal_ if each circuit
in _G_ of length at least 4 has a chord. (A _chord_ is an edge connecting two vertices of
the circuit that do not form two neighbours in the circuit.)

For any set _A_ of vertices let _N_(_A_) denote the set of vertices not in _A_ that are
adjacent to at least one vertex in _A_. Call a vertex _v simplicial_ if _N_({_v_}) is a clique
in _G_.
Dirac [1961] showed the following basic property of chordal graphs:

**Theorem 7.9.** _Each chordal graph G contains a simplicial vertex._

**Proof.** We may assume that _G_ has at least two nonadjacent vertices _a, b_. Let _A_
be a maximal nonempty subset of _V_ such that _G_|_A_ is connected and such that
_A_∪_N_(_A_) [^6]:= _V_. Such a subset _A_ exists as _G_|{_a_} is connected and {_a_}∪_N_({_a_}) [^6]:= _V_.
Let _B_ := _V_ \(_A_∪_N_(_A_)). Then each vertex _v_ in _N_(_A_) is adjacent to each vertex in
_B_, since otherwise we could increase _A_ by _v_. Moreover, _N_(_A_) is a clique, for suppose
that _u, w_ ∈ _N_(_A_) are nonadjacent. Choose _v_ ∈ _B_. Let _P_ be a shortest path in
_A_ ∪ _N_(_A_) connecting _u_ and _w_. Then _P_ ∪{_u, v, w_} would form a circuit of length at
least 4 without chords, a contradiction.
Now inductively we know that _G_|_B_ contains a vertex _v_ that is simplicial in _G_|_B_.
Since _N_(_A_) is a clique and since each vertex in _B_ is connected to each vertex in _N_(_A_),
_v_ is also simplicial in _G_.

```
This implies a result of Hajnal and Sur ́anyi [1958]:
```

**Theorem 7.10.** _The complement of any chordal graph is perfect._

**Proof.** Let _G_ = (_V, E_) be a chordal graph. Since the class of chordal graphs is closed
under taking induced subgraphs, it suffices to show _ω_(_G_) ≥ _χ_(_G_).
By Theorem 7.1, _G_ has a simplicial vertex _v_. So _K_ := {_v_}∪ _N_({_v_}) is a clique.
Let _G_′ be the subgraph of _G_ induced by _V_ \_K_. By induction we have _ω_(_G_′) = _χ_(_G_′).
Now _ω_(_G_) ≥ _ω_(_G_′) + 1, since we can add _v_ to any clique of _G_′. Similarly, _χ_(_G_) ≤
_χ_(_G_′) + 1, since we can add _K_ to any colouring of _G_′. Hence _ω_(_G_) ≥ _χ_(_G_).

```
With Lov ́asz’s perfect graph theorem, this implies the result of Berge [1960]:
```

**Corollary 7.10a.** _Any chordal graph is perfect._

**Proof.** Directly from Theorem 7.10 and the perfect graph theorem (Corollary 7.7a).

We can characterize chordal graphs in terms of subtrees of a tree _T_. Let S be a
collection of nonempty subtrees of a tree _T_. The _intersection graph_ of S is the graph
with vertex set S, where two vertices _S, S_′ are adjacent if and only if they intersect
(in at least one vertex).
The class of graphs obtained in this way coincides with the class of chordal graphs.
To see this, we first show the following elementary lemma:

**Lemma 7.1.** _Let_ S _be a collection of pairwise intersecting subtrees of a tree T. Then_
_there is a vertex of T contained in all subtrees in_ S_._

**Proof.** By induction on |_V T_|. If |_V T_| = 1 the lemma is trivial, so assume |_V T_|≥ 2.
Let _t_ be an end vertex of _T_. If there exists a subtree in S consisting only of _t_, the
lemma is trivial. Hence we may assume that each subtree in S containing _t_ also
contains the neighbour of _t_. So deleting _t_ from _T_ and from all subtrees in S gives the
lemma by induction.

```
Then:
```

**Theorem 7.11.** _A graph is chordal if and only if it is isomorphic to the intersection_
_graph of a collection of subtrees of some tree._

**Proof.** _Necessity._ Let _G_ = (_V, E_) be chordal. By Theorem 7.9, _G_ contains a simplicial
vertex _v_. By induction, the subgraph _G_−_v_ of _G_ is the intersection graph of a collection
S of subtrees of some tree _T_. Let S′ be the subcollection of S corresponding to the
set _N_ of neighbours of _v_ in _G_. As _N_ is a clique, S′ consists of pairwise intersecting
subtrees. Hence, by Lemma 7.1 these subtrees have a vertex _t_ of _T_ in common. Now
we extend _T_ and all subtrees in S′ with a new vertex _t_′ and a new edge _tt_′. Moreover,
we introduce a new subtree {_t_′} representing _v_. In this way we obtain a subtree
representation for _G_.
_Sufficiency._ Let _G_ be the intersection graph of some collection S of subtrees
of some tree _T_. Suppose that _G_ contains a chordless circuit _Ck_ with _k_ ≥ 4. Let
_Ck_ be the intersection graph of _S_[^1]:_,... , Sk_ ∈ S, with _S_[^1]: and _S_[^2]: intersecting. Then
_S_[^1]:_, S_[^2]:_, S_[^3]:∪···∪_Sk_ are three subtrees of _T_ that are pairwise intersecting. So by Lemma
7.1, _T_ has a vertex _v_ contained in each of these three subtrees. So _v_ ∈ _S_[^1]: ∩ _S_[^2]: ∩ _Si_
for some _i_ ∈{[^3]:_,... , k_}. This yields a chord in _Ck_.

This theorem enables us to interpret the perfectness of chordal graphs in terms of
trees:

**Corollary 7.11a.** _Let_ S _be a collection of nonempty subtrees of a tree T. Then_
_the maximum number of pairwise vertex-disjoint trees in_ S _is equal to the minimum_
_number of vertices of T intersecting each tree in_ S_._

**Proof.** Directly from Theorems 7.10 and 7.11, using Lemma 7.1.

```
Similarly we have:
```

**Corollary 7.11b.** _Let_ S _be a collection of subtrees of a tree T. Let k be the max-_
_imum number of times that any vertex of T is covered by trees in_ S_. Then_ S _can_

_be partitioned into subcollections_ S[^1]:_,... ,_S_k such that each_ S_i consists of pairwise_

_vertex-disjoint trees._

**Proof.** Directly from Corollary 7.10a and Theorem 7.11, again using Lemma 7.1.

**Exercises**

7.28. Show that a graph _G_ = (_V,E_) is chordal if and only if each induced subgraph has a
simplicial vertex.

7.29. Show that a graph is an interval graph if and only if it is chordal and its complement
is a comparability graph.

7.30. Derive from the proof of Theorem 7.9 that each chordal graph is either a clique or
contains two nonadjacent simplicial vertices.

7.31. Let _G_ be a chordal graph. Derive from the proof of Theorem 7.9 that each vertex
_v_ that is nonadjacent to at least one vertex _w_ [^6]:= _v_, is nonadjacent to at least one
simplicial vertex _w_ [^6]:= _v_.

7.32. Show that a graph _G_ = (_V,E_) is chordal if and only if the edges of _G_ can be oriented
so as to obtain a directed graph _D_ = (_V,A_) with the following properties:

```
(26) (i)D is acyclic;
(ii)if (u,v) and (u,w) belong to A then (v,w) or (w,v) belongs to A.
```

## 8. Integer linear programming and

## totally unimodular matrices

### 8.1. Integer linear programming

Many combinatorial optimization problems can be described as maximizing a linear
function _cTx_ over the _integer_ vectors in some polyhedron _P_ = {_x_ | _Ax_ ≤ _b_}. (A
vector _x_ ∈ R_n_ is called _integer_ if each component is an integer, i.e., if _x_ belongs to
Z_n_.)
So this type of problems can be described as:

(1) max{_cTx_ | _Ax_ ≤ _b_;_x_ ∈ Z_n_}_._

Such problems are called _integer linear programming_ problems. They consist of max-
imizing a linear function over the intersection _P_ ∩ Z_n_ of a polyhedron _P_ with the set
Z_n_ of integer vectors.

**Example.** Consider a graph _G_ = (_V, E_). Then the problem of finding a matching
of maximum cardinality can be described as follows. Let _A_ be the _V_ × _E_ incidence
matrix of _G_. So the rows of _A_ are indexed by the vertices of _G_, while the columns of
_A_ are indexed by the edges of _G_ and for any _v_ ∈ _V_ and _e_ ∈ _E_:

(2) _Av,e_ := 1 if _v_ ∈ _e_;
:= 0 if _v_ 6∈ _e_.

Now finding a maximum-cardinality matching is equivalent to:

(3) maximize

```
∑
```

```
e∈E
```

```
xe
```

```
subject to
```

```
∑
```

```
e∋v
```

```
xe ≤ 1 for each v ∈ V ,
```

```
xe ≥ 0 for each e ∈ E,
xe ∈ Z for each e ∈ E.
```

This is the same as:

(4) max{[^1]:_Tx_ | _x_ ≥ 0;_Ax_ ≤ [^1]:;_x_ integer}_,_

where [^1]: denotes an all-one vector, of appropriate size.

```
Clearly, always the following holds:
```

(5) max{_cTx_ | _Ax_ ≤ _b_;_x_ integer}≤ max{_cTx_ | _Ax_ ≤ _b_}_._

The above example, applied to the graph _K_[^3]: shows that strict inequality can hold.
This implies, that generally one will have strict inequality in the following duality
relation:

(6) max{_cTx_ | _Ax_ ≤ _b_;_x_ integer}≤ min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_;_y_ integer}_._

A polytope _P_ is called _integer_ if each of its vertices is an integer vector. Clearly,
if a polytope _P_ = {_x_ | _Ax_ ≤ _b_} is integer, then the _LP-_problem

(7) max{_cTx_ | _Ax_ ≤ _b_}

has an integer optimum solution. So in that case,

(8) max{_cTx_ | _Ax_ ≤ _b_;_x_ integer} = max{_cTx_ | _Ax_ ≤ _b_}_._

In Exercise 8.5 below we shall see that in a sense also the converse holds.

No polynomial-time algorithm is known to exist for solving an integer linear pro-
gramming problem in general. In fact, the general integer linear programming prob-
lem is NP-complete, and it is conjectured that no polynomial-time algorithm exists.
However, for special classes of integer linear programming problems, polynomial-
time algorithms have been found. These classes often come from combinatorial prob-
lems, like the matching problem above.

**Exercises**

```
8.1. Let P be a polytope. Prove that the set conv.hull(P ∩ Zn) is again a polytope.
```

```
8.2. Let P = {x | Ax ≤ b} be a polyhedron, where A is a rational matrix. Show that the
set conv.hull(P ∩ Zn) is again a polyhedron.
```

```
8.3. Let G = (V,E) be a graph. Describe the problem of finding a vertex cover of minimum
cardinality as an integer linear programming problem.
```

```
8.4. Let G = (V,E) be a graph. Describe the problem of finding a clique (= complete
subgraph) of maximum cardinality as an integer linear programming problem.
```

```
8.5. Show that a polytope P is integer if and only if for each vector c, the linear program-
ming problem max{cTx | Ax ≤ b} has an integer optimum solution.
```

### 8.2. Totally unimodular matrices

Total unimodularity of matrices turns out to form an important tool in studying
integer vectors in polyhedra.
A matrix _A_ is called _totally unimodular_ if each square submatrix of _A_ has determi-
nant equal to 0, +1, or −1. In particular, each entry of a totally unimodular matrix
is 0, +1, or −1.
A link between total unimodularity and integer linear programming is given by
the following fundamental result.

**Theorem 8.1.** _Let A be a totally unimodular m_ × _n matrix and let b_ ∈ Z_m. Then_
_each vertex of the polyhedron_

(9) _P_ := {_x_ | _Ax_ ≤ _b_}

_is an integer vector._

**Proof.** Let _A_ have order _m_ × _n_. Let _z_ be a vertex of _P_. By Theorem 2.2, the
submatrix _Az_ has rank _n_. So _Az_ has a nonsingular _n_ × _n_ submatrix _A_′. Let _b_′ be the
part of _b_ corresponding to the rows of _A_ that occur in _A_′.
Since, by definition, _Az_ is the set of rows _ai_ of _A_ for which _aiz_ = _bi_, we know
_A_′_z_ = _b_′. Hence _z_ = (_A_′)−[^1]_b_′. However, since |det_A_′| = 1, all entries of the matrix
(_A_′)−[^1] are integer. Therefore, _z_ is an integer vector.

As a direct corollary we have a similar result for polyhedra in general (not neces-
sarily having vertices). Define a polyhedron _P_ to be _integer_ if for each vector _c_ for
which

(10) max{_cTx_ | _x_ ∈ _P_}

is finite, the maximum is attained by some integer vector. So:

(11) if _P_ = {_x_ | _Ax_ ≤ _b_} where _A_ is an _m_×_n_ matrix of rank _n_, then _P_ is integer
if and only if each vertex of _P_ is integer.

```
Then we have:
```

**Corollary 8.1a.** _Let A be a totally unimodular m_ × _n matrix and let b_ ∈ Z_m. Then_
_the polyhedron_

(12) _P_ := {_x_ | _Ax_ ≤ _b_}

_is an integer polyhedron._

**Proof.** Let _x_∗ be an optimum solution of (10). Choose integer vectors _d_′_, d_′′ ∈ Z_n_
such that _d_′ ≤ _x_∗ ≤ _d_′′. Consider the polyhedron

(13) _Q_ := {_x_ ∈ R_n_ | _Ax_ ≤ _b_;_d_′ ≤ _x_ ≤ _d_′′}_._

So _Q_ is bounded.
Moreover, _Q_ is the set of all vectors _x_ satisfying

(14)

```


```

```
A
−I
I
```

```

```

```
x ≤
```

```


```

```
b
−d′
d′′
```

```

```

```
.
```

Now the matrix here is again totally unimodular (this follows easily from the total
unimodularity of _A_). Hence by Theorem 8.1, _Q_ is an integer polytope. This implies
that the linear programming problem max{_cTx_ | _x_ ∈ _Q_} is attained by some integer
vector ̃_x_.
But then ̃_x_ is also an optimum solution for the original LP-problem max{_cTx_ |
_Ax_ ≤ _b_}. Indeed, ̃_x_ satisfies _Ax_ ̃ ≤ _b_, as ̃_x_ belongs to _Q_. Moreover,

(15) _cTx_ ̃ ≥ _cTx_∗ = max{_cTx_ | _Ax_ ≤ _b_}_,_

implying that ̃_x_ is an optimum solution.

It follows that each linear programming problem with integer data and totally
unimodular constraint matrix has integer optimum primal and dual solutions:

**Corollary 8.1b.** _Let A be a totally unimodular m_ × _n matrix, let b_ ∈ Z_m and let_
_c_ ∈ Z_n. Then both problems in the LP-duality equation:_

(16) max{_cTx_ | _Ax_ ≤ _b_} = min{_yTb_ | _y_ ≥ 0;_yTA_ = _cT_}

_have integer optimum solutions (if the optima are finite)._

**Proof.** Directly from Corollary 8.1a, using the fact that with _A_ also the matrix

(17)

```


```

```
−I
AT
−AT
```

```


```

is totally unimodular.

```
Hoffman and Kruskal [1956] showed, as we shall see below, that the above property
```

more or less characterizes total unimodularity.
To derive this result, define an _m_ × _n_ matrix _A_ to be _unimodular_ if it has rank _m_
and each _m_ × _m_ submatrix has determinant equal to 0, +1, or −1. It is easy to see
that a matrix _A_ is totally unimodular if and only if the matrix [_I A_] is unimodular.
We follow the proof of Hoffman and Kruskal’s result given by Veinott and Dantzig
[1968]. As a preparation one first shows:

**Theorem 8.2.** _Let A be an integer m_ × _n matrix of rank m. Then A is unimodular_
_if and only if for each integer vector b the polyhedron_

(18) _P_ = {_x_ | _x_ ≥ 0;_Ax_ = _b_}

_is integer._

**Proof.** _Necessity._ First suppose that _A_ is unimodular. Let _b_ be an integer vector.
Let _D_ be the matrix

(19) _D_ :=

```


```

```
−I
A
−A
```

```

```

```
 and f :=
```

```


```

```
0
b
−b
```

```

```

```
.
```

Note that the system _x_ ≥ [^0]:_, Ax_ = _b_ is the same as _Dx_ ≤ _f_.
Since _D_ has rank _n_, we know that for each _c_ ∈ R_n_, the linear programming
problem

(20) max{_cTx_ | _x_ ≥ 0;_Ax_ = _b_} = max{_cTx_ | _Dx_ ≤ _f_}

is attained by a _vertex z_ of _P_ (if the optima are finite).
Now consider the matrix _Dz_. By definition, this is the submatrix of _D_ consisting
of those rows _Di_ of _D_ which have equality in _Dz_ ≤ _f_.
Clearly, _Dz_ contains all rows of _D_ that are in _A_ and in −_A_. Since _A_ has rank _m_,
this implies that _Dz_ contains a nonsingular _n_×_n_ matrix _B_ that fully contains _A_ and
moreover, part of −_I_. Since _A_ is unimodular, det_B_ equals +1 or −1. Let _f_′ be the
part of _f_ corresponding to _B_. So _Bz_ = _f_′, and hence _z_ = _B_−[^1]_f_′. As |det_B_| = 1, it
follows that _z_ is an integer vector.

_Sufficiency._ Suppose that _P_ = {_x_ | _x_ ≥ 0;_Ax_ = _b_} is integer, for each choice of
an integer vector _b_. Let _B_ be an _m_ × _m_ nonsingular submatrix of _A_. We must show
that det_B_ equals +1 or −1.
Without loss of generality, we may assume that _B_ consists of the first _m_ columns
of _A_.
It suffices to show that _B_−[^1]_v_ is an integer vector for each choice of an integer
vector _v_. (This follows from the fact that then _B_−[^1] itself is an integer matrix, and

hence (det_B_)−[^1]=det(_B_−[^1]) is an integer. This implies that det_B_ equals +1 or −1.)
So let _v_ be an integer vector. Then there exists an integer vector _u_ ∈ R_m_ such
that

(21) _z_ := _u_ + _B_−[^1]_v >_ [^0]:_._

Define

(22) _b_ := _Bz._

So _b_ = _Bz_ = _Bu_ + _BB_−[^1]_v_ = _Bu_ + _v_ is an integer vector.
Let _z_′ arise from _z_ by adding zero-components to _z_ so as to obtain a vector in R_n_.
So

(23) _z_′ =

```
(
z
0
```

```
)
,
```

where [^0]: is the all-zero vector in R_n_−_m_.
Then _z_′ is a vertex of the polyhedron _P_ (since _z_′ ∈ _P_ and since there are _n_ linearly
independent rows in the matrix _D_ for which _Dz_ ≤ _f_ holds with equality).
So _z_′ is integer, and hence

(24) _B_−[^1]_v_ = _z_ − _u_

is an integer vector.

```
This gives the result of Hoffman and Kruskal [1956]:
```

**Corollary 8.2a** (Hoffman-Kruskal theorem)**.** _Let A be an integer m_ × _n matrix._
_Then A is totally unimodular if and only if for each integer vector b the polyhedron_

(25) _P_ = {_x_ | _x_ ≥ 0;_Ax_ ≤ _b_}

_is integer._

**Proof.** _Necessity._ Directly from Corollary 8.1a.

_Sufficiency._ Let _P_ be an integer polyhedron, for each choice of an integer vector
_b_. We show that, for each choice of _b_ ∈ Z_m_, each vertex _z_ of the polyhedron

(26) _Q_ := {_z_ | _z_ ≥ 0; [_I A_]_z_ = _b_}_._

is integer. Indeed, _z_ can be decomposed as

(27) _z_ =

```
(
z′
z′′
```

```
)
,
```

where _z_′ ∈ R_m_ and _z_′′ ∈ R_n_. So _z_′ = _b_ − _Az_′′.
Then _z_′′ is a vertex of _P_. [This follows from the fact that if _z_′′ would be equal to
1
[^2]:(_v_ + _w_) for two other points _v, w_ in _P_, then

(28) _z_′ = _b_ − _Az_′′ =

```
1
2
```

```
(b − Av) +
```

```
1
2
```

```
(b − Aw).
```

Hence

(29) _z_ =

```
(
z′
z′′
```

```
)
=
```

```
1
2
```

```
(
b − Av
v
```

```
)
+
```

```
1
2
```

```
(
b − Aw
w
```

```
)
```

_._

This contradicts the fact that _z_ is a vertex of _Q_.]
So, by assumption, _z_′′ is integer. Hence also _z_′ = _b_ − _Az_′′ is integer, and hence _z_
is integer.
So for each choice of _b_ in Z_m_, the polyhedron _Q_ is integer. Hence, by Theorem
8.2, the matrix [_I A_] is unimodular. This implies that _A_ is totally unimodular.

**Exercises**

```
8.6. Show that an integer matrix A is totally unimodular if and only if for all integer
vectors b and c, both sides of the linear programming duality equation
```

```
(30) max{cTx | x ≥ 0;Ax ≤ b} = min{yTb | y ≥ 0;yTA ≥ cT}
```

```
are attained by integer optimum solutions x and y (if the optima are finite).
```

```
8.7. Give an example of an integer matrix A and an integer vector b such that the poly-
hedron P := {x | Ax ≤ b} is integer, while A is not totally unimodular.
```

```
8.8. Let A be a totally unimodular matrix. Show that the columns of A can be split
into two classes such that the sum of the columns in one class, minus the sum of the
columns in the other class, gives a vector with entries 0, +1, and −1 only.
```

```
8.9. Let A be a totally unimodular matrix and let b be an integer vector. Let x be an
integer vector satisfying x ≥ 0;Ax ≤ 2b. Show that there exist integer vectors x′ ≥ 0
```

### 8.3. Totally unimodular matrices from bipartite

### graphs

Let _A_ be the _V_ × _E_ incidence matrix of a graph _G_ = (_V, E_) (cf. (2)). The matrix

_A_ generally is not totally unimodular. E.g., if _G_ is the complete graph _K_[^3]: on three
vertices, then the determinant of _A_ is equal to +2 or −2.
However, the following can be proved:

**Theorem 8.3.** _Graph G is bipartite if and only if its incidence matrix A is totally_
_unimodular._

**Proof.** _Sufficiency._ Let _A_ be totally unimodular. Suppose _G_ is not bipartite. Then
_G_ contains an odd circuit, say with vertices _v_[^1]:_,... , vk_ and edges _e_[^1]:_,... , ek_. The sub-
matrix of _A_ on the rows indexed by _v_[^1]:_,... , vk_ and the columns indexed by _e_[^1]:_,... , ek_,
is of type

(31)

```

```

```

```

```
1 1 0 ··· ··· 0 0
0 1 1 ··· ··· 0 0
0 0 1 ··· ··· 0 0
```

..................
...
...
.........
...
0 0 0 ··· ··· 1 1
1 0 0 ··· ··· 0 1

```

```

```

```

```
,
```

up to permutation of rows and columns.
It is not difficult to see that matrix (31) has determinant 2. This contradicts the
total unimodularity of _A_.

_Necessity._ Let _G_ be bipartite. Let _B_ be a square submatrix of _A_, of order _t_ × _t_,
say. We show that det_B_ equal 0 or ±1 by induction on _t_. If _t_ = 1, the statement is
trivial.
So let _t >_ 1. We distinguish three cases.
**Case 1.** _B has a column with only_ [^0]:_’s._ Then det_B_=0.
**Case 2.** _B has a column with exactly one_ [^1]:_._ In that case we can write (possibly
after permuting rows or columns):

(32) _B_ =

```
(
1 bT
0 B′
```

```
)
,
```

for some matrix _B_′ and vector _b_, where [^0]: denotes the all-zero vector in R_t_−[^1]. By the
induction hypothesis, det_B_′ ∈{[^0]:_,_±[^1]:}. Hence, by (32), det_B_ ∈{[^0]:_,_±[^1]:}.

**Case 3.** _Each column of B contains exactly two_ [^1]:_’s._ Then, since _G_ is bipartite,
we can write (possibly after permuting rows):

(33) _B_ =

```
(
B′
B′′
```

```
)
,
```

in such a way that each column of _B_′ contains exactly one 1 and each column of
_B_′′ contains exactly one 1. So adding up all rows in _B_′ gives the all-one vector, and
also adding up all rows in _B_′′ gives the all-one vector. Therefore, the rows of _B_ are
linearly dependent, and hence det_B_=0.

As direct corollaries of this theorem, together with Corollary 8.1b, we obtain some
theorems of K ̋onig. First:

**Corollary 8.3a** (K ̋onig’s matching theorem)**.** _Let G be a bipartite graph. Then the_
_maximum cardinality of a matching in G is equal to the minimum cardinality of a_
_vertex cover in G._

**Proof.** Clearly, the maximum cannot be larger than the minimum. To see that
equality holds, let _A_ be the _V_ × _E_ incidence matrix of _G_. Then by Corollary 8.1b,
both optima in the LP-duality equation

(34) max{[^1]:_Tx_ | _x_ ≥ 0;_Ax_ ≤ [^1]:} = min{_yT_[^1]: | _y_ ≥ 0;_yTA_ ≥ [^1]:}

are attained by integer optimum solutions _x_∗ and _y_∗.
Since _x_∗ is an integer vector satisfying _x_ ≥ 0;_Ax_ ≤ [^1]:, _x_∗ is a {[^0]:_,_[^1]:} vector. Let
_M_ be the set of edges _e_ of _G_ for which _x_∗
_e_ = 1. Then _M_ is a matching, since _Ax_∗ ≤[^1]
holds, implying that for each vertex _v_ there is at most one edge _e_ with _x_∗
_e_ = 1.
Moreover, the cardinality |_M_| of _M_ satisfies |_M_| = [^1]:_Tx_∗. So |_M_| is equal to the
maximum in (34).
On the other hand, as vector _y_∗ attains the minimum in (34), it should be a {[^0]:_,_[^1]:}
vector. (If some component would be 2 or larger, we could reduce it to 1, without
violating _yTA_ ≥ [^1]: but decreasing _yT_[^1]:. This contradicts the fact that _y_∗ attains the
minimum.)
Let _W_ be the set of vertices of _G_ for which _y_∗
_v_ = 1. Then _W_ is a vertex cover,
since _y_∗_TA_ ≥ [^1]: holds, implying that for each edge _e_ of _G_ there is at least one vertex
_v_ with _y_∗
_v_ = 1. Moreover, the cardinality |_W_| of _W_ satisfies |_W_| = _y_∗_T_[^1]. So |_W_| is
equal to the minimum in (34).

```
One similarly derives:
```

**Corollary 8.3b** (K ̋onig’s edge cover theorem)**.** _Let G be a bipartite graph. Then the_

_maximum cardinality of a stable set in G is equal to the minimum cardinality of an_
_edge cover in G._

**Proof.** Similar to the proof of Corollary 8.1a (now with _AT_ instead of _A_).

One can also derive weighted versions of these two min-max relations. Let _X_ be
some finite set and let _w_ : _X_ → R be a ‘weight’ function on _X_. The _weight w_(_Y_ ) of
some subset _Y_ ⊆ _X_ is, by definition:

(35) _w_(_Y_ ) :=

```
∑
```

```
x∈Y
```

```
w(x).
```

Then:

**Corollary 8.3c.** _Let G_ = (_V, E_) _be a bipartite graph and let w_ : _E_ → Z+ _be a weight_
_function on E. Then:_

```
(i) ∑The maximum weight of a matching in G is equal to the minimum value of
v∈V f(v), where f ranges over all functions f : V → Z+ such that f(u) +
f(v) ≥ w({u, v}) for each edge {u, v} of G;
```

```
(ii) ∑The minimum weight of an edge cover in G is equal to the maximum value of
v∈V f(v), where f ranges over all functions f : V → Z+ such that f(u) +
f(v) ≤ w({u, v}) for each edge {u, v} of G.
```

**Proof.** The statements are equivalent to both sides in

(36) max{_wTx_ | _x_ ≥ 0;_Ax_ ≤ [^1]:} = min{_yT_[^1]: | _y_ ≥ 0;_yTA_ ≥ _w_}

and in

(37) min{_wTx_ | _x_ ≥ 0;_Ax_ ≥ [^1]:} = max{_yT_[^1]: | _y_ ≥ 0;_yTA_ ≤ _w_}

having integer optimum solutions. These facts follow from Theorem 8.3 and Corollary
8.1b.

Similarly one has min-max relations for the maximum weight of a stable set and
the minimum weight of a vertex cover in bipartite graphs (cf. Exercises 8.10 and 8.11).
Another corollary is as follows. For any finite set _X_ and any subset _Y_ of _X_, define
the _incidence vector χY_ ∈ R_X_ of _Y_ as:

(38) _χY_
_x_ := 1 if _x_ ∈ _Y_ ;
:= 0 if _x_ 6∈ _Y_.

Now let _G_ = (_V, E_) be a graph. The _matching polytope P_matching(_G_) of _G_ is, by
definition, the convex hull (in R_E_) of the incidence vectors of all matchings in _G_.
That is:

(39) _P_matching(_G_) = conv.hull{_χM_ | _M_ matching in _G_}_._

Now with Theorem 8.3 we can give the linear inequalities describing _P_matching(_G_):

**Corollary 8.3d.** _If G is bipartite, the matching polytope P_matching(_G_) _of G is equal_
_to the set of vectors x in_ R_E satisfying:_

(40) (i) _xe_ ≥ [^0]: _for each e_ ∈ _E;_
(ii)

```
∑
```

```
e∋v
```

```
xe ≤ 1 for each v ∈ V.
```

**Proof.** Let _Q_ be the polytope defined by (40). Clearly, _P_matching(_G_) ⊆ _Q_, since the
incidence vector _χM_ of any matching _M_ satisfies (40).
To see that _Q_ ⊆ _P_matching(_G_), observe that _Q_ satisfies

(41) _Q_ = {_x_ | _x_ ≥ 0;_Ax_ ≤ [^1]:}_,_

where _A_ is the incidence matrix of _A_.
Since _A_ is totally unimodular (Theorem 8.3), we know that _Q_ is integer, i.e., that
each vertex of _Q_ is an integer vector (Corollary 8.1a). So _Q_ is the convex hull of the
integer vectors contained in _Q_. Now each integer vector in _Q_ is equal to the incidence
vector _χM_ of some matching _M_ in _G_. So _Q_ must be contained in _P_matching(_G_).

Again, one cannot delete the bipartiteness condition here, as for any odd circuit
there exists a vector satisfying (40) but not belonging to the matching polytope
_P_matching(_G_).
Similarly, let the _perfect matching polytope P_perfect matching(_G_) of _G_ be defined as
the convex hull of the incidence vectors of the _perfect_ matchings in _G_. Then we have:

**Corollary 8.3e.** _If G is bipartite, the perfect matching polytope P_perfect matching(_G_) _of_
_G is equal to the set of vectors x in_ R_E satisfying:_

(42) (i) _xe_ ≥ [^0]: _for each e_ ∈ _E;_
(ii)

```
∑
```

```
e∋v
```

```
xe = 1 for each v ∈ V.
```

**Proof.** Similarly as above.

**Exercises**

8.10. Give a min-max relation for the maximum weight of a stable set in a bipartite graph.

8.11. Give a min-max relation for the minimum weight of a vertex cover in a bipartite
graph.

8.12. Let _G_ = (_V,E_) be a nonbipartite graph. Show that the inequalities (40) are not
enough to define the matching polytope of _G_.

8.13. The _edge cover polytope P_edge cover(_G_) of a graph is the convex hull of the incidence
vectors of the edge covers in _G_. Give a description of the linear inequalities defining
the edge cover polytope of a bipartite graph.

8.14. The _stable set polytope P_stable set(_G_) of a graph is the convex hull of the incidence
vectors of the stable sets in _G_. Give a description of the linear inequalities defining
the stable set polytope of a bipartite graph.

8.15. The _vertex cover polytope P_vertex cover(_G_) of a graph is the convex hull of the incidence
vectors of the vertex covers in _G_. Give a description of the linear inequalities defining
the vertex cover polytope of a bipartite graph.

8.16. Derive from Corollary 8.3e that for each doubly stochastic matrix _M_ there exist
permutation matrices _P_[^1]:_,...,Pm_ and reals _λ_[^1]:_,...,λm_ ≥ 0 such that _λ_[^1]:+···+_λm_ = 1
and

```
(43) M = λ1P1 + ···λmPm.
```

```
(A matrix M is called doubly stochastic if each row sum and each column sum is equal
to 1. A matrix P is called a permutation matrix if it is a {0,1} matrix, with in each
row and in each column exactly one 1.)
```

### 8.4. Totally unimodular matrices from directed graphs

A second class of totally unimodular matrices can be derived from directed graphs.
Let _D_ = (_V, A_) be a directed graph. The _V_ × _A incidence matrix M_ of _D_ is defined
by:

(44) _Mv,a_ := +1 if _a_ leaves _v_,
:= −1 if _a_ enters _v_,
:= 0 otherwise.

So each column of _M_ has exactly one +1 and exactly one −1, while all other entries
are 0.
Now we have:

**Theorem 8.4.** _The incidence matrix M of any directed graph D is totally unimodu-_
_lar._

**Proof.** Let _B_ be a square submatrix of _M_, of order _t_ say. We prove that det_B_ ∈
{[^0]:_,_±[^1]:} by induction on _t_, the case _t_ = 1 being trivial.
Let _t >_ 1. We distinguish three cases.
**Case 1.** _B has a column with only zeros._ Then det_B_ = 0.
**Case 2.** _B has a column with exactly one nonzero._ Then we can write (up to
permuting rows and columns):

(45) _B_ =

```
(
±1 bT
0 B′
```

```
)
,
```

for some vector _b_ and matrix _B_′.
Now by our induction hypothesis, det_B_′ ∈{[^0]:_,_±[^1]:}, and hence det_B_ ∈{[^0]:_,_±[^1]:}.
**Case 3.** _Each column of B contains two nonzeros._ Then each column of _B_
contains one +1 and one −1, while all other entries are 0. So the rows of _B_ add up
to an all-zero vector, and hence det_B_ = 0.

The incidence matrix _M_ of a directed graph _D_ = (_V, A_) relates to flows and
circulations in _D_. Indeed, any vector _x_ ∈ R_A_ can be considered as a function defined
on the arcs of _D_. Then the condition

(46) _M x_ = 0

is just the ‘flow conservation law’. That is, it says:

(47)

```
∑
```

```
a∈δout(v)
```

```
x(a) =
```

```
∑
```

```
a∈δin(v)
```

```
x(a) for each v ∈ V.
```

So we can derive from Theorem 8.4:

**Corollary 8.4a.** _Let D_ = (_V, A_) _be a directed graph and let c_ : _A_ → Z _and d_ : _A_ → Z_._
_If there exists a circulation x on A with c_ ≤ _x_ ≤ _d, then there exists an integer_
_circulation x on A with c_ ≤ _x_ ≤ _d._

**Proof.** If there exists a circulation _x_ with _c_ ≤ _x_ ≤ _d_, then the polytope

(48) _P_ := {_x_ | _c_ ≤ _x_ ≤ _d_;_M x_ = 0}

is nonempty. So it has at least one vertex _x_∗. Then, by Corollary 8.1a, _x_∗ is an
integer circulation satisfying _c_ ≤ _x_∗ ≤ _d_.

In fact, one can derive Hoffman’s circulation theorem— see Exercise 8.17. Another
theorem that can be derived is the max-flow min-cut theorem.

**Corollary 8.4b** (max-flow min-cut theorem)**.** _Let D_ = (_V, A_) _be a directed graph,_
_let s and t be two of the vertices of D, and let c_ : _A_ → R+ _be a ‘capacity’ function_
_on A. Then the maximum value of an s_−_t flow subject to c is equal to the minimum_
_capacity of an s_ − _t cut._

**Proof.** Since the maximum clearly cannot exceed the minimum, it suffices to show
that there exists an _s_ − _t_ flow _x_ ≤ _c_ and an _s_ − _t_ cut, the capacity of which is not
more than the value of _x_.
Let _M_ be the incidence matrix of _D_ and let _M_′ arise from _M_ by deleting the rows
corresponding to _s_ and _t_. So the condition _M_′_x_ = 0 means that the flow conservation
law should hold in any vertex _v_ [^6]:= _s, t_.
Let _w_ be the row of _M_ corresponding to vertex _s_. So _wa_ = +1 if arc _a_ leaves _s_
and _wa_ = −1 if arc _a_ enters _s_, while _wa_ = 0 for all other arcs _a_.
Now the maximum value of an _s_ − _t_ flow subject to _c_ is equal to

(49) max{_wTx_ | [^0]: ≤ _x_ ≤ _c_;_M_′_x_ = 0}_._

By LP-duality, this is equal to

(50) min{_yTc_ | _y_ ≥ 0;_yT_ + _zTM_′ ≥ _w_}_._

The inequality system in (50) is:

(51) (_yT zT_)

```
(
I I
0 M′
```

```
)
≥ (0 w).
```

The matrix here is totally unimodular, by Theorem 8.4.
Since _w_ is an integer vector, this implies that the minimum (50) is attained by
_integer_ vectors _y_ and _z_.
Now define

(52) _W_ := {_v_ ∈ _V_ \{_s, t_}| _zv_ ≤−[^1]:}∪{_s_}_._

So _W_ is a subset of _V_ containing _s_ and not containing _t_.
It suffices now to show that

(53) _c_(_δ_out(_W_)) ≤ _yTc,_

since _yTc_ is not more than the maximum flow value (49).

```
To prove (53) it suffices to show that
```

(54) if _a_ = (_u, v_) ∈ _δ_out(_W_) then _ya_ ≥ 1.

Define ̃_zr_ := −1, ̃_zs_ := 0, and ̃_zu_ = _zu_ for all other _u_. Then _yT_ + ̃_zTM_ ≥ 0. Hence
for all _a_ = (_u, v_) ∈ _δ_out(_W_) one has _ya_ + ̃_zu_ − _z_ ̃_v_ ≥ 0, implying _ya_ ≥ _z_ ̃_v_ − _z_ ̃_u_ ≥ 1. This
proves (54).

Similarly as in Corollary 8.4a it follows that if all capacities are integers, then
there exists a maximum _integer_ flow.
Next define a matrix to be an _interval matrix_ if each entry is 0 or 1 and each row
is of type

(55) (0_,... ,_[^0]:_,_[^1]:_,... ,_[^1]:_,_[^0]:_,... ,_0)_._

**Corollary 8.4c.** _Each interval matrix is totally unimodular._

**Proof.** Let _M_ be an interval matrix and let _B_ be a _t_ × _t_ submatrix of _M_. Then _B_
is again an interval matrix. Let _N_ be the _t_ × _t_ matrix given by:

(56) _N_ :=

```

```

```

```

```
1 −1 0 ··· ··· 0 0
0 1 −1 ··· ··· 0 0
0 0 1 ··· ··· 0 0
```

...
...
.........
...
..................
0 0 0 ··· ··· [^1]: −1
0 0 0 ··· ··· 0 1

```

```

```

```

_._

Then the matrix _N_ · _BT_ is a {[^0]:_,_±[^1]:} matrix, with at most one +1 and at most one
−1 in each column.
So it is a submatrix of the incidence matrix of some directed graph. Hence by
Theorem 8.4, det(_N_ · _BT_) ∈ {[^0]:_,_±[^1]:}. Moreover, det_N_ = 1. So det_B_ = det_BT_ ∈
{[^0]:_,_±[^1]:}.

**Exercises**

8.17. Derive Hoffman’s circulation theorem (Theorem 4.6) from Theorem 8.4.

8.18. Derive Dilworth’s decomposition theorem (Theorem 7.6) from Theorem 8.4.

8.19. Let _D_ = (_V,A_) be a directed graph and let _T_ = (_V,A_′) be a directed spanning tree
on _V_.

Let _C_ be the _A_′ × _A_ matrix defined as follows. Take _a_′ ∈ _A_′ and _a_ = (_u,v_) ∈ _A_,
and define _Ca_′_,a_ := +1 if _a_′ occurs in forward direction in the _u_ − _v_ path in _T_ and
_Ca_′_,a_ := −1 if _a_′ occurs in backward direction in the _u_ − _v_ path in _T_. For all other
_a_′ ∈ _A_′ and _a_ ∈ _A_ set _Ca_′_,a_ := 0.

```
(i) Prove that C is totally unimodular.
(Hint: Use a matrix similar to matrix N in Corollary 8.4c.)
(ii) Show that interval matrices and incidence matrices of directed graphs are special
cases of such a matrix C.
```

## 9. Multicommodity flows and

## disjoint paths

### 9.1. Introduction

```
The problem of finding a maximum flow from one ‘source’ s to one ‘sink’ t is highly
tractable. There is a very efficient algorithm, which outputs an integer maximum
flow if all capacities are integer. Moreover, the maximum flow value is equal to
the minimum capacity of a cut separating s and t. If all capacities are equal to 1,
the problem reduces to finding arc-disjoint paths. Some direct transformations give
similar results for vertex capacities and for vertex-disjoint paths.
Often in practice however, one is not interested in connecting only one pair of
source and sink by a flow or by paths, but several pairs of sources and sinks simulta-
neously. One may think of a large communication or transportation network, where
several messages or goods must be transmitted all at the same time over the same
network, between different pairs of terminals. A recent application is the design of
very large-scale integrated (VLSI) circuits, where several pairs of pins must be inter-
connected by wires on a chip, in such a way that the wires follow given ‘channels’ and
that the wires connecting different pairs of pins do not intersect each other.
Mathematically, these problems can be formulated as follows. First, there is the
multicommodity flow problem (or k-commodity flow problem):
```

(1) given: a directed graph _G_ = (_V, E_)_,_ pairs (_s_[^1]:_, t_[^1]:)_,... ,_(_sk, tk_) of vertices of _G_, a
‘capacity’ function _c_ : _E_ → Q+, and ‘demands’ _d_[^1]:_,... , dk,_
find: for each _i_ = 1_,... , k,_ an _si_ − _ti_ flow _xi_ ∈ Q_E_
+ so that _xi_ has value _di_ and so
that for each arc _e_ of _G_:

```
∑k
```

```
i=1
```

```
xi(e) ≤ c(e).
```

```
The pairs (si, ti) are called the commodities or the nets. (We assume si 6= ti through-
out.)
If we require each xi to be an integer flow, the problem is called the integer
multicommodity flow problem or integer k-commodity flow problem. (To distinguish
from the integer version of this problem, one sometimes adds the adjective fractional
to the name of the problem if no integrality is required.)
The problem has a natural analogue to the case where G is undirected. We replace
each undirected edge e = {v, w} by two opposite arcs (v, w) and (w, v) and ask for
flows x1,... , xk of values d1,... , dk, respectively, so that for each edge e = {v, w} of
G:
```

```
(2)
```

```
∑k
```

```
i=1
```

```
(xi(v, w) + xi(w, v)) ≤ c(e).
```

```
Thus we obtain the undirected multicommodity flow problem or undirected k-commodity
flow problem. Again, we add integer if we require the xi to be integer flows.
If all capacities and demands are 1, the integer multicommodity flow problem is
equivalent to the arc- or edge-disjoint paths problem:
```

(3) given: a (directed or undirected) graph _G_ = (_V, E_), pairs (_s_[^1]:_, t_[^1]:)_,... ,_ (_sk, tk_) of
vertices of _G_,
find: pairwise edge-disjoint paths _P_[^1]:_,... , Pk_ where _Pi_ is an _si_ − _ti_ path (_i_ =
[^1]:_,... , k_).

```
Related is the vertex-disjoint paths problem:
```

(4) given: a (directed or undirected) graph _G_ = (_V, E_), pairs (_s_[^1]:_, t_[^1]:)_,... ,_ (_sk, tk_) of
vertices of _G_,

```
find: pairwise vertex-disjoint paths P1,... , Pk where Pi is an si − ti path (i =
1,... , k).
```

```
We leave it as an exercise (Exercise 9.1) to check that the vertex-disjoint paths
problem can be transformed to the directed edge-disjoint paths problem.
The (fractional) multicommodity flow problem can be easily described as one of
solving a system of linear inequalities in the variables xi(e) for i = 1,... , k and
e ∈ E. The constraints are the flow conservation laws for each flow xi separately,
together with the inequalities given in (1). Therefore, the fractional multicommod-
ity flow problem can be solved in polynomial time with any polynomial-time linear
programming algorithm.
In fact, the only polynomial-time algorithm known for the fractional multicom-
modity flow problem is any general linear programming algorithm. Ford and Fulker-
son [1958] designed an algorithm based on the simplex method, with column genera-
tion — see Section 9.6.
The following cut condition trivially is a necessary condition for the existence of
a solution to the fractional multicommodity flow problem (1):
```

```
(5) for each W ⊆ V the capacity of δout
E (W) is not less than the demand of
δout
R (W),
```

```
where R := {(s1, t1),... ,(sk, tk)}. However, this condition is in general not sufficient,
even not in the two simple cases given in Figure 9.1 (taking all capacities and demands
equal to 1).
One may derive from the max-flow min-cut theorem that the cut condition is
```

```
s =t
```

```
2
```

```
2 1
```

```
t
```

```
s1
```

```
s1=t2 s2=t1
```

```
Figure 9.1
```

sufficient if _s_[^1]: = _s_[^2]: = ··· = _sk_ (similarly if _t_[^1]: = _t_[^2]: = ··· = _tk_) — see Exercise 9.3.
Similarly, in the undirected case a necessary condition is the following cut condi-
tion:

(6) for each _W_ ⊆ _V,_ the capacity of _δE_(_W_) is not less than the demand of
_δR_(_W_)

(taking _R_ := {{_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}}). In the special case of the edge-disjoint paths
problem (where all capacities and demands are equal to 1), the cut condition reads:

(7) for each _W_ ⊆ _V,_|_δE_(_W_)|≥|_δR_(_W_)|_._

Figure 9.2 shows that this condition again is not sufficient.

```
=s
```

```
t1
t
```

```
1
```

```
=s2
s4 4
```

```
t3
```

```
t2=s3
Figure 9.2
```

However, Hu [1963] showed that the cut condition is sufficient for the existence
of a fractional multicommodity flow, in the undirected case with _k_ = 2 commodities.
He gave an algorithm that yields a half-integer solution if all capacities and demands
are integer. This result was extended by Rothschild and Whinston [1966]. We discuss
these results in Section 9.2.
Similar results were obtained by Okamura and Seymour [1981] for arbitrary _k_,
provided that the graph is planar and all terminals _si, ti_ are on the boundary of the
unbounded face — see Section 9.5.

The integer multicommodity flow problem is NP-complete, even in the undirected
case with _k_ = 2 commodities and all capacities equal to 1, with arbitrary demands
_d_[^1]:_, d_[^2]: (Even, Itai, and Shamir [1976]). This implies that the undirected edge-disjoint
paths problem is NP-complete, even if |{{_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}}| = 2.
In fact, the disjoint paths problem is NP-complete in all modes (directed/undirected,
vertex/edge disjoint), even if we restrict the graph _G_ to be planar (D.E. Knuth (see
Karp [1975]), Lynch [1975], Kramer and van Leeuwen [1984]). For general directed
graphs the arc-disjoint paths problem is NP-complete even for _k_ = 2 ‘opposite’ com-
modities (_s, t_) and (_t, s_) (Fortune, Hopcroft, and Wyllie [1980]).
On the other hand, it is a deep result of Robertson and Seymour [1995] that
the undirected vertex-disjoint paths problem is polynomially solvable for any fixed
number _k_ of commodities. Hence also the undirected edge-disjoint paths problem is
polynomially solvable for any fixed number _k_ of commodities.
Robertson and Seymour observed that if the graph _G_ is planar and all termi-

nals _si, ti_ are on the boundary of the unbounded face, there is an easy ‘greedy-type’
algorithm for the vertex-disjoint paths problem — see Section 9.4.
It is shown by Schrijver [1994] that for each fixed _k_, the _k_ disjoint paths problem
is solvable in polynomial time for directed planar graphs. For the directed planar arc-
disjoint version, the complexity is unknown. That is, there is the following research
problem:

**Research problem**. Is the directed arc-disjoint paths problem polynomially solvable
for planar graphs with _k_ = 2 commodities? Is it NP-complete?

**Application 9.1: Multicommodity flows.** Certain goods or messages must be trans-
ported through the same network, where the goods or messages may have different sources
and sinks.
This is a direct special case of the problems described above.

**Application 9.2: VLSI-routing.** On a chip certain modules are placed, each containing
a number of ’pins’. Certain pairs of pins should be connected by an electrical connection
(a ‘wire’) on the chip, in such a way that each wire follows a certain (very fine) grid on the
chip and that wires connecting different pairs of pins are disjoint.
Determining the routes of the wires clearly is a special case of the disjoint paths prob-
lem.

**Application 9.3: Routing of railway stock.** An extension of Application 4.5 is as
follows. The stock of the railway company NS for the Amsterdam–Vlissingen line now
consists of two types (1 and 2 say) of units, with a different number of seats _s_[^1]: and _s_[^2]: and
different length _l_[^1]: and _l_[^2]:. All units (also of different types) can be coupled with each other.
Again there is a schedule given, together with for each segment a minimum number of
seats and a maximum length of the train. Moreover, the price _pi_ of buying any unit of type
_i_ is given.
Now the company wishes to determine the minimum costs of buying units of the two

types so that the schedule can be performed and so that the total cost is minimized.
This can be considered as a ‘min-cost integer multicommodity circulation problem’.
That is we make the directed graph _D_ as in Application 4.5. For each arc _a_ corresponding
to a segment we define _d_(_a_) to be the minimum number of seats that should be offered on
that segment, and _c_(_a_) to be the maximum length possible at that segment. For all other
arcs _a_ we define _d_(_a_) := 0 and _c_(_a_) := ∞.
One should find two integer-valued circulations _f_[^1]: and _f_[^2]: in _D_ such that

(8) _s_[^1]:_f_[^1]:(_a_) + _s_[^2]:_f_[^2]:(_a_) ≥ _d_(_a_) and _l_[^1]:_f_[^1]:(_a_) + _l_[^2]:_f_[^2]:(_a_) ≤ _c_(_a_)

for each arc _a_ and such that the sum

∑
(_p_[^1]:_f_[^1]:(_a_) + _p_[^2]:_f_[^2]:(_a_)) is minimized, where _a_ ranges
over all ‘overnight’ arcs. Then _fi_(_a_) denotes the number of units of type _i_ that should go
on segment _a_.
Again several variations are possible, incorporating for instance the kilometer costs and
maximum capacities of stock areas.

**Exercises**

```
9.1. Show that each of the following problems (a), (b), (c) can be reduced to problems
(b), (c), (d), respectively:
```

```
(a) the undirected edge-disjoint paths problem,
(b) the undirected vertex-disjoint paths problem,
(c) the directed vertex-disjoint paths problem,
(d) the directed arc-disjoint paths problem.
```

```
9.2. Show that the undirected edge-disjoint paths problem for planar graphs can be re-
duced to the directed arc-disjoint paths problem for planar graphs.
```

```
9.3. Derive from the max-flow min-cut theorem that the cut condition (5) is sufficient for
the existence of a fractional multicommodity flow if s1 = ··· = sk.
9.4. Show that if the undirected graph G = (V,E) is connected and the cut condition (7)
is violated, then it is violated by some W ⊆ V for which both W and V \ W induce
connected subgraphs of G.
```

```
9.5. (i) Show with Farkas’ lemma: the fractional multicommodity flow problem (1) has
a solution if and only if for each ‘length’ function l : E → Q+ one has:
```

```
(9)
```

```
∑k
```

```
i=1
```

```
di · distl(si,ti) ≤
```

```
∑
```

```
e∈E
```

```
l(e)c(e).
```

```
(Here distl(s,t) denotes the length of a shortest s − t path with respect to l.)
(ii) Interprete the cut condition (5) as a special case of this condition.
```

### 9.2. Two commodities

Hu [1963] gave a direct combinatorial method for the undirected two-commodity flow
problem and he showed that in this case the cut condition suffices. In fact, he showed
that if the cut condition holds and all capacities and demands are integer, there exists
a half-integer solution. We first give a proof of this result due to Sakarovitch [1973].

Consider a graph _G_ = (_V, E_), with commodities {_s_[^1]:_, t_[^1]:} and {_s_[^2]:_, t_[^2]:}, a capacity
function _c_ : _E_ → Z+ and demands _d_[^1]: and _d_[^2]:.

**Theorem 9.1** (Hu’s two-commodity flow theorem)**.** _The undirected two-commodity_
_flow problem, with integer capacities and demands, has a half-integer solution if and_
_only if the cut condition_ (6) _is satisfied._

**Proof.** Suppose that the cut condition holds. Orient the edges of _G_ arbitrarily,
yielding the directed graph _D_ = (_V, A_). For any _a_ ∈ _A_ we denote by _c_(_a_) the
capacity of the underlying undirected edge.
Define for any _x_ ∈ _RA_ and any _v_ ∈ _V_ :

(10) _f_(_x, v_) :=

```
∑
```

```
a∈δout(v)
```

```
x(a) −
```

```
∑
```

```
a∈δin(v)
```

```
x(a).
```

So _f_(_x, v_) is the ‘net loss’ of _x_ in vertex _v_.
By the max-flow min-cut theorem there exists a function _x_′ : _A_ → Z satisfying:

(11) _f_(_x_′_, s_[^1]:) = _d_[^1]:_, f_(_x_′_, t_[^1]:) = −_d_[^1]:_, f_(_x_′_, s_[^2]:) = _d_[^2]:_, f_(_x_′_, t_[^2]:) = −_d_[^2]:_,_
_f_(_x_′_, v_) = 0 for each other vertex _v_,
|_x_′(_a_)|≤ _c_(_a_) for each arc _a_ of _D._

This can be seen by extending the undirected graph _G_ by adding two new vertices _s_′
and _t_′ and four new edges {_s_′_, s_[^1]:}_,_{_t_[^1]:_, t_′} (both with capacity _d_[^1]:) and {_s_′_, s_[^2]:}_,_{_t_[^2]:_, t_′}
(both with capacity _d_[^2]:) as in Figure 9.3.

```
1
```

```
t’
```

```
s
```

```
1
```

```
G
```

```
t2
```

```
s’
```

```
s2 t
```

```
Figure 9.3
```

Then the cut condition for the two-commodity flow problem implies that the
minimum capacity of any _s_′ −_t_′ cut in the extended graph is equal to _d_[^1]: +_d_[^2]:. Hence,
by the max-flow min-cut theorem, there exists an integer-valued _s_′ − _t_′ flow in the

extended graph of value _d_[^1]: + _d_[^2]:. This gives _x_′ satisfying (11).
Similarly, the max-flow min-cut theorem implies the existence of a function _x_′′ :
_A_ → Z satisfying:

(12) _f_(_x_′′_, s_[^1]:) = _d_[^1]:_, f_(_x_′′_, t_[^1]:) = −_d_[^1]:_, f_(_x_′′_, s_[^2]:) = −_d_[^2]:_, f_(_x_′′_, t_[^2]:) = _d_[^2]:_,_
_f_(_x_′′_, v_) = 0 for each other vertex _v_,
|_x_′′(_a_)|≤ _c_(_a_) for each arc _a_ of _D_.

To see this we extend _G_ with vertices _s_′′ and _t_′′ and edges {_s_′′_, s_[^1]:}_,_{_t_[^1]:_, t_′′} (both with
capacity _d_[^1]:) and {_s_′′_, t_[^2]:}_,_{_s_[^2]:_, t_′′} (both with capacity _d_[^2]:) (cf. Figure 9.4).

```
G
```

```
s1 t2
```

```
s"
```

```
s2 t1
```

```
t"
Figure 9.4
```

After this we proceed as above.
Now consider the vectors

(13) _x_[^1]: :=[^1]
[^2]:(_x_′ + _x_′′) and _x_[^2] :=[^1]
[^2]:(_x_′ − _x_′′)_._

Since _f_(_x_[^1]:_, v_) =[^1]
[^2]:(_f_(_x_′_, v_) + _f_(_x_′′_, v_)) for each _v_, we see from (11) and (12) that _x_[^1]
satisfies:

(14) _f_(_x_[^1]:_, s_[^1]:) = _d_[^1]:_, f_(_x_[^1]:_, t_[^1]:) = −_d_[^1]:_, f_(_x_[^1]:_, v_) = 0 for all other _v._

So _x_[^1]: gives a half-integer _s_[^1]: − _t_[^1]: flow in _G_ of value _d_[^1]:. Similarly, _x_[^2]: satisfies:

(15) _f_(_x_[^2]:_, s_[^2]:) = _d_[^2]:_, f_(_x_[^2]:_, t_[^2]:) = −_d_[^2]:_, f_(_x_[^2]:_, v_) = 0 for all other _v._

So _x_[^2]: gives a half-integer _s_[^2]: − _t_[^2]: flow in _G_ of value _d_[^2]:.
Moreover, _x_[^1]: and _x_[^2]: together satisfy the capacity constraint, since for each edge
_a_ of _D_:

(16) |_x_[^1]:(_a_)| + |_x_[^2]:(_a_)| =[^1]
[^2]:|_x_′(_a_) + _x_′′(_a_)| +[^1]
[^2]:|_x_′(_a_) − _x_′′(_a_)|
= max{|_x_′(_a_)|_,_|_x_′′(_a_)|}≤ _c_(_a_)_._

(Note that[^1]
[^2]:|_α_ + _β_| +[^1]
[^2]:|_α_ − _β_| = max{|_α_|_,_|_β_|} for all reals _α, β_.)
So we have a half-integer solution to the two-commodity flow problem.

This proof also directly gives a polynomial-time algorithm for finding a half-integer
flow.
The cut condition is not enough to derive an _integer_ solution, as is shown by
Figure 9.5 (taking all capacities and demands equal to 1).

```
t
```

```
1 2
```

```
s2 1
```

```
s t
```

```
Figure 9.5
```

Moreover, as mentioned, the undirected integer two-commodity flow problem is NP-
complete (Even, Itai, and Shamir [1976]).
However, Rothschild and Whinston [1966] showed that an integer solution exists
if the cut condition holds, provided that the following _Euler condition_ is satisfied:

(17)

```
∑
e∈δ(v) c(e) ≡ 0 (mod 2) if v[^6]= s[^1], t[^1], s[^2], t[^2],
≡ d1 (mod 2) if v = s1, t1,
≡ d2 (mod 2) if v = s2, t2.
```

(Equivalently, the graph obtained from _G_ by replacing each edge _e_ by _c_(_e_) parallel

edges and by adding _di_ parallel edges connecting _si_ and _ti_ (_i_ = 1_,_2), should be an
Eulerian graph.)

**Theorem 9.2.** _If all capacities and demands are integer and the cut condition and_
_the Euler condition are satisfied, then the undirected two-commodity flow problem has_
_an integer solution._

**Proof.** If the Euler condition holds, we can take _x_′ in the proof of Theorem 9.1 so
that the following further condition is satisfied:

(18) _x_′(_a_) ≡ _c_(_a_) (mod 2) for each _a_ ∈ _A._

```
To see this, let x′ satisfy (11) and let
```

(19) _A_′ := {_a_ ∈ _A_ | _x_′(_a_) 6≡ _c_(_a_) (mod 2)}_._

Then each vertex _v_ is incident with an even number _δ_ of arcs in _A_′, since

(20) _δ_ ≡ _f_(_x_′_, v_) − _f_(_c, v_) ≡ [^0]: (mod 2)_,_

by (11) and (17). So if _A_′ [^6]:= ∅ then _A_′ contains an (undirected) circuit. Increasing
and decreasing _x_′ by 1 on the arcs along this circuit (depending on whether the arc
is forward or backward), we obtain a function again satisfying (11). Repeating this,
we finally obtain a function _x_′ satisfying (18).
Similarly, we can take _x_′′ so that

(21) _x_′′(_a_) ≡ _c_(_a_) (mod 2) for each _a_ ∈ _A._

Conditions (18) and (21) imply that _x_′(_a_) ≡ _x_′′(_a_) (mod 2) for each _a_ ∈ _A_.
Hence _x_[^1]: =[^1]
[^2]:(_x_′ + _x_′′) and _x_[^2] =[^1]
[^2]:(_x_′ − _x_”) are integer vectors.

This proof directly yields a polynomial-time algorithm for finding the integer
solution.

**Exercises**

```
9.6. Derive from Theorem 9.1 the following max-biflow min-cut theorem of Hu: Let G =
(V,E) be a graph, let s1,t1,s2,t2 be distinct vertices, and let c : E → Q+ be a
capacity function. Then the maximum value of d1+d2 so that there exist si−ti flows
xi of value di (i = 1,2), together satisfying the capacity constraint, is equal to the
minimum capacity of a cut both separating s1 and t1 and separating s2 and t2.
```

```
9.7. Derive from Theorem 9.1 that the cut condition suffices to have a half-integer solu-
tion to the undirected k-commodity flow problem (with all capacities and demands
integer), if there exist two vertices u and w so that each commodity {si,ti} intersects
{u,w}. (Dinits (cf. Adel’son-Vel’ski ̆ı, Dinits, and Karzanov [1975]).)
```

```
9.8. Derive the following from Theorem 9.2. Let G = (V,E) be a Eulerian graph and
let s1,t1,s2,t2 be distinct vertices. Then the maximum number t of pairwise edge-
disjoint paths P1,...,Pt, where each Pj connects either s1 and t1 or s2 and t2, is
```

```
equal to the minimum cardinality of a cut both separating s1 and t1 and separating
s2 and t2.
```

### 9.3. Disjoint paths in acyclic directed graphs

```
Fortune, Hopcroft, and Wyllie [1980] showed that the vertex-disjoint paths problem
is NP-complete for directed graphs, even when fixing the number of paths to k = 2.
On the other hand they proved that if D is acyclic, then for each fixed k, the k
vertex-disjoint paths problem can be solved in polynomial time. (A directed graph is
called acyclic if it does not contain any directed circuit.)
The algorithm is contained in the proof of the following theorem:
```

```
Theorem 9.3. For each fixed k there exists a polynomial-time algorithm for the k
vertex-disjoint paths problem for acyclic directed graphs.
```

```
Proof. Let D = (V, A) be an acyclic digraph and let s1, t1,... , sk, tk be vertices of
D, all distinct. In order to solve the vertex-disjoint paths problem we may assume
that each si is a source and each ti is a sink. (Here a source is a vertex with indegree
0, and a sink is a vertex with outdegree 0.)
Make an auxiliary digraph D′ = (V ′, A′) as follows. The vertex set V ′ consists of
all k-tuples (v1,... , vk) of distinct vertices of D. In D′ there is an arc from (v1,... , vk)
to (w1,... , wk) if and only if there exists an i ∈{1,... , k} such that:
```

(22) (i) _vj_ = _wj_ for all _j_ [^6]:= _i_;
(ii) (_vi, wi_) is an arc of _D_;
(iii) for each _j_ [^6]:= _i_ there is no directed path in _D_ from _vj_ to _vi_.

```
Now the following holds:
```

```
(23) D contains k vertex-disjoint directed paths P1,... , Pk such that Pi runs
from si to ti (i = 1,... , k)
⇐⇒ D′ contains a directed path P from (s1,... , sk) to (t1,... , tk).
```

```
To see =⇒, let Pi follow the vertices vi,0, vi,1,... , vi,pi for i = 1,... , k. So vi,0 = si
and vi,pi = ti for each i. Choose j1,... , jk such that 0 ≤ ji ≤ pi for each i and such
that:
```

(24) (i) _D_′ contains a directed path from (_s_[^1]:_,... , sk_) to (_v_[^1]:_,j_[^1]:_,... , vk,jk_),
(ii) _j_[^1]: + ··· + _jk_ is as large as possible.

```
Let I := {i | ji < pi}. If I = ∅ we are done, so assume I 6= ∅. Then by the
definition of D′ and the maximality of j1 +···+jk there exists for each i ∈ I an i′ 6= i
```

such that there is a directed path in _D_ from _vi_′_,ji_′ to _vi,ji_. Since _ti_′ is a sink we know
that _vi_′_,ji_′ [^6]:= _si_′ and that hence _i_′ belongs to _I_. So each vertex in {_vi,ji_ | _i_ ∈ _I_} is
end vertex of a directed path in _D_ starting at another vertex in {_vi,ji_ | _i_ ∈ _I_}. This
contradicts the fact that _D_ is acyclic.
To see ⇐= in (23), let _P_ be a directed path from (_s_[^1]:_,... , sk_) to (_t_[^1]:_,... , tk_) in _D_′.
Let _P_ follow the vertices (_v_[^1]:_,j,... , vk,j_) for _j_ = 0_,... , t_. So _vi,_[^0]: = _si_ and _vi,t_ = _ti_
for _i_ = 1_,... , k_. For each _i_ = 1_,... , k_ let _Pi_ be the path in _D_ following _vi,j_ for
_j_ = 0_,... , t_, taking repeated vertices only once. So _Pi_ is a directed path from _si_ to _ti_.
Moreover, _P_[^1]:_,... , Pk_ are pairwise disjoint. For suppose that _P_[^1]: and _P_[^2]: (say) have
a vertex in common. That is _v_[^1]:_,j_ = _v_[^2]:_,j_′ for some _j_ [^6]:= _j_′. Without loss of generality,
_j < j_′ and _v_[^1]:_,j_ [^6]:= _v_[^1]:_,j_+1. By definition of _D_′, there is no directed path in _D_ from _v_[^2]:_,j_
to _v_[^1]:_,j_. However, this contradicts the facts that _v_[^1]:_,j_ = _v_[^2]:_,j_′ and that there exists a
directed path in _D_ from _v_[^2]:_,j_ to _v_[^2]:_,j_′.

One can derive from this that for fixed _k_ also the _k arc_-disjoint paths problem is
solvable in polynomial time for acyclic directed graphs (Exercise 9.9).

**Application 9.4: Routing airplanes.** This application extends Application 4.1. The
data are similar, except that legal rules now prescribe the exact day of the coming week at
which certain airplanes should be at the home basis for maintenance.
Again at Saturday 18.00h the company determines the exact routing for the next week.
One can make the same directed graph as in Application 4.1. Now however it is prescribed
that some of the paths _Pi_ should start at a certain (_c,t_) (where _c_ is the city where airplane
_ai_ will be first after Saturday 18.00h) and that they should traverse the arc corresponding
to maintenance on a prescribed day of the coming week (for instance Wednesday).
Now if for each airplane _ai_ which should be home for maintenance next week we can
find this path _Pi_ such that it traverses the for that plane required maintenance arc and in
such a way that paths found for different airplanes are arc disjoint, then it is easy to see
that these paths can be extended to paths _P_[^1]:_,...,Pn_ such that each arc is traversed exactly
once.
As the directed graph _D_ is acyclic, the problem can be solved with the algorithm
described in the proof of Theorem 9.3, provided that the number of airplanes that should
be home for maintenance the coming week is not too large.

**Exercises**
9.9. Derive from Theorem 9.3 that for each fixed _k_ the _k_ arc-disjoint paths problem is
solvable in polynomial time for acyclic directed graphs.
9.10. Show that for fixed _k_, the following problem is solvable in polynomial time:

```
(25) given:an acyclic directed graph D = (V,A), pairs s1,t1,...,sk,tk of
vertices, and subsets A1,...,Ak of A;
find:pairwise arc-disjoint directed paths P1,...,Pk, where Pi runs from
si to ti and traverses only arcs in Ai (i = 1,...,k).
```

### 9.4. Vertex-disjoint paths in planar graphs

```
Finding disjoint paths in planar graphs is of interest not only for planar communica-
tion or transportation networks, but especially also for the design of VLSI-circuits.
The routing of wires should follow certain channels on layers of the chip. On each
layer, these channels form a planar graph.
Unfortunately, even for planar graphs disjoint paths problems are in general hard.
However, for some special cases, polynomial-time algorithms have been found. Such
algorithms can be used, for example, as subroutines when solving any hard problem
by decomposition. In Sections 9.4 and 9.5 we discuss some of these algorithms.
Let G = (V, E) be a planar graph, embedded in the plane R[^2] and let {s1, t1},... ,{sk, tk}
be pairwise disjoint pairs of vertices. Robertson and Seymour [1986] observed that
there is an easy greedy-type algorithm for the vertex-disjoint paths problem if all
vertices s1, t1,... , sk, tk belong to the boundary of one face I of G. That is, there
exists a polynomial-time algorithm for the following problem:[^21]
```

(26) given: a planar graph _G_ = (_V, E_) embedded in R[^2], a face _I_ of _G_, pairs {_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}
of vertices on bd(_I_),

```
find: pairwise vertex-disjoint paths P1,... , Pk in G, where Pi connects si and ti
(i = 1,... , k).
```

```
In fact, we may assume without loss of generality that I is the unbounded face.
Let us first describe the simple intuitive idea of the method, by explaining the
recursive step in the ‘ideal’ case where G is connected and where bd(I) is a simple
circuit.
We say that {s, t} and {s′, t′} cross (around I) if s, s′, t, t′ are distinct and occur
in this order cyclically around bd(I), clockwise or anti-clockwise (see Figure 9.6).
r
```

```
s
```

```
s’ r’
```

```
r
```

```
s
```

```
r’ s’
```

```
Figure 9.6
```

```
If any {si, ti} and {sj, tj} cross around I (for some i 6= j), problem (26) clearly
has no solution. So we may assume that no pair of commodities crosses. This implies
that there exists an i so that at least one of the si − ti paths along bd(I) does not
contain any sj or tj for j 6= i: just choose i so that the shortest si − ti path along
bd(I) is shortest among all i = 1,... , k.
```

```
21bd(I) denotes the boundary of I.
```

Without loss of generality, _i_ = _k_. Let _Q_ be the shortest _sk_ − _tk_ path along bd(_I_).
Delete from _G_ all vertices in _Q_, together with all edges incident with them. De-
note the new graph by _G_′. Next solve the vertex-disjoint paths problem for input
_G_′_,_ {_s_[^1]:_, t_[^1]:}_,... ,_{_sk_−[^1]:_, tk_−[^1]:}. If this gives a solution _P_[^1]:_,... , Pk_−[^1]:, then _P_[^1]:_,... , Pk_−[^1]:_, Q_
forms a solution to the original problem (trivially).
If the reduced problem turns out to have no solution, then the original problem
also has no solution. This follows from the fact that if _P_[^1]:_,... , Pk_−[^1]:_, Pk_ would be
a solution to the original problem, we may assume without loss of generality that
_Pk_ = _Q_, since we can ‘push’ _Pk_ ‘against’ the border bd(_I_). Hence _P_[^1]:_,... , Pk_−[^1]: would
form a solution to the reduced problem.
Although this might give a suggestive sketch of the algorithm, it is not completely
accurate, since the ideal situation need not be preserved throughout the iteration
process. Even if we start with a highly connected graph, after some iterations the
reduced graph might have cut vertices or be disconnected. So one should be more
precise.

Let us call a sequence (_v_[^1]:_,... , vn_) of vertices of a connected planar graph _G_ a _border_
_sequence_ if it is the sequence of vertices traversed when following the boundary of _G_
clockwise. Thus the graph in Figure 9.7 has border sequence (_a, b, c, d, e, c, f, c, g, b_).

```
c
```

```
d
```

```
e
f
```

```
h
```

```
g
```

```
b
```

```
a
```

```
Figure 9.7
```

In fact, each cyclic permutation of a border sequence is again a border sequence.
Note that no border sequence will contain_... r... s... r... s..._ for any two dis-
tinct vertices. Hence for any two vertices _s_ and _t_ on the boundary of _G_ there is a
unique sequence

(27) _P_(_s, t_) = (_s, w_[^1]:_,... , wt, t_)

with the properties that _P_(_s, t_) is part of a border sequence of _G_ and that _w_[^1]:_,... , wt_
all are distinct from _s_ and _t_. Trivially, the vertices in _P_(_s, t_) contain an _s_ − _t_ path.
We say that two disjoint pairs {_s, t_} and {_s_′_, t_′} cross (around _G_) if_... s... s_′_... t... t_′_..._
or_... s... t_′_... t... s_′_..._ occur in some border sequence of _G_. So the following _cross-_
_freeness condition_ is a necessary condition for (26) to have a solution:

(28) No two disjoint commodities {_si, ti_}_,_{_sj, tj_} cross (around the same com-
ponent of _G_).

Now the algorithm can be described more precisely as follows. First check the cross-
freeness condition. If it is violated, (26) has no solution. If it is satisfied, apply the
following iterative step:

(29) Check for each _i_ = 1_,... , k_ if _si_ and _ti_ belong to the same component of _G_.
If not, the problem has no solution.
If so, choose _i_ ∈ {[^1]:_,... , k_} for which the shortest among _P_(_si, ti_) and
_P_(_ti, si_) is as short as possible. Without loss of generality, _i_ = _k_ and
_P_(_sk, tk_) is shortest. Take for _Pk_ any _sk_ − _tk_ path using the vertices in
_P_(_sk, tk_) only.
If _k_ = 1, stop. If _k >_ 1, let _G_′ be the graph obtained from _G_ by deleting
all vertices occurring in _P_(_sk, tk_). Repeat this iterative step for _G_′ and
{_s_[^1]:_, t_[^1]:}_,... ,_{_sk_−[^1]:_, tk_−[^1]:}.
If it gives a solution _P_[^1]:_,... , Pk_−[^1]:, then _P_[^1]:_,... , Pk_−[^1]:_, Pk_ is a solution to
the original problem. If it gives no solution, the original problem has no
solution.

We leave it as a (technical) exercise to show the correctness of this algorithm. (The
correctness can be derived also from the proof of Theorem 9.4 below.) It clearly is
a polynomial-time method. Recently, Ripphausen-Lipa, Wagner, and Weihe [1997]
found a linear-time algorithm.
Moreover, the method implies a characterization by means of a cut condition for
the existence of a solution to (26). A _simple closed curve C_ in R[^2] is by definition
a one-to-one continuous function from the unit circle to R[^2]. We will identify the
function _C_ with its image.
We say that _C separates_ the pair {_s, t_} if each curve connecting _s_ and _t_ intersects
_C_. Now the following _cut condition_ clearly is necessary for the existence of a solution
to the vertex-disjoint paths problem in planar graphs:

(30) each simple closed curve in R[^2] intersects _G_ at least as often as it separates
pairs {_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}.

```
Robertson and Seymour [1986] showed with this method:
```

**Theorem 9.4.** _Let G_ = (_V, E_) _be a planar graph embedded in_ R[^2] _and let_ {_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}
_be pairs of vertices on the boundary of G. Then there exist pairwise vertex-disjoint_
_paths P_[^1]:_,... , Pk where Pi connects si and ti (i_ = 1_,... , k) if and only if the cross-_
_freeness condition_ (28) _and the cut condition_ (30) _hold._

**Proof.** Necessity of the conditions is trivial. We show sufficiency by induction on _k_,

the case _k_ = 0 being trivial. Let _k >_ 1 and let (28) and (30) be satisfied. Suppose

paths _P_[^1]:_,... , Pk_ as required do not exist. Trivially, {_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_} are pairwise
disjoint (otherwise there would exist a simple closed curve _C_ with |_C_ ∩ _G_| = 1 and
intersecting two commodities, thus violating the cut condition).
The induction is based on the iterative step (29). To simplify the argument, we
first show that we may assume that _G_ is 2-connected.
First, we may assume that _G_ is connected, as we can decompose _G_ into its compo-
nents. (If some _si_ and _ti_ would belong to different components, there trivially exists
a closed curve _C_ violating the cut condition.)
Knowing that _G_ is connected, the case _k_ = 1 is trivial. So we may assume that
_k_ ≥ 2. Suppose _G_ contains a cut vertex _v_. We may assume that each component
of _G_ − _v_ intersects {_s_[^1]:_, t_[^1]:_,... , sk, tk_} (otherwise we could delete it from _G_ without
violating the cut condition). This implies that we can extend _G_ planarly by an edge
_e_ connecting some vertices _u_′ and _u_′′ in different components of _G_ − _v_, in such a way
that _u_′ ∈{_si_′_, ti_′} and _u_′′ ∈{_si_′′_, ti_′′} for some _i_′ [^6]:= _i_′′ and that _s_[^1]:_, t_[^1]:_,... , sk, tk_ are still
on the boundary of _G_∪_e_. The cut condition holds for _G_∪_e_ (a fortiori), but pairwise
vertex-disjoint _si_ −_ti_ paths (_i_ = 1_,... , k_) do not exist in _G_∪_e_ (since we cannot make
use of edge _e_, as _i_′ [^6]:= _i_′′). Repeating this we end up with a 2-connected graph.
If _G_ is 2-connected, the boundary of _G_ is a simple circuit. Now we apply the itera-
tive step (29). That is, we find, without loss of generality, that the path _P_(_sk, tk_) from
_sk_ to _tk_ clockwise along the boundary of _G_ does not contain any _s_[^1]:_, t_[^1]:_,... , sk_−[^1]:_, tk_−[^1]:.
Let _Pk_ be the corresponding _sk_ − _tk_ path.
Again, let _G_′ be the graph obtained from _G_ by deleting all vertices in _Pk_, together
with all edges incident with them. Let _I_ and _I_′ denote the unbounded faces of _G_ and
_G_′, respectively (we take _I_ and _I_′ as _open_ regions). So _I_ ⊆ _I_′.
Now _G_′ does not contain pairwise vertex-disjoint _si_ − _ti_ paths (_i_ = 1_,... , k_ − 1),
since by assumption _G_ does not contain pairwise vertex-disjoint _si_ − _ti_ paths (_i_ =
[^1]:_,... , k_). Hence, by the induction hypothesis, there exists a simple closed curve _C_
with |_C_ ∩ _G_′| smaller than the number of pairs {_s_[^1]:_, t_[^1]:}_,... ,_{_sk_−[^1]:_, tk_−[^1]:} separated by
_C_. We may assume that _C_ traverses each of the connected regions _I_′_, I_ and _I_′ \ _I_
at most once. That is, each of _C_ ∩ _I_′_, C_ ∩ _I_ and _C_ ∩ (_I_′ \ _I_) is connected (possibly
empty).
If _C_ ∩(_I_′\_I_) is empty, then _C_ ∩_G_ = _C_ ∩_G_′ and hence _C_ violates the cut condition
also for _G_. If _C_ ∩ _I_ is empty, then _C_ does not separate any {_si, ti_} except for those
intersected by _C_. Then _C_ cannot violate the cut condition for _G_′.
If both _C_∩_I_ and _C_∩(_I_′\_I_) are nonempty, we may assume that |_C_∩_G_| = |_C_∩_G_′|+1
and that _C_ separates {_sk, tk_} (since each face of _G_ contained in _I_′ is incident with at
least one vertex on _Pk_). It follows that _C_ violates the cut condition for _G_.

**Application 9.5: VLSI-routing.** The VLSI-routing problem asks for the routes that
wires should make on a chip so as to connect certain pairs of pins and so that wires con-
necting different pairs of pins are disjoint.
Since the routes that the wires potentially can make form a graph, the problem to be

```
1 2
3
4
```

```
11 12 1
```

```
7
13
```

```
9
16
15
11 14
```

```
5 6
7
8
10
```

```
6 5
14
```

```
9
```

```
3
```

```
11
```

```
4 13
```

```
8
```

```
16 10
15 2
12
```

```
Figure 9.8
```

solved can be modeled as a disjoint paths problem. Consider an example of such a problem
as in Figure 9.8 — relatively simple, since generally the number of pins to be connected
is of the order of several thousands. The grey areas are ‘modules’ on which the pins are
located. Points with the same label should be connected.

```
1 2
3
4
```

```
11 12 1
```

```
7
13
```

```
9
16
15
11 14
```

```
5 6
7
8
10
```

```
6 5
14
```

```
9
```

```
3
```

```
11
```

```
4 13
```

```
8
```

```
16 10
15 2
12
```

```
Figure 9.9
```

In the example, the graph is a ‘grid graph’, which is typical in VLSI-design since it
facilitates the manufacturing of the chip and it ensures a certain minimum distance between
disjoint wires. But even for such graphs the disjoint paths problem is NP-complete.
Now the following two-step approach was proposed by Pinter [1983]. First choose the
‘homotopies’ of the wires; for instance like in Figure 9.9. That is, for each _i_ one chooses a
curve _Ci_ in the plane connecting the two vertices _i_, in such a way that they are pairwise
disjoint, and such that the modules are not traversed (Figure 9.9).

Second, try to find disjoint paths _P_[^1]:_,...,Pk_ in the graph such that _Pi_ is homotopic to
_Ci_, in the space obtained from the plane by taking out the rectangles forming the modules;
that is, the paths _Pi_ should be obtained from the curves _Ci_ by shifting _Ci_ over the surface,
but not over any module, fixing the end points of the curve. In Figure 9.10 such a solution
is given.

```
1 2
3
4
```

```
11 12 1
```

```
7
13
```

```
9
16
15
11 14
```

```
5 6
7
8
10
```

```
6 5
14
```

```
9
```

```
3
```

```
11
```

```
4 13
```

```
8
```

```
16 10
15 2
12
```

```
Figure 9.10
```

It was shown by Leiserson and Maley [1985] that this second step can be performed
in polynomial time. So the hard part of the problem is the first step: finding the right
topology of the layout.
Cole and Siegel [1984] proved a Menger-type cut theorem characterizing the existence of
a solution in the second step. That is, if there is no solution for the disjoint paths problem
given the homotopies, there is an ‘oversaturated’ cut: a curve _D_ connecting two holes in
the plane and intersecting the graph less than the number of times _D_ necessarily crosses
the curves _Ci_.
This can be used in a heuristic practical algorithm for the VLSI-routing problem: first
guess the homotopies of the solution; second try to find disjoint paths of the guessed ho-
motopies; if you find them you can stop; if you don’t find them, the oversaturated cut will
indicate a bottleneck in the chosen homotopies; amend the bottleneck and repeat.
Similar results hold if one wants to pack trees instead of paths (which is generally
the case at VLSI-design), and the result can be extended to any planar graph (Schrijver
[1991]). As a theoretical consequence one has that for each fixed number of modules, the
planar VLSI-routing problem can be solved in polynomial time.

**Exercises**

9.11. Extend the algorithm and Theorem 9.4 to the directed case.

9.12. Extend the algorithm and Theorem 9.4 to the following _vertex-disjoint trees problem_:

(31) given:a planar graph _G_ = (_V,E_), sets _R_[^1]:_,...,Rk_ of vertices on the
boundary of _G_,
find:pairwise vertex-disjoint subtrees _T_[^1]:_,...,Tk_ of _G_ so that _Ti_ covers
_Ri_ (_i_ = 1_,...,k_)_._
9.13. Extend the algorithm and Theorem 9.4 to the following problem:

(32) given:a planar graph _G_ = (_V,E_), pairs {_s_[^1]:_,t_[^1]:}_,...,_{_sk,tk_} of vertices
on the boundary of _G_, subgraphs _G_[^1]:_,...,Gk_ of _G_,
find:pairwise vertex-disjoint paths _P_[^1]:_,...,Pk_ where _Pi_ connects _si_ and
_ti_ and where _Pi_ is in _Gi_ (_i_ = 1_,...,k_)_._
9.14. (i) Reduce the edge-disjoint paths problem where all commodities are on the bound-
ary of a planar graph so that the cross-freeness condition is satisfied, to the
vertex-disjoint paths problem(26).
(ii) Show that the cut condition (7) is sufficient in this case of the edge-disjoint
paths problem.

### 9.5. Edge-disjoint paths in planar graphs

The trivially necessary cross-freeness condition for the commodities if they are on
the boundary of a planar graph, turned out to be of great help in handling the
_vertex_-disjoint paths problem: it gives an ordering of the commodities, allowing us to
handling them one by one.
As we saw in Exercise 9.14, the edge-disjoint analogue can be handled in the same
way if the cross-freeness condition holds. In that case, the cut condition (7) is again
sufficient. However, Figure 9.5 shows that without cross-freeness, the cut condition
is not sufficient. That simple example shows that we may not hope for many other
interesting cases where the cut condition is sufficient.
In fact, the complexity of the edge-disjoint paths problem for planar graphs with
all commodities on the boundary, is open. Therefore, we put:
**Research problem**. Is the undirected edge-disjoint paths problem polynomially
solvable for planar graphs with all commodities on the boundary? Is it
NP-complete?
Okamura and Seymour [1981] showed that the problem is polynomially solvable
if we pose the following _Euler condition_:

(33) the graph (_V, E_ ∪{{_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_}}) is Eulerian.

(We have parallel edges if some {_si, ti_} coincide or form an edge of _G_.) Moreover,
they showed that with the Euler condition, the cut condition is a sufficient condition.
(Thus we have an analogue to Rothschild and Whinston’s theorem (Theorem 9.2).)

```
We here observe that the Euler condition (33) implies that for each U ⊆ V :
```

(34) |_δE_(_U_)|≡ number of _i_ with |_U_ ∩{_si, ti_}| = 1 (mod 2).

**Theorem 9.5** (Okamura-Seymour theorem)**.** _Let G_ = (_V, E_) _be a planar graph and_
_let_ {_s_[^1]:_, t_[^1]:}_,... ,_{_sk, tk_} _be pairs of vertices on the boundary of G such that the Euler_
_condition_ (33) _holds. Then the edge-disjoint paths problem has a solution if and only_
_if the cut condition holds._

**Proof.** Necessity of the cut condition being trivial, we show sufficiency. The cut
condition implies that |_R_| ≤ |_E_| (assuming that each _r_ ∈ _R_ consists of two distinct
vertices), since

(35) [^2]:|_R_| =

```
∑
```

```
v∈V
```

```
degR(v) ≤
```

```
∑
```

```
v∈V
```

```
degE(v) = 2|E|.
```

So we can consider a counterexample with 2|_E_|−|_R_| minimal. Then

(36) _G_ is 2-connected.

Indeed, if _G_ is disconnected, we can deal with the components separately. Suppose
next that _G_ is connected and has a cut vertex _v_. We may assume that for no _r_ =
_st_ ∈ _R_, the vertices _s_ and _t_ belong to different components of _G_ − _v_, since otherwise
we can replace _r_ by _sv_ and _vt_, without violating the Euler or cut condition. For any
component _K_ of _G_ − _v_ consider the graph induced by _K_ ∪{_v_}. Again, the Euler
and cut conditions hold (with respect to those nets contained in _K_ ∪{_v_}). So by the
minimality of 2|_E_|−|_R_|, we know that paths as required exist in _K_ ∪{_v_}. As this is
the case for each component of _G_ − _v_, we have paths as required in _G_. This proves
(36).
Let _C_ be the circuit formed by the outer boundary of _G_. If some _r_ ∈ _R_ has the
same ends as some edge _e_ of _G_, we can delete _e_ from _G_ and _r_ from _R_, and obtain a
smaller counterexample. So no such _r_ exists.
Call a subset _X_ of _V tight_ if _dE_(_X_) = _dR_(_X_). Then

(37) there exists a tight subset _X_ of _V_ such that _δE_(_X_) intersects _EC_ in precisely
two edges.

Indeed, if there is no tight set _X_ with ∅ 6= _X_ [^6]:= _V_ , we can choose an edge _e_ ∈ _EC_,
and replace _E_ and _R_ by _E_ \{_e_} and _R_∪{_e_}. This does not violate the cut condition,
and hence would give a smaller counterexample.
So there exists a tight proper nonempty subset _X_ of _V_. Choose _X_ with |_δE_(_X_)|
minimal. Then _G_[_X_] and _G_ − _X_ are connected. For suppose that (say) _G_[_X_] is not

connected. Let _K_ be a component of _G_[_X_]. Then

(38) |_δE_(_K_)| + |_δE_(_X_ \ _K_)|≥|_δR_(_K_)| + |_δR_(_X_ \ _K_)|≥|_δR_(_X_)|
= |_δE_(_X_)| = |_δE_(_K_)| + |_δE_(_X_ \ _K_)|.

So _K_ is tight, while |_δE_(_K_)| _<_ |_δE_(_X_)|, contradicting the minimality assumption.
Hence _G_[_X_] and _G_ − _X_ are connected, implying (37).
Choose a set _X_ as in (37) with |_X_| minimal. Let _e_ be one of the two edges in _EC_
that belong to _δE_(_X_). Say _e_ = _uw_ with _u_ 6∈ _X_ and _w_ ∈ _X_.
Since _dR_(_X_) = _dE_(_X_) ≥ 2, we know _δR_(_X_) [^6]:= ∅. For each _r_ ∈ _δR_(_X_), let _sr_ be
the vertex in _r_ ∩ _X_, and _tr_ the vertex in _r_ \ _X_. Choose _r_ ∈ _δR_(_X_) such that _tr_ is as
close as possible to _u_ in the graph _C_ − _X_.
Since _sr_ and _tr_ are nonadjacent, we know that {_sr, tr_} 6= {_u, w_}. So we can
choose _v_ ∈ {_u, w_}\{_sr, tr_}. Let _R_′ := (_R_ \{_r_}) ∪{_srv, vtr_}. Trivially the Euler
condition is maintained. We show that also the cut condition is maintained, yielding
a contradiction as 2|_E_|−|_R_′| _<_ [^2]:|_E_|−|_R_| and as a solution for _R_′ yields a solution
for _R_.
To see that the cut condition is maintained, suppose to the contrary that there is
a _Y_ ⊆ _V_ satisfying

(39) _dE_(_Y_ ) _< dR_′(_Y_ )_._

By choosing _Y_ under the additional condition that _dE_(_Y_ ) is as small as possible, we
have that _G_[_Y_ ] and _G_−_Y_ are connected. So _δE_(_Y_ ) has two edges on _C_. By symmetry
we can assume that _tr_ 6∈ _Y_. By the Euler condition, (39) implies _dE_(_Y_ ) ≤ _dR_′(_Y_ )−2.
So

(40) _dR_′(_Y_ ) ≥ _dE_(_Y_ ) + 2 ≥ _dR_(_Y_ ) + 2 ≥ _dR_′(_Y_ )_._

Hence we have equality throughout. So _δR_′(_Y_ ) contains both _srv_ and _vtr_, that is,
_sr, tr_ 6∈ _Y_ and _v_ ∈ _Y_. Moreover, _dE_(_Y_ ) = _dR_(_Y_ ).
By the choice of _r_, there is no pair _r_′ in _R_ connecting _X_ \ _Y_ and _Y_ \ _X_ (since
then _tr_′ ∈ _Y_ \ _X_, and hence _tr_′ is closer than _tr_ to _u_ in _C_ − _X_). This implies

(41) _dR_(_X_ ∩ _Y_ ) + _dR_(_X_ ∪ _Y_ ) = _dR_(_X_) + _dR_(_Y_ )_._

Moreover,

(42) _dE_(_X_ ∩ _Y_ ) + _dE_(_X_ ∪ _Y_ ) ≤ _dE_(_X_) + _dE_(_Y_ )_._

As the cut condition holds for _X_ ∩ _Y_ and _X_ ∪ _Y_ , we have equality in (42), and
therefore _X_ ∩ _Y_ is tight. Since _sr_ ∈ _X_ \ _Y_ , we know |_X_ ∩ _Y_ | _<_ |_X_|. So by the

minimality of _X_ we have _X_ ∩_Y_ = ∅. So _w_ 6∈ _Y_ , hence _u_ = _v_ ∈ _Y_. Then edge _e_ = _uw_
connects _X_ \ _Y_ and _Y_ \ _X_, contradicting equality in (42).

Clearly, this method gives a polynomial-time algorithm for finding the paths,
since we can determine a minimum-cardinality cut containing _e_′ and _e_′′, for any pair
of edges _e_′_, e_′′ on the boundary of _G_ (cf. Exercise 9.16).
Becker and Mehlhorn [1986] and Matsumoto, Nishizeki, and Saito [1985] gave
implementations with running time of order _O_(|_E_|[^2]). Recently, Wagner and Weihe
[1995] found a linear-time algorithm.

**Exercises**

9.15. Let _G_ = (_V,E_) be a finite subgraph of the rectangular grid graph in R[^2], such that
each bounded face of _G_ is a square of area 1. Let {_s_[^1]:_,t_[^1]:}_,...,_{_sk,tk_} be pairs of
vertices on the boundary of _G_ such that each vertex of (_V,E_∩{{_s_[^1]:_,t_[^1]:}_,...,_{_sk,tk_}})
has degree even and at most 4. A cut is called a _1-bend cut_ if it is the set of edges
crossed by the union of some horizontal and some vertical half-line with one common
end vertex.
Show that the cut condition holds whenever it holds for all 1-bend cuts.

9.16. Let _G_ be a planar graph and let _e_′ and _e_′′ be edges on the boundary of _G_. Reduce
the problem of finding a minimum-cardinality cut containing _e_′ and _e_′′ to a shortest
path problem.

### 9.6. A column generation technique for multicom-

### modity flows

The fractional multicommodity flow problem (1) asks for flows _x_[^1]:_,... , xk_ of given
values _d_[^1]:_,... , dk_ such that the total amount of flow through any arc _e_ does not
exceed the capacity of _e_. So it amounts to finding a solution to the following system
of linear inequalities in the _k_|_E_| variables _xi_(_e_) (_i_ = 1_,... , k_; _e_ ∈ _E_):

(43) (i)

```
∑
```

```
e∈δout(v)
```

```
xi(e) −
```

```
∑
```

```
e∈δin(v)
```

```
xi(e) = 0 (i = 1,... , k; v ∈ V, v 6= si, ti),
```

```
(ii)
```

```
∑
```

```
e∈δout(si)
```

```
xi(e) −
```

```
∑
```

```
e∈δin(si)
```

```
xi(e) = di (i = 1,... , k),
```

```
(iii)
```

```
∑k
```

```
i=1
```

```
xi(e) ≤ c(e) (e ∈ E),
```

```
(iv) xi(e) ≥ 0 (i = 1,... , k; e ∈ E).
```

Thus any linear programming method can solve the multicommodity flow problem.
In particular, the problem is solvable in polynomial time.

Since for each fixed _i_ = 1_,... , k_, a solution _xi_ to (43) is an _si_ − _ti_ flow, we can
decompose _xi_ as a nonnegative combination of _si_−_ti_ paths. That is, there exist _si_−_ti_
paths _Pi_[^1]:_,... , Pini_ and nonnegative reals _zi_[^1]:_,... , zini_ satisfying:

(44) (i)

```
∑ni
```

```
j=1
```

```
zijXPij(e) = xj(e) (e ∈ E),
```

```
(ii)
```

```
∑ni
```

```
j=1
```

```
zij = di.
```

Here X_P_ denotes the _incidence vector_ of _P_ in Q_E_, that is, X_P_(_e_) = 1 if _P_ traverses
_e_, and = 0 otherwise.

Hence the multicommodity flow problem amounts to finding paths _Pij_ and non-
negative reals _zij_, where _Pij_ is an _si_ − _ti_ path, such that:

(45) (i)

```
∑ni
```

```
j=1
```

```
zij = di (i = 1,... , k),
```

```
(ii)
```

```
∑k
```

```
i=1
```

```
∑ni
```

```
j=1
```

```
zijXPij(e) ≤ c(e) (e ∈ E).
```

This formulation applies to both the directed and the undirected problems.
Solving (45) again amounts to solving a system of linear inequalities, albeit with
an enormous number of variables: one variable for each _i_ = 1_,... , k_ and each _si_ − _ti_
path.
Ford and Fulkerson [1958] showed that this large number of variables can be
avoided when solving the problem with the simplex method. The variables can be
handled implicitly by using a _column generation technique_ as follows.
First convert the problem to a maximization problem. To this end, add, for each
_i_ = 1_,... , k_, a vertex _s_′
_i_ and an arc _s_′
_isi_, with capacity equal to _di_. Then we can
delete the constraint (45)(i), and maximize

∑
_i,j zij_ over the remaining constraints
(replacing _si_ by _s_′
_i_). If the maximum value is equal to

∑
_i di_ we have a solution to
(45). If the maximum value is less, then (45) has no nonnegative solution _zij_.
Having this reduction, we see that the problem is equivalent to the following LP-
problem. Let P be the collection of all _si_ − _ti_ paths for all _i_ = 1_,... , k_. Then:

(46) maximize:
_P_∈P

```
zP
```

```
subject to: (i)
```

```
∑
```

```
P∈P
```

```
zPXP(e) ≤ c(e) (e ∈ E),
```

```
(ii) zP ≥ 0 (P ∈P).
```

When solving (46) with the simplex method we first should add a slack variable _ze_
for each _e_ ∈ _E_. Thus if _A_ denotes the _E_ ×P matrix with the incidence vectors of
all paths in P as its columns (in some order) and _w_ is the vector in RP × R_E_ with
_wP_ = 1 (_P_ ∈P) and _we_ = 0 (_e_ ∈ _E_), we solve:

(47)

```
maximize: wTz
subject to: [A I]z = c,
z ≥ 0.
```

Now each simplex tableau is completely determined by the set of variables in the
current basis. So knowing subsets P′ of P and _E_′ of _E_, giving the indices of variables
in the basis, is enough to know implicitly the whole tableau. Note that |P′|+|_E_′| = _E_.
So although the tableau is exponentially large, it can be represented in a concise way.

Let _B_ be the matrix consisting of those columns of [_A I_] corresponding to P′
and _E_′. So the rows of _B_ are indexed by _E_ and the columns by P′ ∪ _E_′. The basic
solution corresponding to _B_ is easily computed: the vector _B_−[^1]_c_ gives the values for
_zP_ if _P_ ∈ P′ and for _ze_ if _e_ ∈ _E_′, while we set _zP_ := 0 if _P_ 6∈ P′ and _ze_ := 0 if
_e_ 6∈ _E_′. (Initially, _B_ = _I_, that is P′ = ∅ and _E_′ = _E_, so that _zP_ = 0 for all _P_ ∈ P
and _ze_ = _c_(_e_) for all _e_ ∈ _E_.)
Now we should describe pivoting (that is, finding variables leaving and entering
the basis) and checking optimality. Interestingly, it turns out that this can be done
by solving a set of shortest path problems.
First consider the dual variable corresponding to an edge _e_. It has value (in the
current tableau):

(48) _wBB_−[^1]_εe_ − _we_ = _wB_(_B_−[^1])_e_

where as usual _wB_ denotes the part of vector _w_ corresponding to _B_ (that is, corre-
sponding to P′ and _E_′) and where _εe_ denotes the _e_-th unit basis vector in R_E_ (which
is the column corresponding to _e_ in [_A I_]). Note that the columns of _B_−[^1] are in-
dexed by _E_; then (_B_−[^1])_e_ is the column corresponding to _e_. Note also that _we_ = 0 by
definition.
Similarly, the dual variable corresponding to a path _P_ in P has value:

(49) _wBB_−[^1]X_P_ − _wP_ = [

```
∑
```

```
e∈P
```

```
wB(B−[^1])e] − 1.
```

(Note that X_P_ is the column in [_A I_] corresponding to _P_.)
In order to pivot, we should identify a negative dual variable. To this end, we
first check if (48) is negative for some edge _e_. If so, we choose such an edge _e_ and
take _ze_ as the variable entering the basis. Selecting the variable leaving the basis now
belongs to the standard simplex routine. We only have to consider that part of the
tableau corresponding to P′_, E_′ and _e_. We select an element _f_ in P′ ∪ _E_′ for which
the quotient _zf/_(_B_−[^1])_f e_ has positive denominator and is as small as possible. Then
_zf_ is the variable leaving the basis.
Suppose next that (48) is nonnegative for each edge _e_. We consider _wB_(_B_−[^1])_e_ as
the length _l_(_e_) of _e_. Then for any path _P_,

(50)

```
∑
```

```
e∈P
```

```
wB(B−[^1])e
```

is equal to the length

∑
_e_∈_P l_(_e_) of _P_. Hence, finding a dual variable (49) of negative
value is the same as finding a path in P of length less than 1.
Such a path can be found by applying any shortest path algorithm: for each
_i_ = 1_,... , k_, we find a shortest _si_ − _ti_ path (with respect to _l_). If each of these
shortest paths has length at least 1, we know that all dual variables have nonnegative
value, and hence the current basic solution is optimum.
If we find some _si_ − _ti_ path _P_ of length less than 1, we choose _zP_ as variable
entering the basis. Again selecting a variable leaving the basis is standard: we select
an element _f_ in P′∪_E_′ for which the quotient _zf/_(_B_−[^1]X_P_)_f_ has positive denominator
and is as small as possible.
This describes pivoting. In order to avoid “cycling”, we apply a lexicographic rule
for selecting the variable leaving the basis. We order the edges of _G_ arbitrarily. Now
in case there is a tie in selecting the _f_ ∈P′ ∪_E_′ for which the corresponding quotient
is as small as possible, we choose the _f_ ∈P′ ∪ _E_′ for which the vector

(51) (_B_−[^1])_f/_(_B_−[^1])_f e_ (if _e_ enters the basis)_,_
(_B_−[^1])_f/_(_B_−[^1]X_P_)_f_ (if _P_ enters the basis)_,_
is lexicographically as small as possible. In Exercise 9.17 we will see that this avoids
cycling.

**Exercises**

9.17. (i) Apply the lexicographic rule above, and consider a simplex tableau, correspond-
ing to P′ and _E_′ say. Show that for each _f_ ∈ P′ ∪ _E_′: if _zf_ = 0 then the first
nonzero entry in the vector (_B_−[^1])_f_ is positive. (Use induction on the number
of pivot steps performed.)
(ii) Derive from (i) that, when applying the lexicographic rule, at each pivot iter-
ation, if the objective value of the solution does not increase, then the vector
_wBB_−[^1] increases lexicographically.

```
(iii) Derive that the lexicographic rule leads to termination of the method.
```

9.18. Modify the column generation technique to solve the following problem: given a
directed graph _G_ = (_V,E_), a capacity function _c_ : _E_ → Q+, commodities (_s_[^1]:_,t_[^1]:)_,...,_
(_sk,tk_) and ‘profits’ _p_[^1]:_,...,pk_ ∈ Q+, find vectors _x_[^1]:_,...,xk_ in Q_E_ and rationals
_d_[^1]:_,...,dk_ so that:

```
(52) (i)xi is an si − ti flow of value di (i = 1,...,k),
```

```
(ii)
```

```
∑k
```

```
i=1
```

```
xi(e) ≤ c(e) (e ∈ E),
```

```
(iii)
```

```
∑k
```

```
i=1
```

```
pidi is as large as possible.
```

9.19. Let _Pij_ and _zij >_ 0 form a solution to the undirected form of (45) and let _W_ ⊆ _V_ be
so that the capacity of _δE_(_W_) is equal to the demand of _δR_(_W_). Show that each _Pij_
intersects _δE_(_W_) at most once.

9.20. Show that if the multicommodity flow problem has no solution, then Ford and Fulk-
erson’s column generation technique yields a length function _l_ violating (9).

## 10. Matroids

### 10.1. Matroids and the greedy algorithm

```
Let G = (V, E) be a connected undirected graph and let w : E → Z be a ‘weight’
function on the edges. In Section 1.4 we saw that a minimum-weight spanning tree
can be found quite straightforwardly with Kruskal’s so-called greedy algorithm.
The algorithm consists of selecting successively edges e1, e2,... , er. If edges e1,... , ek
have been selected, we select an edge e ∈ E so that:
```

(1) (i) _e_ 6∈{_e_[^1]:_,... , ek_} and {_e_[^1]:_,... , ek, e_} is a forest,
(ii) _w_(_e_) is as small as possible among all edges _e_ satisfying (i).

```
We take ek+1 := e. If no e satisfying (1)(i) exists, that is, if {e1,... , ek} forms
a spanning tree, we stop, setting r := k. Then {e1,... , er} is a spanning tree of
minimum weight.
By replacing ‘as small as possible’ in (1)(ii) by ‘as large as possible’, one obtains
a spanning tree of maximum weight.
It is obviously not true that such a greedy approach would lead to an optimal
solution for any combinatorial optimization problem. We could think of such an
approach to find a matching of maximum weight. Then in (1)(i) we replace ‘forest’
by ‘matching’ and ‘small’ by ‘large’. Application to the weighted graph in Figure 10.1
would give e1 = cd, e2 = ab.
```

```
1
```

```
3
```

```
a b
```

```
d c
```

```
3
```

```
4
Figure 10.1
```

```
However, ab and cd do not form a matching of maximum weight.
It turns out that the structures for which the greedy algorithm does lead to an
optimal solution, are the matroids. It is worth studying them, not only because it
enables us to recognize when the greedy algorithm applies, but also because there
exist fast algorithms for ‘intersections’ of two different matroids.
```

```
The concept of matroid is defined as follows. Let X be a finite set and let I be a
collection of subsets of X. Then the pair (X,I) is called a matroid if it satisfies the
following conditions:
```

(2) (i) ∅∈I,
(ii) if _Y_ ∈I and _Z_ ⊆ _Y_ then _Z_ ∈I,
(iii) if _Y, Z_ ∈I and |_Y_ | _<_ |_Z_| then _Y_ ∪{_x_}∈I for some _x_ ∈ _Z_ \ _Y_.

```
For any matroid M = (X,I), a subset Y of X is called independent if Y belongs
to I, and dependent otherwise.
Let Y ⊆ X. A subset B of Y is called a basis of Y if B is an inclusionwise maximal
independent subset of Y. That is, for any set Z ∈I with B ⊆ Z ⊆ Y one has Z = B.
It is not difficult to see that condition (2)(iii) is equivalent to:
```

```
(3) for any subset Y of X, any two bases of Y have the same cardinality.
```

```
(Exercise 10.1.) The common cardinality of the bases of a subset Y of X is called the
rank of Y , denoted by rM(Y ).
We now show that if G = (V, E) is a graph and I is the collection of forests in
G, then (E,I) indeed is a matroid. Conditions (2)(i) and (ii) are trivial. To see
that condition (3) holds, let E′ ⊆ E. Then, by definition, each basis Y of E′ is an
inclusionwise maximal forest contained in E′. Hence Y forms a spanning tree in each
component of the graph (V, E′). So Y has |V |− k elements, where k is the number
of components of (V, E′). So each basis of E′ has |V |− k elements, proving (3).
A set is called simply a basis if it is a basis of X. The common cardinality of all
bases is called the rank of the matroid. If I is the collection of forests in a connected
graph G = (V, E), then the bases of the matroid (E,I) are exactly the spanning trees
in G.
```

```
We next show that the matroids indeed are those structures for which the greedy
algorithm leads to an optimal solution. Let X be some finite set and let I be a
collection of subsets of X satisfying (2)(i) and (ii).
For any weight function w : X → R we want to find a set Y in I maximizing
```

```
(4) w(Y ) :=
```

```
∑
```

```
y∈Y
```

```
w(y).
```

```
The greedy algorithm consists of selecting y1,... , yr successively as follows. If y1,... , yk
have been selected, choose y ∈ X so that:
```

(5) (i) _y_ 6∈{_y_[^1]:_,... , yk_} and {_y_[^1]:_,... , yk, y_}∈I,
(ii) _w_(_y_) is as large as possible among all _y_ satisfying (i).

```
We stop if no y satisfying (5)(i) exist, that is, if {y1,... , yk} is a basis.
Now:
```

```
Theorem 10.1. The pair (X,I) satisfying (2)(i) and (ii) is a matroid if and only if
```

_the greedy algorithm leads to a set Y in_ I _of maximum weight w_(_Y_ )_, for each weight_
_function w_ : _X_ → R+_._
**Proof.** _Sufficiency._ Suppose that the greedy algorithm leads to an independent set
of maximum weight for each weight function _w_. We show that (_X,_I) is a matroid.
Conditions (2)(i) and (ii) are satisfied by assumption. To see condition (2)(iii),
let _Y, Z_ ∈I with |_Y_ | _<_ |_Z_|. Suppose that _Y_ ∪{_z_}6∈I for each _z_ ∈ _Z_ \ _Y_.
Consider the following weight function _w_ on _X_. Let _k_ := |_Y_ |. Define:

(6) _w_(_x_) := _k_ + 2 if _x_ ∈ _Y_ ,
_w_(_x_) := _k_ + 1 if _x_ ∈ _Z_ \ _Y_ ,
_w_(_x_) := 0 if _x_ ∈ _X_ \ (_Y_ ∪ _Z_).

Now in the first _k_ iterations of the greedy algorithm we find the _k_ elements in
_Y_. By assumption, at any further iteration, we cannot choose any element in _Z_ \ _Y_.
Hence any further element chosen, has weight 0. So the greedy algorithm will yield a
basis of weight _k_(_k_ + 2).
However, any basis containing _Z_ will have weight at least |_Z_ ∩ _Y_ |(_k_ + 2) + |_Z_ \
_Y_ |(_k_ + 1) ≥|_Z_|(_k_ + 1) ≥ (_k_ + 1)(_k_ + 1) _> k_(_k_ + 2). Hence the greedy algorithm does
not give a maximum-weight independent set.

_Necessity._ Now let (_X,_I) be a matroid. Let _w_ : _X_ → R+ be any weight function on
_X_. Call an independent set _Y greedy_ if it is contained in a maximum-weight basis. It
suffices to show that if _Y_ is greedy, and _x_ is an element in _X_\_Y_ such that _Y_ ∪{_x_}∈I
and such that _w_(_x_) is as large as possible, then _Y_ ∪{_x_} is greedy.
As _Y_ is greedy, there exists a maximum-weight basis _B_ ⊇ _Y_. If _x_ ∈ _B_ then
_Y_ ∪{_x_} is greedy again. If _x_ 6∈ _B_, then there exists a basis _B_′ containing _Y_ ∪{_x_}
and contained in _B_ ∪{_x_}. So _B_′ = (_B_ \{_x_′}) ∪{_x_} for some _x_′ ∈ _B_ \ _Y_. As _w_(_x_)
is chosen maximum, _w_(_x_) ≥ _w_(_x_′). Hence _w_(_B_′) ≥ _w_(_B_), and therefore _B_′ is a
maximum-weight basis. So _Y_ ∪{_x_} is greedy.

Note that by replacing “as large as possible” in (5) by “as small as possible”, one
obtains an algorithm for finding a _minimum_-weight basis in a matroid. Moreover,
by ignoring elements of negative weight, the algorithm can be adapted to yield an
independent set of maximum weight, for any weight function _w_ : _X_ → R.

**Exercises**

10.1. Show that condition (3) is equivalent to condition (2)(iii) (assuming (2)(i) and (ii)).

10.2. Let _M_ = (_X,_I) be a matroid. Two elements _x,y_ of _X_ are called _parallel_ if {_x,y_} is
a circuit. Show that if _x_ and _y_ are parallel and _Y_ is an independent set with _x_ ∈ _Y_ ,
then also (_Y_ \{_x_}) ∪{_y_} is independent.

```
10.3. Let M = (X,I) be a matroid, with X = {x1,...,xm}. Define
```

```
(7) Y := {xi | rM({x1,...,xi}) > rM({x1,...,xi−1})}.
```

```
Prove that Y belongs to I.
```

### 10.2. Equivalent axioms for matroids

```
The definition of the notion of matroid given in the previous section is given by
‘axioms’ in terms of the independent sets. There are several other axioms that char-
acterize matroids. In this section we give a number of them.
Let X be a finite set, and let I be a nonempty down-monotone collection of
subsets of X; that is, if F ∈ I and F′ ⊆ F, then F′ ∈ I. Let B be the collection of
inclusionwise maximal sets in I, and let C be the collection of inclusionwise minimimal
sets that are not in I. Finally, for any subset Y of X, define
```

```
(8) r(Y ) := max{|Z|| Z ⊆ Y, Z ∈I}.
```

```
Obviously, knowing one of the objects I, B, C, r, we know all the other. Moreover,
any nonempty antichain[^22] B arises in this way from some nonempty down-monotone
collection I of subsets. Similarly, any antichain C consisting of nonempty sets arises
in this way. Finally, r arises in this way if and only if
```

(9) (i) _r_(∅) = 0,
(ii) if _Z_ ⊆ _Y_ ⊆ _X_ then _r_(_Z_) ≤ _r_(_Y_ ).

```
We can now characterize when such objects arise from a matroid (X,I). That is,
we obtain the following equivalent characterizations of matroids.
```

```
Theorem 10.2. Let I, B, C, and r be as above. Then the following are equivalent:
```

```
(i) if F, F′ ∈I and |F′| > |F|, then F ∪{x}∈I for some x ∈ F′ \ F;
(ii) if B, B′ ∈B and x ∈ B′ \ B, then (B′ \{x}) ∪{y}∈B for some y ∈ B \ B′;
(iii) if B, B′ ∈B and x ∈ B′ \ B, then (B \{y}) ∪{x}∈B for some y ∈ B \ B′;
(iv) if C, C′ ∈C with C 6= C′ and x ∈ C ∩ C′, then (C ∪ C′) \{x} contains a set in
C;
(v) if C, C′ ∈C, x ∈ C ∩ C′, and y ∈ C \ C′, then (C ∪ C′) \{x} contains a set in
C containing y;
(vi) for all Y, Z ⊆ X one has
```

```
22An antichain is a collection of sets no two of which are contained in each other.
```

```
(10) r(Y ∩ Z) + r(Y ∪ Z) ≤ r(Y ) + r(Z).
```

**Proof.** (i)⇒(ii): (i) directly implies that all sets in B have the same size. Now let
_B, B_′ ∈ B and _x_ ∈ _B_′ \ _B_. Since _B_′ \{_x_} ∈ I, by (i) there exists a _y_ ∈ _B_ \ _B_′ such
that _B_′′ := (_B_′ \{_x_}) ∪{_y_}∈I. Since |_B_′′| = |_B_′|, we know _B_′′ ∈B.

(iii)⇒(i): Let _F, F_′ form a counterexample to (i) with |_F_ ∩_F_′| as large as possible.
Consider sets _B, B_′ in B with _F_ ⊆ _B_ and _F_′ ⊆ _B_′.
As _F, F_′ is a counterexample, we know _F_ 6⊆ _B_′. Choose _x_ ∈ _F_ \ _B_′. Then by
(iii), (_B_′ \{_y_}) ∪{_x_} for some _y_ ∈ _B_′ \ _B_. Hence replacing _F_′ by (_F_′ \{_y_}) ∪{_x_} we
would keep a counterexample but increase |_F_ ∩ _F_′|, a contradiction.
(ii)⇒(iii): By the foregoing we know that (iii) implies (ii). Now axioms (ii) and
(iii) interchange if we replace B by the collection of complements of sets in B. Hence
also the implication (ii)⇒(iii) holds.
(i)⇒(v): If (i) holds, then by the foregoing, also (ii) holds. Let _C, C_′ ∈ C and
_x_ ∈ _C_ ∩ _C_′, _y_ ∈ _C_ \ _C_′. We can assume that _X_ = _C_ ∪ _C_′. Let _B, B_′ ∈ B with
_B_ ⊇ _C_ \{_y_} and _B_′ ⊇ _C_′ \{_x_}. Then _y_ 6∈ _B_ and _x_ 6∈ _B_′ (since _C_ 6⊆ _B_ and _C_′ 6⊆ _B_′).
We can assume that _y_ 6∈ _B_′. Otherwise, _y_ ∈ _B_′ \_B_, and hence by (ii), there exists
a _z_ ∈ _B_ \ _B_′ with _B_′′ := (_B_′ \{_y_}) ∪{_z_}∈B. Then _z_ [^6]:= _x_, since otherwise _C_′ ⊆ _B_′′.
Hence, replacing _B_′ by _B_′′ gives _y_ 6∈ _B_′.
As _y_ 6∈ _B_′, we know _B_′ ∪{_y_} 6∈ I, and hence there exists a _C_′′ ∈ C contained in
_B_′ ∪{_y_}. As _C_′′ 6⊆ _B_′, we know _y_ ∈ _C_′′. Moreover, as _x_ 6∈ _B_′ we know _x_ 6∈ _C_′′.
(v)⇒(iv): is trivial.
(iv)⇒(i): Let _F, F_′ form a counterexample to (i) with |_F_ ∩ _F_′| maximal. Then
_F_ 6⊆ _F_′, and so we can choose _y_ ∈ _F_ \ _F_′. By the maximality of |_F_ ∩ _F_′|, we know
_F_′ ∪{_x_} 6∈ I. So there is a _C_ ∈ C contained in _F_′ ∪{_x_}. As _C_ 6⊆ _F_′ we know
_x_ ∈ _C_. Then _C_ is the unique set in C contained in _F_′ ∪{_x_}. For suppose there is
another, _C_′ say. Again, _x_ ∈ _C_′, and hence by (iv) there exists a _C_′′ ∈C contained in
(_C_ ∪ _C_′) \{_x_}. But then _C_′′ ⊆ _F_′, a contradiction.
As _C_ 6⊆ _F_, _C_ intersects _F_′\_F_. Choose _y_ ∈ _C_∩(_F_′\_F_). Then _F_′′ := (_F_′∪{_x_})\{_y_}
does not contain any set in C (as _C_ is the only set in C contained in _F_′ ∪ {_x_}).
Then replacing _F_′ by _F_′′, we would keep a counterexample while increasing |_F_′ ∩ _F_|,
contradicting our assumption.
(i)⇒(vi): Choose _Y, Z_ ⊆ _X_. Let _F_ be an inclusionwise maximal set in I with
_F_ ⊆ _Y_ ∩ _Z_, and let _F_′ be an inclusionwise maximal set in I with _F_ ⊆ _F_ ⊆ _Y_ ∪ _Z_.
By (i) we know that _r_(_Y_ ∩ _Z_) = |_F_| and _r_(_Y_ ∪ _Z_) = |_F_′|. Then

(11) |_F_′ ∩ _Y_ | + |_F_′ ∩ _Z_| = |_F_′ ∩ (_Y_ ∩ _Z_)| + |_F_′ ∩ (_Y_ ∪ _Z_)|≥|_F_| + |_F_′|_,_

and hence we have (10).

(vi)⇒(i): Let _F, F_′ ∈ I with |_F_| _<_ |_F_′|. Let _U_ be the largest subset of _F_′ \ _F_
with _r_(_F_ ∪ _U_) = |_F_|. Then _U_ [^6]:= _F_′ \ _F_, since _r_(_F_ ∪ _F_′) ≥|_F_′| _>_ |_F_|. So there exists
an _x_ ∈ _F_′ \ _F_ ∪ _U_. If _F_ ∪{_x_}∈I we are done, so we can assume that _F_ ∪{_x_}6∈I;
equivalently, _r_(_F_ ∪{_x_}) = |_F_|. Let _U_′ := _U_ ∪{_x_}. Then by (10),

(12) _r_(_F_ ∪ _U_′) ≤ _r_(_F_ ∪ _U_) + _r_(_F_ ∪{_x_}) − _r_(_F_) = |_F_|_,_

contradicting the maximality of _U_.

Given a matroid _M_ = (_X,_I), any in B is called a _basis_ and any set in C a _circuit_
of _M_. The function _r_ is called _rank function_ of _M_ (often denoted by _rM_), and _r_(_Y_ )
the _rank_ of _Y_.
The symmetry of (ii) and (iii) in Theorem 10.2 immediately implies the following.
Define

(13) B∗ := {_X_ \ _B_ | _B_ ∈B}_._

Then

**Corollary 10.2a.** _If_ B _is the collection of bases of some matroid M, then_ B∗ _also is_
_the collection of bases of some matroid on X, denoted by M_∗_._

**Proof.** Directly from the equivalence of (ii) and (iii) in Theorem 10.2.

The matroid _M_∗ is called the _dual matroid_ of _M_. Since (B∗)∗ = B, we know
(_M_∗)∗ = _M_.

**Theorem 10.3.** _The rank function rM_∗ _of the dual matroid M_∗ _satisfies:_

(14) _rM_∗(_Y_ ) = |_Y_ | + _rM_(_X_ \ _Y_ ) − _rM_(_X_)_._

**Proof.**

(15) _rM_∗(_Y_ ) = max{|_A_ ∩ _Y_ || _A_ ∈B∗} =
= |_Y_ |− min{|_B_ ∩ _Y_ || _B_ ∈B} = |_Y_ |− _rM_(_X_) + max{|_B_ \ _Y_ || _B_ ∈B} =
|_Y_ |− _rM_(_X_) + _rM_(_X_ \ _Y_ )_._

Another way of constructing matroids from matroids is by ‘deletion’ and ‘contrac-
tion’. Let _M_ = (_X,_I) be a matroid and let _Y_ ⊆ _X_. Define

(16) I′ := {_Z_ | _Z_ ⊆ _Y, Z_ ∈I}_._

Then _M_′ = (_Y,_I′) is a matroid again, as one easily checks. _M_′ is called the _restriction_
of _M_ to _Y_. If _Y_ = _X_ \ _Z_ with _Z_ ⊆ _X_, we say that _M_′ arises by _deleting Z_, and
denote _M_′ by _M_ \ _Z_.
_Contracting Z_ means replacing _M_ by (_M_∗ \ _Z_)∗. This matroid is denoted by
_M/Z_. One may check that if _G_ is a graph and _e_ is an edge of _G_, then contracting
edge {_e_} in the cycle matroid _M_(_G_) of _G_ corresponds to contracting _e_ in the graph.
That is, _M_(_G_)_/_{_e_} = _M_(_G/_{_e_}), where _G/_{_e_} denotes the graph obtained from _G_
by contracting _e_.
If matroid _M_′ arises from _M_ by a series of deletions and contractions, _M_′ is called
a _minor_ of _M_.

**Exercises**

10.4. (i) Let _X_ be a finite set and let _k_ be a natural number. Let I := {_Y_ ⊆ _X_ ||_Y_ |≤ _k_}.
Show that (_X,_I) is a matroid. Such matroids are called _k-uniform matroids_.
(ii) Show that _k_-uniform matroids are transversal matroids. Give an example of a
_k_-uniform matroid that is neither graphic nor cographic.

10.5. Let _M_ = (_X,_I) be a matroid and let _k_ be a natural number. Define I′ := {_Y_ ∈ I |
|_Y_ |≤ _k_}. Show that (_X,_I′) is again a matroid (called the _k-truncation_ of _M_).

10.6. Let _M_ = (_X,_I) be a matroid, let _U_ be a set disjoint from _X_, and let _k_ ≥ 0. Define

```
(17) I′ := {U′ ∪ Y | U′ ⊆ U,Y ∈I,|U′ ∪ Y |≤ k}.
```

```
Show that (U ∪ X,I′) is again a matroid.
```

10.7. Let _M_ = (_X,_I) be a matroid and let _x_ ∈ _X_.

```
(i) Show that if x is not a loop, then a subset Y of X \{x} is independent in the
contracted matroid M/{x} if and only if Y ∪{x} is independent in M.
(ii) Show that if x is a loop, then M/{x} = M \{x}.
(iii) Show that for each Y ⊆ X : rM/{x}(Y ) = rM(Y ∪{x}) − rM({x}).
```

10.8. Let _M_ = (_X,_I) be a matroid and let _Y_ ⊆ _X_.

```
(ii) Let B be a basis of Y. Show that a subset U of X \ Y is independent in the
contracted matroid M/Y if and only if U ∪ B is independent in M.
(ii) Show that for each U ⊆ X \ Y
```

```
(18) rM/Y (U) = rM(U ∪ Y ) − rM(Y ).
```

```
10.9. Let M = (X,I) be a matroid and let Y,Z ⊆ X. Show that (M \Y )/Z = (M/Z)\Y.
(That is, deletion and contraction commute.)
```

10.10. Let _M_ = (_X,_I) be a matroid, and suppose that we can test in polynomial time if
any subset _Y_ of _X_ belongs to I. Show that then the same holds for the dual matroid
_M_∗.

### 10.3. Examples of matroids

```
In this section we describe some classes of examples of matroids.
```

```
I. Graphic matroids. As a first example we consider the matroids described in
Section 10.1.
Let G = (V, E) be a graph. Let I be the collection of all forests in G. Then
M = (E,I) is a matroid, as we saw in Section 10.1.
The matroid M is called the cycle matroid of G, denoted by M(G). Any matroid
obtained in this way, or isomorphic to such a matroid, is called a graphic matroid.
Note that the bases of M(G) are exactly those forests F of G for which the graph
(V, F) has the same number of components as G. So if G is connected, the bases are
the spanning trees.
Note also that the circuits of M(G), in the matroid sense, are exactly the circuits
of G, in the graph sense.
```

```
II. Cographic matroids. There is an alternative way of obtaining a matroid from
a graph G = (V, E). It is in fact the matroid dual of the graphic matroid.
Let B be the set of subsets J of E such that E \ J is an inclusionwise maximal
forest. By Corollary 10.2a, B forms the collection of bases of a matroid. Its collection
I of independent sets consists of those subsets J of E for which
```

```
(19) κ(V, E \ J) = κ(V, E).
```

```
where, for any graph H, let κ(H) denote the number of components of H.
The matroid (E,I) is called the cocycle matroid of G, denoted by M∗(G). Any
matroid obtained in this way, or isomorphic to such a matroid, is called a cographic
matroid.
By definition, a subset C of E is a circuit of M∗(G) if it is an inclusionwise minimal
set with the property that (V, E \ C) has more components than G. Hence C is a
circuit of M∗(G) if and only if C is an inclusionwise minimal nonempty cutset in G.
```

```
III. Linear matroids. Let A be an m × n matrix. Let X = {1,... , n} and let I
be the collection of all those subsets Y of X so that the columns with index in Y are
```

linearly independent. That is, so that the submatrix of _A_ consisting of the columns
with index in _Y_ has rank |_Y_ |.
Now:

**Theorem 10.4.** (_X,_I) _is a matroid._

**Proof.** Again, conditions (2)(i) and (ii) are easy to check. To see condition (2)(iii), let
_Y_ and _Z_ be subsets of _X_ so that the columns with index in _Y_ are linearly independent,
and similarly for _Z_, and so that |_Y_ | _<_ |_Z_|.
Suppose that _Y_ ∪{_x_}6∈I for each _x_ ∈ _Z_ \ _Y_. This means that each column with
index in _Z_ \ _Y_ is spanned by the columns with index in _Y_. Trivially, each column
with index in _Z_ ∩ _Y_ is spanned by the columns with index in _Y_. Hence each column
with index in _Z_ is spanned by the columns with index in _Y_. This contradicts the fact
that the columns indexed by _Y_ span an |_Y_ |-dimensional space, while the columns
indexed by _Z_ span an |_Z_|-dimensional space, with |_Z_| _>_ |_Y_ |.

Any matroid obtained in this way, or isomorphic to such a matroid, is called a
_linear matroid_.
Note that the rank _rM_(_Y_ ) of any subset _Y_ of _X_ is equal to the rank of the matrix
formed by the columns indexed by _Y_.

**IV. Transversal matroids.** Let _X_[^1]:_,... , Xm_ be subsets of the finite set _X_. A set
_Y_ = {_y_[^1]:_,... , yn_} is called a _partial transversal_ (_of X_[^1]:_,... , Xm_), if there exist distinct

indices _i_[^1]:_,... , in_ so that _yj_ ∈ _Xij_ for _j_ = 1_,... , n_. A partial transversal of cardinality
_m_ is called a _transversal_ (or a _system of distinct representatives_, or an _SDR_).
Another way of representing partial transversals is as follows. Let G be the bipar-
tite graph with vertex set V := {[^1]:_,... , m_}∪ _X_ and with edges all pairs {_i, x_} with

_i_ ∈{[^1]:_,... , m_} and _x_ ∈ _Xi_. (We assume here that {[^1]:_,... , m_}∩ _X_ = ∅.)
For any matching _M_ in G, let _ρ_(_M_) denote the set of those elements in _X_ that
belong to some edge in _M_. Then it is not difficult to see that:

(20) _Y_ ⊆ _X_ is a partial transversal if and only if _Y_ = _ρ_(_M_) for some matching
_M_ in G.

```
Now let I be the collection of all partial transversals for X1,... , Xm. Then:
```

**Theorem 10.5.** (_X,_I) _is a matroid._

**Proof.** Again, conditions (2)(i) and (ii) are trivial. To see (2)(iii), let _Y_ and _Z_ be
partial transversals with |_Y_ | _<_ |_Z_|. Consider the graph G constructed above. By
(20) there exist matchings _M_ and _M_′ in G so that _Y_ = _ρ_(_M_) and _Z_ = _ρ_(_M_′). So
|_M_| = |_Y_ | _<_ |_Z_| = |_M_′|.
Consider the union _M_ ∪ _M_′ of _M_ and _M_′. Each component of the graph (V_, M_ ∪

```
M′) is either a path, or a circuit, or a singleton vertex. Since |M′| > |M|, at least
one of these components is a path P with more edges in M′ than in M. The path
consists of edges alternatingly in M′ and in M, with end edges in M′.
Let N and N′ denote the edges in P occurring in M and M′, respectively. So
|N′| = |N| + 1. Since P has odd length, exactly one of its end vertices belongs
to X; call this end vertex x. Then x ∈ ρ(M′) = Z and x 6∈ ρ(M) = Y. Define
M′′ := (M \ N) ∪ N′. Clearly, M′′ is a matching with ρ(M′′) = Y ∪{x}. So Y ∪{x}
belongs to I.
```

```
Any matroid obtained in this way, or isomorphic to such a matroid, is called a
transversal matroid. If the sets X1,... , Xm form a partition of X, one speaks of a
partition matroid.
```

```
These four classes of examples show that the greedy algorithm has a wider appli-
cability than just for finding minimum-weight spanning trees. There are more classes
of matroids (like ‘algebraic matroids’, ‘gammoids’), for which we refer to Welsh [1976].
```

```
Exercises
```

10.11. Show that a partition matroid is graphic, cographic, and linear.

10.12. Let _M_ = (_V,_I) be the transversal matroid derived from subsets _X_[^1]:_,...,Xm_ of _X_ as
in Example IV.

```
(i) Show with K ̋onig’s matching theorem that:
```

```
(21) rM(X) = min
J⊆{1,...,m}
```

```
(
∣∣
```

```
⋃
```

```
j∈J
```

```
Xj
∣∣ + m −|J|).
```

```
(ii) Derive a formula for rM(Y ) for any Y ⊆ X.
```

10.13. Let _G_ = (_V,E_) be a graph. Let I be the collection of those subsets _Y_ of _E_ so that _F_
has at most one circuit. Show that (_E,_I) is a matroid.

10.14. Let _G_ = (_V,E_) be a graph. Call a collection C of circuits a _circuit basis_ of _G_ if each
circuit of _G_ is a symmetric difference of circuits in C. (We consider circuits as edge
sets.)

```
Give a polynomial-time algorithm to find a circuit basis∑ C of G that minimizes
C∈C |C|.
(The running time of the algorithm should be bounded by a polynomial in |V |+|E|.)
```

10.15. Let _G_ = (_V,E_) be a connected graph. For each subset _E_′ of _E_, let _κ_(_V,E_′) denote
the number of components of the graph (_V,E_′). Show that for each _E_′ ⊆ _E_:

```
(i) rM(G)(E′) = |V |− κ(V,E′);
```

```
(ii) rM∗(G)(E′) = |E′|− κ(V,E \ E′) + 1.
```

10.16. Let _G_ be a planar graph and let _G_∗ be a planar graph dual to _G_. Show that the cycle
matroid _M_(_G_∗) of _G_∗ is isomorphic to the cocycle matroid _M_∗(_G_) of _G_.

10.17. Show that the dual matroid of a linear matroid is again a linear matroid.

10.18. Let _G_ = (_V,E_) be a loopless undirected graph. Let _A_ be the matrix obtained from
the _V_ ×_E_ incidence matrix of _G_ by replacing in each column, exactly one of the two
1’s by −1.

```
(i) Show that a subset Y of E is a forest if and only if the columns of A with index
in Y are linearly independent.
(ii) Derive that any graphic matroid is a linear matroid.
(iii) Derive (with the help of Exercise 10.17) that any cographic matroid is a linear
matroid.
```

### 10.4. Two technical lemmas

```
In this section we prove two technical lemmas as a preparation to the coming sections
on matroid intersection.
Let M = (X,I) be a matroid. For any Y ∈ I define a bipartite graph H(M, Y )
as follows. The graph H(M, Y ) has vertex set X, with colour classes Y and X \ Y.
Elements y ∈ Y and x ∈ X \ Y are adjacent if and only if
```

```
(22) (Y \{y}) ∪{x}∈I.
```

```
Then we have:
```

```
Lemma 10.1. Let M = (X,I) be a matroid and let Y, Z ∈ I with |Y | = |Z|. Then
H(M, Y ) contains a perfect matching on Y △Z.[^23]
```

```
Proof. Suppose not. By K ̋onig’s matching theorem there exist a subset S of Y \ Z
and a subset S′ of Z \ Y such that for each edge {y, z} of H(M, Y ) satisfying z ∈ S′
one has y ∈ S and such that |S| < |S′|.
As |(Y ∩ Z) ∪ S| < |(Y ∩ Z) ∪ S′|, there exists an element z ∈ S′ such that
T := (Y ∩ Z) ∪ S ∪{z} belongs to I. This implies that there exists an U ∈ I such
that T ⊆ U ⊆ T ∪ Y and |U| = |Y |. So U = (Y \{x}) ∪{z} for some x 6∈ S. As
{x, z} is an edge of H(M, Y ) this contradicts the choice of S and S′.
```

```
The following forms a counterpart:
```

```
23A perfect matching on a vertex set U is a matching M with
```

```
⋃
M = U.
```

```
Lemma 10.2. Let M = (X,I) be a matroid and let Y ∈I. Let Z ⊆ X be such that
|Y | = |Z| and such that H(M, Y ) contains a unique perfect matching N on Y △Z.
Then Z belongs to I.
```

```
Proof. By induction on k := |Z \ Y |, the case k = 0 being trivial. Let k ≥ 1.
By the unicity of N there exists an edge {y, z}∈ N, with y ∈ Y \Z and z ∈ Z \Y ,
with the property that there is no z′ ∈ Z \ Y such that z′ 6= z and {y, z′} is an edge
of H(M, Y ).
Let Z′ := (Z \{z})∪{y} and N′ := N \{{y, z}}. Then N′ is the unique matching
in H(M, Y ) with union Y △Z′. Hence by induction, Z′ belongs to I.
There exists an S ∈ I such that Z′ \{y} ⊆ S ⊆ (Y \{y}) ∪ Z and |S| = |Y |
(since (Y \{y}) ∪ Z = (Y \{y}) ∪{z}∪ Z′ and since (Y \{y}) ∪{z} belongs to I).
Assuming Z 6∈ I, we know z 6∈ S and hence r((Y ∪ Z′) \{y}) = |Y |. Hence there
exists an z′ ∈ Z′ \ Y such that (Y \{y}) ∪{z′} belongs to I. This contradicts the
choice of y.
```

```
Exercises
```

10.19. Let _M_ = (_X,_I) be a matroid, let _B_ be a basis of _M_, and let _w_ : _X_ → R be a weight
function. Show that _B_ is a basis of maximum weight if and only if _w_(_B_′) ≤ _w_(_B_) for
every basis _B_′ with |_B_′ \ _B_| = 1.

10.20. Let _M_ = (_X,_I) be a matroid and let _Y_ and _Z_ be independent sets with |_Y_ | = |_Z_|.
For any _y_ ∈ _Y_ \ _Z_ define _δ_(_y_) as the set of those _z_ ∈ _Z_ \ _Y_ which are adjacent to _y_
in the graph _H_(_M,Y_ ).

```
(i) Show that for each y ∈ Y \ Z the set (Z \ δ(y)) ∪{y} belongs to I.
(Hint: Apply inequality (10) to X′ := (Z \ δ(y)) ∪{y} and X′′ := (Z \ δ(y)) ∪
(Y \{y}).)
(ii) Derive from (i) that for each y ∈ Y \ Z there exists an z ∈ Z \ Y so that {y,z}
is an edge both of H(M,Y ) and of H(M,Z).
```

### 10.5. Matroid intersection

```
Edmonds [1970] discovered that the concept of matroid has even more algorithmic
power, by showing that there exist fast algorithms also for intersections of matroids.
Let M1 = (X,I1) and M2 = (X,I2) be two matroids, on the same set X. Consider
the collection I1 ∩I2 of common independent sets. The pair (X,I1 ∩I2) is generally
not a matroid again (cf. Exercise 10.21).
What Edmonds showed is that, for any weight function w on X, a maximum-
weight common independent set can be found in polynomial time. In particular, a
common independent set of maximum cardinality can be found in polynomial time.
We consider first some applications.
```

**Example 10.5a.** Let _G_ = (_V, E_) be a bipartite graph, with colour classes _V_[^1]: and
_V_[^2]:, say. Let I[^1]: be the collection of all subsets _F_ of _E_ so that no two edges in _F_ have
a vertex in _V_[^1]: in common. Similarly, let I[^2]: be the collection of all subsets _F_ of _E_ so
that no two edges in _F_ have a vertex in _V_[^2]: in common. So both (_X,_I[^1]:) and (_X,_I[^2]:)
are partition matroids.

Now I[^1]: ∩ I[^2]: is the collection of matchings in _G_. Finding a maximum-weight
common independent set amounts to finding a maximum-weight matching in _G_.

**Example 10.5b.** Let _X_[^1]:_,... , Xm_ and _Y_[^1]:_,... , Ym_ be subsets of _X_. Let _M_[^1]: = (_X,_I[^1]:)
and _M_[^2]: = (_X,_I[^2]:) be the corresponding transversal matroids.
Then common independent sets correspond to common partial transversals. The
collections (_X_[^1]:_,... , Xm_) and (_Y_[^1]:_,... , Ym_) have a common transversal if and only if
the maximum cardinality of a common independent set is equal to _m_.

**Example 10.5c.** Let _D_ = (_V, A_) be a directed graph. Let _M_[^1]: = (_A,_I[^1]:) be the cycle
matroid of the underlying undirected graph. Let I[^2]: be the collection of subsets _Y_ of
_A_ so that each vertex of _D_ is entered by at most one arc in _Y_. So _M_[^2]: := (_A,_I[^2]:) is a
partition matroid.
Now the common independent sets are those subsets _Y_ of _A_ with the property
that each component of (_V, Y_ ) is a rooted tree. Moreover, _D_ has a rooted spanning
tree if and only if the maximum cardinality of a set in I[^1]: ∩I[^2]: is equal to |_V_ |− 1.

**Example 10.5d.** Let _G_ = (_V, E_) be a connected undirected graph. Then _G_ has
two edge-disjoint spanning trees if and only if the maximum cardinality of a common
independent set in the cycle matroid _M_(_G_) of _G_ and the cocycle matroid _M_∗(_G_) of
_G_ is equal to |_V_ |− 1.

In this section we describe an algorithm for finding a maximum-cardinality com-
mon independent sets in two given matroids. In the next section we consider the
more general maximum-weight problem.
For any two matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) and any _Y_ ∈ I[^1]: ∩I[^2]:, we
define a directed graph _H_(_M_[^1]:_, M_[^2]:_, Y_ ) as follows. Its vertex set is _X_, while for any
_y_ ∈ _Y, x_ ∈ _X_ \ _Y_ ,

(23) (_y, x_) is an arc of _H_(_M_[^1]:_, M_[^2]:_, Y_ ) if and only if (_Y_ \{_y_}) ∪{_x_}∈I[^1]:,
(_x, y_) is an arc of _H_(_M_[^1]:_, M_[^2]:_, Y_ ) if and only if (_Y_ \{_y_}) ∪{_x_}∈I[^2]:.

These are all arcs of _H_(_M_[^1]:_, M_[^2]:_, Y_ ). In fact, this graph can be considered as the union
of directed versions of the graphs _H_(_M_[^1]:_, Y_ ) and _H_(_M_[^2]:_, Y_ ) defined in Section 10.4.
The following is the basis for finding a maximum-cardinality common independent
set in two matroids.

**Cardinality common independent set augmenting algorithm**

**input:** matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) and a set _Y_ ∈I[^1]: ∩I[^2]:;
**output:** a set _Y_ ′ ∈I[^1]: ∩I[^2]: with |_Y_ ′| _>_ |_Y_ |, if it exists.
**description of the algorithm:** We assume that _M_[^1]: and _M_[^2]: are given in such a way
that for any subset _Z_ of _X_ we can check in polynomial time if _Z_ ∈I[^1]: and if _Z_ ∈I[^2]:.
Consider the sets

(24) _X_[^1]: := {_y_ ∈ _X_ \ _Y_ | _Y_ ∪{_y_}∈I[^1]:},
_X_[^2]: := {_y_ ∈ _X_ \ _Y_ | _Y_ ∪{_y_}∈I[^2]:}.

Moreover, consider the directed graph _H_(_M_[^1]:_, M_[^2]:_, Y_ ) defined above. There are two
cases.

**Case 1.** _There exists a directed path P in H_(_M_[^1]:_, M_[^2]:_, Y_ ) _from some vertex in X_[^1]: _to_
_some vertex in X_[^2]:_._ (Possibly of length 0 if _X_[^1]: ∩ _X_[^2]: [^6]:= ∅.)
We take a shortest such path _P_ (that is, with a minimum number of arcs). Let _P_
traverse the vertices _y_[^0]:_, z_[^1]:_, y_[^1]:_,... , zm, ym_ of _H_(_M_[^1]:_, M_[^2]:_, Y_ ), in this order. By construc-
tion of the graph _H_(_M_[^1]:_, M_[^2]:_, Y_ ) and the sets _X_[^1]: and _X_[^2]:, this implies that _y_[^0]:_,... , ym_
belong to _X_ \ _Y_ and _z_[^1]:_,... , zm_ belong to _Y_.
Now output

(25) _Y_ ′ := (_Y_ \{_z_[^1]:_,... , zm_}) ∪{_y_[^0]:_,... , ym_}.

**Case 2.** _There is no directed path in H_(_M_[^1]:_, M_[^2]:_, Y_ ) _from any vertex in X_[^1]: _to any_
_vertex vertex in X_[^2]:_._ Then _Y_ is a maximum-cardinality common independent set.

This finishes the description of the algorithm. The correctness of the algorithm is
given in the following two theorems.

**Theorem 10.6.** _If Case 1 applies, then Y_ ′ ∈I[^1]: ∩I[^2]:_._

**Proof.** Assume that Case 1 applies. By symmetry it suffices to show that _Y_ ′ belongs
to I[^1]:.
To see that _Y_ ′\{_y_[^0]:} belongs to I[^1]:, consider the graph _H_(_M_[^1]:_, Y_ ) defined in Section
10.4. Observe that the edges {_zj, yj_} form the only matching in _H_(_M_[^1]:_, Y_ ) with union
equal to {_z_[^1]:_,... , zm, y_[^1]:_,... , ym_} (otherwise _P_ would have a shortcut). So by Lemma
10.2, _Y_ ′ \{_y_[^0]:} = (_Y_ \{_z_[^1]:_,... , zm_}) ∪{_y_[^1]:_,... , ym_} belongs to I[^1]:.
To show that _Y_ ′ belongs to I[^1]:, observe that _rM_[^1]:(_Y_ ∪_Y_ ′) ≥ _rM_[^1]:(_Y_ ∪{_y_[^0]:}) = |_Y_ |+1,
and that, as (_Y_ ′ \{_y_[^0]:}) ∩ _X_[^1]: = ∅, _rM_[^1]:((_Y_ ∪ _Y_ ′) \{_y_[^0]:}) = |_Y_ |. As _Y_ ′ \{_y_[^0]:}∈I[^1]:, we
know _Y_ ′ ∈I[^1]:.

**Theorem 10.7.** _If Case 2 applies, then Y is a maximum-cardinality common inde-_

_pendent set._

**Proof.** As Case 2 applies, there is no directed _X_[^1]: −_X_[^2]: path in _H_(_M_[^1]:_, M_[^2]:_, Y_ ). Hence
there is a subset _U_ of _X_ containing _X_[^2]: such that _U_ ∩ _X_[^1]: = ∅ and such that no arc of
_H_(_M_[^1]:_, M_[^2]:_, Y_ ) enters _U_. (We can take for _U_ the set of vertices that are not reachable
by a directed path from _X_[^1]:.)
We show

(26) _rM_[^1]:(_U_) + _rM_[^2]:(_X_ \ _U_) = |_Y_ |_._

```
To this end, we first show
```

(27) _rM_[^1]:(_U_) = |_Y_ ∩ _U_|_._

Clearly, as _Y_ ∩_U_ ∈I[^1]:, we know _rM_[^1]:(_U_) ≥|_Y_ ∩_U_|. Suppose _rM_[^1]:(_U_) _>_ |_Y_ ∩_U_|. Then
there exists an _x_ in _U_ \_Y_ so that (_Y_ ∩_U_)∪{_x_}∈I[^1]:. Since _Y_ ∈I[^1]:, this implies that
there exists a set _Z_ ∈ I[^1]: with |_Z_| ≥ |_Y_ | and (_Y_ ∩ _U_) ∪{_x_} ⊆ _Z_ ⊆ _Y_ ∪{_x_}. Then
_Z_ = _Y_ ∪{_x_} or _Z_ = (_Y_ \{_y_}) ∪{_x_} for some _y_ ∈ _Y_ \ _U_.
In the first alternative, _x_ ∈ _X_[^1]:, contradicting the fact that _x_ belongs to _U_. In the
second alternative, (_y, x_) is an arc of _H_(_M_[^1]:_, M_[^2]:_, Y_ ) entering _U_. This contradicts the
definition of _U_ (as _y_ 6∈ _U_ and _x_ ∈ _U_).
This shows (27). Similarly we have that _rM_[^2]:(_X_ \ _U_) = |_Y_ \ _U_|. Hence we have
(26).
Now (26) implies that for any set _Z_ in I[^1]: ∩I[^2]: one has

(28) |_Z_| = |_Z_ ∩ _U_| + |_Z_ \ _U_|≤ _rM_[^1]:(_U_) + _rM_[^2]:(_X_ \ _U_) = |_Y_ |_._

So _Y_ is a common independent set of maximum cardinality.

The algorithm clearly has polynomially bounded running time, since we can con-
struct the auxiliary directed graph _H_(_M_[^1]:_, M_[^2]:_, Y_ ) and find the path _P_ (if it exists),
in polynomial time.
It implies the result of Edmonds [1970]:

**Theorem 10.8.** _A maximum-cardinality common independent set in two matroids_
_can be found in polynomial time._

**Proof.** Directly from the above, as we can find a maximum-cardinality common inde-
pendent set after applying at most |_X_| times the common independent set augmenting
algorithm.

```
The algorithm also yields a min-max relation for the maximum cardinality of a
```

```
common independent set, as was shown again by Edmonds [1970].
```

```
Theorem 10.9 (Edmonds’ matroid intersection theorem). Let M1 = (X,I1) and
M2 = (X,I2) be matroids. Then
```

```
(29) max
Y∈I1∩I2
```

```
|Y | = min
U⊆X(rM[^1](U) + rM[^2](X \ U)).
```

```
Proof. The inequality ≤ follows similarly as in (28). The reverse inequality follows
from the fact that if the algorithm stops with set Y , we obtain a set U for which (26)
holds. Therefore, the maximum in (29) is at least as large as the minimum.
```

```
Exercises
```

10.21. Give an example of two matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) so that (_X,_I[^1]:∩I[^2]:)
is _not_ a matroid.

10.22. Derive K ̋onig’s matching theorem from Edmonds’ matroid intersection theorem.

10.23. Let (_X_[^1]:_,...,Xm_) and (_Y_[^1]:_,...,Ym_) be subsets of the finite set _X_. Derive from Ed-
monds’ matroid intersection theorem: (_X_[^1]:_,...,Xm_) and (_Y_[^1]:_,...,Ym_) have a common
transversal if and only if

```
(30)
∣∣(
```

```
⋃
```

```
i∈I
```

```
Xi) ∩ (
```

```
⋃
```

```
j∈J
```

```
Yj)
∣∣ ≥|I| + |J|− m
```

```
for all subsets I and J of {1,...,m}.
```

10.24. Reduce the problem of finding a Hamiltonian cycle in a directed graph to the problem
of finding a maximum-cardinality common independent set in _three_ matroids.

10.25. Let _G_ = (_V,E_) be a graph and let the edges of _G_ be coloured with |_V_ |− 1 colours.
That is, we have partitioned _E_ into classes _X_[^1]:_,...,X_|_V_|−[^1]:, called _colours_. Show that
there exists a spanning tree with all edges coloured differently if and only if (_V,E_′)
has at most |_V_ |− _t_ components, for any union _E_′ of _t_ colours, for any _t_ ≥ 0.

10.26. Let _M_ = (_X,_I) be a matroid and let _X_[^1]:_,...,Xm_ be subsets of _X_. Then (_X_[^1]:_,...,Xm_)
has an independent transversal if and only if the rank of the union of any _t_ sets among
_X_[^1]:_,...,Xm_ is at least _t_, for any _t_ ≥ 0. (Rado [1942].)

10.27. Let _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) be matroids. Define

```
(31) I1 ∨I2 := {Y1 ∪ Y2 | Y1 ∈I1,Y2 ∈I2}.
```

```
(i) Show that the maximum cardinality of a set in I1 ∨I2 is equal to
```

```
(32) min
U⊆X(rM[^1](U) + rM[^2](U) + |X \ U|).
```

```
(Hint: Apply the matroid intersection theorem to M1 and M∗
2.)
(ii) Derive that for each Y ⊆ X:
(33) max{|Z|| Z ⊆ Y,Z ∈I1 ∨I2} =
min
U⊆Y(rM[^1](U) + rM[^2](U) + |Y \ U|).
(iii) Derive that (X,I1 ∨I2) is again a matroid.
(Hint: Use axiom (vi) in Theorem 10.2.)
This matroid is called the union of M1 and M2, denoted by M1∨M2. (Edmonds
and Fulkerson [1965], Nash-Williams [1967].)
(iv) Let M1 = (X,I1),...,Mk = (X,Ik) be matroids and let
```

```
(34) I1 ∨ ... ∨Ik := {Y1 ∪ ... ∪ Yk | Y1 ∈I1,...,Yk ∈Ik}.
```

```
Derive from (iii) that M1 ∨...∨Mk := (X,I1 ∨...∨Ik) is again a matroid and
give a formula for its rank function.
```

10.28. (i) Let _M_ = (_X,_I) be a matroid and let _k_ ≥ 0. Show that _X_ can be covered by _k_
independent sets if and only if |_U_|≤ _k_ · _rM_(_U_) for each subset _U_ of _X_.
(_Hint:_ Use Exercise 10.27.) (Edmonds [1965b].)

(ii) Show that the problem of finding a minimum number of independent sets cov-
ering _X_ in a given matroid _M_ = (_X,_I), is solvable in polynomial time.
10.29. Let _G_ = (_V,E_) be a graph and let _k_ ≥ 0. Show that _E_ can be partitioned into _k_
forests if and only if each nonempty subset _W_ of _V_ contains at most _k_(|_W_|−1) edges
of _G_.
(_Hint:_ Use Exercise 10.28.) (Nash-Williams [1964].)

10.30. Let _X_[^1]:_,...,Xm_ be subsets of _X_ and let _k_ ≥ 0.

```
(i) Show that X can be partitioned into k partial transversals of (X1,...,Xm) if
and only if
```

```
(35) k(m −|I|) ≥
∣∣X \
```

```
⋃
```

```
i∈I
```

```
Xi
∣∣
```

```
for each subset I of {1,...,m}.
(ii) Derive from (i) that {1,...,m} can be partitioned into classes I1,...,Ik so that
(Xi | i ∈ Ij) has a transversal, for each j = 1,...,k if and only if Y contains at
most k|Y | of the Xi as a subset, for each Y ⊆ X.
(Hint: Interchange the roles of {1,...,m} and X.) (Edmonds and Fulkerson
[1965].)
```

10.31. (i) Let _M_ = (_X,_I) be a matroid and let _k_ ≥ 0. Show that there exist _k_ pairwise
disjoint bases of _M_ if and only if _k_(_rM_(_X_) − _rM_(_U_)) ≥|_X_ \ _U_| for each subset
_U_ of _X_.

```
(Hint: Use Exercise 10.27.) (Edmonds [1965b].)
(ii) Show that the problem of finding a maximum number of pairwise disjoint bases
in a given matroid, is solvable in polynomial time.
```

10.32. Let _G_ = (_V,E_) be a connected graph and let _k_ ≥ 0. Show that there exist _k_ pairwise
edge-disjoint spanning trees if and only if for each _t_, for each partition (_V_[^1]:_,...,Vt_) of
_V_ into _t_ classes, there are at least _k_(_t_ − 1) edges connecting different classes of this
partition.
(_Hint:_ Use Exercise 10.31.) (Nash-Williams [1961], Tutte [1961].)

10.33. Let _M_[^1]: and _M_[^2]: be matroids so that, for _i_ = 1_,_2, we can test in polynomial time if a
given set is independent in _Mi_. Show that the same holds for the union _M_[^1]: ∨ _M_[^2]:.

10.34. Let _M_ = (_X,_I) be a matroid and let _B_ and _B_′ be two disjoint bases. Let _B_ be
partitioned into sets _Y_[^1]: and _Y_[^2]:. Show that there exists a partition of _B_′ into sets _Z_1
and _Z_[^2]: so that both _Y_[^1]: ∪ _Z_[^1]: ∪ _Z_[^2]: and _Z_[^1]: ∪ _Y_[^2]: are bases of _M_.
(_Hint:_ Assume without loss of generality that _X_ = _B_ ∪ _B_′. Apply the matroid
intersection theorem to the matroids (_M_ \ _Y_[^1]:)_/Y_[^2]: and (_M_∗ \ _Y_[^1]:)_/Y_[^2]:.)

10.35. The following is a special case of a theorem of Nash-Williams [1985]:
Let _G_ = (_V,E_) be a simple, connected graph and let _b_ : _V_ → Z+. Call a graph
_G_ ̃ = (_V ,_ ̃ _E_ ̃) a _b-detachment_ of _G_ if there is a function _φ_ : _V_ ̃ → _V_ such that |_φ_−[^1](_v_)| =
_b_(_v_) for each _v_ ∈ _V_ , and such that there is a one-to-one function _ψ_ : _E_ ̃ → _E_ with
_ψ_(_e_) = {_φ_(_v_)_,φ_(_w_)} for each edge _e_ = {_v,w_} of _G_ ̃.
Then there exists a connected _b_-detachment if and only if for each _U_ ⊆ _V_ the number
of components of the graph induced by _V_ \ _U_ is at most |_EU_|− _b_(_U_) + 1.
Here _EU_ denotes the set of edges intersecting _U_.

### 10.6. Weighted matroid intersection

```
We next consider the problem of finding a maximum-weight common independent
set, in two given matroids, with a given weight function. The algorithm, again due
to Edmonds [1970], is an extension of the algorithm given in Section 10.5. In each
iteration, instead of finding a path P with a minimum number of arcs in H, we will
now require P to have minimum length with respect to some length function defined
on H.
To describe the algorithm, if matroid M1 = (S,I1) and M2 = (S,I2) and a weight
function w : S → R are given, call a set Y ∈I1 ∩I2 extreme if w(Z) ≤ w(Y ) for each
Z ∈I1 ∩I2 satisfying |Z| = |Y |.
```

**Weighted common independent set augmenting algorithm**

**input:** matroids _M_[^1]: = (_S,_I[^1]:) and _M_[^2]: = (_S,_I[^2]:), a weight function _w_ : _S_ → Q, and
an extreme common independent set _Y_ ;
**output:** an extreme common independent set _Y_ ′ with |_Y_ ′| = |_Y_ | + 1, if it exists

**description of the algorithm:** Consider again the sets _X_[^1]: and _X_[^2]: and the directed
graph _H_(_M_[^1]:_, M_[^2]:_, Y_ ) on _S_ as in the cardinality case.
For any _x_ ∈ _S_ define the ‘length’ _l_(_x_) of _x_ by:

(36) _l_(_x_) := _w_(_x_) if _x_ ∈ _Y_ ,
_l_(_x_) := −_w_(_x_) if _x_ 6∈ _Y._

The _length_ of a path _P_, denoted by _l_(_P_), is equal to the sum of the lengths of the
vertices traversed by _P_, counting multiplicities.
We consider two cases.

**Case 1.** _H_(_M_[^1]:_, M_[^2]:_, Y_ ) _has an X_[^1]: −_X_[^2]: _path P._ We choose _P_ so that _l_(_P_) is minimal
and so that it has a minimum number of arcs among all minimum-length _X_[^1]: − _X_2
paths. Set _Y_ ′ := _Y_ △_V P_.

**Case 2.** _H_(_M_[^1]:_, M_[^2]:_, Y_ ) _has no X_[^1]: − _X_[^2]: _path._ Then _Y_ is a maximum-size common
independent set.

This finishes the description of the algorithm. The correctness of the algorithm if
Case 2 applies follows directly from Theorem 10.7. In order to show the correctness
if Case 1 applies, we first prove the following basic property of the length function _l_.

**Theorem 10.10.** _Let C be a directed circuit in H_(_M_[^1]:_, M_[^2]:_, Y_ )_, and let t_ ∈ _V C._
_Define Z_ := _Y_ △_V C. If Z_ 6∈ I[^1]: ∩ I[^2]: _then there exists a directed cycle C_′ _with_
_V C_′ ⊂ _V C such that l_(_C_′) _<_ [^0]:_, or l_(_C_′) ≤ _l_(_C_) _and t_ ∈ _V C_′_._

**Proof.** By symmetry we can assume that _Z_ 6∈ I[^1]:. Let _N_[^1]: and _N_[^2]: be the sets of
arcs in _C_ belonging to _H_(_M_[^1]:_, Y_ ) and _H_(_M_[^2]:_, Y_ ) respectively. If _Z_ 6∈ I[^1]:, there is, by
Lemma 10.2 a matching _N_′
[^1]: in _H_(_M_[^1]_, Y_ ) on _V C_ with _N_′
[^1]: [^6]:= _N_[^1]. Consider the directed
graph _D_ = (_V C, A_) formed by the arcs in _N_[^1]:, _N_′
[^1]: (taking arcs in _N_[^1] ∩ _N_′
[^1]: multiple),
and by the arcs in _N_[^2]: taking each of them twice (parallel). Now each vertex in _V C_
is entered and left by exactly two arcs of _D_. Moreover, since _N_′
[^1]: [^6]:= _N_[^1], _D_ contains
a directed circuit _C_[^1]: with _V C_[^1]: ⊂ _V C_. We can extend this to a decomposition of _A_
into directed circuits _C_[^1]:_,... , Ck_. Then

(37) _χV C_[^1] + ··· + _χV Ck_ = 2 · _χV C_.

Since _V C_[^1]: [^6]:= _V C_ we know that _V Cj_ = _V C_ for at most one _j_. If, say _V Ck_ = _V C_,

then (37) implies that either _l_(_V Cj_) _<_ 0 for some _j < k_ or _l_(_V Cj_) ≤ _l_(_V C_) for all
_j < k_, implying the proposition.
If _V Cj_ [^6]:= _V C_ for all _j_, then _l_(_V Cj_) _<_ 0 for some _j_ ≤ _k_ or _l_(_V Cj_) ≤ _l_(_V C_) for all
_j_ ≤ _k_, again implying the proposition.

```
This implies:
```

**Theorem 10.11.** _Let Y_ ∈I[^1]: ∩I[^2]:_. Then Y is extreme if and only if H_(_M_[^1]:_, M_[^2]:_, Y_ )
_has no directed cycle of negative length._

**Proof.** To see necessity, suppose _H_(_M_[^1]:_, M_[^2]:_, Y_ ) has a cycle _C_ of negative length.
Choose _C_ with |_V C_| minimal. Consider _Z_ := _Y_ △_V C_. Since _w_(_Z_) = _w_(_Y_ ) − _l_(_C_) _>_
_w_(_Y_ ), while |_Z_| = |_Y_ |, we know that _Z_ 6∈ I[^1]: ∩I[^2]:. Hence by Proposition 10.10,
_H_(_M_[^1]:_, M_[^2]:_, Y_ ) has a negative-length directed cycle covering fewer than |_V C_| vertices,
contradicting our assumption.
To see sufficiency, consider a _Z_ ∈ I[^1]: ∩I[^2]: with |_Z_| = |_Y_ |. By Lemma 10.1, both
_H_(_M_[^1]:_, Y_ ) and _H_(_M_[^2]:_, Y_ ) have a perfect matching on _Y_ △_Z_. These two matchings
together form a disjoint union of a number of directed cycles _C_[^1]:_,... , Ct_. Then

(38) _w_(_Y_ ) − _w_(_Z_) =

```
∑t
```

```
j=1
```

```
l(Cj) ≥ 0,
```

implying _w_(_Z_) ≤ _w_(_Y_ ). So _Y_ is extreme.

This theorem implies that we can find in the algorithm a shortest path _P_ in
polynomial time (with the Bellman-Ford method).
This also gives:

**Theorem 10.12.** _If Case 1 applies, Y_ ′ _is an extreme common independent set._

**Proof.** We first show that _Y_ ′ ∈ I[^1]: ∩I[^2]:. To this end, let _t_ be a new element, and
extend (for each _i_ = 1_,_2), _Mi_ to a matroid _M_′
_i_ = (_S_ +_t,_I′
_i_), where for each _T_ ⊆ _S_ +_t_:

(39) _T_ ∈I′
_i_ if and only if _T_ − _t_ ∈I_i_.

Note that _H_(_M_′
[^1]:_, M_′
[^2]:_, Y_ + _t_) arises from _H_(_M_[^1]_, M_[^2]_, Y_ ) by extending it with a new
vertex _t_ and adding arcs from each vertex in _X_[^1]: to _t_, and from _t_ to each vertex in
_X_[^2]:.
Let _P_ be the path found in the algorithm. Define

(40) _w_(_t_) := _l_(_t_) := −_l_(_P_)_._

```
As P is a shortest X1 − X2 path, this makes that H(M′
1, M′
2, Y + t) has no negative-
length directed cycle. Hence, by Theorem 10.11, Y + t is an extreme common inde-
pendent set in M′
1 and M′
2.
Let P run from z1 ∈ X1 to z2 ∈ X2. Extend P by the arcs (t, z1) and (z2, t) to a
directed cycle C. So Z = (Y + t)△V C. As P has a minimum number of arcs among
all shortest X1 − X2 paths, and as H(M′
1, M′
2, Y + t) has no negative-length directed
circuits, by Proposition 10.10 we know that Z ∈I1 ∩I2.
Moreover, Z is extreme, since Y + t is extreme and w(Z) = w(Y + t).
```

```
So the weighted common independent set augmenting algorithm is correct. It
obviously has polynomially bounded running time. Therefore:
```

```
Theorem 10.13. A maximum-weight common independent set in two matroids can
be found in polynomial time.
```

```
Proof. Starting with the extreme common independent set Y0 := ∅ we can find
iteratively extreme common independent sets Y0, Y1,... , Yk, where |Yi| = i for i =
0,... , k and where Yk is a maximum-size common independent set. Taking one among
Y0,... , Yk of maximum weight, we have an extreme common independent set.
```

```
Exercises
```

10.36. Give an example of two matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) and a weight
function _w_ : _X_ → Z+ so that there is no maximum-weight common independent set
which is also a maximum-cardinality common independent set.

10.37. Reduce the problem of finding a maximum-weight common basis in two matroids to
the problem of finding a maximum-weight common independent set.

10.38. Let _D_ = (_V,A_) be a directed graph, let _r_ ∈ _V_ , and let _l_ : _A_ → Z+ be a length
function. Reduce the problem of finding a minimum-length rooted tree with root
_r_, to the problem of finding a maximum-weight common independent set in two
matroids.

10.39. Let _B_ be a common basis of the matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) and let
_w_ : _X_ → Z be a weight function. Define length function _l_ : _X_ → Z by _l_(_x_) := _w_(_x_) if
_x_ ∈ _B_ and _l_(_x_) := −_w_(_x_) if _x_ 6∈ _B_.

```
Show that B has maximum-weight among all common bases of M1 and M2 if and
only if H(M1,M2,B) has no directed circuit of negative length.
```

### 10.7. Matroids and polyhedra

The algorithmic results obtained in the previous sections have interesting conse-
quences for polyhedra associated with matroids.
Let _M_ = (_X,_I) be a matroid. The _matroid polytope P_(_M_) of _M_ is, by definition,
the convex hull of the incidence vectors of the independent sets of _M_. So _P_(_M_) is a
polytope in R_X_.
Each vector _z_ in _P_(_M_) satisfies the following linear inequalities:

(41) _z_(_x_) ≥[^0] for _x_ ∈ _X_,
_z_(_Y_ ) ≤ _rM_(_Y_ ) for _Y_ ⊆ _X_.

This follows from the fact that the incidence vector _χY_ of any independent set _Y_ of
_M_ satisfies (41).
Note that if _z_ is an integer vector satisfying (41), then _z_ is the incidence vector of
some independent set of _M_.
Edmonds [1970] showed that system (41) in fact fully determines the matroid
polytope _P_(_M_). It means that for each weight function _w_ : _X_ → R, the linear
programming problem

(42) maximize _wTz,_
subject to _z_(_x_) ≥ [^0]: (_x_ ∈ _X_)
_z_(_Y_ ) ≤ _rM_(_Y_ ) (_Y_ ⊆ _X_)

has an integer optimum solution _z_. This optimum solution necessarily is the incidence
vector of some independent set of _M_. In order to prove this, we also consider the
LP-problem dual to (42):

(43) minimize

```
∑
```

```
Y⊆X
```

```
yY rM(Y ),
```

```
subject to ∑ yY ≥ 0 (Y ⊆ X)
```

```
Y⊆X,x∈Y
```

```
yY ≥ w(x) (x ∈ X).
```

We show:

**Theorem 10.14.** _If w is integer, then_ (42) _and_ (43) _have integer optimum solutions._

**Proof.** Order the elements of _X_ as _y_[^1]:_,... , ym_ in such a way that _w_(_y_[^1]:) ≥ _w_(_y_[^2]:) ≥

_... w_(_ym_). Let _n_ be the largest index for which _w_(_yn_) ≥ 0. Define _Xi_ := {_y_[^1]:_,... , yi_}
for _i_ = 0_,... , m_ and

(44) _Y_ := {_yi_ | _i_ ≤ _n_;_rM_(_Xi_) _> rM_(_Xi_−[^1]:)}_._

Then _Y_ belongs to I (cf. Exercise 10.3). So _z_ := _χY_ is an integer feasible solution of
(42).
Moreover, define a vector _y_ in RP(_X_) by:

(45)

```
yY := w(yi) − w(yi+1) if Y = Xi for some i = 1,... , n − 1,
yY := w(yn) if Y = Xn,
yY := 0 for all other Y ⊆ X
```

Then _y_ is an integer feasible solution of (43).
We show that _z_ and _y_ have the same objective value, thus proving the theorem:

```
wTz = w(Y ) =
```

```
∑
```

```
x∈Y
```

```
w(x) =
```

```
∑n
```

```
i=1
```

```
(46) w(yi) · (rM(Xi) − rM(Xi−1))
```

```
= w(yn) · rM(Xn) +
```

```
∑n−[^1]
```

```
i=1
```

```
(w(yi) − w(yi+1)) · rM(Xi) =
```

```
∑
```

```
Y⊆X
```

```
yY rM(Y ).
```

```
So system (41) is totally dual integral. This directly implies:
```

**Corollary 10.14a.** _The matroid polytope P_(_M_) _is determined by_ (41)_._

**Proof.** Immediately from Theorem 10.14.

An even stronger phenomenon occurs at intersections of matroid polytopes. It
turns out that the intersection of two matroid polytopes gives exactly the convex hull
of the common independent sets, as was shown again by Edmonds [1970].
To see this, we first derive a basic property:

**Theorem 10.15.** _Let M_[^1]: = (_X,_I[^1]:) _and M_[^2]: = (_X,_I[^2]:) _be matroids, let w_ : _X_ → Z
_be a weight function and let B be a common basis of maximum weight w_(_B_)_. Then_
_there exist functions w_[^1]:_, w_[^2]: : _X_ → Z _so that w_ = _w_[^1]: + _w_[^2]:_, and so that B is both a_
_maximum-weight basis of M_[^1]: _with respect to w_[^1]: _and a maximum-weight basis of M_2
_with respect to w_[^2]:_._

**Proof.** Consider the directed graph _H_(_M_[^1]:_, M_[^2]:_, B_) with length function _l_ as defined
in Exercise 10.39. Since _B_ is a maximum-weight basis, _H_(_M_[^1]:_, M_[^2]:_, B_) has no directed
circuits of negative length. Hence there exists a function _φ_ : _X_ → Z so that _φ_(_y_) −
_φ_(_x_) ≤ _l_(_y_) for each arc (_x, y_) of _H_(_M_[^1]:_, M_[^2]:_, B_). Using the definition of _H_(_M_[^1]:_, M_[^2]:_, B_)
and _l_, this implies that for _y_ ∈ _B, x_ ∈ _X_ \ _B_:

(47) _φ_(_x_) − _φ_(_y_) ≤ −_w_(_x_) if (_B_ \{_y_}) ∪{_x_}∈I[^1],
_φ_(_y_) − _φ_(_x_) ≤ _w_(_x_) if (_B_ \{_y_}) ∪{_x_}∈I[^2]:.

Now define

(48) _w_[^1](_y_) := _φ_(_y_)_, w_[^2](_y_) := _w_(_y_) − _φ_(_y_) for _y_ ∈ _B_
_w_[^1]:(_x_) := _w_(_x_) + _φ_(_x_)_, w_[^2]:(_x_) := −_φ_(_x_) for _x_ ∈ _X_ \ _B_.

Then _w_[^1]:(_x_) ≤ _w_[^1]:(_y_) whenever (_B_ \{_y_}) ∪{_x_} ∈ I[^1]:. So by Exercise 10.19, _B_ is a
maximum-weight basis of _M_[^1]: with respect to _w_[^1]:. Similarly, _B_ is a maximum-weight
basis of _M_[^2]: with respect to _w_[^2]:.

Note that if _B_ is a maximum-weight basis of _M_[^1]: with respect to some weight
function _w_, then also after adding a constant function to _w_ this remains the case.
This observation will be used in showing that a theorem similar to Theorem 10.15
holds for independent sets instead of bases.

**Theorem 10.16.** _Let M_[^1]: = (_X,_I[^1]:) _and M_[^2]: = (_X,_I[^2]:) _be matroids, let w_ : _X_ → Z
_be a weight function, and let Y be a maximum-weight common independent set. Then_
_there exist weight functions w_[^1]:_, w_[^2]: : _X_ → Z _so that w_ = _w_[^1]:+_w_[^2]: _and so that Y is both_
_a maximum-weight independent set of M_[^1]: _with respect to w_[^1]: _and a maximum-weight_
_independent set of M_[^2]: _with respect to w_[^2]:_._
**Proof.** Let _U_ be a set of cardinality |_X_| + 2 disjoint from _X_. Define

(49) J[^1]: := {_Y_ ∪ _W_ | _Y_ ∈I[^1]:_, W_ ⊆ _U,_|_Y_ ∪ _W_|≤|_X_| + 1}_,_
J[^2]: := {_Y_ ∪ _W_ | _Y_ ∈I[^2]:_, W_ ⊆ _U,_|_Y_ ∪ _W_|≤|_X_| + 1}_._

Then _M_′
[^1]: := (_X_∪_U,_J[^1]) and _M_[^2] := (_X_∪_U,_J[^2]) are matroids again. Define ̃_w_ : _X_ → Z
by

(50) _w_ ̃(_x_) := _w_(_x_) if _x_ ∈ _X_,
_w_ ̃(_x_) := 0 if _x_ ∈ _U_.

Let _W_ be a subset of _U_ of cardinality |_X_ \_Y_ |+1. Then _Y_ ∪_W_ is a common basis
of _M_′
[^1]: and _M_′
[^2]:. In fact, _Y_ ∪ _W_ is a maximum-weight common basis with respect to
the weight function ̃_w_. (Any basis _B_ of larger weight would intersect _X_ in a common
independent set of _M_[^1]: and _M_[^2]: of larger weight than _Y_ .)
So by Theorem 10.15, there exist functions ̃_w_[^1]:_,w_ ̃[^2]: : _X_ → Z so that ̃_w_[^1]: + ̃_w_[^2]: = ̃_w_
and so that _Y_ ∪ _W_ is both a maximum-weight basis of _M_′
[^1]: with respect to ̃_w_[^1] and a
maximum-weight basis of _M_′
[^2]: with respect to ̃_w_[^2].
Now, ̃_w_[^1]:(_u_′′) ≤ _w_ ̃[^1]:(_u_′) whenever _u_′ ∈ _W, u_′′ ∈ _U_ \ _W_. Otherwise we can replace _u_′
in _Y_ ∪_W_ by _u_′′ to obtain a basis of _M_′
[^1]: of larger ̃_w_[^1]-weight. Similarly, ̃_w_[^2](_u_′′) ≤ _w_ ̃[^2](_u_′)
whenever _u_′ ∈ _W, u_′′ ∈ _U_ \ _W_.
Since ̃_w_[^1]:(_u_) + ̃_w_[^2]:(_u_) = ̃_w_(_u_) = 0 for all _u_ ∈ _U_, this implies that ̃_w_[^1]:(_u_′′) = ̃_w_[^1]:(_u_′)
whenever _u_′ ∈ _W, u_′′ ∈ _U_ \_W_. As ∅6= _W_ [^6]:= _U_, it follows that ̃_w_[^1]: and ̃_w_[^2]: are constant

on _U_. Since we can add a constant function to ̃_w_[^1]: and subtracting the same function
from ̃_w_[^2]: without spoiling the required properties, we may assume that ̃_w_[^1]: and ̃_w_[^2]: are
0 on _U_.

Now define _w_[^1]:(_x_) := ̃_w_[^1]:(_x_) and _w_[^2]:(_x_) := ̃_w_[^2]:(_x_) for each _x_ ∈ _X_. Then _Y_ is both a
maximum-weight independent set of _M_[^1]: with respect to _w_[^1]: and a maximum-weight
independent set of _M_[^2]: with respect to _w_[^2]:.

Having this theorem, it is quite easy to derive that the intersection of two matroid
polytopes has integer vertices, being incidence vectors of common independent sets.
By Theorem 10.14 the intersection _P_(_M_[^1]:) ∩ _P_(_M_[^2]:) of the matroid polytopes as-
sociated with the matroids _M_[^1]: = (_X,_I[^1]:) and _M_[^2]: = (_X,_I[^2]:) is determined by:

(51) _z_(_x_) ≥ [^0]: (_x_ ∈ _X_)_,_
_z_(_Y_ ) ≤ _rM_[^1]:(_Y_ ) (_Y_ ⊆ _X_)_,_
_z_(_Y_ ) ≤ _rM_[^2]:(_Y_ ) (_Y_ ⊆ _X_)_,_

The corresponding linear programming problem is, for any _w_ : _X_ → R:

(52) maximize _wTz_,
subject to _z_(_x_) ≥ [^0]: (_x_ ∈ _X_)_,_
_z_(_Y_ ) ≤ _rM_[^1]:(_Y_ ) (_Y_ ⊆ _X_)_,_
_z_(_Y_ ) ≤ _rM_[^2]:(_Y_ ) (_Y_ ⊆ _X_)_._

Again we consider the dual linear programming problem:

(53) minimize

```
∑
```

```
Y⊆X
```

```
(y′
Y rM[^1](Y ) + y′′
Y rM[^2](Y ))
```

```
subject to y′
Y ≥[^0] (Y ⊆ X),
y′′
∑ Y ≥[^0] (Y ⊆ X),
```

```
Y⊆X,x∈Y
```

```
(y′
Y + y′′
Y ) ≥ w(x) (x ∈ X).
```

Now

**Theorem 10.17.** _If w is integer, then_ (52) _and_ (53) _have integer optimum solutions._

**Proof.** Let _Y_ be a common independent set of maximum weight _w_(_Y_ ). By Theorem
10.15, there exist functions _w_[^1]:_, w_[^2]: : _X_ → Z so that _w_[^1]: + _w_[^2]: = _w_ and so that _Y_ is a
maximum-weight independent set in _Mi_ with respect to _wi_ (_i_ = 1_,_2).
Applying Theorem 10.14 to _w_[^1]: and _w_[^2]:, respectively, we know that there exist
integer optimum solutions _y_′ and _y_′′, respectively, for problem (43) with respect to
_M_[^1]:_, w_[^1]: and _M_[^2]:_, w_[^2]:, respectively. One easily checks that _y_′_, y_′′ forms a feasible solution
of (53). Optimality follows from:

```
(54) w(Z) = w1(Z) + w2(Z) =
Y⊆X
```

```
y′
Y rM[^1](Y ) +
Y⊆X
```

```
y′′
Y rM[^2](Y )
```

```
=
```

```
∑
```

```
Y⊆X
```

```
(y′
Y rM[^1](Y ) + y′′
Y rM[^2](Y )).
```

```
So system (51) is totally dual integral. Again, this directly implies:
```

```
Corollary 10.17a. The convex hull of the common independent sets of two matroids
M1 and M2 is determined by (51).
Proof. Directly from Theorem 10.17.
```

```
Exercises
```

10.40. Give an example of three matroids _M_[^1]:, _M_[^2]:, and _M_[^3]: on the same set _X_ so that the
intersection _P_(_M_[^1]:)∩_P_(_M_[^2]:)∩_P_(_M_[^3]:) is _not_ the convex hull of the common independent
sets.

10.41. Derive Edmonds’ matroid intersection theorem (Theorem 10.9) from Theorem 10.17.

## References

[1975] G.M. Adel’son-Vel’ski ̆ı, E.A. Dinits, A.V. Karzanov, _Potokovye algoritmy_ [Russian;
Flow Algorithms], Izdatel’stvo “Nauka”, Moscow, 1975.

[1974] A.V. Aho, J.E. Hopcroft, J.D. Ullman, _The Design and Analysis of Computer Algo-_
_rithms_, Addison-Wesley, Reading, Massachusetts, 1974.

[1993] R.K. Ahuja, T.L. Magnanti, J.B. Orlin, _Network Flows — Theory, Algorithms, and_
_Applications_, Prentice Hall, Englewood Cliffs, New Jersey, 1993.

[1977] K. Appel, W. Haken, Every planar map is four colorable Part I: discharging, _Illinois_
_Journal of Mathematics_ 21 (1977) 429–490.

[1977] K. Appel, W. Haken, J. Koch, Every planar map is four colorable Part II: reducibility,
_Illinois Journal of Mathematics_ 21 (1977) 491–567.

[1969] M.L. Balinski, Labelling to obtain a maximum matching, in: _Combinatorial Math-_
_ematics and Its Applications_ (Proceedings Conference Chapel Hill, North Carolina,
1967; R.C. Bose, T.A. Dowling, eds.), The University of North Carolina Press, Chapel
Hill, North Carolina, 1969, pp. 585–602.

[1957] T.E. Bartlett, An algorithm for the minimum number of transport units to maintain
a fixed schedule, _Naval Research Logistics Quarterly_ 4 (1957) 139–149.

[1986] M. Becker, K. Mehlhorn, Algorithms for routing in planar graphs, _Acta Informatica_
23 (1986) 163–176.

[1958] R. Bellman, On a routing problem, _Quarterly of Applied Mathematics_ 16 (1958) 87–

90.

[1958] C. Berge, Sur le couplage maximum d’un graphe, _Comptes Rendus Hebdomadaires_
_des S ́eances de l’Acad ́emie des Sciences [Paris]_ 247 (1958) 258–259.

[1960] C. Berge, Les probl`emes de coloration en th ́eorie des graphes, _Publications de l’Institut_
_de Statistique de l’Universit ́e de Paris_ 9 (1960) 123–160.

[1961] C. Berge, F ̈arbung von Graphen, deren s ̈amtliche bzw. deren ungerade Kreise starr
sind, _Wissenschaftliche Zeitschrift der Martin-Luther-Universit ̈at Halle-Wittenberg,_
_Mathematisch-Naturwissenschaftliche Reihe_ 10 (1961) 114–115.

[1963] C. Berge, Some classes of perfect graphs, in: _Six Papers on Graph Theory_ [related to
a series of lectures at the Research and Training School of the Indian Statistical Insti-
tute, Calcutta, March-April 1963], Research and Training School, Indian Statistical
Institute, Calcutta, [1963,] pp. 1–21.

[1946] G. Birkhoff, Tres observaciones sobre el algebra lineal, _Revista Facultad de Ciencias_
_Exactas, Puras y Aplicadas Universidad Nacional de Tucuman, Serie A (Matematicas_
_y Fisica Teorica)_ 5 (1946) 147–151.

```
[1926] O. Bor ̊uvka, O jist ́em probl ́emu minim ́aln ́ım [Czech, with German summary; On a
minimal problem], Pr ́ace Moravsk ́e Pˇr ́ırodovˇedeck ́e Spoleˇcnosti Brno [Acta Societatis
Scientiarum Naturalium Moravi[c]ae] 3 (1926) 37–58.
[1941] R.L. Brooks, On colouring the nodes of a network, Proceedings of the Cambridge
Philosophical Society 37 (1941) 194–197.
[1911] C. Carath ́eodory, Uber den Variabilit ̈atsbereich der Fourierschen Konstant ̈ en von
positiven harmonischen Funktionen, Rendiconti del Circolo Matematico di Palermo 32
(1911) 193–217 [reprinted in: Constantin Carath ́eodory, Gesammelte Mathematische
Schriften, Band III (H. Tietze, ed.), C.H. Beck’sche Verlagsbuchhandlung, M ̈unchen,
1955, pp. 78–110].
[1976] N. Christofides, Worst-Case Analysis of a New Heuristic for the Travelling Salesman
Problem, Management Sciences Research Report 388, Graduate School of Industrial
Administration, Carnegie-Mellon University, Pittsburgh, Pennsylvania, 1976.
[1984] R. Cole, A. Siegel, River routing every which way, but loose, in: 25th Annual Sympo-
sium on Foundations of Computer Science (25th FOCS, Singer Island, Florida, 1984),
IEEE, New York, 1984, pp. 65–73.
[1971] S.A. Cook, The complexity of theorem-proving procedures, in: Conference Record of
Third Annual ACM Symposium on Theory of Computing (3rd STOC, Shaker Heights,
Ohio, 1971), The Association for Computing Machinery, New York, 1971, pp. 151–
```

158.
[1978] W.H. Cunningham, A.B. Marsh, III, A primal algorithm for optimum matching,
[in: _Polyhedral Combinatorics — Dedicated to the Memory of D.R. Fulkerson_ (M.L.
Balinski, A.J. Hoffman, eds.)] _Mathematical Programming Study_ 8 (1978) 50–72.
[1951a] G.B. Dantzig, Application of the simplex method to a transportation problem, in: _Ac-_
_tivity Analysis of Production and Allocation — Proceedings of a Conference_ (Proceed-
ings Conference on Linear Programming, Chicago, Illinois, 1949; Tj.C. Koopmans,
ed.), Wiley, New York, 1951, pp. 359–373.

[1951b] G.B. Dantzig, Maximization of a linear function of variables subject to linear in-
equalities, in: _Activity Analysis of Production and Allocation — Proceedings of a_
_Conference_ (Proceedings Conference on Linear Programming, Chicago, Illinois, 1949;
Tj.C. Koopmans, ed.), Wiley, New York, 1951, pp. 339–347.

```
[1988] D. De Caen, On a theorem of K ̋onig on bipartite graphs, Journal of Combinatorics,
Information & System Sciences 13 (1988) 127.
[1959] E.W. Dijkstra, A note on two problems in connexion with graphs, Numerische Math-
ematik 1 (1959) 269–271.
[1950] R.P. Dilworth, A decomposition theorem for partially ordered sets, Annals of Math-
ematics (2) 51 (1950) 161–166 [reprinted in: The Dilworth Theorems — Selected
Papers of Robert P. Dilworth (K.P. Bogart, R. Freese, J.P.S. Kung, eds.), Birkh ̈auser,
Boston, Massachusetts, 1990, pp. 7–12].
```

```
[1970] E.A. Dinits, Algoritm resheniya zadachi o maksimal’nom potoke v seti so stepenno ̆ı
otsenko ̆ı [Russian], Doklady Akademii Nauk SSSR 194 (1970) 754–757 [English trans-
lation: Algorithm for solution of a problem of maximum flow in a network with power
estimation Soviet Mathematics Doklady 11 (1970) 1277–1280].
```

```
[1961] G.A. Dirac, On rigid circuit graphs, Abhandlungen aus dem Mathematischen Seminar
der Universit ̈at Hamburg 25 (1961) 71–76.
```

[1965a] J. Edmonds, Maximum matching and a polyhedron with 0,1-vertices, _Journal of_
_Research National Bureau of Standards Section B_ 69 (1965) 125–130.

[1965b] J. Edmonds, Minimum partition of a matroid into independent subsets, _Journal of_
_Research National Bureau of Standards Section B_ 69 (1965) 67–72.

```
[1965c] J. Edmonds, Paths, trees, and flowers, Canadian Journal of Mathematics 17 (1965)
449–467.
```

```
[1970] J. Edmonds, Submodular functions, matroids, and certain polyhedra, in: Combinato-
rial Structures and Their Applications (Proceedings Calgary International Conference
on Combinatorial Structures and Their Applications, Calgary, Alberta, 1969; R. Guy,
H. Hanani, N. Sauer, J. Sch ̈onheim, eds.), Gordon and Breach, New York, 1970, pp.
69–87.
```

```
[1965] J. Edmonds, D.R. Fulkerson, Transversals and matroid partition, Journal of Research
National Bureau of Standards Section B 69 (1965) 147–153.
```

```
[1972] J. Edmonds, R.M. Karp, Theoretical improvements in algorithmic efficiency for net-
work flow problems, Journal of the Association for Computing Machinery 19 (1972)
248–264.
```

```
[1931] J. Egerv ́ary, Matrixok kombinatorius tulajdons ́agair ́ol [Hungarian, with German sum-
mary], Matematikai ́es Fizikai Lapok 38 (1931) 16–28 [English translation [by H.W.
Kuhn]: On combinatorial properties of matrices, Logistics Papers, George Washing-
ton University, issue 11 (1955), paper 4, pp. 1–11].
```

```
[1976] S. Even, A. Itai, A. Shamir, On the complexity of timetable and multicommodity
flow problems, SIAM Journal on Computing 5 (1976) 691–703.
```

```
[1975] S. Even, O. Kariv, An O(n[^2].[^5]) algorithm for maximum matching in general graphs, in:
16th Annual Symposium on Foundations of Computer Science (16th FOCS, Berkeley,
California, 1975), IEEE, New York, 1975, pp. 100–112.
```

```
[1894] Gy. Farkas, A Fourier-f ́ele mechanikai elv alkalmaz ́asai [Hungarian], Mathematikai
́es Term ́eszettudom ́anyi Ertesit ̋o ́ 12 (1894) 457–472 [German translation, with slight
alterations: J. Farkas, Uber die Anwendungen des mechanischen Princips von Fourier ̈ ,
Mathematische und naturwissenschaftliche Berichte aus Ungarn 12 (1895) 263–281].
```

```
[1896] Gy. Farkas, A Fourier-f ́ele mechanikai elv alkalmaz ́as ́anak algebrai alapja [Hungarian;
On the algebraic foundation of the applications of the mechanical principle of Fourier],
```

```
Mathematikai ́es Physikai Lapok 5 (1896) 49–54 [German translation, with some ad-
ditions: Section I of: J. Farkas, Die algebraischen Grundlagen der Anwendungen des
Fourier’schen Princips in der Mechanik, Mathematische und naturwissenschaftliche
Berichte aus Ungarn 15 (1897-99) 25–40].
```

[1898] Gy. Farkas, A Fourier-f ́ele mechanikai elv algebrai alapja [Hungarian], _Math ́ematikai_
_́es Term ́eszettudom ́anyi Ertesit ̋o ́_ 16 (1898) 361–364 [German translation: J. Farkas,
Die algebraische Grundlage der Anwendungen des mechanischen Princips von Fourier,
_Mathematische und naturwissenschaftliche Berichte aus Ungarn_ 16 (1898) 154–157].

[1957] G.J. Feeney, The empty boxcar distribution problem, _Proceedings of the First Interna-_
_tional Conference on Operational Research (Oxford 1957)_ (M. Davies, R.T. Eddison,
T. Page, eds.), Operations Research Society of America, Baltimore, Maryland, 1957,
pp. 250–265.

[1955] A.R. Ferguson, G.B. Dantzig, The problem of routing aircraft, _Aeronautical Engi-_
_neering Review_ 14 (1955) 51–55.

[1956] L.R. Ford, Jr, _Network Flow Theory_, Paper P-923, The RAND Corporation, Santa
Monica, California, [August 14,] 1956.

[1956] L.R. Ford, Jr, D.R. Fulkerson, Maximal flow through a network, _Canadian Journal_
_of Mathematics_ 8 (1956) 399–404.

[1957] L.R. Ford, Jr, D.R. Fulkerson, A simple algorithm for finding maximal network flows
and an application to the Hitchcock problem, _Canadian Journal of Mathematics_ 9
(1957) 210–218.

[1958] L.R. Ford, Jr, D.R. Fulkerson, A suggested computation for maximal multi-commodity
network flows, _Management Science_ 5 (1958) 97–101.

[1962] L.R. Ford, Jr, D.R. Fulkerson, _Flows in Networks_, Princeton University Press, Prince-
ton, New Jersey, 1962.

[1980] S. Fortune, J. Hopcroft, J. Wyllie, The directed subgraph homeomorphism problem,
_Theoretical Computer Science_ 10 (1980) 111–121.

[1984] M.L. Fredman, R.E. Tarjan, Fibonacci heaps and their uses in improved network
optimization algorithms, in: _25th Annual Symposium on Foundations of Computer_
_Science_ (25th FOCS, Singer Island, Florida, 1984), IEEE, New York, 1984, pp. 338–

346.

[1917] G. Frobenius, Uber zerlegbare Determinanten, ̈ _Sitzungsberichte der K ̈oniglich Preuß-_
_ischen Akademie der Wissenschaften zu Berlin_ (1917) 274–277 [reprinted in: _Fer-_
_dinand Georg Frobenius, Gesammelte Abhandlungen_, Band III (J.-P. Serre, ed.),
Springer, Berlin, 1968, pp. 701–704].

[1973] H.N. Gabow, _Implementation of Algorithms for Maximum Matching on Nonbipar-_
_tite Graphs_, Ph.D. Thesis, Department of Computer Science, Stanford University,
Stanford, California, 1973.

[1990] H.N. Gabow, Data structures for weighted matching and nearest common ancestors
with linking, in: _Proceedings of the First Annual ACM-SIAM Symposium on Dis-_
_crete Algorithms_ (San Francisco, California, 1990), Society for Industrial and Applied
Mathematics, Philadelphia, Pennsylvania, 1990, pp. 434–443.

[1986] Z. Galil, S. Micali, H. Gabow, An _O_(_EV_ log_V_ ) algorithm for finding a maximal
weighted matching in general graphs, _SIAM Journal on Computing_ 15 (1986) 120–

130.

[1958] T. Gallai, Maximum-minimum S ̈atze ̈uber Graphen, _Acta Mathematica Academiae_
_Scientiarum Hungaricae_ 9 (1958) 395–434.

[1959] T. Gallai, Uber extreme Punkt- und Kantenmengen, ̈ _Annales Universitatis Scien-_
_tiarum Budapestinensis de Rolando E ̈otv ̈os Nominatae, Sectio Mathematica_ 2 (1959)
133–138.

[1979] M.R. Garey, D.S. Johnson, _Computers and Intractability — A Guide to the Theory_
_of NP-Completeness_, Freeman, San Francisco, California, 1979.

[1996] G.S. Gasparian, Minimal imperfect graphs: a simple approach, _Combinatorica_ 16
(1996) 209–212.

[1990] A.V. Goldberg, E. Tardos, R.E. Tarjan, Network flow algorithms, in: ́ _Paths, Flows,_
_and VLSI-Layout_ (B. Korte, L. Lov ́asz, H.J. Pr ̈omel, A. Schrijver, eds.), Springer,
Berlin, 1990, pp. 101–164.

[1988] A.V. Goldberg, R.E. Tarjan, A new approach to the maximum-flow problem, _Journal_
_of the Association for Computing Machinery_ 35 (1988) 921–940.

[1990] A.V. Goldberg, R.E. Tarjan, Finding minimum-cost circulations by successive ap-
proximation, _Mathematics of Operations Research_ 15 (1990) 430–466.

[1873] P. Gordan, Ueber die Aufl ̈osung linearer Gleichungen mit reellen Coefficienten, _Math-_
_ematische Annalen_ 6 (1873) 23–28.

[2000] F. G ̈oring, Short proof of Menger’s theorem, _Discrete Mathematics_ 219 (2000) 295–

296.

[1981] M. Gr ̈otschel, L. Lov ́asz, A. Schrijver, The ellipsoid method and its consequences in
combinatorial optimization, _Combinatorica_ 1 (1981) 169–197 [corrigendum: _Combi-_
_natorica_ 4 (1984) 291–295].

[1960] M.-g. Guan, Graphic programming using odd or even points [in Chinese], _Acta Math-_
_ematica Sinica_ 10 (1960) 263–266 [English translation: _Chinese Mathematics_ 1 (1962)
273–277].

[1943] H. Hadwiger, Uber eine Klassifikation der Streckenkomplexe, ̈ _Vierteljahrsschrift der_
_naturforschenden Gesellschaft in Z ̈urich_ 88 (1943) 133–142.

[1958] A. Hajnal, J. Sur ́anyi, Uber die Aufl ̈osung von Graphen in vollst ̈andige Teilgraphe ̈ n,
_Annales Universitatis Scientiarum Budapestinensis de Rolando E ̈otv ̈os Nominatae —_
_Sectio Mathematica_ 1 (1958) 113–121.

[1935] P. Hall, On representatives of subsets, _The Journal of the London Mathematical So-_
_ciety_ 10 (1935) 26–30 [reprinted in: _The Collected Works of Philip Hall_ (K.W. Gru-
enberg, J.E. Roseblade, eds.), Clarendon Press, Oxford, 1988, pp. 165–169].

[1960] A.J. Hoffman, Some recent applications of the theory of linear inequalities to extremal
combinatorial analysis, in: _Combinatorial Analysis_ (New York, 1958; R. Bellman, M.
Hall, Jr, eds.) [Proceedings of Symposia in Applied Mathematics, Volume X], Ameri-
can Mathematical Society, Providence, Rhode Island, 1960, pp. 113–127 [reprinted in:
_Selected Papers of Alan Hoffman — With Commentary_ (C.A. Micchelli, ed.), World
Scientific, Singapore, 2003, pp. 244–248].

[1956] A.J. Hoffman, J.B. Kruskal, Integral boundary points of convex polyhedra, in: _Lin-_
_ear Inequalities and Related Systems_ (H.W. Kuhn, A.W. Tucker, eds.) [Annals of
Mathematics Studies 38], Princeton University Press, Princeton, New Jersey, 1956,
pp. 223–246 [reprinted in: _Selected Papers of Alan Hoffman — With Commentary_
(C.A. Micchelli, ed.), World Scientific, Singapore, 2003, pp. 220–243].

[1973] J. Hopcroft, R.M. Karp, An _n_[^5]_/_[^2] algorithm for maximum matchings in bipartite
graphs, _SIAM Journal on Computing_ 2 (1973) 225–231.

[1961] T.C. Hu, The maximum capacity route problem, _Operations Research_ 9 (1961) 898–

900.

[1963] T.C. Hu, Multi-commodity network flows, _Operations Research_ 11 (1963) 344–360.

[1977] D.B. Johnson, Efficient algorithms for shortest paths in sparse networks, _Journal of_
_the Association for Computing Machinery_ 24 (1977) 1–13.

[1984] N. Karmarkar, A new polynomial-time algorithm for linear programming, _Combina-_
_torica_ 4 (1984) 373–395.

[1972] R.M. Karp, Reducibility among combinatorial problems, in: _Complexity of Computer_
_Computations_ (Proceedings of a symposium on the Complexity of Computer Com-
putations, IBM Thomas J. Watson Research Center, Yorktown Heights, New York,
1972; R.E. Miller, J.W. Thatcher, eds.), Plenum Press, New York, 1972, pp. 85–103.

[1975] R.M. Karp, On the computational complexity of combinatorial problems, _Networks_
5 (1975) 45–68.

[1974] A.V. Karzanov, Nakhozhdenie maksimal’nogo potoka v seti metodom predpotokov
[Russian; Determining the maximal flow in a network by the method of preflows],
_Doklady Akademii Nauk SSSR_ 215 (1974) 49–52 [English translation: _Soviet Mathe-_
_matics Doklady_ 15 (1974) 434–437].

[1979] L.G. Khachiyan, Polinomialny ̆ı algoritm v line ̆ınom programmirovanii [Russian], _Dok-_
_lady Akademii Nauk SSSR_ 244 (1979) 1093–1096 [English translation: A polynomial
algorithm in linear programming, _Soviet Mathematics Doklady_ 20 (1979) 191–194].
[1980] L.G. Khachiyan, Polinomial’nye algoritmy v line ̆ınom programmirovanii [Russian],
_Zhurnal Vychislitel’no ̆ı Matematiki i Matematichesko ̆ı Fiziki_ 20 (1980) 51–86 [English
translation: Polynomial algorithms in linear programming, _U.S.S.R. Computational_
_Mathematics and Mathematical Physics_ 20 (1980) 53–72].
[1968] D.E. Knuth, _The Art of Computer Programming, Volume I Fundamental Algorithms_,
Addison-Wesley, Reading, Massachusetts, 1968.
[1916] D. K ̋onig, Graphok ́es alkalmaz ́asuk a determin ́ansok ́es a halmazok elm ́elet ́ere [Hun-
garian], _Mathematikai ́es Term ́eszettudom ́anyi Ertesit ̋o ́_ 34 (1916) 104–119 [German
translation: Uber Graphen und ihre Anwendung auf Determinantentheorie u ̈ nd Men-
genlehre, _Mathematische Annalen_ 77 (1916) 453–465].
[1931] D. K ̋onig, Graphok ́es matrixok [Hungarian; Graphs and matrices], _Matematikai ́es_
_Fizikai Lapok_ 38 (1931) 116–119.
[1932] D. K ̋onig, Uber trennende Knotenpunkte in Graphen (nebst Anwendungen ̈ auf Deter-
minanten und Matrizen), _Acta Litterarum ac Scientiarum Regiae Universitatis Hun-_
_garicae Francisco-Josephinae, Sectio Scientiarum Mathematicarum [Szeged]_ 6 (1932-
34) 155–179.
[1948] Tj.C. Koopmans, Optimum utilization of the transportation system, in: _The Econo-_
_metric Society Meeting_ (Washington, D.C., 1947; D.H. Leavens, ed.) [Proceedings of
the International Statistical Conferences — Volume V], 1948, pp. 136–146 [reprinted
in: _Econometrica_ 17 (Supplement) (1949) 136–146] [reprinted in: _Scientific Papers of_
_Tjalling C. Koopmans_, Springer, Berlin, 1970, pp. 184–193].
[1984] M.R. Kramer, J. van Leeuwen, The complexity of wire-routing and finding minimum
area layouts for arbitrary VLSI circuits, in: _VLSI-Theory_ (F.P. Preparata, ed.) [Ad-
vances in Computing Research, Volume 2], JAI Press, Greenwich, Connecticut, 1984,
pp. 129–146.
[1956] J.B. Kruskal, Jr, On the shortest spanning subtree of a graph and the traveling
salesman problem, _Proceedings of the American Mathematical Society_ 7 (1956) 48–

50.
[1955] H.W. Kuhn, The Hungarian method for the assignment problem, _Naval Research_
_Logistics Quarterly_ 2 (1955) 83–97.

[1976] E.L. Lawler, _Combinatorial Optimization: Networks and Matroids_, Holt, Rinehart
and Winston, New York, 1976.

[1985] C.E. Leiserson, F.M. Maley, Algorithms for routing and testing routability of planar
VLSI layouts, in: _Proceedings of the Seventeenth Annual ACM Symposium on Theory_
_of Computing_ (17th STOC, Providence, Rhode Island, 1985), The Association for
Computing Machinery, New York, 1985, pp. 69–78.

[1972a] L. Lov ́asz, A characterization of perfect graphs, _Journal of Combinatorial Theory,_
_Series B_ 13 (1972) 95–98.

[1972b] L. Lov ́asz, Normal hypergraphs and the perfect graph conjecture, _Discrete Mathe-_
_matics_ 2 (1972) 253–267 [reprinted as: Normal hypergraphs and the weak perfect
graph conjecture, in: _Topics on Perfect Graphs_ (C. Berge, V. Chv ́atal, eds.) [Annals
of Discrete Mathematics 21], North-Holland, Amsterdam, 1984, pp. 29–42].

```
[1979] L. Lov ́asz, Graph theory and integer programming, in: Discrete Optimization I (Pro-
ceedings Advanced Research Institute on Discrete Optimization and Systems Appli-
cations and Discrete Optimization Symposium, Banff, Alta, and Vancouver, B.C.,
Canada, 1977; P.L. Hammer, E.L. Johnson, B.H. Korte, eds.) [Annals of Discrete
Mathematics 4], North-Holland, Amsterdam, 1979, pp. 141–158.
```

```
[1986] L. Lov ́asz, M.D. Plummer, Matching Theory, Akad ́emiai Kiad ́o, Budapest [also:
North-Holland Mathematics Studies Volume 121, North-Holland, Amsterdam], 1986.
```

```
[1975] J.F. Lynch, The equivalence of theorem proving and the interconnection problem,
(ACM) SIGDA Newsletter 5:3 (1975) 31–36.
```

```
[1978] V.M. Malhotra, M.P. Kumar, S.N. Maheshwari, An O(|V |[^3]) algorithm for finding
maximum flows in networks, Information Processing Letters 7 (1978) 277–278.
```

```
[1985] K. Matsumoto, T. Nishizeki, N. Saito, An efficient algorithm for finding multicom-
modity flows in planar networks, SIAM Journal on Computing 14 (1985) 289–302.
```

```
[1927] K. Menger, Zur allgemeinen Kurventheorie, Fundamenta Mathematicae 10 (1927)
96–115.
```

```
[1980] S. Micali, V.V. Vazirani, An O(
```

```
√
|v||E|) algorithm for finding maximum matching
in general graphs, in: 21st Annual Symposium on Foundations of Computer Science
(21st FOCS, Syracuse, New York, 1980), IEEE, New York, 1980, pp. 17–27.
```

```
[1784] G. Monge, M ́emoire sur la th ́eorie des d ́eblais et des remblais, Histoire de l’Acad ́emie
Royale des Sciences [ann ́ee 1781. Avec les M ́emoires de Math ́ematique & de Physique,
pour la mˆeme Ann ́ee] (2e partie) (1784) [Histoire: 34–38, M ́emoire:] 666–704.
```

```
[1936] T.S. Motzkin, Beitr ̈age zur Theorie der linearen Ungleichungen, Inaugural Disserta-
tion Basel, Azriel, Jerusalem, 1936 [English translation: Contributions to the theory
of linear inequalities, RAND Corporation Translation 22, The RAND Corporation,
Santa Monica, California, 1952 [reprinted in: Theodore S. Motzkin: Selected Papers
(D. Cantor, B. Gordon, B. Rothschild, eds.), Birkh ̈auser, Boston, Massachusetts,
1983, pp. 1–80]].
```

```
[1961] C.St.J.A. Nash-Williams, Edge-disjoint spanning trees of finite graphs, The Journal
of the London Mathematical Society 36 (1961) 445–450.
```

```
[1964] C.St.J.A. Nash-Williams, Decomposition of finite graphs into forests, The Journal of
the London Mathematical Society 39 (1964) 12.
```

[1967] C.St.J.A. Nash-Williams, An application of matroids to graph theory, in: _Theory_
_of Graphs — International Symposium — Th ́eorie des graphes — Journ ́ees interna-_
_tionales d’ ́etude_ (Rome, 1966; P. Rosenstiehl, ed.), Gordon and Breach, New York,
and Dunod, Paris, 1967, pp. 263–265.

[1985] C.St.J.A. Nash-Williams, Connected detachments of graphs and generalized Euler
trails, _The Journal of the London Mathematical Society_ (2) 31 (1985) 17–29.

[1947] J. von Neumann, _Discussion of a maximum problem_, typescript dated November 15–
16, 1947, Institute for Advanced Study, Princeton, New Jersey, 1947 [reprinted in:
_John von Neumann, Collected Works, Volume VI_ (A.H. Taub, ed.), Pergamon Press,
Oxford, 1963, pp. 89–95].

[1953] J. von Neumann, A certain zero-sum two-person game equivalent to the optimal
assignment problem, in: _Contributions to the Theory of Games Volume II_ (H.W.
Kuhn, A.W. Tucker, eds.) [Annals of Mathematics Studies 28], Princeton University
Press, Princeton, New Jersey, 1953, pp. 5–12 [reprinted in: _John von Neumann,_
_Collected Works, Volume VI_ (A.H. Taub, ed.), Pergamon Press, Oxford, 1963, pp.
44–49].

[1968] A.R.D. Norman, M.J. Dowling, _Railroad Car Inventory: Empty Woodrack Cars on_
_the Louisville and Nashville Railroad_, Technical Report 320-2926, IBM New York
Scientific Center, New York, 1968.

[1981] H. Okamura, P.D. Seymour, Multicommodity flows in planar graphs, _Journal of Com-_
_binatorial Theory, Series B_ 31 (1981) 75–81.

[1988] J.B. Orlin, A faster strongly polynomial minimum cost flow algorithm, in: _Proceedings_
_of the Twentieth Annual ACM Symposium on Theory of Computing_ (20th STOC,
Chicago, Illinois, 1988), The Association for Computing Machinery, New York, 1988,
pp. 377–387.

[1993] J.B. Orlin, A faster strongly polynomial minimum cost flow algorithm, _Operations_
_Research_ 41 (1993) 338–350.

[1994] C.H. Papadimitriou, _Computational Complexity_, Addison-Wesley, Reading, Massachusetts,

1994.

[1983] R.Y. Pinter, River routing: methodology and analysis, in: _Third CalTech Confer-_
_ence on Very Large Scale Integration_ (Pasadena, California, 1983; R. Bryant, ed.),
Springer, Berlin, 1983, pp. 141–163.

[1957] R.C. Prim, Shortest connection networks and some generalizations, _The Bell System_
_Technical Journal_ 36 (1957) 1389–1401.

[1942] R. Rado, A theorem on independence relations, _The Quarterly Journal of Mathemat-_
_ics (Oxford)_ (2) 13 (1942) 83–89.

[1965] J.W.H.M.T.S.J. van Rees, Een studie omtrent de circulatie van materieel, _Spoor- en_
_Tramwegen_ 38 (1965) 363–367.

[1997] H. Ripphausen-Lipa, D. Wagner, K. Weihe, The vertex-disjoint Menger problem in
planar graphs, _SIAM Journal on Computing_ 26 (1997) 331–349.

[1956] J.T. Robacker, _Min-Max Theorems on Shortest Chains and Disjunct Cuts of a Net-_
_work_, Research Memorandum RM-1660, The RAND Corporation, Santa Monica,
California, [12 January] 1956.

[1986] N. Robertson, P.D. Seymour, Graph minors. VI. Disjoint paths across a disc, _Journal_
_of Combinatorial Theory, Series B_ 41 (1986) 115–138.

[1995] N. Robertson, P.D. Seymour, Graph minors. XIII. The disjoint paths problem, _Jour-_
_nal of Combinatorial Theory, Series B_ 63 (1995) 65–110.

[1993] N. Robertson, P. Seymour, R. Thomas, Hadwiger’s conjecture for _K_[^6]:-free graphs,
_Combinatorica_ 13 (1993) 279–361.

[1966] B. Rothschild, A. Whinston, Feasibility of two commodity network flows, _Operations_
_Research_ 14 (1966) 1121–1129.

[1973] M. Sakarovitch, Two commodity network flows and linear programming, _Mathemat-_
_ical Programming_ 4 (1973) 1–20.

[1991] A. Schrijver, Disjoint homotopic paths and trees in a planar graph, _Discrete & Com-_
_putational Geometry_ 6 (1991) 527–574.

[1994] A. Schrijver, Finding _k_ disjoint paths in a directed planar graph, _SIAM Journal on_
_Computing_ 23 (1994) 780–788.

[2003] A. Schrijver, _Combinatorial Optimization — Polyhedra and Efficiency_, Springer, Berlin,

2003.

[1915] E. Stiemke, Uber positive L ̈osungen homogener linearer Gleichungen, ̈ _Mathematische_
_Annalen_ 76 (1915) 340–342.

[1985] E. Tardos, A strongly polynomial minimum cost circulation a ́ lgorithm, _Combinatorica_
5 (1985) 247–255.

[1974] R. Tarjan, Finding dominators in directed graphs, _SIAM Journal on Computing_ 3
(1974) 62–89.

[1984] R.E. Tarjan, A simple version of Karzanov’s blocking flow algorithm, _Operations_
_Research Letters_ 2 (1984) 265–268.

[1950] R.L. Thorndike, The problem of the classification of personnel, _Psychometrika_ 15
(1950) 215–235.

[1937] A.M. Turing, On computable numbers, with an application to the Entscheidungsprob-
lem, _Proceedings of the London Mathematical Society_ (2) 42 (1937) 230–265 [correc-
tion: 43 (1937) 544–546] [reprinted in: _The Undecidable — Basic Papers on Undecid-_
_able Propositions, Unsolvable Problems and Computable Functions_ (M. Davis, ed.),
Raven Press, Hewlett, New York, 1965, pp. 116–154].

[1947] W.T. Tutte, The factorization of linear graphs, _The Journal of the London Mathemat-_
_ical Society_ 22 (1947) 107–111 [reprinted in: _Selected Papers of W.T. Tutte Volume I_
(D. McCarthy, R.G. Stanton, eds.), The Charles Babbage Research Centre, St. Pierre,
Manitoba, 1979, pp.89–97].

[1961] W.T. Tutte, On the problem of decomposing a graph into _n_ connected factors, _The_
_Journal of the London Mathematical Society_ 36 (1961) 221–230 [reprinted in: _Selected_
_Papers of W.T. Tutte Volume I_ (D. McCarthy, R.G. Stanton, eds.), The Charles
Babbage Research Centre, St. Pierre, Manitoba, 1979, pp. 214–225].

[1968] A.F. Veinott, Jr, G.B. Dantzig, Integral extreme points, _SIAM Review_ 10 (1968)
371–372.

[1964] V.G. Vizing, Ob otsenke khromaticheskogo klassa _p_-grafa [Russian; On an estimate
of the chromatic class of a _p_-graph], _Diskretny ̆ı Analiz_ 3 (1964) 25–30.

[1937] K. Wagner, Uber eine Eigenschaft der ebenen Komplexe, ̈ _Mathematische Annalen_ 114
(1937) 570–590.

[1995] D. Wagner, K. Weihe, A linear-time algorithm for edge-disjoint paths in planar
graphs, _Combinatorica_ 15 (1995) 135–150.

[1976] D.J.A. Welsh, _Matroid Theory_, Academic Press, London, 1976.

[1969] W.W. White, A.M. Bomberault, A network algorithm for empty freight car allocation,
_IBM Systems Journal_ 8 (1969) 147–169.

## Name index

Adel’son-Vel’ski ̆ı, G.M. 156, 199
Aho, A.V. 98, 110, 199
Ahuja, R.K. 68, 73, 199
Appel, K. 112, 199

Balinski, M.L. 83, 199
Bartlett, T.E. 76, 199
Becker, M. 168, 199
Bellman, R.E. 13, 16-18, 48, 73, 192, 199
Berge, C. 2, 78-79, 94, 96, 113, 125-126,
129, 199
Birkhoff, G. 44, 199
Bomberault, A.M. 74, 209
Bor ̊uvka, O. 19, 200
Brooks, R.L. 112, 200

Carath ́eodory, C. 24, 30, 33, 200
Christofides, N. 89-91, 200
Chudnovsky, M. 113, 126
Cole, R.J. 164, 200
Cook, S.A. 98, 105-106, 200
Cunningham, W.H. 94, 96, 200

Dantzig, G.B. 34, 60, 76, 136, 200, 202,
209
De Caen, D. 41, 200
Dijkstra, E.W. 6, 8-10, 19-21, 48, 200
Dilworth, R.P. 121-124, 127-128, 146,
200
Dinitz, Y. 66, 156, 199, 201
Dirac, G.A. 129, 201
Dowling, M.J. 74, 207

Edmonds, J.R. 66, 79, 81, 83, 85, 88,
91-94, 96, 184, 187-190, 194-195, 198,
201
Egerv ́ary, J. 47, 52, 201
Euler, L. 166-167
Even, S. 83, 151, 155, 201

Farkas, Gy. 2, 23, 30-34, 152, 201-202
Feeney, G.J. 74, 202
Ferguson, A.R. 76, 202
Ford, Jr, L.R. 13, 16-18, 48, 59-60, 62,

```
73, 149, 169, 172, 192, 202
Fortune, S. 151, 157, 202
Fredman, M.L. 10, 202
Frobenius, F.G. 41, 202
Fulkerson, D.R. 59-60, 62, 73, 149, 169,
172, 189, 201-202
```

```
Gabow, H.N. 88, 202-203
Galil, Z. 88, 203
Gallai, T. 2, 39, 79, 111, 127, 203
Garey, M.R. 98, 203
Gasparyan, G.S. 125, 203
Goldberg, A.V. 68, 203
Gordan, P. 33, 203
G ̈oring, F. 54, 203
Gr ̈otschel, M. 128, 203
Guan, M.-g. 89, 203
```

```
Hadwiger, H. 113, 203
Hajnal, A. 129, 204
Haken, W. 112, 199
Hall, P. 43, 204
Hoffman, A.J. 51, 68-69, 135-137,
145-146, 204
Hopcroft, J.E. 45, 98, 110, 151, 157, 199,
202, 204
Hu, T.C. 21, 150, 153, 156, 204
```

```
Itai, A. 151, 155, 201
```

```
Johnson, D.B. 10, 204
Johnson, D.S. 98, 203
```

```
Kariv, O. 83, 201
Karmarkar, N. 34, 204
Karp, R.M. 45, 66, 98, 106, 151, 201, 204
Karzanov, A.V. 67, 156, 199, 204
Khachiyan, L.G. 34, 205
Knuth, D.E. 67, 151, 205
Koch, J. 112, 199
K ̋onig, D. 2, 41-42, 46, 52-53, 115-117,
120, 123, 126-127, 140, 182-183, 188,
205
Koopmans, Tj.C. 73-74, 205
```

Kramer, M.R. 151, 205
Kruskal, Jr, J.B. 20-21, 51, 135-137, 173,
204-205
Kuhn, H.W. 47, 205
Kumar, M.P. 67, 206

Lawler, E.L. 88, 205
Leeuwen, J. van 151, 205
Leiserson, C.E. 164, 205
Lov ́asz, L. 79, 88, 125-126, 128-129, 203,
206
Lynch, J.F. 151, 206

Magnanti, T.L. 68, 73, 199
Maheshwari, S.N. 67, 206
Maley, F.M. 164, 205
Malhotra, V.M. 67, 206
Marsh, III, A.B. 94, 96, 200
Matsumoto, K. 168, 206
Mehlhorn, K. 168, 199
Menger, K. 2, 54-55, 58-60, 206
Micali, S. 83, 88, 203, 206
Monge, G. 49, 206
Motzkin, T.S. 33, 206

Nash-Williams, C.St.J.A. 189-190,
206-207
Neumann, J. von 34, 44, 207
Nishizeki, T. 168, 206
Norman, A.R.D. 74, 207

Okamura, H. 150, 165-166, 207
Orlin, J.B. 68, 73, 199, 207

Papadimitriou, C.H. 98, 110, 207
Pinter, R.Y. 163, 207
Plummer, M.D. 88, 206
Prim, R.C. 19-21, 207

Rado, R. 188, 207
Rees, J.W.H.M.T.S.J. van 76, 207
Ripphausen-Lipa, H. 161, 208
Robacker, J.T. 5, 208
Robertson, G.N. 113, 126, 151, 159, 161,
208
Rothschild, B. 150, 155, 165, 208

```
Saito, N. 168, 206
Sakarovitch, M. 153, 208
Schrijver, A. 83, 128, 151, 164, 203, 208
Seymour, P.D. 113, 126, 150-151, 159,
161, 165-166, 207-208
Shamir, A. 151, 155, 201
Siegel, A. 164, 200
Stiemke, E. 33, 208
Sur ́anyi, J. 129, 204
```

```
Tardos, E. 68, 73, 203, 208 ́
Tarjan, R.E. 10, 67-68, 202-203, 208
Thomas, R. 113, 126, 208
Thorndike, R.L. 49, 208
Thue, A. 100
Turing, A.M. 3, 100, 108-110, 208
Tutte, W.T. 2, 78-80, 94, 96, 103, 190,
209
```

```
Ullman, J.D. 98, 110, 199
```

```
Vazirani, V.V. 83, 206
Veinott, Jr, A.F. 136, 209
Vizing, V.G. 116, 209
```

```
Wagner, D. 161, 168, 208-209
Wagner, K. 113, 209
Weihe, K. 161, 168, 208-209
Welsh, D.J.A. 182, 209
Whinston, A. 150, 155, 165, 208
White, W.W. 74, 209
Wyllie, J. 151, 157, 202
```

## Subject index

accepts word
algorithm **101**
Turing machine **109**
acyclic digraph **157**
affine halfspace **24**
affine hyperplane **23**
airline timetabling 88-89
airplane crew pairing 88
airplane routing 7-8, 56-57, 76, 158
airport terminal assignment 123
algorithm [^100]:-101
polynomial-time **101**
allows sequence of words
algorithm **101**
alphabet 98
alternating forest
_M_- **86**
alternating walk
_M_- **81**
antichain [^121]:-124, **176**
arc-disjoint **54**
arc-disjoint paths **55**
arc-disjoint paths problem [^149]:-152, 158
arc-disjoint _s_ − _t_ paths 55
arc-disjoint _s_ − _t_ paths/min-max 55
assignment
bungalow 122-123
frequency 114
job 45-46, 48-49, 83-84, 122
platform 123
room 83, 88, 114, 123
seat 83, 88
terminal 123
assignment problem 45-46
optimal 48-50
augmenting path
flow **61**
_M_- [^40]:-45, 47-48, **82**

_b_-detachment **190**
_b_-matching [^44]:, [^80]:-81

```
basic solution 30
basis 178
basis in a matroid 174
basis of a matrix 30
Bellman-Ford method 13-14
bend cut
1- 168
bipartite matching 41-53, 185
cardinality 41-46
weighted 47-53
bipartite matching algorithm
cardinality 45
weighted 47-48
blocking flow 67-68
blossom
M- 82
boolean expression 103
box cars
routing empty railway 74
bridge 80
Brooks’ theorem 112
bungalow assignment 122-123
bus station platform assignment 123
```

```
capacity of a cut 58
Carath ́eodory’s theorem 30
cardinality bipartite matching 41-46
cardinality bipartite matching algorithm
45
cardinality common independent set
algorithm 185-187
cardinality common independent set
augmenting algorithm 185-187
cardinality common independent set
problem 184-190
cardinality matching 41-46, 78-85, 132
cardinality matroid intersection 184-190
cardinality matroid intersection algorithm
185-187
cardinality nonbipartite matching 78-85,
132
```

cardinality nonbipartite matching
algorithm 81-83
certificate 97, [^101]:-103
chain [^121]:-124
maximal **123**
child **9**
Chinese postman problem [^89]:-91
chord **128**
chordal graph [^128]:-131
Christofides’ approximative algorithm for
the traveling salesman problem
89-91
chromatic number [^111]:-115, 125-128
edge- [^116]:-117
vertex- [^111]:-115, 125-128
circuit **178**
Hamiltonian **89**
circuit basis **182**
circulation [^68]:-70, 144-146
min-cost 73
minimum-cost 73
circulation theorem
Hoffman’s [^69]:-70, 145-146
class scheduling 117-118
clique [^111]:-112, 125-128
clique number [^111]:-112, 125-128
co-NP [^102]:-103
COCLIQUE **111**
cocycle matroid **180**
cographic matroid [^180]:, 182-183
colour [^111]:, **115**
colourable
_k_- **111**
3- 112-115
colouring [^111]:-115
edge- [^115]:-117
map 113
vertex- [^111]:-115
colouring number [^111]:-115, 125-128
edge- [^116]:-117
vertex- [^111]:-115, 125-128
colouring theorem
K ̋onig’s edge- [^116]:, 120-127
column generation technique 168-172

```
commodity 148
commodity flow problem
fractional k- 148-152, 168-172
integer k- 148-151, 155-156
integer undirected k- 149-151
k- 148-152, 168-172
undirected k- 149-151
common independent set 184
extreme 190
common independent set algorithm
cardinality 185-187
weighted 191-193
common independent set augmenting
algorithm
cardinality 185-187
weighted 191-193
common independent set problem
184-193
cardinality 184-190
weighted 190-193
common SDR 43, 57-58, 70, 120-121,
185, 188
common system of distinct representatives
43, 57-58, 70, 120-121, 185, 188
common transversal 43, 57-58, 70,
120-121, 185, 188
comparability graph 124, 127-128
complement 125
complementary graph 125
complete
NP- 97-98, 103
component
odd 78
component of a collection of sets 80
cone
convex 29
finitely generated convex 29-30
conservation law
flow 58
contraction in a matroid 179-180
convex cone 29
finitely generated 29-30
convex hull 23
convex set 23-24
```

Cook’s theorem 105
cost **71**
cover
edge [^39]:-40
vertex [^39]:-40
cover number
edge [^39]:-40, 79-80
vertex [^39]:-40
cover polytope
edge **143**
vertex **143**
CPM 14-16, 122
crew pairing
airplane 88
Critical Path Method 14-16, 122
cross-free **90**
cross-freeness condition [^160]:-161
Cunningham-Marsh formula [^94]:-96
cut **5**
1-bend **168**
_s_ − _t_ **55**
_s_ − _t_ vertex- **55**
cut condition [^149]:, 152, [^161]:, 168
cut/minimum-size
_s_ − _t_ 55
_s_ − _t_ vertex- 55
cut/minimum-size/min-max
_s_ − _t_ 55
_s_ − _t_ vertex- 55
cycle matroid **180**

decomposition theorem
Dilworth’s [^121]:-122, 124, 127-128,
146
deletion in a matroid [^179]:-180
dependent set in a matroid **174**
deshrinking **85**
detachment
_b_- **190**
Dijkstra-Prim method 19-20
Dilworth’s decomposition theorem
[^121]:-122, 124, 127-128, 146
DIRECTED HAMILTONIAN CYCLE
**107**

```
directed Hamiltonian cycle problem 107
disconnecting vertex set
S − T 54
disconnecting vertex
set/minimum-size/min-max
S − T 54-55
disjoint
arc- 54
internally vertex- 54
vertex- 54
disjoint paths
arc- 55
internally vertex- 55
disjoint paths problem
arc- 149-152, 158
edge- 149-152, 156-168
vertex- 149-152, 157-162
disjoint s − t paths
arc- 55
internally 55
internally vertex- 55
disjoint s − t paths/min-max
arc- 55
internally 55
internally vertex- 55
disjoint S − T paths/min-max 54-55
disjoint spanning trees problem
edge- 185, 190
disjoint trees problem
vertex- 164-165
distance 5-6
distinct representatives
common system of 43, 57-58, 70,
120-121, 185, 188
partial system of 43
system of 42-43, 46, 188
doubly stochastic matrix 44, 143
down-monotone 176
dual LP-problem 34
dual matroid 178, 182-183
dual of a matroid 180
duality theorem of linear programming
34-37
```

dynamic programming 8

edge-chromatic number [^116]:-117
edge-colouring [^115]:-117
edge-colouring number [^116]:-117
edge-colouring theorem
K ̋onig’s [^116]:, 120-127
edge cover [^39]:-40
edge cover number [^39]:-40, 79-80
edge cover polytope **143**
edge cover theorem
K ̋onig’s [^42]:, 115, 123, 126-127,
140-141
edge-disjoint paths problem [^149]:-152,
156-168
edge-disjoint spanning trees problem 185,
190
Edmonds’ matching polytope theorem
[^91]:-[^93]:-94, 96
Edmonds’ matroid intersection theorem
[^188]:, 198
ellipsoid method 34
empty railway box cars
routing 74
end vertex **5**
Euler condition [^155]:, [^165]:-167
extreme common independent set **190**
extreme flow **71**
extreme matching **47**

factor
1- [^78]:-80
factor theorem
Tutte’s 1- **79**
Farkas’ lemma 30-32
Fibonacci forest [^10]:-11
Fibonacci heap [^11]:-12
finitely generated convex cone [^29]:-30
flow 144-146
blocking [^67]:-68
_s_ − _t_ **58**
flow algorithm
Ford-Fulkerson maximum 60-68
maximum 60-68
minimum-cost [^70]:-73

```
flow augmenting algorithm 60-62
flow augmenting path 61
flow conservation law 58
flow problem
fractional k-commodity 148-152,
168-172
fractional multicommodity
148-152, 168-172
integer k-commodity 148-151,
155-156
integer multicommodity 148-151,
155-156
integer two-commodity 155-156
integer undirected k-commodity
149-151
integer undirected multicommodity
149-151
k-commodity 148-152, 168-172
maximum 58-68
min-cost 70-73
minimum-cost 70-73
multicommodity 148-152, 168-172
undirected k-commodity 149-151
undirected multicommodity
149-151
flow theorem
integer 60, 146
follows from word
word 100
Ford-Fulkerson maximum flow algorithm
60-68
forest
Fibonacci 10-11
M-alternating 86
rooted 9
four-colour conjecture 112
four-colour theorem 112
4CC 112
4CT 112
fractional k-commodity flow problem
148-152, 168-172
fractional multicommodity flow problem
148-152, 168-172
frequency assignment 114
```

Gallai’s theorem 39-40
good characterization **103**
good forest **22**
goods
storage of 113-114
Gordan’s theorem 33
graphic matroid [^180]:, 182-183
greedy algorithm 173-176
greedy forest **19**
grid graph 168

Hadwiger’s conjecture [^113]:, 115
halfspace **24**
affine **24**
Hall’s marriage theorem 43
Hamiltonian circuit **89**
HAMILTONIAN CYCLE
DIRECTED **107**
UNDIRECTED [^107]:-108
Hamiltonian cycle problem
directed **107**
undirected [^107]:-108
heap **9**
Fibonacci [^11]:-12
2- [^9]:-10
Hoffman-Kruskal theorem [^137]:-138
Hoffman’s circulation theorem [^69]:-70,
145-146
hull
convex **23**
Hungarian method [^47]:-48
Hu’s two-commodity flow theorem
[^153]:-155
hyperplane **23**
affine **23**

incidence function **91**
incidence matrix of a directed graph **143**
incidence vector [^50]:, [^91]:, [^124]:, [^141]:, **169**
independent set algorithm
cardinality common [^185]:-187
weighted common [^191]:-193
independent set augmenting algorithm
cardinality common [^185]:-187

```
independent set in a matroid 174
independent set problem
cardinality common 184-190
common 184-193
weighted common 190-193
induced subgraph 113
integer flow theorem 60, 146
integer k-commodity flow problem
148-151, 155-156
integer linear programming 132-147
integer multicommodity flow problem
148-151, 155-156
integer polyhedron 133-134-138
integer polytope 133
integer two-commodity flow problem
155-156
integer undirected k-commodity flow
problem 149-151
integer undirected multicommodity flow
problem 149-151
integer vector 132
integrity theorem 60, 62
interior point method 34
internally disjoint s − t paths 55
internally disjoint s−t paths/min-max 55
internally vertex-disjoint 54
internally vertex-disjoint paths 55
internally vertex-disjoint s − t paths 55
internally vertex-disjoint s − t
paths/min-max 55
intersection graph 129
interval matrix 146
```

```
job assignment 45-46, 48-49, 83-84, 122
join
T- 91
```

```
k-commodity flow problem 148-152,
168-172
fractional 148-152, 168-172
integer 148-151, 155-156
integer undirected 149-151
undirected 149-151
k-truncation of a matroid 179
k-uniform matroid 179
```

knapsack problem 14
K ̋onig’s edge-colouring theorem [^116]:,
120-127
K ̋onig’s edge cover theorem [^42]:, 115, 123,
126-127, 140-141
K ̋onig’s matching theorem [^41]:-42, 46,
52-53, 115, 127, 140, 188
Kruskal’s method 20

length of a walk **5**
linear matroid [^180]:-181, 183
linear programming 33-37
duality theorem of [^34]:-37
integer 132-147
literal **111**
Lov ́asz’s perfect graph theorem 125-[^126]:,
128
LP 33-37

_M_-alternating forest **86**
_M_-alternating walk **81**
_M_-augmenting path [^40]:-45, 47-48, **82**
_M_-blossom **82**
map colouring 113
marriage theorem
Hall’s 43
matching [^39]:-53, 78-91, 132
_b_- [^44]:, [^80]:-81
bipartite 41-53, 185
cardinality 41-46, 78-85, 132
cardinality bipartite 41-46
cardinality nonbipartite 78-85, 132
nonbipartite 78-91, 132
perfect [^39]:, 42, 44, 50-[^51]:, [^78]:-80
weighted bipartite 47-53
weighted nonbipartite 85-91
matching algorithm
cardinality bipartite 45
weighted bipartite 47-48
matching number [^39]:-40, 78-79
matching polytope 50-53, 91-[^93]:-94,
[^142]:-143
perfect [^51]:, [^91]:-94, **142**
matching polytope theorem
Edmonds’ [^91]:-[^93]:-94, 96

```
matching problem
weighted 52
matching theorem
K ̋onig’s 41-42, 46, 52-53, 115, 127,
140, 188
matroid 173-198
matroid intersection 184-193
cardinality 184-190
weighted 190-193
matroid intersection algorithm
cardinality 185-187
weighted 191-193
matroid intersection theorem
Edmonds’ 188, 198
matroid polytope 194-198
max-biflow min-cut theorem 156
max-flow min-cut theorem 59-60, 62,
145-146
maximal chain 123
maximum flow algorithm 60-68
Ford-Fulkerson 60-68
maximum flow problem 58-68
maximum reliability problem 21
maximum-size matching 78-79
Menger’s theorem 54-56
directed arc-disjoint version of 55
directed internally vertex-disjoint
version of 55
directed vertex-disjoint version of
54-55
min-cost circulation 73
min-cost flow problem 70-73
minimum-cost circulation 73
minimum-cost flow algorithm 70-73
minimum-cost flow problem 70-73
minimum-cost transportation problem
73-74
minimum-size s − t cut 55
minimum-size s − t cut/min-max 55
minimum-size S − T disconnecting vertex
set/min-max 54-55
minimum-size s − t vertex-cut 55
minimum-size s − t vertex-cut/min-max
55
```

minimum spanning tree 19-22
minor of a graph **113**
minor of a matroid **179**
Motzkin’s theorem 33
multicommodity flow problem [^148]:-152,
168-172
fractional [^148]:-152, 168-172
integer [^148]:-151, 155-156
integer undirected [^149]:-151
undirected [^149]:-151

nested family **85**
net **148**
nonbipartite matching 78-91, 132
cardinality 78-85, 132
weighted 85-91
NP 97-98, [^101]:-103
co- [^102]:-103
NP-complete 97-98, **103**

odd component **78**
Okamura-Seymour theorem [^165]:-168
1-bend cut **168**
1-factor [^78]:-80
1-factor theorem
Tutte’s **79**
optimal assignment problem 48-50
optimal pairing 88
ordered set
partially [^121]:-124

P 97-98, **101**
pairing 83
airplane crew 88
optimal 88
parallel elements in a matroid **175**
parent **9**
partial SDR **43**
partial system of distinct representatives
**43**
partial transversal **43**
partially ordered set [^121]:-124
PARTITION [^106]:-107
partition matroid **182**
partition problem [^106]:-107

```
path 5
M-augmenting 40-45, 47-48, 82
s − t 5
shortest 5-19, 91
path problem
shortest 5-19
paths
arc-disjoint s − t 55
internally disjoint s − t 55
internally vertex-disjoint 55
internally vertex-disjoint s − t 55
paths/min-max
arc-disjoint s − t 55
disjoint S − T 54-55
internally disjoint s − t 55
internally vertex-disjoint s − t 55
perfect graph 125-128
perfect graph theorem 125-126, 128
Lov ́asz’s 125-126, 128
strong 113, 126
perfect matching 39, 42, 44, 50-51, 78-80
perfect matching polytope 51, 91-94, 142
PERT 14-16, 122
PERT-CPM 14-16, 122
planar graph 159-168
platform assignment 123
polyhedron 25-29
integer 133-134-138
polynomial time
Turing machine solves problem in
109
polynomial-time algorithm 101
polytope 25-29
edge cover 143
integer 133
matching 50-53, 91-93-94, 142-143
matroid 194-198
perfect matching 51, 91-94, 142
stable set 143
vertex cover 143
polytope theorem
Edmonds’ matching 91-93-94, 96
postman problem
Chinese 89-91
```

primality testing 103
prize equilibrium 16-17
problem **100**
processor
two- **84**
processor scheduling
two- 83-84
Program Evaluation and Review
Technique 14-16, 122
project scheduling 122

railway box cars
routing empty 74
railway platform assignment 123
railway stock routing 74-76, 151-152
rank **178**
rank function **178**
rank of a matroid **174**
reliability problem
maximum 21
representatives
common system of distinct [^43]:,
57-58, 70, 120-121, 185, 188
partial system of distinct **43**
system of distinct [^42]:-43, 46, 188
restriction in a matroid **179**
rigid circuit graph [^128]:-131
room assignment 83, 88, 114, 123
root [^5]:, **9**
rooted forest **9**
rooted tree [^5]:, [^9]:, 185, 193
routing
airplane 7-8, 56-57, 76, 158
railway stock 74-76, 151-152
ship 7-8, 73-74
vehicle 7-8
VLSI- 151, 162-164
routing empty freighters 73-74
routing empty railway box cars 74

_s_ − _t_ cut/minimum-size 55
_s_ − _t_ cut/minimum-size/min-max 55
_S_ − _T_ disconnecting vertex set **54**
_S_ − _T_ disconnecting vertex
set/minimum-size/min-max 54-55

```
S − T path 54
s − t paths
arc-disjoint 55
internally disjoint 55
internally vertex-disjoint 55
S − T paths/min-max
disjoint 54-55
s − t paths/min-max
arc-disjoint 55
internally disjoint 55
internally vertex-disjoint 55
s − t vertex-cut 55
s − t vertex-cut/minimum-size 55
s − t vertex-cut/minimum-size/min-max
55
salesman problem
Christofides’ approximative
algorithm for the traveling
89-91
traveling 89-90, 108
salesman tour
traveling 89
SAT 103-105
3- 106
satisfiability problem 103-105
3- 106
scheduling
class 117-118
project 122
two-processor 83-84
SDR 42-43, 46, 188
common 43, 57-58, 70, 120-121,
185, 188
partial 43
seat assignment 83, 88
separate 23
separates pair
curve 161
ship routing 7-8, 73-74
shortest path 5-19, 91
shortest path problem 5-19
shrinking 81
simplex method 34
simplicial vertex 129
```

sink **157**
size of a word **99**
solves problem
algorithm **101**
Turing machine **109**
solves problem in polynomial time
Turing machine **109**
source **157**
spanning tree
minimum 19-22
spanning trees problem
edge-disjoint 185, 190
stable set [^39]:-40, 125-128
stable set number [^39]:-40, 111-112, 125-128
stable set polytope **143**
starting vertex **5**
Stiemke’s theorem 33
stops at word
algorithm **101**
storage of goods 113-114
strong perfect graph theorem [^113]:, **126**
subgraph
induced **113**
subject to capacity
flow **58**
subtrees of a tree 129-131
system of distinct representatives [^42]:-43,
46, 188
common [^43]:, 57-58, 70, 120-121,
185, 188
partial **43**

_T_-join **91**
terminal assignment 123
3-SAT **106**
3-satisfiability problem **106**
Thue system 100-101
tight subset **166**
timetabling
airline 88-89
totally unimodular matrix [^134]:-147
tour
traveling salesman **89**
transportation problem 49-50, 63, 65

```
minimum-cost 73-74
transversal 42-43, 46, 188
common 43, 57-58, 70, 120-121,
185, 188
partial 43
transversal matroid 181-182
traveling salesman problem 89-90, 108
Christofides’ approximative
algorithm for the 89-91
traveling salesman tour 89
tree
minimum spanning 19-22
rooted 5, 9, 185, 193
trees problem
edge-disjoint spanning 185, 190
vertex-disjoint 164-165
triangulated graph 128-131
truncation of a matroid 179
k- 179
TSP 108
Turing machine 100, 108-110
Tutte-Berge formula 78-79, 96
Tutte’s 1-factor theorem 79
two-commodity flow problem
integer 155-156
two-commodity flow theorem
Hu’s 153-155
2-heap 9-10
two-processor 84
two-processor scheduling 83-84
```

```
under capacity
flow 58
UNDIRECTED HAMILTONIAN CYCLE
107-108
undirected Hamiltonian cycle problem
107-108
undirected k-commodity flow problem
149-151
integer 149-151
undirected multicommodity flow problem
149-151
integer 149-151
uniform matroid 179
```

_k_- **179**
unimodular matrix [^136]:-137
union of matroids **189**

value of a flow **58**
vehicle routing 7-8
vertex-chromatic number [^111]:-115,
125-128
VERTEX-COLOURING [^111]:-112
vertex-colouring [^111]:-115
vertex-colouring number [^111]:-115,
125-128
vertex cover [^39]:-40
vertex cover number [^39]:-40
vertex-cover number 111
vertex cover polytope **143**
vertex-cut
_s_ − _t_ **55**
vertex-cut/minimum-size
_s_ − _t_ 55
vertex-cut/minimum-size/min-max
_s_ − _t_ 55
vertex-disjoint **54**
internally **54**
vertex-disjoint paths
internally **55**
vertex-disjoint paths problem [^149]:-152,
157-162
vertex-disjoint _s_ − _t_ paths
internally 55
vertex-disjoint _s_ − _t_ paths/min-max
internally 55
vertex-disjoint trees problem [^164]:-165
vertex of a convex set **25**
Vizing’s theorem [^116]:-117
VLSI-routing 151, 162-164

_W_ − _v_ walk **82**
walk **5**
_M_-alternating **81**
_s_ − _t_ **5**
_W_ − _v_ **82**
weighted bipartite matching 47-53
weighted bipartite matching algorithm
47-48

```
weighted bipartite matching problem 77
weighted common independent set
algorithm 191-193
weighted common independent set
augmenting algorithm 191-193
weighted common independent set
problem 190-193
weighted matching 47-53, 85-91
weighted matching problem 52
weighted matroid intersection 190-193
weighted matroid intersection algorithm
191-193
weighted nonbipartite matching 85-91
word 99
```

