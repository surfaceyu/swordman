<template>
    <p style="white-space: pre-wrap;">
        「游戏介绍」
        <br>
        <a href="javascript:void(0)" @click="aboutGameClick('equip')">装备</a>|
        <a href="javascript:void(0)">道具</a>|
        <a href="javascript:void(0)">玩法</a>|
        <a href="javascript:void(0)" @click="aboutGameClick('map')">地图</a>|
        <a href="javascript:void(0)">门派</a>|
        <a href="javascript:void(0)">私聊</a>|
        <a href="javascript:void(0)">副本</a>|
        <a href="javascript:void(0)">其他</a>|
        <a href="javascript:void(0)">说明</a>
        <br>
        {{ gameIntroduction }}
    </p>
</template>

<script setup lang="ts">
import { ref } from 'vue';
interface stringMap {
    [id:string]: string,
}
const introductionMap: stringMap = {
    "equip":`一、装备介绍:
1.武器技能:
剑:致盲伤害,致盲状态下一定90%几率无法命中。
刀:冰冻伤害:冰冻状态下玩家90%几率无法移动。
棍:虚弱伤害:虚弱状态下对方无法回复生命值。
斧:燃烧伤害:燃烧状态下,每分钟掉下40%。
弓:中毒伤害:中毒状态下,每分钟掉下40%。且无法回复内力。PVE特殊效果 攻击触发负面状态时 双倍伤害
武器相克:剑克刀克棍克斧克弓克剑,当攻击被克武器的对手时攻击提升30%,被克制时伤害降低20%
鞋子阻挡:鞋子属性高低决定阻挡仇人几率。

反扑算法:反扑几率=30-己方装备等级+对方装备等级 最大40 最小0
暴击几率算法:武器副属性每提升4点增加1%暴击几率
暴击伤害算法:剑 弓 刀 三档暴击 10%,30%,50% 棍三档暴击 10%,20%,30% 斧 四挡暴击 10%,30%,50%,80%
伪随机算法说明:PVP暴击几率 反扑几率 均使用PRD伪随机算法`,

    "map": `地图介绍:
1.自由大陆:由海岛 丛林 平原 山脉 村庄 峡谷 深渊 组成。
2.自由城:坐落在大陆中心,主要城镇 自由大陆的主要势力及核心生活区域
3.村庄:分布在自由城南侧(新手历练的好去处1-40级)
4:落日森林:在自由城的西侧,大陆的修行者经常进入落日森林,历练获取宝物.充满危险的气息。(40-80)
5:失落平原:落日森林北侧,荒无人烟的地方全是异兽和妖魔(80=120)。
6:魔兽山脉:自由城东侧听说是上古时代,诸神大战 留下的战场遗迹,里面全是高阶魔兽 120-200。
6:海岛:传说自由大陆一共有6座海岛,单目前只找到了其中五座海岛被五大门派占领,剩下一个海岛一直在寻找。`
}
let gameIntroduction = ref(introductionMap["equip"])
function aboutGameClick(name: string) {
    gameIntroduction.value = introductionMap[name] ? introductionMap[name] : introductionMap.value
}
</script>