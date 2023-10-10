package parser

import (
	"fmt"
	"testing"
)

func Test_byteToStr(t *testing.T) {
	fmt.Println(string([]byte{0x6d, 0x65, 0x74, 0x61}))
}

func TestDecodeMetaOutScript(t *testing.T) {
	//script := "006a046d6574612231514545337a6a6637646d346a4358535a3774616548725436533861345a434177424065653761643138663464393837656561366634393532633235323861323332313836306534373663306230633839343932343138663332393239303535333537064d65746149440e53616d706c6550726f746f636f6c05302e302e390130417b22636f6e74656e74223a22e8bf99e698afe4b880e4b8aae6b58be8af95e58685e5aeb9222c2263726561746554696d65223a313538373631303632323132337d0d31353837363132373439393539"

	//ROOT
	//script := "006a046d65746122314b4658726447516637555341336b755664504566655743584a59414b4239364335044e554c4c064d657461494404526f6f7405302e302e390130044e554c4c0d31353837363132373331363737"
	//INFO
	//script := "006a046d65746122314e7142796f62616e6f4d334252763748566d5a4a7641685a4c5859507a6e6835704039626166323063393663383761643566366162323865636434323735623737333231323339373931356365646634333336383366303465353531663835353938064d657461494404496e666f05302e302e390130044e554c4c0d31353837363132373334343730"
	//NAME
	//script := "006a046d65746122314a3679345252373564744146434c4e316d743351337a744e6a7143476152504c344064303634643262663431616262613063663966383037393430356437356230356133373930386338353330613030373131323361356131336233623532646165064d6574614944046e616d6505302e302e39013005416c6963650d31353837363132373338313835"
	//protocol
	//script := "006a046d6574612231467434386a795876386342576a7a594a4c72533956586b7a5963375a52637a72384039626166323063393663383761643566366162323865636434323735623737333231323339373931356365646634333336383366303465353531663835353938064d65746149440a2350726f746f636f6c7305302e302e390130044e554c4c0d31353837363132373432353130"
	//sample_protocol
	//script := "006a046d6574612231427a353134586f726d63673752555535544364635a317535784e39724e366941384063623363363162373062316139633039353661363635663139666664343661643034323634623561373737623634303335333662636234626562616432313364064d65746149440f2353616d706c6550726f746f636f6c05302e302e3901300c3030316434366666656266340d31353837363132373436363737"
	//sample_detail
	//script := "006a046d6574612231514545337a6a6637646d346a4358535a3774616548725436533861345a434177424065653761643138663464393837656561366634393532633235323861323332313836306534373663306230633839343932343138663332393239303535333537064d65746149440e53616d706c6550726f746f636f6c05302e302e390130417b22636f6e74656e74223a22e8bf99e698afe4b880e4b8aae6b58be8af95e58685e5aeb9222c2263726561746554696d65223a313538373631303632323132337d0d31353837363132373439393539"

	//Text
	//script := "006a046d657461223135557878536759694e656f5136614c3670665863443344505531435a653436476b40313138316535383837653534303266313465353633373061323366383862373835313539356363376130303961633934396366396238626330353137663134310a5465737453686f7749441553686f77546578742330303030303030303030303003312e300130487b227469746c65223a227469746c65222c22636f6e74656e74223a2253686f775465787420e58685e5aeb932222c22636f6e74656e7454797065223a22706c61696e54657874227d0a31353837363236313833"
	//null data
	//script := "006a046d6574612231504d344e4c6e74655777737145726e5a6e687a77545a78785477426a4e787236364064303634643262663431616262613063663966383037393430356437356230356133373930386338353330613030373131323361356131336233623532646165"
	//DATA : 1
	//script := "006a046d6574612231504d344e4c6e74655777737145726e5a6e687a77545a78785477426a4e7872363640643036346432626634316162626130636639663830373934303564373562303561333739303863383533306130303731313233613561313362336235326461650131"

	//script := "006a046d657461223132634c72757371777571365761727970327535626d765a366a736b59416e6232574066363738373232383230623861613831333662663666343132323864386536663439383363373734376631346230313461363164663264643264316134316662064d6574614944223132634c72757371777571365761727970327535626d765a366a736b59416e623257417b22636f6e74656e74223a22e8bf99e698afe4b880e4b8aae6b58be8af95e58685e5aeb9222c2263726561746554696d65223a313538373631303632323132337d013004302e30390a746578742f706c61696e055554462d38"
	//script := "006a046d65746122313766593439734d59537a4b72547278527135364c4831365557446d38627058626a40333536376633663964656638653837303465326533663533376631363333343363323562626434643330383565623465373939363336383030343739616438360a5465737453686f7749440853686f77546578744c9e7b227469746c65223a227469746c65222c22636f6e74656e74223a2268656c6c6f222c22636f6e74656e7454797065223a22706c61696e54657874222c224d6574614964223a2239643766356264653735663266316537373930383033363437303530366632646166633338663163346366656166623232643535656161313932623530323339222c2274696d65223a313538383034313634303538337d013005302e302e390a746578742f706c61696e055554462d38"

	//script := ""
	//script := "006a046d65746142303265346531646639633565303838373433636164393964633066386636316336623033313933613662636266386336643139336562323461386161386439633731044e554c4c0a5465737453686f77494404526f6f74044e756c6c0130044e756c6c0a746578742f706c61696e055554462d38"
	//script := ""
	//script := "006a3f7065657267616d652e636f6d7c53454e447c4c477c4253567c7b2276223a2230222c2264223a5b7b2273223a224c222c2261223a302e3032303030307d5d7d"

	//err
	//script := "006a4ce585ace5858332303230e5b9b4e698afe4baba5b40656d6f6a693d5c75393835455d5b40656d6f6a693d5c75364237375de58fb2e79a84e58886e6b0b45b40656d6f6a693d5c75354442415d"
	//script := "006a046d6574612231514142394e4e456843616f364c66424e3741325a4c447843514551544e65563761044e554c4c0a5465737453686f77494404526f6f7403312e30013009526f6f742d746573740a31353837333639303534"
	//script := ""
	//script := "5869616f446f6765436f696e0000000000000000584443000000000000000008a3b41cb3c9138112b1714c7f5fe5bed51bfb83f8406f400100000000a56a0aa476c0cdb616d83aca6cdd916538f9bc72eff59aed795ff696ab583c0c000000000100000073656e7369626c65"
	//script := ""
	//script := ""
	//script := ""
	//script := ""
	script := "006a036d76634230323338663839303664326536356262313831626437326630336334313461373966646237333233386166643963353334393931396132663962373765346237363440633261333739326236333833376363636466636661656432633666646566643961376366316331393865623162343564663034326230366133343437613230620a746573746d657461696413467449737375652d30323338663839303664324d45017b2274797065223a226d657461636f6e7472616374222c2267656e657369734964223a2261636230633564663538373066333666313239386434306136383366643334383837613434396333222c2273656e7369626c654964223a22663962643531363263353939616538333464316564613438643335613862326136363263363535653035343864323063343737313931396134633262643633323030303030303030222c22746f6b656e416d6f756e74223a223330303030303030303030303030222c2267656e6573697341646472657373223a226e32645248724c4272424a536169663775675a3254484e544c44596d754d6d4d4742222c2261646472657373223a226e32645248724c4272424a536169663775675a3254484e544c44596d754d6d4d4742222c22616c6c6f77496e637265617365497373756573223a747275657d013005312e302e30106170706c69636174696f6e2f6a736f6e055554462d38"

	data, _, err := DecodeMetaOutScript(script, "")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("data:")
	fmt.Println(len(data))
	fmt.Println(data)
	da := PartsToDataPart(data)
	fmt.Println(da.MetaIdTag)
}

func TestSHA256(t *testing.T) {
	nodePubKey := "02eb4e1fa39f7a168a70b03a6bc4a9909b40d263a1b7002dd7dc28ef93eca73f1a"
	nodeTxId := "a5c05146de8fa49d8641cf34ae556c982048cd74a22731285dd744a48b5ed400"
	fmt.Println(MakeMetanetId(nodePubKey, nodeTxId))
}
