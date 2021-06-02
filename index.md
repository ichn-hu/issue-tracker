

If you are interested in database development, or you are a TiDB user, no matter what, if you want to contribute to TiDB and learn about how a distributed HTAP database worked, here is the right place.

Let's get started by solving some bugs! Here is a curated list of some easy-to-go bugs, pick the one that you want to smash!

* [sig/planner](#sig/planner)
* [sig/execution](#sig/execution)
* [sig/transaction](#sig/transaction)
* [sig/DDL](#sig/DDL)

Note: currently the issues are classified by their SIG owners, such as sig/planner and sig/execution, which stands for special interests groups that focus on SQL planning and SQL execution, to know more about TiDB community, see [the community repository](https://github.com/pingcap/community). We also host discussions on slack, if you are not in the corresponding slack channel, we highly recommend you to join so that you could ask questions and get responses immediately from these SIGs members. [Join TiDB Community slack workspace now!](https://join.slack.com/t/tidbcommunity/shared_invite/enQtNzc0MzI4ODExMDc4LWYwYmIzMjZkYzJiNDUxMmZlN2FiMGJkZjAyMzQ5NGU0NGY0NzI3NTYwMjAyNGQ1N2I2ZjAxNzc1OGUwYWM0NzE)

If there is a &#x2757; before the issue link, it means there is no one assigned, nor a PR linked, nor picked, and it is for the maintainers to track the progress of each issue, it is also a notation of "welcome to take a look".

Feel free to comment on issues that interest you, and ask whatever questions you have on how to get started working on them!

<h2 name="sig/planner">sig/planner</h2>

|                             ISSUE                              | PRIORITY |                              ASSIGNEE                               |                          PR                          | HINT |
|----------------------------------------------------------------|----------|---------------------------------------------------------------------|------------------------------------------------------|------|
| [#8190](https://github.com/pingcap/tidb/issues/8190)           | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#9373](https://github.com/pingcap/tidb/issues/9373)           | moderate | @morgo                                                              |                                                      |      |
| [#10151](https://github.com/pingcap/tidb/issues/10151)         | moderate | <sub>@wjhuang2016</sub>                                             |                                                      |      |
| [#12182](https://github.com/pingcap/tidb/issues/12182)&#x2757; | moderate |                                                                     |                                                      |      |
| [#13167](https://github.com/pingcap/tidb/issues/13167)         | moderate | @fzhedu                                                             |                                                      |      |
| [#13856](https://github.com/pingcap/tidb/issues/13856)         | moderate | @winoros                                                            |                                                      |      |
| [#15514](https://github.com/pingcap/tidb/issues/15514)&#x2757; | minor    |                                                                     |                                                      |      |
| [#16473](https://github.com/pingcap/tidb/issues/16473)         | minor    | @winoros                                                            |                                                      |      |
| [#16764](https://github.com/pingcap/tidb/issues/16764)         | moderate | @lzmhhh123                                                          |                                                      |      |
| [#16788](https://github.com/pingcap/tidb/issues/16788)         | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#16909](https://github.com/pingcap/tidb/issues/16909)         | moderate | @lzmhhh123                                                          |                                                      |      |
| [#17731](https://github.com/pingcap/tidb/issues/17731)         | moderate |                                                                     | [#22416](https://github.com/pingcap/tidb/pull/22416) |      |
| [#17852](https://github.com/pingcap/tidb/issues/17852)         | moderate | @XuHuaiyu                                                           |                                                      |      |
| [#18216](https://github.com/pingcap/tidb/issues/18216)         | moderate | @fzhedu                                                             |                                                      |      |
| [#19713](https://github.com/pingcap/tidb/issues/19713)         | minor    | @winoros                                                            |                                                      |      |
| [#19802](https://github.com/pingcap/tidb/issues/19802)&#x2757; | minor    |                                                                     |                                                      |      |
| [#20019](https://github.com/pingcap/tidb/issues/20019)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21098](https://github.com/pingcap/tidb/issues/21098)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21454](https://github.com/pingcap/tidb/issues/21454)&#x2757; | minor    |                                                                     |                                                      |      |
| [#21475](https://github.com/pingcap/tidb/issues/21475)&#x2757; | moderate |                                                                     |                                                      |      |
| [#21625](https://github.com/pingcap/tidb/issues/21625)&#x2757; | minor    |                                                                     |                                                      |      |
| [#21677](https://github.com/pingcap/tidb/issues/21677)         | major    | <sub><sup>@xuyifangreeneyes</sup></sub>                             |                                                      |      |
| [#22082](https://github.com/pingcap/tidb/issues/22082)&#x2757; | moderate |                                                                     |                                                      |      |
| [#22096](https://github.com/pingcap/tidb/issues/22096)&#x2757; | moderate |                                                                     |                                                      |      |
| [#22301](https://github.com/pingcap/tidb/issues/22301)         | minor    | <sub>@Reminiscent</sub></br><sub><sup>@xuyifangreeneyes</sup></sub> |                                                      |      |
| [#22535](https://github.com/pingcap/tidb/issues/22535)&#x2757; | minor    |                                                                     |                                                      |      |
| [#22694](https://github.com/pingcap/tidb/issues/22694)         | moderate | <sub><sup>@xuyifangreeneyes</sup></sub>                             | [#22722](https://github.com/pingcap/tidb/pull/22722) |      |
| [#22934](https://github.com/pingcap/tidb/issues/22934)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23391](https://github.com/pingcap/tidb/issues/23391)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23396](https://github.com/pingcap/tidb/issues/23396)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23459](https://github.com/pingcap/tidb/issues/23459)&#x2757; | minor    |                                                                     |                                                      |      |
| [#23499](https://github.com/pingcap/tidb/issues/23499)         | major    | @winoros                                                            |                                                      |      |
| [#23506](https://github.com/pingcap/tidb/issues/23506)         | major    | <sub>@xiongjiwei</sub>                                              | [#23760](https://github.com/pingcap/tidb/pull/23760) |      |
| [#23539](https://github.com/pingcap/tidb/issues/23539)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23582](https://github.com/pingcap/tidb/issues/23582)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23633](https://github.com/pingcap/tidb/issues/23633)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23659](https://github.com/pingcap/tidb/issues/23659)&#x2757; | moderate |                                                                     |                                                      |      |
| [#23802](https://github.com/pingcap/tidb/issues/23802)         | major    | @Howie59                                                            | [#23849](https://github.com/pingcap/tidb/pull/23849) |      |
| [#23907](https://github.com/pingcap/tidb/issues/23907)&#x2757; | major    |                                                                     |                                                      |      |
| [#23965](https://github.com/pingcap/tidb/issues/23965)&#x2757; | moderate |                                                                     |                                                      |      |
| [#24010](https://github.com/pingcap/tidb/issues/24010)         | moderate | @rebelice                                                           |                                                      |      |
| [#24013](https://github.com/pingcap/tidb/issues/24013)         | moderate | @rebelice                                                           |                                                      |      |
| [#24014](https://github.com/pingcap/tidb/issues/24014)         | moderate | @rebelice                                                           |                                                      |      |
| [#24015](https://github.com/pingcap/tidb/issues/24015)         | moderate | @rebelice                                                           |                                                      |      |
| [#24057](https://github.com/pingcap/tidb/issues/24057)         | moderate | @winoros                                                            |                                                      |      |
| [#24309](https://github.com/pingcap/tidb/issues/24309)         | moderate |                                                                     | [#24913](https://github.com/pingcap/tidb/pull/24913) |      |
| [#24318](https://github.com/pingcap/tidb/issues/24318)         | major    | @qw4990                                                             |                                                      |      |
| [#24323](https://github.com/pingcap/tidb/issues/24323)         | minor    | @sylzd                                                              |                                                      |      |
| [#24324](https://github.com/pingcap/tidb/issues/24324)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#24449](https://github.com/pingcap/tidb/issues/24449)         | minor    | @zoomxi                                                             | [#24921](https://github.com/pingcap/tidb/pull/24921) |      |
| [#24452](https://github.com/pingcap/tidb/issues/24452)         | moderate | @winoros                                                            |                                                      |      |
| [#24473](https://github.com/pingcap/tidb/issues/24473)         | critical | @sylzd                                                              | [#24775](https://github.com/pingcap/tidb/pull/24775) |      |
| [#24512](https://github.com/pingcap/tidb/issues/24512)&#x2757; | moderate |                                                                     |                                                      |      |
| [#24526](https://github.com/pingcap/tidb/issues/24526)         | major    | @winoros                                                            |                                                      |      |
| [#24550](https://github.com/pingcap/tidb/issues/24550)         | moderate | @eurekaka                                                           | [#24663](https://github.com/pingcap/tidb/pull/24663) |      |
| [#24563](https://github.com/pingcap/tidb/issues/24563)         | major    | @eurekaka                                                           | [#24663](https://github.com/pingcap/tidb/pull/24663) |      |
| [#24567](https://github.com/pingcap/tidb/issues/24567)&#x2757; | major    |                                                                     |                                                      |      |
| [#24594](https://github.com/pingcap/tidb/issues/24594)         | major    | @winoros                                                            |                                                      |      |
| [#24622](https://github.com/pingcap/tidb/issues/24622)&#x2757; | moderate |                                                                     |                                                      |      |
| [#24636](https://github.com/pingcap/tidb/issues/24636)         | critical | @mjonss</br>@qw4990                                                 |                                                      |      |
| [#24667](https://github.com/pingcap/tidb/issues/24667)         | moderate | <sub><sup>@time-and-fate</sup></sub>                                | [#24669](https://github.com/pingcap/tidb/pull/24669) |      |
| [#24679](https://github.com/pingcap/tidb/issues/24679)         | moderate | @rebelice                                                           |                                                      |      |
| [#24706](https://github.com/pingcap/tidb/issues/24706)         | major    | @lysu                                                               | [#24828](https://github.com/pingcap/tidb/pull/24828) |      |
| [#24722](https://github.com/pingcap/tidb/issues/24722)         | moderate | @eurekaka                                                           | [#24753](https://github.com/pingcap/tidb/pull/24753) |      |
| [#24743](https://github.com/pingcap/tidb/issues/24743)         | moderate | @eurekaka                                                           | [#24793](https://github.com/pingcap/tidb/pull/24793) |      |
| [#24855](https://github.com/pingcap/tidb/issues/24855)         | major    | <sub><sup>@time-and-fate</sup></sub>                                |                                                      |      |
| [#24865](https://github.com/pingcap/tidb/issues/24865)         | critical | @eurekaka                                                           | [#24937](https://github.com/pingcap/tidb/pull/24937) |      |
| [#24893](https://github.com/pingcap/tidb/issues/24893)         | critical | @winoros                                                            |                                                      |      |
| [#24954](https://github.com/pingcap/tidb/issues/24954)         | major    | @eurekaka                                                           | [#24994](https://github.com/pingcap/tidb/pull/24994) |      |
| [#25010](https://github.com/pingcap/tidb/issues/25010)&#x2757; | major    |                                                                     |                                                      |      |
| [#25013](https://github.com/pingcap/tidb/issues/25013)&#x2757; | major    |                                                                     |                                                      |      |
| [#25014](https://github.com/pingcap/tidb/issues/25014)&#x2757; | major    |                                                                     |                                                      |      |
| [#25041](https://github.com/pingcap/tidb/issues/25041)         | moderate | @eurekaka                                                           |                                                      |      |
| [#25043](https://github.com/pingcap/tidb/issues/25043)&#x2757; | major    |                                                                     |                                                      |      |
| [#25045](https://github.com/pingcap/tidb/issues/25045)         | major    | @zhuo-zhi                                                           |                                                      |      |


<h2 name="sig/execution">sig/execution</h2>

|                             ISSUE                              | PRIORITY |        ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|----------|-------------------------|------------------------------------------------------|------|
| [#8205](https://github.com/pingcap/tidb/issues/8205)&#x2757;   | minor    |                         |                                                      |      |
| [#11866](https://github.com/pingcap/tidb/issues/11866)         | moderate | @dragonly               |                                                      |      |
| [#11932](https://github.com/pingcap/tidb/issues/11932)         | moderate | @wshwsh12               |                                                      |      |
| [#13018](https://github.com/pingcap/tidb/issues/13018)         | minor    | @qw4990                 |                                                      |      |
| [#13157](https://github.com/pingcap/tidb/issues/13157)         | moderate | <sub>@wjhuang2016</sub> |                                                      |      |
| [#13440](https://github.com/pingcap/tidb/issues/13440)         | minor    | <sub>@SunRunAway</sub>  |                                                      |      |
| [#14399](https://github.com/pingcap/tidb/issues/14399)&#x2757; | moderate |                         |                                                      |      |
| [#15234](https://github.com/pingcap/tidb/issues/15234)         | moderate |                         | [#20015](https://github.com/pingcap/tidb/pull/20015) |      |
| [#15608](https://github.com/pingcap/tidb/issues/15608)         | moderate | @ichn-hu                |                                                      |      |
| [#15884](https://github.com/pingcap/tidb/issues/15884)&#x2757; | moderate |                         |                                                      |      |
| [#17083](https://github.com/pingcap/tidb/issues/17083)         | moderate | @mmyj</br>@qw4990       |                                                      |      |
| [#17489](https://github.com/pingcap/tidb/issues/17489)&#x2757; | minor    |                         |                                                      |      |
| [#17677](https://github.com/pingcap/tidb/issues/17677)         | moderate | @qw4990                 |                                                      |      |
| [#17751](https://github.com/pingcap/tidb/issues/17751)&#x2757; | minor    |                         |                                                      |      |
| [#17832](https://github.com/pingcap/tidb/issues/17832)         | minor    | @qw4990                 |                                                      |      |
| [#17993](https://github.com/pingcap/tidb/issues/17993)&#x2757; | minor    |                         |                                                      |      |
| [#18488](https://github.com/pingcap/tidb/issues/18488)         | moderate | @morgo                  | [#19029](https://github.com/pingcap/tidb/pull/19029) |      |
| [#18493](https://github.com/pingcap/tidb/issues/18493)         | minor    | @fzhedu                 |                                                      |      |
| [#19025](https://github.com/pingcap/tidb/issues/19025)         | moderate | <sub>@SunRunAway</sub>  | [#19029](https://github.com/pingcap/tidb/pull/19029) |      |
| [#20411](https://github.com/pingcap/tidb/issues/20411)         | moderate | @qw4990                 |                                                      |      |
| [#20563](https://github.com/pingcap/tidb/issues/20563)         | moderate | <sub>@wjhuang2016</sub> | [#1129](https://github.com/pingcap/parser/pull/1129) |      |
| [#21307](https://github.com/pingcap/tidb/issues/21307)&#x2757; | minor    |                         |                                                      |      |
| [#21374](https://github.com/pingcap/tidb/issues/21374)         | minor    |                         | [#21508](https://github.com/pingcap/tidb/pull/21508) |      |
| [#21584](https://github.com/pingcap/tidb/issues/21584)&#x2757; | major    |                         |                                                      |      |
| [#21653](https://github.com/pingcap/tidb/issues/21653)         | moderate |                         | [#21230](https://github.com/pingcap/tidb/pull/21230) |      |
| [#21787](https://github.com/pingcap/tidb/issues/21787)&#x2757; | minor    |                         |                                                      |      |
| [#22088](https://github.com/pingcap/tidb/issues/22088)&#x2757; | moderate |                         |                                                      |      |
| [#22132](https://github.com/pingcap/tidb/issues/22132)         | major    | @wshwsh12               | [#22347](https://github.com/pingcap/tidb/pull/22347) |      |
| [#22206](https://github.com/pingcap/tidb/issues/22206)         | moderate |                         | [#22616](https://github.com/pingcap/tidb/pull/22616) |      |
| [#22227](https://github.com/pingcap/tidb/issues/22227)&#x2757; | minor    |                         |                                                      |      |
| [#22386](https://github.com/pingcap/tidb/issues/22386)&#x2757; | moderate |                         |                                                      |      |
| [#22394](https://github.com/pingcap/tidb/issues/22394)         | moderate |                         | [#22407](https://github.com/pingcap/tidb/pull/22407) |      |
| [#22399](https://github.com/pingcap/tidb/issues/22399)&#x2757; | minor    |                         |                                                      |      |
| [#22423](https://github.com/pingcap/tidb/issues/22423)&#x2757; | minor    |                         |                                                      |      |
| [#22525](https://github.com/pingcap/tidb/issues/22525)         | major    | @lzmhhh123              |                                                      |      |
| [#22592](https://github.com/pingcap/tidb/issues/22592)&#x2757; | major    |                         |                                                      |      |
| [#22598](https://github.com/pingcap/tidb/issues/22598)&#x2757; | moderate |                         |                                                      |      |
| [#22604](https://github.com/pingcap/tidb/issues/22604)&#x2757; | major    |                         |                                                      |      |
| [#22665](https://github.com/pingcap/tidb/issues/22665)         | major    |                         | [#22666](https://github.com/pingcap/tidb/pull/22666) |      |
| [#22703](https://github.com/pingcap/tidb/issues/22703)         | minor    | @xhebox                 | [#23519](https://github.com/pingcap/tidb/pull/23519) |      |
| [#22735](https://github.com/pingcap/tidb/issues/22735)&#x2757; | critical |                         |                                                      |      |
| [#22749](https://github.com/pingcap/tidb/issues/22749)&#x2757; | minor    |                         |                                                      |      |
| [#22791](https://github.com/pingcap/tidb/issues/22791)         | major    | @johan-j                | [#22823](https://github.com/pingcap/tidb/pull/22823) |      |
| [#22975](https://github.com/pingcap/tidb/issues/22975)         | major    | <sub>@hanfei1991</sub>  |                                                      |      |
| [#23101](https://github.com/pingcap/tidb/issues/23101)&#x2757; | moderate |                         |                                                      |      |
| [#23110](https://github.com/pingcap/tidb/issues/23110)         | moderate | @jyz0309                |                                                      |      |
| [#23159](https://github.com/pingcap/tidb/issues/23159)         | moderate |                         | [#23206](https://github.com/pingcap/tidb/pull/23206) |      |
| [#23344](https://github.com/pingcap/tidb/issues/23344)&#x2757; | moderate |                         |                                                      |      |
| [#23366](https://github.com/pingcap/tidb/issues/23366)         | minor    | @wshwsh12               |                                                      |      |
| [#23387](https://github.com/pingcap/tidb/issues/23387)         | major    | <sub>@tiancaiamao</sub> | [#23403](https://github.com/pingcap/tidb/pull/23403) |      |
| [#23411](https://github.com/pingcap/tidb/issues/23411)         | moderate | @wshwsh12               |                                                      |      |
| [#23501](https://github.com/pingcap/tidb/issues/23501)&#x2757; | moderate |                         |                                                      |      |
| [#23508](https://github.com/pingcap/tidb/issues/23508)         | major    | <sub>@wjhuang2016</sub> | [#23559](https://github.com/pingcap/tidb/pull/23559) |      |
| [#23531](https://github.com/pingcap/tidb/issues/23531)&#x2757; | moderate |                         |                                                      |      |
| [#23865](https://github.com/pingcap/tidb/issues/23865)         | moderate | @xhebox                 |                                                      |      |
| [#23869](https://github.com/pingcap/tidb/issues/23869)         | moderate | @ichn-hu                |                                                      |      |
| [#23897](https://github.com/pingcap/tidb/issues/23897)         | major    | @wshwsh12               | [#24179](https://github.com/pingcap/tidb/pull/24179) |      |
| [#23898](https://github.com/pingcap/tidb/issues/23898)         | major    | @lzmhhh123              |                                                      |      |
| [#23952](https://github.com/pingcap/tidb/issues/23952)&#x2757; | moderate |                         |                                                      |      |
| [#24044](https://github.com/pingcap/tidb/issues/24044)&#x2757; | minor    |                         |                                                      |      |
| [#24134](https://github.com/pingcap/tidb/issues/24134)&#x2757; | moderate |                         |                                                      |      |
| [#24227](https://github.com/pingcap/tidb/issues/24227)&#x2757; | moderate |                         |                                                      |      |
| [#24271](https://github.com/pingcap/tidb/issues/24271)&#x2757; | minor    |                         |                                                      |      |
| [#24284](https://github.com/pingcap/tidb/issues/24284)&#x2757; | minor    |                         |                                                      |      |
| [#24319](https://github.com/pingcap/tidb/issues/24319)&#x2757; | minor    |                         |                                                      |      |
| [#24387](https://github.com/pingcap/tidb/issues/24387)&#x2757; | moderate |                         |                                                      |      |
| [#24502](https://github.com/pingcap/tidb/issues/24502)         | major    | <sub>@xiongjiwei</sub>  |                                                      |      |
| [#24627](https://github.com/pingcap/tidb/issues/24627)&#x2757; | major    |                         |                                                      |      |
| [#24690](https://github.com/pingcap/tidb/issues/24690)         | minor    |                         | [#24691](https://github.com/pingcap/tidb/pull/24691) |      |
| [#24725](https://github.com/pingcap/tidb/issues/24725)&#x2757; | moderate |                         |                                                      |      |
| [#24756](https://github.com/pingcap/tidb/issues/24756)&#x2757; |          |                         |                                                      |      |
| [#24761](https://github.com/pingcap/tidb/issues/24761)         | major    | @wshwsh12               |                                                      |      |
| [#24768](https://github.com/pingcap/tidb/issues/24768)&#x2757; | major    |                         |                                                      |      |
| [#24771](https://github.com/pingcap/tidb/issues/24771)         | major    | @lzmhhh123              |                                                      |      |
| [#24819](https://github.com/pingcap/tidb/issues/24819)&#x2757; | major    |                         |                                                      |      |
| [#24868](https://github.com/pingcap/tidb/issues/24868)         | major    | @XuHuaiyu               |                                                      |      |
| [#24888](https://github.com/pingcap/tidb/issues/24888)         | minor    |                         | [#24889](https://github.com/pingcap/tidb/pull/24889) |      |
| [#24917](https://github.com/pingcap/tidb/issues/24917)&#x2757; | moderate |                         |                                                      |      |
| [#24933](https://github.com/pingcap/tidb/issues/24933)         | critical | @XuHuaiyu               |                                                      |      |
| [#24944](https://github.com/pingcap/tidb/issues/24944)&#x2757; | critical |                         |                                                      |      |
| [#24949](https://github.com/pingcap/tidb/issues/24949)         | major    | @ichn-hu                |                                                      |      |
| [#24951](https://github.com/pingcap/tidb/issues/24951)&#x2757; | moderate |                         |                                                      |      |
| [#24953](https://github.com/pingcap/tidb/issues/24953)&#x2757; | moderate |                         |                                                      |      |
| [#24955](https://github.com/pingcap/tidb/issues/24955)&#x2757; | moderate |                         |                                                      |      |
| [#24967](https://github.com/pingcap/tidb/issues/24967)&#x2757; | major    |                         |                                                      |      |
| [#24969](https://github.com/pingcap/tidb/issues/24969)&#x2757; | moderate |                         |                                                      |      |
| [#24972](https://github.com/pingcap/tidb/issues/24972)&#x2757; | major    |                         |                                                      |      |
| [#24997](https://github.com/pingcap/tidb/issues/24997)&#x2757; |          |                         |                                                      |      |
| [#25012](https://github.com/pingcap/tidb/issues/25012)&#x2757; | major    |                         |                                                      |      |
| [#25038](https://github.com/pingcap/tidb/issues/25038)         | major    | @ichn-hu                |                                                      |      |
| [#25053](https://github.com/pingcap/tidb/issues/25053)&#x2757; | moderate |                         |                                                      |      |


<h2 name="sig/transaction">sig/transaction</h2>

|                             ISSUE                              | PRIORITY |        ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|----------|-------------------------|------------------------------------------------------|------|
| [#9762](https://github.com/pingcap/tidb/issues/9762)&#x2757;   | major    |                         |                                                      |      |
| [#10657](https://github.com/pingcap/tidb/issues/10657)         | major    | @nolouch                |                                                      |      |
| [#13958](https://github.com/pingcap/tidb/issues/13958)         | moderate | @fzhedu                 |                                                      |      |
| [#14914](https://github.com/pingcap/tidb/issues/14914)         | minor    | <sub>@tiancaiamao</sub> |                                                      |      |
| [#17797](https://github.com/pingcap/tidb/issues/17797)&#x2757; | moderate |                         |                                                      |      |
| [#17798](https://github.com/pingcap/tidb/issues/17798)&#x2757; | moderate |                         |                                                      |      |
| [#18048](https://github.com/pingcap/tidb/issues/18048)         | moderate | @qw4990                 |                                                      |      |
| [#20028](https://github.com/pingcap/tidb/issues/20028)         | major    | <sub>@tiancaiamao</sub> | [#21148](https://github.com/pingcap/tidb/pull/21148) |      |
| [#20949](https://github.com/pingcap/tidb/issues/20949)&#x2757; | minor    |                         |                                                      |      |
| [#20990](https://github.com/pingcap/tidb/issues/20990)&#x2757; | moderate |                         |                                                      |      |
| [#21335](https://github.com/pingcap/tidb/issues/21335)         | minor    | @you06                  | [#22146](https://github.com/pingcap/tidb/pull/22146) |      |
| [#21355](https://github.com/pingcap/tidb/issues/21355)&#x2757; | moderate |                         |                                                      |      |
| [#21506](https://github.com/pingcap/tidb/issues/21506)         | moderate | @cfzjywxk               |                                                      |      |
| [#21618](https://github.com/pingcap/tidb/issues/21618)&#x2757; | major    |                         |                                                      |      |
| [#21655](https://github.com/pingcap/tidb/issues/21655)&#x2757; | major    |                         |                                                      |      |
| [#21688](https://github.com/pingcap/tidb/issues/21688)         | moderate | @you06                  | [#21878](https://github.com/pingcap/tidb/pull/21878) |      |
| [#22244](https://github.com/pingcap/tidb/issues/22244)&#x2757; | moderate |                         |                                                      |      |
| [#22345](https://github.com/pingcap/tidb/issues/22345)         | minor    | @lysu                   | [#22372](https://github.com/pingcap/tidb/pull/22372) |      |
| [#22356](https://github.com/pingcap/tidb/issues/22356)&#x2757; | moderate |                         |                                                      |      |
| [#22393](https://github.com/pingcap/tidb/issues/22393)&#x2757; | moderate |                         |                                                      |      |
| [#22783](https://github.com/pingcap/tidb/issues/22783)         | minor    | <sub>@tiancaiamao</sub> | [#24696](https://github.com/pingcap/tidb/pull/24696) |      |
| [#22927](https://github.com/pingcap/tidb/issues/22927)&#x2757; | critical |                         |                                                      |      |
| [#23015](https://github.com/pingcap/tidb/issues/23015)&#x2757; | minor    |                         |                                                      |      |
| [#23180](https://github.com/pingcap/tidb/issues/23180)&#x2757; | minor    |                         |                                                      |      |
| [#23235](https://github.com/pingcap/tidb/issues/23235)         | minor    | @tangenta               |                                                      |      |
| [#23331](https://github.com/pingcap/tidb/issues/23331)         | major    | @cfzjywxk               | [#23342](https://github.com/pingcap/tidb/pull/23342) |      |
| [#23423](https://github.com/pingcap/tidb/issues/23423)&#x2757; | moderate |                         |                                                      |      |
| [#23542](https://github.com/pingcap/tidb/issues/23542)         | minor    |                         | [#24140](https://github.com/pingcap/tidb/pull/24140) |      |
| [#23797](https://github.com/pingcap/tidb/issues/23797)         | major    | @tangenta               |                                                      |      |
| [#24195](https://github.com/pingcap/tidb/issues/24195)&#x2757; | major    |                         |                                                      |      |
| [#24411](https://github.com/pingcap/tidb/issues/24411)&#x2757; | minor    |                         |                                                      |      |
| [#24586](https://github.com/pingcap/tidb/issues/24586)&#x2757; | major    |                         |                                                      |      |
| [#24589](https://github.com/pingcap/tidb/issues/24589)&#x2757; | moderate |                         |                                                      |      |
| [#24741](https://github.com/pingcap/tidb/issues/24741)         | critical | @cfzjywxk               |                                                      |      |
| [#24858](https://github.com/pingcap/tidb/issues/24858)&#x2757; | minor    |                         |                                                      |      |
| [#25003](https://github.com/pingcap/tidb/issues/25003)&#x2757; | major    |                         |                                                      |      |
| [#25029](https://github.com/pingcap/tidb/issues/25029)&#x2757; | minor    |                         |                                                      |      |


<h2 name="sig/DDL">sig/DDL</h2>

|                             ISSUE                              | PRIORITY |        ASSIGNEE         |                          PR                          | HINT |
|----------------------------------------------------------------|----------|-------------------------|------------------------------------------------------|------|
| [#982](https://github.com/pingcap/tidb/issues/982)             | major    | @tangenta               | [#20708](https://github.com/pingcap/tidb/pull/20708) |      |
| [#3644](https://github.com/pingcap/tidb/issues/3644)&#x2757;   | moderate |                         |                                                      |      |
| [#3876](https://github.com/pingcap/tidb/issues/3876)&#x2757;   | moderate |                         |                                                      |      |
| [#7545](https://github.com/pingcap/tidb/issues/7545)&#x2757;   | moderate |                         |                                                      |      |
| [#10260](https://github.com/pingcap/tidb/issues/10260)&#x2757; | moderate |                         |                                                      |      |
| [#11410](https://github.com/pingcap/tidb/issues/11410)&#x2757; | minor    |                         |                                                      |      |
| [#11648](https://github.com/pingcap/tidb/issues/11648)         | moderate |                         | [#21237](https://github.com/pingcap/tidb/pull/21237) |      |
| [#11952](https://github.com/pingcap/tidb/issues/11952)         | major    | @bb7133                 | [#21564](https://github.com/pingcap/tidb/pull/21564) |      |
| [#14241](https://github.com/pingcap/tidb/issues/14241)&#x2757; | minor    |                         |                                                      |      |
| [#15567](https://github.com/pingcap/tidb/issues/15567)&#x2757; | moderate |                         |                                                      |      |
| [#17460](https://github.com/pingcap/tidb/issues/17460)         | minor    | @zimulala               |                                                      |      |
| [#17686](https://github.com/pingcap/tidb/issues/17686)         | moderate | @qw4990                 |                                                      |      |
| [#17745](https://github.com/pingcap/tidb/issues/17745)&#x2757; | moderate |                         |                                                      |      |
| [#17808](https://github.com/pingcap/tidb/issues/17808)         | major    | <sub>@Reminiscent</sub> | [#21497](https://github.com/pingcap/tidb/pull/21497) |      |
| [#18336](https://github.com/pingcap/tidb/issues/18336)         | moderate | @AilinKid               |                                                      |      |
| [#18907](https://github.com/pingcap/tidb/issues/18907)&#x2757; | moderate |                         |                                                      |      |
| [#19697](https://github.com/pingcap/tidb/issues/19697)         | moderate | <sub>@xiongjiwei</sub>  |                                                      |      |
| [#20592](https://github.com/pingcap/tidb/issues/20592)&#x2757; | moderate |                         |                                                      |      |
| [#21063](https://github.com/pingcap/tidb/issues/21063)         | moderate |                         | [#21064](https://github.com/pingcap/tidb/pull/21064) |      |
| [#21835](https://github.com/pingcap/tidb/issues/21835)         | moderate |                         | [#21845](https://github.com/pingcap/tidb/pull/21845) |      |
| [#21943](https://github.com/pingcap/tidb/issues/21943)         | moderate | @bb7133                 | [#21953](https://github.com/pingcap/tidb/pull/21953) |      |
| [#22117](https://github.com/pingcap/tidb/issues/22117)&#x2757; | critical |                         |                                                      |      |
| [#22453](https://github.com/pingcap/tidb/issues/22453)         | major    | @tangenta               | [#22458](https://github.com/pingcap/tidb/pull/22458) |      |
| [#22704](https://github.com/pingcap/tidb/issues/22704)&#x2757; |          |                         |                                                      |      |
| [#22708](https://github.com/pingcap/tidb/issues/22708)&#x2757; | moderate |                         |                                                      |      |



---

updated at 2021-06-02T09:47:07-00:00


