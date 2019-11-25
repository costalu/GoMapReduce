### Group Members:

1. Andrew Stanley - 991431494
2. Fernando Costa - 991447247
3. Andrew Bordin - 991436444

### Sequential Map Reduce (language of choice: Golang)

*Output:* 
```
#nbafinals -> 40704
rt -> 31578
#whateverittakes -> 13744
@nba: -> 11340
oh -> 9918
he -> 9597
yes -> 8568
did! -> 8489
https://t.co/axun0gomok -> 8452
lebron -> 7238
@bleacherreport: -> 4683
????? -> 4562
game -> 4298
shook. -> 4159
https://t.co/9jvtjcpel0 -> 4152
???? -> 3543
3 -> 2657
?? -> 2612
james -> 2540
@kingjames -> 2455
???????? -> 2398
??? -> 2354
his -> 2291
de -> 2173
#dubnation -> 1879
@cavs -> 1816
i -> 1737
cavs -> 1511
| -> 1457
@arabic1_nba: -> 1392
warriors -> 1377
up -> 1374
que -> 1344
go -> 1327
#nbaonabc -> 1282
o -> 1260
?????? -> 1244
love -> 1200
you -> 1199
????????? -> 1137
king -> 1134
???????????????? -> 1082
kevin -> 1054
all -> 1045
https:/. -> 1039
: -> 1026
????????????? -> 1020
para -> 1008
@nbaontnt: -> 1007
james! -> 990
el -> 974
just -> 967
get -> 957
? -> 948
at -> 948
#nba -> 937
be -> 930
draymond -> 908
en -> 906
my, -> 895
https://t.co/zoehximfee -> 893
into -> 881
@nbatv: -> 866
#cavs -> 858
e -> 850
out -> 841
cavaliers -> 827
nba -> 821
do -> 821
taking -> 812
no -> 811
own -> 805
off -> 778
matters -> 759
first -> 759
hands! -> 756
https://t.co/yqz4pvpbwb -> 755
green -> 754
la -> 752
&amp; -> 738
after -> 737
3. -> 737
! -> 718
let's -> 709
rebounds -> 705
5th -> 696
#thisiswhyweplay -> 690
himself -> 689
- -> 687
cleveland -> 683
moving -> 682
congrats -> 662
http. -> 661
from -> 657
list! -> 651
@warriors -> 650
was -> 648
have -> 640
not -> 636
#nbanaespn -> 632
5 -> 625
my -> 618
```

*Time:* 465.346894ms

### Distributed/Parallel Map Reduce (language of choice: Golang)

Approach we took: We used go routines and channels.

*Output:*

```
#nbafinals -> 40704
rt -> 31578
#whateverittakes -> 13744
@nba: -> 11340
oh -> 9918
he -> 9597
yes -> 8568
did! -> 8489
https://t.co/axun0gomok -> 8452
lebron -> 7238
@bleacherreport: -> 4683
????? -> 4562
game -> 4298
shook. -> 4159
https://t.co/9jvtjcpel0 -> 4152
???? -> 3543
3 -> 2657
?? -> 2612
james -> 2540
@kingjames -> 2455
???????? -> 2398
??? -> 2354
his -> 2291
de -> 2173
#dubnation -> 1879
@cavs -> 1816
i -> 1737
cavs -> 1511
| -> 1457
@arabic1_nba: -> 1392
warriors -> 1377
up -> 1374
que -> 1344
go -> 1327
#nbaonabc -> 1282
o -> 1260
?????? -> 1244
love -> 1200
you -> 1199
????????? -> 1137
king -> 1134
???????????????? -> 1082
kevin -> 1054
all -> 1045
https:/. -> 1039
: -> 1026
????????????? -> 1020
para -> 1008
@nbaontnt: -> 1007
james! -> 990
el -> 974
just -> 967
get -> 957
? -> 948
at -> 948
#nba -> 937
be -> 930
draymond -> 908
en -> 906
my, -> 895
https://t.co/zoehximfee -> 893
into -> 881
@nbatv: -> 866
#cavs -> 858
e -> 850
out -> 841
cavaliers -> 827
nba -> 821
do -> 821
taking -> 812
no -> 811
own -> 805
off -> 778
matters -> 759
first -> 759
hands! -> 756
https://t.co/yqz4pvpbwb -> 755
green -> 754
la -> 752
&amp; -> 738
after -> 737
3. -> 737
! -> 718
let's -> 709
rebounds -> 705
5th -> 696
#thisiswhyweplay -> 690
himself -> 689
- -> 687
cleveland -> 683
moving -> 682
congrats -> 662
http. -> 661
from -> 657
list! -> 651
@warriors -> 650
was -> 648
have -> 640
not -> 636
#nbanaespn -> 632
5 -> 625
my -> 618
```

*Time:* 449.074827ms
  
