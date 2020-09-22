package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unsafe"
)

type Property []byte

func (p Property) value(index int) int16 {
	return Int16(p, index)
}

func (p Property) updateValue(index int, v int16) Property {
	SetInt16(p, index, v)
	return p
}

func (p Property) Body() int16 {
	return p.value(0x36E)
}

func (p Property) UpdateBody(i int16) Property {
	return p.updateValue(0x36E, i)
}

func (p Property) Life() int16 {
	return p.value(0x366)
}

func (p Property) UpdateLife(i int16) Property {
	return p.updateValue(0x366, i)
}

func (p Property) MaxLife() int16 {
	return p.value(0x368)
}

func (p Property) UpdateMaxLife(i int16) Property {
	return p.updateValue(0x368, i)
}

func (p Property) DoubleAttack() bool {
	return p.value(0x3B8) == 1
}

func (p Property) UpdateDoubleAttack(b bool) Property {
	if b {
		return p.updateValue(0x3B8, 1)
	} else {
		return p.updateValue(0x3B8, 0)
	}
}

func (p Property) Sense() int16 {
	return p.value(0x3B2)
}

func (p Property) UpdateSense(i int16) Property {
	return p.updateValue(0x3B2, i)
}

func (p Property) Poisonous() int16 {
	return p.value(0x3B6)
}

func (p Property) UpdatePoisonous(i int16) Property {
	return p.updateValue(0x3B6, i)
}

func (p Property) Qualification() int16 {
	return p.value(0x3BC)
}

func (p Property) UpdateQualification(i int16) Property {
	return p.updateValue(0x3BC, i)
}

var Gongfu = []string{"无",
	"野球拳", "武当长拳", "罗汉拳", "灵蛇拳",
	"神王毒掌", "七伤拳", "混元掌", "寒冰绵掌",
	"鹰爪功", "逍遥掌", "铁掌", "幻阴指",
	"寒冰神掌", "千手如来掌", "天山六阳掌", "玄冥神掌",
	"冰蚕毒掌", "龙象般若功", "一阳指", "太极拳",
	"空明拳", "蛤蟆功", "太玄神功", "黯然销魂掌",
	"降龙十八掌", "葵花神功", "化功大法", "吸星大法",
	"北冥神功", "六脉神剑", "躺尸剑法", "青城剑法",
	"冰雪剑法", "恒山剑法", "泰山剑法", "衡山剑法",
	"华山剑法", "嵩山剑法", "全真剑法", "峨嵋剑法",
	"武当剑法", "万花剑法", "泼墨剑法", "雪山剑法",
	"泰山十八盘", "回峰落雁剑法", "两仪剑法", "太岳三青峰",
	"玉女素心剑", "逍遥剑法", "慕容剑法", "倚天剑法",
	"七星剑法", "金蛇剑法", "苗家剑法", "玉箫剑法",
	"玄铁剑法", "太极剑法", "达摩剑法", "辟邪剑法",
	"独孤九剑", "西瓜刀法", "血刀大法", "狂风刀法",
	"反两仪刀法", "火焰刀法", "胡家刀法", "霹雳刀法",
	"神龙双钩", "大轮杖法", "怪异武器", "炼石弹",
	"叫化棍法", "火焰发射器", "鳄鱼剪", "大蜘蛛",
	"毒龙鞭法", "黄沙万里鞭法", "雪怪", "判官笔",
	"持棋盘", "大剪刀", "持瑶琴", "大莽蛇",
	"金花杖法", "神龙鹿杖", "打狗棍法", "五轮大法",
	"松风剑法", "普通攻击", "狮子吼", "九阳神功",
}

var GongfuID = map[string]int{}

func init() {
	for i, s := range Gongfu {
		GongfuID[s] = i
	}
}

func ExitIfError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ExecutablePath() (string, error) {
	execPath, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(filepath.Dir(execPath))
}

func SavesPath(dir string) ([]string, error) {
	file, err := os.Open(dir)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	names, err := file.Readdirnames(-1)
	if err != nil {
		return nil, err
	}
	saves := make([]string, 3)
	for _, name := range names {
		if strings.ToUpper(name) == "R1.GRP" {
			saves[0] = filepath.Join(dir, name)
		} else if strings.ToUpper(name) == "R2.GRP" {
			saves[1] = filepath.Join(dir, name)
		} else if strings.ToUpper(name) == "R3.GRP" {
			saves[2] = filepath.Join(dir, name)
		}
		if len(saves[0]) > 0 && len(saves[1]) > 0 && len(saves[2]) > 0 {
			break
		}
	}
	if len(saves[0]) > 0 && len(saves[1]) > 0 && len(saves[2]) > 0 {
		return saves, nil
	}
	return nil, fmt.Errorf("没有找到存档文件在 %q 中", dir)
}

func Uint16(buf []byte, index int) uint16 {
	return uint16(buf[index]) | uint16(buf[index+1])<<8
}

func Int16(buf []byte, index int) int16 {
	ret := Uint16(buf, index)
	return *((*int16)(unsafe.Pointer(&ret)))
}

func SetUint16(buf []byte, index int, value uint16) {
	buf[index] = byte(value)
	buf[index+1] = byte(value >> 8)
}

func SetInt16(buf []byte, index int, value int16) {
	SetUint16(buf, index, *((*uint16)(unsafe.Pointer(&value))))
}
