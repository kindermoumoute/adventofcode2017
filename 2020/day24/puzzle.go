package main

import (
	"github.com/kindermoumoute/adventofcode/pkg/execute"
)

var tests = execute.TestCases{
	{
		`sesenwnenenewseeswwswswwnenewsewsw
neeenesenwnwwswnenewnwwsewnenwseswesw
seswneswswsenwwnwse
nwnwneseeswswnenewneswwnewseswneseene
swweswneswnenwsewnwneneseenw
eesenwseswswnenwswnwnwsewwnwsene
sewnenenenesenwsewnenwwwse
wenwwweseeeweswwwnwwe
wsweesenenewnwwnwsenewsenwwsesesenwne
neeswseenwwswnwswswnw
nenwswwsewswnenenewsenwsenwnesesenew
enewnwewneswsewnwswenweswnenwsenwsw
sweneswneswneneenwnewenewwneswswnese
swwesenesewenwneswnwwneseswwne
enesenwswwswneneswsenwnewswseenwsese
wnwnesenesenenwwnenwsewesewsesesew
nenewswnwewswnenesenwnesewesw
eneswnwswnwsenenwnwnwwseeswneewsenese
neswnwewnwnwseenwseesewsenwsweewe
wseweeenwnesenwwwswnew`,
		10,
		2208,
	}, {
		puzzle,
		228,
		3672,
	},
}

var puzzle = `nwswswnenwswsweswseswswnwswsweswsw
swseswsenwswnenwsesesenwseseesesenwenw
seswseswseswswseswwswneseswsese
nwwwsewenwwwwswwwe
seswwwseeseseseseseseswseswneneesew
eenewewweseneneneenene
nenenenenwnwwswseenwwneneenwneswnwnw
swneswwswwnewwwswwswwww
seseweewwnwwnwwsewwnwnenewnwne
nwwewewnwnwwwwnwswnwwsenwneww
sewnewswswwsenwsenenwneswswsesweenw
nwwnwnenwnwwnwsewnwseewsenwsesenwnwnw
enwnenwnwswnwnenesenwneswnenwnenwwnenese
neneneneneswneneenwenenenenenenewnwse
nwenwswsenwwnwwnwwnwwnwneww
wnwnwnwnwnenenwnesenwnwnwneswnwnenwsewne
wnwnwnwsesenwwwnenewnwnwnwnww
wnewwwswwwwwsw
nenwnenenewnenwnesenesenewnenenw
eeseeseeesenwesesese
seneseneneeeneewwswenweeneeneenene
wswsesenwseeeeseseseseseswnewseswnwnwse
swnenewwneneneneneneneseenenenenenee
swenwwwwewweswswwswswwswwnww
neswnenenenwswnwnwnwneneneneewnwneenw
eeesewswseenweewnwweeseene
swnwseeswneenwsesenwseeswnw
enwwnwwwswnwnwnenenenenwnenwnesesesenw
swwwwwwnewnewewwnewwsesw
eneseswwneseewnwnweswswwwnenwnenwne
sewsesesesenwnesewnweeswswswneseswwsw
nwneswnwnenenwnenweneswnesenwswnewnw
seneneseseseewsenenwswsewseswswsenesw
nwneseneenesewwnesenene
swwnwwsenwewswneewnweweswewnewse
nenenwseneswnwnwnwsenwwnwnenwnenenenw
neneneswneneenweesweneeeewenenw
swswswsenewseseseewswnesewswswseseswnw
eswwnwnwwswnwwwnwnwewwnw
swneseswenwseswsesesenwneseseseseswsese
nesenesenewnenenenwsenewnewsenwwnwne
swswseswseseswswnwseswsenesweswsw
swnwnwwswwswswswswwswsweneseswswnewne
nenenenenenwswneneswsenwnesenenenwnenenwnee
newnwneneeswneswnenwwsenenwneneenwnenw
swneeesesenesesewseseeesewse
nenwnwswseeenwseswneenewwene
enenenenenenenwnenesw
nwwswwwseneewwwwnwwwswnenweww
sewseneeesenwesesesesewneswseeew
swnwnweneswneneseswnwnenenenenenenwnese
nwsenwneseswwewnweswneewnwnwnwenw
wsenwwnwwnwenwwwnweswnw
swswwewnwnenesesewsesewneesewnene
neswnwwseswenenenwneseneeneenewsesww
nwwwsesenesesesesesenwseneeswnwsenewse
wnwswneswswswsweswne
enesenweneswsewneenenenwswnewnwnwsw
enenwnwswsenenwnwswnwwnwnwnwnwwnwwnw
wswswswwswwsesenwnwswnewsewswswnese
swwswneswswswneswswswsw
swswseswsesesenwseeeswseseswnwseseswnewsw
nwseeseseseswwesesesesesesesenesesenwe
swnwewewwwnwnwwnwnenwwwnwsewnw
swswswswnwswsenwswwweswwswsw
swenwswnesenwswesweseswwnewswenwswswnw
enewswnwwwwswswneewneewsw
nwsewnenwnwwnwnwnwenwnwnwnwnwswsesenwnw
wneeseseswseseesenweeseenwsenesewwne
senenwnwnenwnwnwnewnwnwnwnwenwnwswnesw
seseseseeseseseneseewse
neneeeneewweeneeneesesw
swswnwseswswwneseswswswswnwneswseswwswe
wsenenwwwseseeseswsewneneneneseswse
seeseeneeeswswseseenwneewneeseese
eeseeeeswsenweweenwneeewseene
wnenewnenenenesesenewnwnwnenenenwsenwne
neneneeswswnenenenenwenwnenwneneneswne
nwnwnwnwswwwnwnwwneenwwwswnwnwsenw
eneweneeneneenesweneneneesweswnww
neseseseseswsesesenwseseswseseseseswnew
seseseswswseseswsewne
swnwnewsweswswweswswswswwswswsesenwswsw
neewnwswnwswswewswnenwseswnwswswswsesw
esewnewwwwwswwwwwnwseewwwne
swswswnwswesweswswsweswswswswsewswnwsw
swneswswwnwswswswwswneswsweewswswsw
swwwesesewwswseneseneseseseeswsesesw
nwswnwnwnweswwnwnwenwnwnwnwe
nwwwnenwseseswnewsenesesewneswswwwe
nwnenenenwsenwwnenwnewseneewse
wwnwswnewwwwnw
swneneswnesenenenewnesenwewnewnenene
nenwneeneneesenwnesenewsw
neseseeewswweesesese
seenwsesesenwseeseseseseesesese
wwwwwswnewswswwwwnwswnewesw
swswwswswseswsenwwwseneeneswseswnwneese
nesenweseswesewsenwwsesesesesesee
neswseswsenwsweseneswswswseswswswswswwwnw
sesesewswseneswnwswswswswnwsesenwswesesw
nwwswesenewnwswswnwswwnwenenwne
eeeeneeweeeneneseeeenw
seseneseswwsewnwneseeeseneseswsenesw
nwnwneenwwnewewenwwnwnenwsesenenenw
neneneneswnenenwneneneeneneswneswswnene
wswseewneneseenewsesewwwswnwnwnwnwnw
nwnwnwnwnwnenwnwswnwnenenwsenwe
swnwnweswnwseeeeenwsesewswwnenesw
ewswseswnwwnwwnwnwsweweeenwnenw
wnewwnwneswsenwsewwenewnwnwwwww
nwnenenenewnwsenwnenwnwnesewnenwnwnwnwse
swsenewneswneneneenwnw
sweseswneeeneesenwnwneneenwnesenww
seswswneswswseswwseswnenwseseneswswswswse
nenwsenenwneneewneneneneneneswswene
wsewwwwwnewwwwwwwnw
swswseseeswswswwswsesesesewneneseswnw
nenweswnwnenenwnwnenwnewsewwnenenwe
neeewnenwnwnwnwnwsenwnwnwswwnwnwnwnw
swsenwswswneseswenewewsenwwnesewsenwne
wswwswwweeswswwswswnwnwnwene
nwneneneneseswnenwnesweswnenenwnenesene
wseseswswsewseneeswneewswseseswsene
swsenweeneswswsesesene
weswswsenewnwwewwswswenwsewnwnesw
neswnewsenwewnesenwsweseeseene
seswswseseswwneswseenweswswswswswnwnw
eseswswsweeswseswswsenwswswswswsewsenw
enewnwswsesweneswseswseenwsewnwsee
neseneswseweseenweeweesweeenew
wnwenesenwswnwnewsenwnwswnwnesewnwwese
swnwswswseeseseeneswswswswswswswswsww
wsesewwneseswswswseswnwswseeswneswewsw
nwnenenwnwenwswnwnenewnwnwswnwsesenwnww
eswseseswsewseseswsesesesenese
nwswswseeswswseswswsenweseewnwsesese
seseeeeenewswnwwnwswsenwnenwnwnwwsesw
neeweewnenwneenwnwsenwsenwsesesene
nenwseeseenwenwseeswswnweeeneeswse
swswseswswswswswenwwwswnwswneswswesww
swswneswnwswswswswswseewswewswwnene
nwneeneenwnewswnenwnenwswnenwnwnwenwne
nenenenenwnenesweneseneneneswnwnenenenw
ewseesweenweneseenwesweewwew
eseseeseseseeweesewseesesene
wneeneneenwweswenwwseneeenesesese
nwnwnenesenenenwwnwwnwsenwnwnwnenenwse
senewnwseesenwwseseswesewesesenenesw
seneeewnwwwsewnenwewnwswswse
neswwewwwwwwwswsweewwwww
wwsweswwswewwsenewwswwswswswnw
nwnwnwnwneneswnwsenwenwnwnwneenewnwnw
senwwnwneneeswneneeeweeenenenesw
nweneswneneswneenwnewnwnwnwnwnenenwnesenw
sewneswnweweswnesewnwwwwnewswsww
wneswswswwswsewswwswswswewswsw
senwneseswseeseswneseswneswswswswswsenw
ewwsweneenwsene
seeneeseeenewnwewswnwswswsesenww
wsesewwsewwswwnewswwnewwwwwnew
wnwsewwwnwwnewsenwnwnewnwnesenwnw
swswseseswsenewsesesesesesesesenwsenwse
nwwwwneswwwwwseswewwwwnenew
seewneeeeseenwsewenwsesw
nwnenwwneneseswnwew
neeenwenwesweeeeeeeseesw
nwnwwswewswnwnenwnwnwnwsenene
ewnwsesesesenwswesesesewswseesesese
neenwwsewsenwseeeeewwswneeneswne
swsenwseeswsewnwweneseneseeswsesesee
swsenwnweseswseenwseswseseseenwwsesene
eseenenenwneswnwnwneeeeneseneenew
swnenesweeeswneswenenwswneeeeseswse
sewsesesweeeeweeneeneseswesee
swwwwwwewnwnwwwewwwwnw
seswswenenwnweeneneneneswwenenenenenwne
seseeeseswwneeesenww
nwnwnwswneenwnenwneneswnwswnwsw
swswswswseswneswswswswnwseswsww
wnwwnwwsenwnwwneenwnwnwsesenwnwenww
eseeseeesewseseswesesenwnenwesese
swswnenwnenenwnenwenenenwenenwnwswnwnwswne
newwsewwwnwwnewsew
swswswseseswwseswswneswneneswnwnwswsesw
eeeesenwswseeeeewwseeneeeee
neswswseswneeswswneswswseswswww
neseeesesesewnwseswsesesewseseswwnwsee
nwnwnesenenwnwsesweswswenwnwnwewnwwse
eswswswswswseeeswwseewsenenwwsesww
newwnesesesweeneewwenee
swnwswwnwswseeneeseeneswswenwswwnesw
sesesweeseswseeseseenweseeswnesenenw
wwwwwwseewwwwwwwwswwnesene
swnwnwnenwesweswswsenwwneneenwnwnww
swnwnwewenwenweenwnwswwwnwnwnwww
wsweswswswswnwswswsweswswswswswswswne
neneneswwnwnenwnwneswneenwsesenene
swewswseneswnwswseswswseseenwswswsee
enweeesweseseweeeeeenwswee
sewseseswnenesewenwneeneneneewnewnew
esenwswewneswneswnwnenwnwnwenwenwnw
swneswwswwswswswwwwenwswswsweswsw
newseswseswnwseseneeswsenwwswnw
wwesweewswswnwwewewwwenwnesenw
seswenwseswenwwseewnenenwnesenwsese
wwwwnewwwswswww
nwwnwnwswnesenenenwwenwenwwnwnwsenwwnw
nwsewesesewswsesweseseneseneseswnese
eseeeeenwseeseswe
swseenwseswwneneswsweseswseswnwwnene
eenwewesweseneweswnenwswsewenee
wewwwwwwwwwnewsewwwwsenw
wwwsewwwwneww
senwneneenenwsenwnesenenenenewnewsesw
eweeeeeeee
nwnewneneneseneneenewseneneswsenenwnenw
swsewswnwneeenwseweseenwwwenesww
swswnwswneneswsesewswseswsewswnesese
newwnwnwwwwnwswswnenwnwneenwnwswnw
eenwsenwnewwneneenenwnenwsewwwsee
eseesenweeseseswseenesenweeewswnwe
nwenwenewnwwwnwswwsweswseswnwwe
nwseneeswswseseseswswswseswseeewswnw
seeeenweeweeeeesweeenee
eneweswnesenenewnenewneswnenesesww
nwwnewswnwwwnwe
swswswswswseswswswswwswswnesw
wwwsewnenwwwwnwsewnw
neeeseseesewseseesesewsesesene
nwnwswswneswwswseswswswsweswswseneswswsw
sesesenewsewseseswseseseseswnesesese
seneseswswseswseeeseenwnwseenesewnwse
nwwnwwwwwnwwewe
swseswwsweswswswnwswsweswswswswneswse
swseeswswswswswswswwnwnwswwnesesenwswsw
senwseseswsesenesenenwsesesewsesesesese
ewswwswnwseneswseswnenwswsenwseseswswne
sweswseswseswwwseneneenesenwseswnwsenw
nwwweswnwswnwswswseeseswsenenwsesww
sesewseesenwesweseeneseeneseswnewese
esewseeeesewenwseeweenesewew
eneseneswnesewwneenwneneswewsewnesw
newneweeseenesweneeeseenwneeswe
seswseseneswswswswswswsw
wsesesesesenesesenwseesenwsesesesesese
nenenenesenenewnenenene
seswwnwswwsesesweswsweswenesenwene
neswesenenesenenweeswseeenwnewnwnene
eeeeeeeenweeneswenenwnweswsw
seneenenesweneswneewnwswnene
swswwswswswwswswnwwwnweenwwseswww
swswwswneswseneswswswsenesw
swwswswswswswsweswswswswsw
seneseseseenwnwsweswnwseseswseswwneswswsw
nenwnwswnenwewneswnwnewnenenesenenee
nwseseesenwwswseeenesesesesenwnwew
swnwswseenewswwswswswwwwwnwswwe
nwneenesenwneewnewneseeneneesw
neswnwwnenweswneneneneswnenwwneenwneene
neeneneswwsenwnwneneneseenenenewsene
swswnewseseseswswseswneswnwseeseswsese
swnwwswseswwnwsweeswwnweenewnwsw
nenweeeeswneweseswneswne
weneenwnenenenenwnenwswnewneeneneenesw
swnenwnwwwswwwnenewwnwsewnesewnw
neswswwswswseesweneswswsewswswwswse
nenenenwwsenwnwnenwnwnenwsenenwnwnwsenw
nwseseswseesenwseswswsenwese
wswneswnwneneswseseswwnwseswswnesenwsesew
weeeewsweeeneneneswesenwswee
senwnwwwnwsenwnwwwnwsesenesesenewsene
wwwesewwsenwwwwswweswwswnenww
swwneneseswswwswsewwwsw
enenwneneeneeseneswnwneeneswnesenene
eseesenenenewwneneseenenewnenenewnene
wenwswewnwewnweswseeeneeneenesw
neneneneswnenewsenwneseneenenwneswnenewne
eneswneneeseeswenewneeeeeenwnee
eswswsenwneseeseeeeeswnweeeewne
nwewwwwenwwwwewnwwnwwewsw
nwwwwnwnesenwwseenwswnwewnwswwnewnw
seweeweeenwseeeeenwseeeene
wswsenenewwwwwwwswwswwwneseww
nenenenweeseneneneeeeeweewswse
neneneneenweseeneswswnenenweeneswee
sesweeeswenwnese
enweseneneneneenwneenwneswneswneswnew
senenweseneneswsesenwseswswnwwswnwsewse
nwnwnwenwnwnwnwnwnwnwswnwnwnwsesenwnwnw
wwwnwnwwenwnwnwnwsenwswww
nwnwnweswnenenwnesenwwswnenenenwnesenwe`
