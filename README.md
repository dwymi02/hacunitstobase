# hacunitstobase

### Package hacunitstobase converts Hacash amounts from [:nnn] to base [:248]

In Hacash numbers are represented as follows:
:248 is the base, everything else is a fractional equivalent

When describing the Hacash number system, numbers are represented as follows:
[100:248] -- Where [100:] represents the amount and [:248] represents units

In the examples provided, each corresponding [amount:unit] represents [100:248]

Theoretically, this representation could be expanded further

```
1000000000000000:235
100000000000000:236
10000000000000:237
1000000000000:238
100000000000:239
10000000000:240
1000000000:241
100000000:242
10000000:243
1000000:244
100000:245
10000:246
1000:247
**100:248**
10:249
1:250
0.1:251
0.01:252
0.001:253
0.0001:254
0.00001:255
0.000001:256
0.0000001:257
0.00000001:258
0.000000001:259
0.0000000001:260
0.00000000001:261
```



##### The package hacunitstobase is designed to take a non-base HAC unit converting the amount to a base HAC unit, viz., [nnn:248]

Within this repository there are two directories named [**test**] and [**convert_hacunitstobase**]

The code in [test] provides a simple test harness to ensure that package **hacunitstobase** is working properly

The code in [**convert_hacunitstobase**] is a command line interface (CLI)

Using the (CLI) the user can enter an amount and unit and the code will provide you with the HAC base equivalent

A couple of examples will be provided using a couple of amounts and units

Tests were performed on Linux and Go:

```
lsb_release -a 
No LSB modules are available. 
Distributor ID: Ubuntu 
Description:   Ubuntu 20.04.2 LTS 
Release:     20.04 
Codename:    focal

go version
go version go1.16.7 linux/amd64
```



#### Example #1:

Convert the amount 1134903170.0 unit 240 **1134903170.0:240**

From the command line enter the following command:

```
./convert_hacunitstobase -amount 1134903170.0 -unit 240         
amount: 1.13490317e+09 
unit: 240 

HAC 1134903170.000000:240 converted to 11.349032:248
```



#### Example #2:

Convert the amount 1134903170.0 unit 241 **1134903170.0:241**

From the command line enter the following command:

```
./convert_hacunitstobase -amount 1134903170.0 -unit 241 
amount: 1.13490317e+09 
unit: 241 

HAC 1134903170.000000:241 converted to 113.490317:248
```



#### Example #3:

Convert the amount 1134903170.0 unit 256 **1134903170.0:256**

From the command line enter the following command:

```
./convert_hacunitstobase -amount 1134903170.0 -unit 256 
amount: 1.13490317e+09 
unit: 256 

HAC 1134903170.000000:256 converted to 113490317000000000.000000:248
```
