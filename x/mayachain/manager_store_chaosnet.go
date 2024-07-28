//go:build !testnet && !stagenet && !regtest
// +build !testnet,!stagenet,!regtest

package mayachain

import (
	"gitlab.com/mayachain/mayanode/common"
	"gitlab.com/mayachain/mayanode/common/cosmos"
	"gitlab.com/mayachain/mayanode/constants"
)

func importPreRegistrationMAYANames(ctx cosmos.Context, mgr Manager) error {
	oneYear := fetchConfigInt64(ctx, mgr, constants.BlocksPerYear)
	names, err := getPreRegisterMAYANames(ctx, ctx.BlockHeight()+oneYear)
	if err != nil {
		return err
	}

	for _, name := range names {
		mgr.Keeper().SetMAYAName(ctx, name)
	}
	return nil
}

func migrateStoreV96(ctx cosmos.Context, mgr Manager) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v88", "error", err)
		}
	}()

	err := importPreRegistrationMAYANames(ctx, mgr)
	if err != nil {
		ctx.Logger().Error("fail to migrate store to v88", "error", err)
	}
}

func migrateStoreV102(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v102", "error", err)
		}
	}()

	if err := mgr.BeginBlock(ctx); err != nil {
		ctx.Logger().Error("fail to initialise block", "error", err)
		return
	}

	// Remove stuck txs
	hashes := []string{
		"D19F621FAD0AE81688E4AF40EA9D0CD132A8A4DBFF3EA56F443E2D9083915F17",
		"A03C0A41909D85B2DF2F7E9D5D13F6E0AF89F366F6B580C0CCC13F5CEC0A9872",
		"7B7CC323ED0AD04DCB26DF1DEB46DE02B85345499336D043CBB5582EB77D22DB",
		"CE29D8AD79314333E307265529256304E26FC0B538B19B2D07578BE3D6252CE4",
		"B6EBF457EB1817E852722CB9F51C26E45C35F58B2445048FE4BD38FD1A603894",
		"199838DB755A6199AB401AAB1D56D296C66B0972001CB033B9CDC4217E636270",
		"674BDD72DF068A95EFA5DD94C4691A1D492A3342DA368DF0799ADD4D344D694D",
		"6F9C3D5AD6221159191540CE55704BCBB446626B209C852DC29C5C0AC7A24A82",
		"F3A3041FD304B11B8EBB748C9BE964E1FCCE0004770B109F5F9B72114F7FB9B9",
		"4912D98B5C8D9D090CAD2732754F39FFB324DA7008A19A0235DD77A4AB8EF3E3",
		"C82029C6D3F7D8D226E9B13F09CD05CF30FEA15F6C96BB8D49E20A4E063F6E82",
		"46C783972218F50015281F28222F5DC46FA3926EABB93549A383180C43064F96",
		"081819200E3ACC82CC8D95DFA87A6F0D87704154922022F777FFA5AD82B1BEF0",
		"9C3B5774352256A37CC3B26B82287458C4D4DEDC988342E6A2088A1800ACE992",
		"8B651D92B0374FA4E97834E86D35601940F90E104B800BEA836685E28452A953",
		"EA57B9FC879E981598732F6112255D756593D354DC88712665FEDC354374AD41",
		"672D551F02D6030A77745E25C0C8768347BBDA35DD7AA61C02751C86799D7C18",
		"116237EDC4814A9F684D8FCBC58FB5ADED2A9386B5ED0F1E627BCEFA8246474C",
		"3EE6362906A180279B0B9470221017465A1AA25807EFDB5A7B9342A95E120E2B",
		"1B04FF39247F519BC01F88FB1AC6843223FF351C47DC1D96B0FEC782645463F1",
		"25BC9C71B8F4D071A684A327D6E2657DA1D01D241E419C5C705D3690B5653C2E",
		"758B2A8DF6BC62F1A922DEA5E75F585A9BDB39CFC01152E4C74FBF929C5B5777",
		"A502C210DE19555884464A27408E16C378D9327BDC155EE3076F7D3D8CC8B963",
		"95AD2B7B2EE2E2CBEF272B20AB271400CDF57EE8EC170F8B265554A9FC24542A",
		"279B076B50ADCF2DE06CF129DE6B4917754F56FFB7CF4644FF0429CFC49A0D23",
		"D427493DEF0DFE953194E2E3C633C7EF3AAEC38F77B06B9E1EFAEFFF2071D58C",
		"B6F9C4CB2ABC7FB80B336950E559DD3020CC44C8F92A6AD9D3449612A5A232CC",
		"DDB8A4FD768443BF36187EA6147469A3D4975ED0CD8B4DDB2140EF4B924C7817",
		"C6E972C90798E33317DAC162D7B419AF825540F352A7CF38A5AB1297EAA866E9",
		"313F924DC160D565573F3B9D0A47F378A099606FFE4D059947B5377AE98E9F65",
		"E8340277F7E2310DDD2435A52EC1CC7C07C6D33FCD1F4ABD513FF23B6B19990F",
		"CDFDEBF26E28789F7C272813524C7F3766A9B82AFF55CBAD9AF347121061171B",
		"9F10CA47145E9F6B6EB4297272D9DF9999A75937CD154D6C4DAEC3DBFE14C3D6",
		"1D843810B79E7ED1CDD5424B0FDBD6158ADE77479E1C006F2583E5263E26E667",
		"1A486C051F7478CE845D67E019EE30DEF58D61B8EE408FE43E6CD520DE45518F",
		"EC461F353F95D933723BDAE7945B970A7F45DCA68A671D06B0FE9AA206686EFC",
		"66F228BD65D1A82C6D78C234D1A86F1C7E118A1051D87FE6546F708E401720FA",
		"233D1D0FB660BDA2C3C13C9B6C2BD0E96E81E05EC93C43A526CE0B782CA4ADA1",
		"8C867538F1C5A564C1C82206CBF0B96277B66E630BF13473E51D27BAE8B1994E",
		"BEA5B29954A3634B37CC0D73EB30EB8427ABF58900521A413F0A66C73AD6742A",
		"F3FA329499C42BF258B4D79E43ADEEB1E9C56FB60D4A9390B12F4946A554642A",
		"B26E3D4D8458DD43DB3B6424F1310B457053BDA95DC22D3936FCC373B49C95AC",
		"972E49601D4BF9949C3B91162399249B4AC997ED1BA830DB6DBC7DF44ABBEF3D",
		"D4B8E0F61978046D1205B5DC857BBD887214BB7054113B499FAADF7105F4CFE0",
		"97C8E399272FD9C64C2E2F1E2E32804157BFBA71504B4B838850F2590F87D781",
		"50C0CCD601689011E88B54358001EA2C6B1E8C0AC6794D1A7D8C95A74256071C",
		"640702E326CB6B61CC7285B0ADCE6DEF0694E9CFB629FD32C34A5475B5391E9E",
		"C60AE6A164FE3BE8B2BE87543B25B0F36E199E1CE466CB09482D9ED7D2D78BA9",
		"A58A823D9E467B368713D65090DCCFAAD92D1C8D6F2B57E3933EB8ACB9946031",
		"8EA259B4E7D15FBB6703C0A1248A137BCBAA7255ABEC09CEFAC2FB34DF7BC2F5",
		"B20036A869329CA1CDF966F0443B8B524A2CF6AB4F4ECA7C359D61A0A167F36B",
		"05F77D6640AE44FCFCB30FBDB8E76F82C4FD75E05A6DC48271EC499A1A09C378",
		"8E70E838DCEAD4D1763B4B40E59942EEFE5492B631D9CFB303A1DB0F7075F835",
		"C268C6821C3A8C19B435B4591F216A20E8DAD9AD2C17EF59F7CF9BD4DA2B4536",
		"9E1DF502EE17709E267AEDF673CF94B188698C0CBB9A6FAAAF57EAE20D043495",
		"FA194EDF6818312E6B28AB1D228A44B8623415595ED1716E7B7A92CB3DFCDE36",
		"941D3F4B252B735C2D358A368724DC809ED9CC63D6ED4426E369E75175EDF0F0",
		"24B077C67D4F835D176B701EEC59FA1C14143A0E3ACFF64077632FB3CEBD2851",
		"C4E86318378C561AD16DF9697F09224D254A314BED36EC7AC6C0B7F35FAB5CDB",
		"BEFFC122704DB5525A9511411A942F7F06EECF6386C104BB05622EDAE94D8096",
		"5422336EF4134851F601A74AA30C5E47702CC08111775ADC3944F4F0B467CD4F",
		"308FDED05E0F39E103A6E3898A497A1F28806ED7DEAB2D88F825E95CC4942D53",
		"C41ADECBC9D85D956D3246CEFD350E54CCABDA2B315793FD2625D30BEA0763C4",
		"333C9BC7B7479D4A675307B63AB2372C89C9C21A75C379BB6FE8EA8FB83813A0",
		"A967482A359194C6B3E0045F68B2E11CD275B29FF7F3A7F6129902D90FFA7055",
		"9DD6CFA490E5ED47BAE45E1CEE141329C411D8BAF5642758CCF3749D13862076",
	}
	removeTransactions(ctx, mgr, hashes...)

	// Rebalance asgard vs real balance for RUNE
	vaultPubkey, err := common.NewPubKey("mayapub1addwnpepq0tgksv4kjn0ya5n4gt2546dnasw84nr3zdtdzfud9z984p8pvmnu5t3qsy")
	if err != nil {
		ctx.Logger().Error("fail to get vault pubkey", "error", err)
		return
	}
	vault, err := mgr.Keeper().GetVault(ctx, vaultPubkey)
	if err != nil {
		ctx.Logger().Error("fail to get vault", "error", err)
		return
	}

	vault.SubFunds(common.NewCoins(common.NewCoin(common.RUNEAsset, cosmos.NewUint(3947_32403277))))

	if err = mgr.Keeper().SetVault(ctx, vault); err != nil {
		ctx.Logger().Error("fail to set vault", "error", err)
		return
	}

	// Remove retiring vault
	vaults, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, RetiringVault)
	if err != nil {
		ctx.Logger().Error("fail to get retiring asgard vaults", "error", err)
		return
	}
	for _, v := range vaults {
		runeAsset := v.GetCoin(common.RUNEAsset)
		v.SubFunds(common.NewCoins(runeAsset))
		if err = mgr.Keeper().SetVault(ctx, v); err != nil {
			ctx.Logger().Error("fail to save vault", "error", err)
		}
	}

	// Add LPs from unobserved txs
	lps := []struct {
		MayaAddress string
		ThorAddress string
		TxID        string
		Amount      cosmos.Uint
		Tier        int64
	}{
		// users which don't have an LP position yet
		{
			MayaAddress: "maya142m4adpj57hkrymqe5n8zzcxm5cqccpn3a6w6y",
			ThorAddress: "thor1jzzaw44tr0cxgxaah7h2sen2ck03lllw882wn2",
			TxID:        "73217ACF7F4061089236E29588825603FB4025E40AC5835586ED0B7959BE4A1F",
			Amount:      cosmos.NewUint(2_00000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya15kg7dfew844rdh5esgkrdevp78yhf4fjryjcfu",
			ThorAddress: "thor1hd9p0fllkwkgj9epe3nynr253az7uclxs4g2uw",
			TxID:        "5594D2500BB36F70ADB4063B4D7A331DCE884D2C34373EDBD69022C33E31CD0F",
			Amount:      cosmos.NewUint(1_00000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya17lllslx89rrxu0dh6y9ctz0aa2j82tljnuuy9s",
			ThorAddress: "thor1vmq7vwk8t6sxg730aps5vqetm905ndtmcvdq69",
			TxID:        "10376393CBF1C9E92CCBBDF582FFE9896FC04E82C2E9C641B4CB18A23559E43E",
			Amount:      cosmos.NewUint(2_00000000),
			Tier:        1,
		},
		{
			MayaAddress: "maya1f40wek6sj6uay6nplxpe2c6pj98c5uq78xspa4",
			ThorAddress: "thor1f40wek6sj6uay6nplxpe2c6pj98c5uq783wdt9",
			TxID:        "2A7297AD1EB1F1C53C90241264E78F067DE94F1C80588C208E8B7B5D86B3B9E7",
			Amount:      cosmos.NewUint(5388_00000000),
			Tier:        1,
		},
		{
			MayaAddress: "maya1j8pswr7vpf9jjmhrn0xlwvzla2f9yfxwcwtj0p",
			ThorAddress: "thor1y9h2yk95c6uqp29xglkgyf9kqxqnu28nn6vwwz",
			TxID:        "D5BEA6C8B3170B418ACD67B8C8A44A60CD0A66696B9B691E7A7471E393F5E8B4",
			Amount:      cosmos.NewUint(1_00000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya1k83lm2nyrd7vgl8h9xcjhwu9kr2zecauslje79",
			ThorAddress: "thor1k83lm2nyrd7vgl8h9xcjhwu9kr2zecausgv4g4",
			TxID:        "DEF2BC77DFCDA774C81D921C8846886FFF804D462F0E6BFF78DBAA1ADDF72E68",
			Amount:      cosmos.NewUint(24_83513163),
		},
		{
			MayaAddress: "maya1p3hcnlfdla2647rpersykfatplvhkehd2duspa",
			ThorAddress: "thor1p3hcnlfdla2647rpersykfatplvhkehd26zuhd",
			TxID:        "C4F73CFBAC15565CCAED86B66EB405AE9E36F712F0457F8353D050FF37D636BB",
			Amount:      cosmos.NewUint(1_00000000),
		},
		{
			MayaAddress: "maya1pn03td7tzsftp6xz25r5fas43dgqynpf0lyan5",
			ThorAddress: "thor1pn03td7tzsftp6xz25r5fas43dgqynpf0g639y",
			TxID:        "A85BE46FFDD915D2074EC85C8E5B63B0407EFDD44CC6094CCC9A616A7FFB0494",
			Amount:      cosmos.NewUint(1_00000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya1s0ry4c65c7k020vgpykjfy5rkqv8d7yn60lzx6",
			ThorAddress: "thor1s0ry4c65c7k020vgpykjfy5rkqv8d7yn6cpws2",
			TxID:        "4A1CA0E1D87869C5083F6BBD2042BF5DA5545B01ECE9CD7922F11D8AB715B261",
			Amount:      cosmos.NewUint(1_00000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya1vwslytml73dclz0h4enc2xluf4z03esrt36n6r",
			ThorAddress: "thor1vwslytml73dclz0h4enc2xluf4z03esrtxylvn",
			TxID:        "237CBC3570DA3AE95D15F6E7C04A50EF3799A4106434A9A831A11BEDA8EB0FF6",
			Amount:      cosmos.NewUint(36_00000000),
			Tier:        1,
		},
		{
			MayaAddress: "maya1wlx25u0692nvxllg57tgt45h53hjsgggzlgavn",
			ThorAddress: "thor1cjlsyrzmfpldxhmz4j3yzyc0f6dp57lhv6cm2r",
			TxID:        "C8D1F65C6C6559D4A23E8BB47533E86CC25D8C41FA8382EC2C6FBF868953AB23",
			Amount:      cosmos.NewUint(1_50000000),
			Tier:        3,
		},
		{
			MayaAddress: "maya1zgtzwkd9qaagvwedgnmxeh9tsqc8wdsjwjxf6e",
			ThorAddress: "thor1zgtzwkd9qaagvwedgnmxeh9tsqc8wdsjw9c9vf",
			TxID:        "79A5288200EB347569B7E3707A822E72B2DB1CCD52BC035323DE2B1DC44273B3",
			Amount:      cosmos.NewUint(499_98000000),
			Tier:        1,
		},
		// users which already have an existing LP position
		{
			MayaAddress: "maya10nqg4w30e9dnm0qg7swa8qsyqevuemwx78dpdx",
			ThorAddress: "thor10nqg4w30e9dnm0qg7swa8qsyqevuemwx7sndmk",
			Amount:      cosmos.NewUint(5_58000000),
		},
		{
			MayaAddress: "maya14sanmhejtzxxp9qeggxaysnuztx8f5jra7vedl",
			ThorAddress: "thor14sanmhejtzxxp9qeggxaysnuztx8f5jrafj4m0",
			Amount:      cosmos.NewUint(958_08765797),
		},
		{
			MayaAddress: "maya17w5n2r7akuunq9e296y22qrljh3qqegf6usf5x",
			ThorAddress: "thor17w5n2r7akuunq9e296y22qrljh3qqegf6tw9zk",
			Amount:      cosmos.NewUint(1400_00000000),
		},
		{
			MayaAddress: "maya1a4v8ajttgx5u822k2s8zms3phvytz3at2a7mgj",
			ThorAddress: "thor1a4v8ajttgx5u822k2s8zms3phvytz3at22qh7z",
			Amount:      cosmos.NewUint(1_000000),
		},
		{
			MayaAddress: "maya1fdl7xga4sxhwlfs48fhkgwen88003g3hl006pn",
			ThorAddress: "thor1fdl7xga4sxhwlfs48fhkgwen88003g3hlc3khr",
			Amount:      cosmos.NewUint(1_00000000),
		},
		{
			MayaAddress: "maya1hh03993slyvggmvdl7q4xperg5n7l86pufhkwr",
			ThorAddress: "thor1wlzhcxs0r4yh4pswj8zfqlp7dnp95p4kxn0dcr",
			Amount:      cosmos.NewUint(4_30000000),
		},
		{
			MayaAddress: "maya1j42xpqgfdyagr57pxkxgmryzdfy2z4l65mjzf9",
			ThorAddress: "thor1j42xpqgfdyagr57pxkxgmryzdfy2z4l65vvwl4",
			Amount:      cosmos.NewUint(2_00000000),
		},
		{
			MayaAddress: "maya1j6ep9yljeswft03w2qunqx8my9e2efph5ywhls",
			ThorAddress: "thor1jj4xufkxrjd4d3uswh0ztgr0yan3mdcdxh3tgn",
			Amount:      cosmos.NewUint(2_00000000),
		},
		{
			MayaAddress: "maya1jwq4zu4v3tfktwemwh2lwwnlu3nvvrhuhs6k0h",
			ThorAddress: "thor1jwq4zu4v3tfktwemwh2lwwnlu3nvvrhuh8y6e8",
			Amount:      cosmos.NewUint(285_40743565),
		},
		{
			MayaAddress: "maya1ka2v9exy8ata00pch87wgzf9dsmyag94tq8mug",
			ThorAddress: "thor1ka2v9exy8ata00pch87wgzf9dsmyag94theh2c",
			Amount:      cosmos.NewUint(978_00000000),
		},
		{
			MayaAddress: "maya1mj8yhw3jqljfcggkjd77k9t7jlcw0uur0yfurh",
			ThorAddress: "thor1mj8yhw3jqljfcggkjd77k9t7jlcw0uur0nhs48",
			Amount:      cosmos.NewUint(341_00000000),
		},
		{
			MayaAddress: "maya1ppdzsyugtsdtd6dpvzzg2746qfdfmux7k2ydal",
			ThorAddress: "thor1z9xhmhtxn5gxd4ugfuxk7hg9hp03tw3qtqs3f3",
			Amount:      cosmos.NewUint(1_00000000),
		},
		{
			MayaAddress: "maya1qdhqqlg5kcn9hz7wf8wzw8hj8ujrjplnz669c9",
			ThorAddress: "thor1ru7upan5aj2hmzlevrztd6gn5r5z8jxrcjzmup",
			Amount:      cosmos.NewUint(1_00000000),
		},
		{
			MayaAddress: "maya1qtcst64ea585s7gtek3daj2xe59hgn8q7j0ccl",
			ThorAddress: "thor1qtcst64ea585s7gtek3daj2xe59hgn8q7935w0",
			Amount:      cosmos.NewUint(2998_00000000),
		},
	}

	pool, err := mgr.Keeper().GetPool(ctx, common.RUNEAsset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "error", err)
		return
	}

	var address common.Address
	for _, sender := range lps {
		address, err = common.NewAddress(sender.MayaAddress)
		if err != nil {
			ctx.Logger().Error("fail to parse address", "error", err)
			continue
		}

		var lp LiquidityProvider
		lp, err = mgr.Keeper().GetLiquidityProvider(ctx, common.RUNEAsset, address)
		if err != nil {
			ctx.Logger().Error("fail to get liquidity provider", "error", err)
			continue
		}

		pool.PendingInboundAsset = pool.PendingInboundAsset.Add(sender.Amount)
		lp.PendingAsset = lp.PendingAsset.Add(sender.Amount)
		lp.LastAddHeight = ctx.BlockHeight()
		if sender.TxID != "" {
			var txID common.TxID
			txID, err = common.NewTxID(sender.TxID)
			if err != nil {
				ctx.Logger().Error("fail to parse txID", "error", err)
				continue
			}
			lp.PendingTxID = txID
		}

		if lp.AssetAddress.IsEmpty() {
			var thorAdd common.Address
			thorAdd, err = common.NewAddress(sender.ThorAddress)
			if err != nil {
				ctx.Logger().Error("fail to parse address", "address", sender.MayaAddress, "error", err)
				continue
			}
			lp.AssetAddress = thorAdd
		}

		mgr.Keeper().SetLiquidityProvider(ctx, lp)
		if sender.Tier != 0 {
			if err = mgr.Keeper().SetLiquidityAuctionTier(ctx, lp.CacaoAddress, sender.Tier); err != nil {
				ctx.Logger().Error("fail to set liquidity auction tier", "address", lp.CacaoAddress, "error", err)
				continue
			}
		}

		if err = mgr.Keeper().SetPool(ctx, pool); err != nil {
			ctx.Logger().Error("fail to set pool", "address", pool.Asset, "error", err)
			return
		}

		evt := NewEventPendingLiquidity(pool.Asset, AddPendingLiquidity, lp.CacaoAddress, cosmos.ZeroUint(), lp.AssetAddress, sender.Amount, common.TxID(""), common.TxID(sender.TxID))
		if err = mgr.EventMgr().EmitEvent(ctx, evt); err != nil {
			continue
		}
	}

	// Remove duplicated THOR address LP position
	// https://mayanode.mayachain.info/mayachain/liquidity_auction_tier/thor.rune/maya1dy6c9tmu7qgpd6cw2unumew3sknduwx7s0myr6?height=488436
	// https://mayanode.mayachain.info/mayachain/liquidity_auction_tier/thor.rune/maya1yf0sglxse7jkq0laddtve2fskkrv6vzclu3u6e?height=488436
	add1, err := common.NewAddress("maya1dy6c9tmu7qgpd6cw2unumew3sknduwx7s0myr6")
	if err != nil {
		ctx.Logger().Error("fail to parse address", "error", err)
		return
	}

	lp1, err := mgr.Keeper().GetLiquidityProvider(ctx, common.RUNEAsset, add1)
	if err != nil {
		ctx.Logger().Error("fail to get liquidity provider", "error", err)
		return
	}

	add2, err := common.NewAddress("maya1yf0sglxse7jkq0laddtve2fskkrv6vzclu3u6e")
	if err != nil {
		ctx.Logger().Error("fail to parse address", "error", err)
		return
	}

	lp2, err := mgr.Keeper().GetLiquidityProvider(ctx, common.RUNEAsset, add2)
	if err != nil {
		ctx.Logger().Error("fail to get liquidity provider", "error", err)
		return
	}
	lp2.PendingAsset = lp2.PendingAsset.Add(lp1.PendingAsset)

	mgr.Keeper().SetLiquidityProvider(ctx, lp2)
	if err = mgr.Keeper().SetLiquidityAuctionTier(ctx, lp2.CacaoAddress, 0); err != nil {
		ctx.Logger().Error("fail to set liquidity auction tier", "error", err)
	}
	mgr.Keeper().RemoveLiquidityProvider(ctx, lp1)

	// Mint cacao
	toMint := common.NewCoin(common.BaseAsset(), cosmos.NewUint(9_900_000_000_00000000))
	if err = mgr.Keeper().MintToModule(ctx, ModuleName, toMint); err != nil {
		ctx.Logger().Error("fail to mint cacao", "error", err)
		return
	}

	if err = mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, ReserveName, common.NewCoins(toMint)); err != nil {
		ctx.Logger().Error("fail to send cacao to reserve", "error", err)
		return
	}

	// 150214766379119 de BTC Asgard a reserve
	// 473657580023 de ETH Asgard a reserve
	// 24192844274670 de RUNE asgard a reserve
	for _, asset := range []common.Asset{common.BTCAsset, common.ETHAsset, common.RUNEAsset} {
		pool, err = mgr.Keeper().GetPool(ctx, asset)
		if err != nil {
			ctx.Logger().Error("fail to get pool", "error", err)
			return
		}
		switch asset {
		case common.BTCAsset:
			pool.BalanceCacao = pool.BalanceCacao.Sub(cosmos.NewUint(1_501_734_01773759))
		case common.ETHAsset:
			pool.BalanceCacao = pool.BalanceCacao.Sub(cosmos.NewUint(4736_57580023))
		case common.RUNEAsset:
			pool.BalanceCacao = pool.BalanceCacao.Sub(cosmos.NewUint(211_877_34242261))
		}

		if err = mgr.Keeper().SetPool(ctx, pool); err != nil {
			ctx.Logger().Error("fail to set pool", "error", err)
			return
		}
	}

	// Sum of all the above will be sent
	asgardToReserve := common.NewCoin(common.BaseAsset(), cosmos.NewUint(1_717_347_93596043))
	if err = mgr.Keeper().SendFromModuleToModule(ctx, AsgardName, ReserveName, common.NewCoins(asgardToReserve)); err != nil {
		ctx.Logger().Error("fail to send asgard to reserve", "error", err)
		return
	}

	// 164293529917265 de itzamna a reserve
	itzamnaToReserve := common.NewCoin(common.BaseAsset(), cosmos.NewUint(1_642_935_29917265))
	itzamnaAcc, err := cosmos.AccAddressFromBech32("maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz")
	if err != nil {
		ctx.Logger().Error("fail to parse address", "error", err)
		return
	}

	if err := mgr.Keeper().SendFromAccountToModule(ctx, itzamnaAcc, ReserveName, common.NewCoins(itzamnaToReserve)); err != nil {
		ctx.Logger().Error("fail to send itzamna to reserve", "error", err)
		return
	}

	// FROM RESERVE TXS
	// 8_910_000_500_00000000 from reserve to itzamna
	reserveToItzamna := common.NewCoin(common.BaseAsset(), cosmos.NewUint(8_910_001_000_00000000))
	if err := mgr.Keeper().SendFromModuleToAccount(ctx, ReserveName, itzamnaAcc, common.NewCoins(reserveToItzamna)); err != nil {
		ctx.Logger().Error("fail to send reserve to itzamna", "error", err)
		return
	}

	// Remove Slash points from genesis nodes
	for _, genesis := range GenesisNodes {
		acc, err := cosmos.AccAddressFromBech32(genesis)
		if err != nil {
			ctx.Logger().Error("fail to parse address", "error", err)
			continue
		}

		mgr.Keeper().ResetNodeAccountSlashPoints(ctx, acc)
	}
}

func migrateStoreV104(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v104", "error", err)
		}
	}()

	// Select the least secure ActiveVault Asgard for all outbounds.
	// Even if it fails (as in if the version changed upon the keygens-complete block of a churn),
	// updating the voter's FinalisedHeight allows another MaxOutboundAttempts for LackSigning vault selection.
	activeAsgards, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, ActiveVault)
	if err != nil || len(activeAsgards) == 0 {
		ctx.Logger().Error("fail to get active asgard vaults", "error", err)
		return
	}
	if len(activeAsgards) > 1 {
		signingTransactionPeriod := mgr.GetConstants().GetInt64Value(constants.SigningTransactionPeriod)
		activeAsgards = mgr.Keeper().SortBySecurity(ctx, activeAsgards, signingTransactionPeriod)
	}
	vaultPubKey := activeAsgards[0].PubKey

	// Refund failed synth swaps back to users
	// These swaps were refunded because the target amount set by user was higher than the swap output
	// but because there were a bug in calculating the fee of synth swaps they were treated as zombie coins,
	// and thus we failed to generate the out tx of refund. (keep in mind that the refund event is emitted)
	// Since they are all inbound transactions, we can refund them back to users without deducting fee (see refundTransactions implementation)
	failedSwaps := []adhocRefundTx{
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      300000000,
			inboundHash: "86AC0A216FA3138E3B1EE15D66DEBCBE46D8A62B45EA6D33E07DE044D4BD638E",
		}, {
			toAddr:      "maya1x5979k5wqgq58f4864glr7w2rtgyuqqm6l2zhx",
			asset:       "THOR/RUNE",
			amount:      26142918750,
			inboundHash: "9FC3C8886CD432338B4E4A388DF718B3EE03B257CA2D87792A9D3AFE4AC76DA6",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5047059000,
			inboundHash: "86964E9623839AEBD7D4E74CC777F917AC5DACA850B322F07E7CD6F9A8ACEC1F",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5273550000,
			inboundHash: "1A9D4E7000FE5EF4E292378F1EA075D69DE4DAF2FD5258AC5C2C6E495F28B843",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5366524000,
			inboundHash: "7BADEBA845A889750BF9477B8A01870F109EAAE46E55EF032EF868540F6DB4C1",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      6347294000,
			inboundHash: "C71D1260FD4FE208CFA70440847250716F8C852674592B55EDA390FF840E1C8E",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      6347446000,
			inboundHash: "703E35046FC628B48CA06DD8FE9A95151ECC447C9A55FA7172CC6ED0F97540C4",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5389962000,
			inboundHash: "B1F5B7C9B8AA46A96A72D1E10BD083172810669B912E46BFF2713B4D6237C42C",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5390043000,
			inboundHash: "D55E1265E68D4605118EC02ADE7FB2FD2A91AD35878E701093D0C82B8D624A04",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      1271000,
			inboundHash: "A20F86A30BED39CFC4734EEA1C50680CC32002974E2FE5CE82BB22B26643D618",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      1271000,
			inboundHash: "F3FE1EFE4181E1F81048CCCD366A0E624A98C8C9ED9DC304E3DC32BF2FD3050D",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      1272000,
			inboundHash: "11BFC34721FEBE40CD080432B379F4D9C43DCA147653AFF82849B82838C1B4FD",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      621000,
			inboundHash: "31722DE22A5243DAB294529F3323B6708E8B3040C0205D5602F4F3F5D4218712",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      847000,
			inboundHash: "FFBBAE0420A1F7D1371F837BCA89D697EAC6E7D90835767C70D5B05584F95CD1",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      237000,
			inboundHash: "AD516C4C23A984336DC2BEE188CD7B607F31E342192BD0D05371A0B1AC127234",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      237000,
			inboundHash: "42C110B63651F47B10066C47334296E9E28E006A6481B675A8DAA27946843B81",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      429000,
			inboundHash: "EC192B4327CE11A03611FB5EFEEF3E133C8937040B36B4312A1743BACB4FFA88",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      384000,
			inboundHash: "A43B1F63D2B3092B80F0321DFFE81179BD3EE7209B1EA035D573A83F68EB7177",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      384000,
			inboundHash: "811035D9F7A177199F2BD84B90F82477AC68B2898D0F99B9EEB524766AC914DD",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      429000,
			inboundHash: "F5A01D066C001EB138E8DC4FA21B36917FCA8DCB07289B0E8575FD9B500C4C59",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      429000,
			inboundHash: "EB9F7D4AB93920447EC3423A8D4F1E92C43AFC42B9D18C1362EC5322DFB5ADDB",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      265000,
			inboundHash: "0C1359A466FDB89F450D02AB5C36F1073179D9333D05616FC33D8946058498BC",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      310000,
			inboundHash: "6AB46CD16A3BD9015570C2CD086CEF7BF75ADEBD98C7732149024C45F8458602",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      237000,
			inboundHash: "A45D65ADC584C687DBB696DD10E4E7E9FDC1E81FBA5525BEBF978A03EA2B93BD",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      310000,
			inboundHash: "4554C137395882EB69F01017D64993CCDCA7263AAFB755ADBFE5FDA6A00AB8A4",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "86567487D4E2B1E05B1A86EFE7A7A548849B8F82E17A8325484F855B92633D9B",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "83767F9F779657C0D770975ADFF3A92BCD1EE7A1C999985EA9DB066FBB44610F",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "34DE850E4F1AEFB5EE32C9FA2446D85B76F531E3006463B5F390570439FD96DE",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "18308912F70B0C58FF53BFC7617514C1279EDE4537FA29C822CA3801BBF7C82B",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      265000,
			inboundHash: "7C0EAF6ADE8B9A6DEF3CCF2F462826DDEEFD71C60B29023CE107115616B614BC",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      265000,
			inboundHash: "45B2F0E6338DCD7799026821D3F86A1C794F80A19ED5EE8CA3E5EC649A194B4C",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      265000,
			inboundHash: "2E22F59FC3B69CF871A411B2A057FDCA2DF00469819EFB2A946BE05E22373362",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      310000,
			inboundHash: "3DB087973B69B3A32EA4FA5B16579B3C2EEE6A0E070C03EF8DDE578A12B399FD",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      265000,
			inboundHash: "841A7B59C3E20A58A20C939BCD45800E69D2316F74B84F990B9DCC2E5D43D632",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "E394B2D44226421FC4FCCAD3C0F58D80EE6FE3F70B93E5FA1B699923EBF73588",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "79122693813DB28FC79D9454FA4327523ADAF71F4675BD36BF677301A568090F",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "34068637E11234A6AFC0C85A318CB58666623A8E36626D4E265B10551E1C7166",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "05AEE74CB1B9A2AD4CA3BAC63DB4D4ADA0ADF1EB345D6BC94CCCD0672669E1F8",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "BDBA06245E439EA80C1DEB8295449DF6EF3FC22D0F4D64FAFF3C0095D64413CB",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "A99BAE4BB38CA491D8457D16F2577144879285D97A44E3624848D70F1FD5963B",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      502000,
			inboundHash: "4541912B35D1D6A27A6263B6A7E608AFEB1687984765B2669DAE2612188AD4B9",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      846000,
			inboundHash: "59E5738F0DDB3B1D7F3DF8EFBD691C61480BDB909B57E971C5FFF03054A0EC3B",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      846000,
			inboundHash: "C48B393A050D7983F83C429BD3D17E2C300FAC2EFCB5B18F45E68E42952BC126",
		}, {
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      20021610000,
			inboundHash: "48D8788931772A5566C922C098C579FBBEE2B2793057B487FE3AE2AC2F3C8ED9",
		}, {
			toAddr:      "maya1x5979k5wqgq58f4864glr7w2rtgyuqqm6l2zhx",
			asset:       "ETH/USDT-0XDAC17F958D2EE523A2206206994597C13D831EC7",
			amount:      82605080070,
			inboundHash: "E0320F7459B83A9F86695C9D0DB78B916F69FF5E408F94E521889F3F0C3CE086",
		},
	}
	refundTransactions(ctx, mgr, vaultPubKey.String(), failedSwaps...)

	// 1st user tier fix
	// User with address "maya1dy6c9tmu7qgpd6cw2unumew3sknduwx7s0myr6" and "maya1yf0sglxse7jkq0laddtve2fskkrv6vzclu3u6e" which had
	// should have been allocated an amount during the cacao donation in the last store migration but seems that
	// there was a problem with the migration and the amount was not allocated. So we change his/her tier to 1
	// and allocate the attribution amount manually from reserve.
	// The changes are as the following:
	// 1. Change Tier from 0 -> 1
	// 2. Overwrite LP Units from 0 -> 38089_5898484080 LP Units
	// 3. Pending Asset from 3210_34000000 -> 0
	// 4. Asset Deposit Value from 0 -> 3273_80071698
	// 5. Cacao Deposit Value from 0 -> 38089_5898484080 (Same as LP Units)
	// 6. Move 38827_9343263458 CACAO from Reserve to Asgard module (CACAO deposit value + Change difference between asset deposit value and pending asset with CACAO denom)
	// 7. Increase by 38827_9343263458 the CACAO on Asgard Pool for RUNE (CACAO deposit value + Change difference between asset deposit value and pending asset with CACAO denom)
	// 8. Increase by 38089_5898484080 the LP UNITS on Asgard Pool for RUNE
	// 9. Move 3210_34000000 Asset from Pending_Asset in Asgard Pool for RUNE to Balance_Asset in Asgard Pool for RUNE
	// 10. Emit Add Liquidity Event
	addr1, err := common.NewAddress("maya1yf0sglxse7jkq0laddtve2fskkrv6vzclu3u6e")
	if err != nil {
		ctx.Logger().Error("fail to parse address", "error", err)
		return
	}
	lp1, err := mgr.Keeper().GetLiquidityProvider(ctx, common.RUNEAsset, addr1)
	if err != nil {
		ctx.Logger().Error("fail to get liquidity provider", "error", err)
		return
	}
	lp1.Units = cosmos.NewUint(38089_5898484080)
	lp1.PendingAsset = cosmos.ZeroUint()
	lp1.AssetDepositValue = cosmos.NewUint(3273_80071698)
	lp1.CacaoDepositValue = cosmos.NewUint(38089_5898484080)
	mgr.Keeper().SetLiquidityProvider(ctx, lp1)
	if err = mgr.Keeper().SetLiquidityAuctionTier(ctx, lp1.CacaoAddress, 1); err != nil {
		ctx.Logger().Error("fail to set liquidity auction tier", "error", err)
	}

	reserve2Asgard1 := common.NewCoin(common.BaseAsset(), cosmos.NewUint(38827_9343263458))
	if err = mgr.Keeper().SendFromModuleToModule(ctx, ReserveName, AsgardName, common.NewCoins(reserve2Asgard1)); err != nil {
		ctx.Logger().Error("fail to send reserve to asgard", "error", err)
		return
	}
	pool, err := mgr.Keeper().GetPool(ctx, common.RUNEAsset)
	if err != nil {
		ctx.Logger().Error("fail to get pool", "error", err)
		return
	}
	addedCacao1 := cosmos.NewUint(38827_9343263458)
	pool.BalanceCacao = pool.BalanceCacao.Add(addedCacao1)
	addedLPUnits := cosmos.NewUint(38089_5898484080)
	pool.LPUnits = pool.LPUnits.Add(addedLPUnits)
	pendingAsset2Balance := cosmos.NewUint(3210_34000000)
	pool.PendingInboundAsset = pool.PendingInboundAsset.Sub(pendingAsset2Balance)
	pool.BalanceAsset = pool.BalanceAsset.Add(pendingAsset2Balance)
	evt1 := NewEventAddLiquidity(
		pool.Asset,
		addedLPUnits,
		lp1.CacaoAddress,
		addedCacao1,
		pendingAsset2Balance,
		common.TxID(""),
		common.TxID(""),
		lp1.AssetAddress,
	)

	// 2nd user tier fix
	// 1. Change Tier from 3 -> 1
	// 2. Increase Asset Deposit Value from 333_46986565 to 507_5548724869
	// 3. Increase Cacao Deposit Value from 3879_8117483183 to 5905_2332716199
	// 4. Increase LP UNITS from 3879_8117483183 to 5905_2332716199
	// 5. Move 4050_84304660322 CACAO from Reserve module to Asgard module (twice as much on purpose, to account for asset side, will be armed away)
	// 6. Increase in 4050_84304660322 CACAO the balance_cacao of Asgard Pool for RUNE (twice as much on purpose, to account for asset side, will be arbed away)
	// 7. Increase by 2025_4215233016 the LP UNITS on Asgard Pool for RUNE
	// 8. Emit Add Liquidity Event
	addr2, err := common.NewAddress("maya1jwq4zu4v3tfktwemwh2lwwnlu3nvvrhuhs6k0h")
	if err != nil {
		ctx.Logger().Error("fail to parse address", "error", err)
		return
	}
	lp2, err := mgr.Keeper().GetLiquidityProvider(ctx, common.RUNEAsset, addr2)
	if err != nil {
		ctx.Logger().Error("fail to get liquidity provider", "error", err)
		return
	}
	lp2.AssetDepositValue = cosmos.NewUint(507_5548724869)
	lp2.CacaoDepositValue = cosmos.NewUint(5905_2332716199)
	lp2.Units = cosmos.NewUint(5905_2332716199)
	mgr.Keeper().SetLiquidityProvider(ctx, lp2)
	if err = mgr.Keeper().SetLiquidityAuctionTier(ctx, lp2.CacaoAddress, 1); err != nil {
		ctx.Logger().Error("fail to set liquidity auction tier", "error", err)
	}

	reserve2Asgard2 := common.NewCoin(common.BaseAsset(), cosmos.NewUint(4050_84304660322))
	if err = mgr.Keeper().SendFromModuleToModule(ctx, ReserveName, AsgardName, common.NewCoins(reserve2Asgard2)); err != nil {
		ctx.Logger().Error("fail to send reserve to asgard", "error", err)
		return
	}
	addedCacao2 := cosmos.NewUint(4050_84304660322)
	pool.BalanceCacao = pool.BalanceCacao.Add(addedCacao2)
	addedLPUnits2 := cosmos.NewUint(2025_4215233016)
	pool.LPUnits = pool.LPUnits.Add(addedLPUnits2)
	evt2 := NewEventAddLiquidity(
		pool.Asset,
		addedLPUnits2,
		lp2.CacaoAddress,
		addedCacao2,
		cosmos.ZeroUint(),
		common.TxID(""),
		common.TxID(""),
		common.Address(""),
	)

	err = mgr.Keeper().SetPool(ctx, pool)
	if err != nil {
		ctx.Logger().Error("fail to set pool", "error", err)
		return
	}
	if err := mgr.EventMgr().EmitEvent(ctx, evt1); err != nil {
		ctx.Logger().Error("fail to emit event", "error", err)
		return
	}
	if err := mgr.EventMgr().EmitEvent(ctx, evt2); err != nil {
		ctx.Logger().Error("fail to emit event", "error", err)
		return
	}
}

// migrateStoreV105 is complementory migration to migration v104
// it will refund another 17 failed synth swaps txs back to users
func migrateStoreV105(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v105", "error", err)
		}
	}()

	// Select the least secure ActiveVault Asgard for all outbounds.
	// Even if it fails (as in if the version changed upon the keygens-complete block of a churn),
	// updating the voter's FinalisedHeight allows another MaxOutboundAttempts for LackSigning vault selection.
	activeAsgards, err := mgr.Keeper().GetAsgardVaultsByStatus(ctx, ActiveVault)
	if err != nil || len(activeAsgards) == 0 {
		ctx.Logger().Error("fail to get active asgard vaults", "error", err)
		return
	}
	if len(activeAsgards) > 1 {
		signingTransactionPeriod := mgr.GetConstants().GetInt64Value(constants.SigningTransactionPeriod)
		activeAsgards = mgr.Keeper().SortBySecurity(ctx, activeAsgards, signingTransactionPeriod)
	}
	vaultPubKey := activeAsgards[0].PubKey

	// Refund failed synth swaps back to users
	// These swaps were refunded because the target amount set by user was higher than the swap output
	// but because there were a bug in calculating the fee of synth swaps they were treated as zombie coins,
	// and thus we failed to generate the out tx of refund. (keep in mind that the refund event is emitted)
	// Since they are all inbound transactions, we can refund them back to users without deducting fee (see refundTransactions implementation)
	failedSwaps := []adhocRefundTx{
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      6175167000,
			inboundHash: "8EECEE5C27795B96E8465D3234DEC050219AC591D899D038D2F11A1EFCE00E72",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      1271000,
			inboundHash: "E31EBA09AA7E64DE5F1209656956286C4883196B0E85A075764600ABC57ACDB6",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      384000,
			inboundHash: "6777B04215485FC495A88FA5D76C1873E250756FFF5E23577CA3CEEB4E042B0C",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      384000,
			inboundHash: "E34798700D6034A3D8C82F80E7FCC4AC0F68574FCB7FD018EFA7E90A2594A44F",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      384000,
			inboundHash: "2D129E0E58A762263272DB2548B432912E995F2A09CFF4A6C06A4DF8534290C7",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      846000,
			inboundHash: "0172C67339320D14E477DCEB64F9FC4FABEE67DF233F08A81EB4D061F1820AC1",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "BTC/BTC",
			amount:      846000,
			inboundHash: "86E8363E44B4EF0B32A894FD3011AC6AB8EC7AAE3EA2F65ACD8D0D15DB1299C7",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5314334000,
			inboundHash: "DBBDA76A5315F25787041BF95A65FC19BD2464B637BA9ED322CD8A52C1CE447E",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5927862000,
			inboundHash: "8D40B6E45B676638764FB38A998FAD782514AF2DDDB840A809A6CB65C854DF70",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      8527629000,
			inboundHash: "A860164DFDC3B0E76B871FF93A509B80736486B622AE59D0EE77ECE5F0E39D6A",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      8527450000,
			inboundHash: "3B99680D1927C6A3D909B964378443E8D5C71F9DA2A3E7FF4AFF16C7B6E08FA5",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      20123240000,
			inboundHash: "53BA8317F50DFB97FF30235BA479F3E3F78E29FEFA90BB2C113891F121D79C04",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      7565885000,
			inboundHash: "DD862F71E427F5DE280F2CAA49E007B77A1B64E30896AA93B1EC782374CDAB04",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      20107080000,
			inboundHash: "62D7648EC776A7B68FFAB23844EB5AA2C967F7E7CA97379E03607682B312B33E",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "THOR/RUNE",
			amount:      20103040000,
			inboundHash: "BD90238148001ADF4B485D98D549AB472F5AC1881E8F67DBA70CC5C80E979803",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5252251000,
			inboundHash: "D028821B72FD5A37C092771FF9F5039C7A7E04FFDDEF8793A0FFE0BD73156733",
		},
		{
			toAddr:      "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc",
			asset:       "ETH/USDC-0XA0B86991C6218B36C1D19D4A2E9EB0CE3606EB48",
			amount:      5060695000,
			inboundHash: "89266DF89E689C79DE4ACCB7312FEC85CE57CF92852A222228C3388F6FBDDA57",
		},
		{
			toAddr:      "maya1x5979k5wqgq58f4864glr7w2rtgyuqqm6l2zhx",
			asset:       "THOR/RUNE",
			amount:      72494713125,
			inboundHash: "8671D17BFD6040531470C89D0412116EE2909396BB6C54E037535DFD529E67D2",
		},
	}
	refundTransactions(ctx, mgr, vaultPubKey.String(), failedSwaps...)

	// Refunding USDT coins that mistakenly got sent to the vault (mayapub1addwnpepqwuwsax7p3raecsn2k9uvqyykanlvhw47asz836se2h0nyg6knug6n9hklq) by "transfer" txs back to user
	// transaction hashes are: 0xda4306037c838dcaed92775ecd515441e4a932b1bcbeef1199bf37a29274575d and 0xa6d765192856e982deae51bfc817f612c30344402ca72fbe526e8c534b91d048 on eth mainnet
	maxGas, err := mgr.GasMgr().GetMaxGas(ctx, common.ETHChain)
	if err != nil {
		ctx.Logger().Error("fail to get max gas", "error", err)
		return
	}
	toi := TxOutItem{
		Chain:       common.ETHChain,
		InHash:      common.BlankTxID,
		ToAddress:   common.Address("0x2510d455bF4a9b829C0CfD579543918D793F7762"),
		Coin:        common.NewCoin(common.USDTAssetV1, cosmos.NewUint(191_970_000+96_448_216)),
		MaxGas:      common.Gas{maxGas},
		GasRate:     int64(mgr.GasMgr().GetGasRate(ctx, common.ETHChain).Uint64()),
		VaultPubKey: common.PubKey("mayapub1addwnpepqwuwsax7p3raecsn2k9uvqyykanlvhw47asz836se2h0nyg6knug6n9hklq"),
	}
	if err := mgr.TxOutStore().UnSafeAddTxOutItem(ctx, mgr, toi, ctx.BlockHeight()); err != nil {
		ctx.Logger().Error("fail to save tx out item for refund transfers", "error", err)
		return
	}
}

func migrateStoreV106(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v106", "error", err)
		}
	}()

	danglingInboundTxIDs := []common.TxID{
		"EE6C4711C360C09B88D399E2000F66EDBC9D88243E977E4DA386575801B6C7BD",
		"D870C04715093BA8180705324F4B5F7BBFAF24D2D9F6FD41825EC3DA0A4848D4",
		"625C4E707AC12244DD657CC0465A280E2B5C64DA37C1F61B70F2DC4269E66760",
		"BBF3652682882D05D0B2ACDF8A06ECD1F16CA95877B50AD9EA22A012F0CE22F2",
		"F5D933BC96464024C7B176A699C881D8C158D3019673D6E6F4156B1D5D1C2B92",
		"C02BEB8C8A35D3DEF148FD1BEEA1BE74A6E0C1E437CFCE0690342C0D988D7BDB",
		"E3BFDD6AAA01B1444DD43DD02F82485C4892FBA10620DD8CB3B8371EE98009E7",
		"77342E2C624EEC3A031EA541498401B60679A46CF793470879EEDE9D95E8B062",
		"DA80B9FA56213D8F4176D1D81B1FA056EA360CDC103408CF10946B3626E54DC0",
		"3733965B49CE9655EAA1AA0AFF7EECE6067448B9BC5C6EA39CDD03CBDF6210E5",
		"3495C1DD09808BE80048D1B169855A96509C07C1814FB0CCFB5F7B20314C33A4",
		"41DCE02268892610CFB0FF0C442C9B701F89A871F8BD1DE0973CF86BC539EC85",
		"3920E1FC6362656466D6BB9BF93D754BDFCAAABD0D83EB720D37A35D4483E696",
		"A0306B7FCB6A1E626E797E828C79EAD36EB64A1A382BA44444071F15B83F8601",
		"BB6735BE958AC4C92EF7FD828C39233504E3F4DF2E879AC028924C52A2373FDB",
		"2561147F935F6FB962ACE98019507BC3459F09BC3DE27B13FF7642375DE895B6",
		"8ABF335AC56A27F028435EC1633474E7CF5CC8549CA418E4B9C69E705E1745B8",
		"CB390FAD8D646D541046999A2347AABFE25FA73B00382E07C18F62A03C833420",
		"8BB5C19CD2CF12BA5FCC505BCC4727155DD5B12200AF81E85176CD3D894417B7",
		"C9AE69BB8982F14CDB7EA9135E0EEA71F9EADD5F9260346302547CC39BAC0ED6",
		"87766587E59B40ABBA29B3615B96B8A8743C0401B87E3F7B55C1FFBCBA2B70DF",
		"42987321CD58E6F174A1FB7703F720CE0B7D78EDD2F2D51FE6EDDD400A4EB881",
		"9E01E3CB16D655CF91FE45477C38265F73C06DEF9684BD91696BEB6635873B0E",
		"A51A34C375E093E5B4BA8D9F0330FCE0D959A3B1D237F3E40DFFE0A97A65414D",
		"210088D913A05BCD6166534FD78CE0472C57D8CD4DD6811911C3E4728AD8CC13",
		"F165CD890E63B782E61B497854F6C2E4F12CB1D5BEA22193352586239E513502",
		"2DA12C51226E5F7BDEF5FD9087C72957A71B2D9FF0068AB200CDC28CD590C5A4",
		"44A0A75C227C0B5C197071116EFF318DD913798F33CA602F3EE07B6B043ED7DD",
		"199EAFFEE1E83873C4A35539DB5A209F87C7796AF7BA47C05758C401B6AE72A8",
		"CB0B1034EB9C82D8FB5CEC349FE37C6BE9185EEDB17384294433D360A7A66202",
		"B6730601AF5CFB78E6A040F32BBBCE599F662067D7C3D1D3E825D94CB53FE95A",
		"6129474913FDBDDC7C0A40D260DAA581F162033962E617F5B047492917F4CF96",
		"3EA40E01E372FF97A42366C64198873EE42D2288F4056AF88A4288E8F6ADF16C",
		"D8039ACA5F751FD777BDDF046064B06962693B6964543FF5EFA65C12A2D76026",
		"63E7565BC422ED92693BEB6E4F43254D168199C49F0C154798E31DDEAAAF879A",
		"D4BD757812983E18C839F0C7C071545C53253697C6FF171E479EEDB71D44664C",
		"A30EA261E2ADC64D115A0FA59A98FF1F060543BFBE44A9DF30FE3E9D7A2097B4",
		"6AF92FB712157AD24E6369090FBF0EA80ED0096361BBE2167C5C3AD86253E8FB",
		"4755A67A2910562F047CFB3A22A68610B403B49F70BC63479525E92F861AB8C1",
		"09A26EFCBB2A7F7CFCA58EFFD9E33D6791E9184AB8E9E69C41CBAA55FA4E61C4",
		"3B471BFCB5D8E839F2DE293CE17F4A9C9C1756D72DD5C897A9FEB25B882DBA6E",
		"8096A1316D4BC86CACF76B1F9584A2FAE89660F5B1E5D78CB8716CC2F3C95D33",
		"F0E8EC18F3C2A264E00BF69139B57B4670735FCAEFD80E4AFDF867A586EEA1C5",
		"D2F0BBBE8D17A231A7EF39EF262F368B23C5206096CA1905870F188C0BDFA14B",
		"9AE32B3D9B7F4325C605B0166B04BF3DB805A54FAEA7180AFF01BAC713960FA2",
		"782304A4A0381EA89187D72958E4F12B4CEDEB0E1292538535834FE91EAB8301",
		"1EDF41B23F38B18C5178CB47CFA75D804F6C07CF2CA9E2453FABDE7E7DDDD8D8",
		"12A925AB62E42776FB9E483040B197AB7777A8353454A36733C76A1F5027AB18",
		"51EDA1F86D2A62121709BCE3BE12DD0E81D8A8CD61F5951CEA740FBDCCB82427",
		"9D73AA660DDDBCF2F376CAB3B2A45DB9210CB74828C1BA7F5403C4B95546DE34",
		"08F4B26D032202EEBE1DC5529C03D65A141DF27F87A2EE24844CACAF5380C442",
		"40B62AD6D4A64C2DABDE1649FFB2AC216583749CF71E531C8BFFEEF6A82A9D00",
		"E9EB28E234D15938F8CB13E7312FF7F3DA62CE0D81369FE2CF13B6A9E3B70C60",
		"998E6B367B43FBF729A702C297E0A89D17041610CA8EFD61A1A9D5B802C9D769",
		"47B49FEEE7E5C1C81F953D4445F8999545E640E6F9AF8D0930FA0D205F3238A0",
		"90BFD95971540E0704CB517D262F0A8C526C0242BCA49A80ACDA2EE4AA06C2A8",
		"D4C8299B5A537AE92716125F7706CDF8B7A7E4C8796E2EE6D3235847419957FE",
	}
	requeueDanglingActions(ctx, mgr, danglingInboundTxIDs)

	spentTxs := []common.TxID{
		"356D59F3211F03C15667470A1AC31255C14FC2840C099F5B5250612C7D07F9FE",
		"15FDE171F250356EBE416D203D5141702B0983ECAA0B270AC9DA2E5C95202C53",
	}
	for _, spentTx := range spentTxs {
		voter, err := mgr.Keeper().GetObservedTxInVoter(ctx, spentTx)
		if err != nil {
			ctx.Logger().Error("fail to get observed tx voter", "error", err)
			continue
		}
		txOut, err := mgr.Keeper().GetTxOut(ctx, voter.OutboundHeight)
		if err != nil {
			ctx.Logger().Error("fail to get tx out array from key value store", "error", err)
			continue
		}
		outTxId := common.BlankTxID
		for _, outTx := range voter.OutTxs {
			if outTx.ID != common.BlankTxID {
				outTxId = outTx.ID
				break
			}
		}
		if outTxId != common.BlankTxID {
			for i := 0; i < len(txOut.TxArray); i++ {
				if txOut.TxArray[i].InHash.Equals(spentTx) {
					txOut.TxArray[i].OutHash = outTxId
				}
			}
		}
		err = mgr.Keeper().SetTxOut(ctx, txOut)
		if err != nil {
			ctx.Logger().Error("fail to save tx out item", "error", err)
			continue
		}
	}
}

func migrateStoreV107(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v106", "error", err)
		}
	}()

	toRemoveTxs := []string{
		"08F4B26D032202EEBE1DC5529C03D65A141DF27F87A2EE24844CACAF5380C442",
		"09A26EFCBB2A7F7CFCA58EFFD9E33D6791E9184AB8E9E69C41CBAA55FA4E61C4",
		"12A925AB62E42776FB9E483040B197AB7777A8353454A36733C76A1F5027AB18",
		"15FDE171F250356EBE416D203D5141702B0983ECAA0B270AC9DA2E5C95202C53",
		"199EAFFEE1E83873C4A35539DB5A209F87C7796AF7BA47C05758C401B6AE72A8",
		"1CD1FF94BD318864E6F5A50D44E4FB3E27378A15CF131EDCEC3367151DE4789C",
		"1EDF41B23F38B18C5178CB47CFA75D804F6C07CF2CA9E2453FABDE7E7DDDD8D8",
		"210088D913A05BCD6166534FD78CE0472C57D8CD4DD6811911C3E4728AD8CC13",
		"2561147F935F6FB962ACE98019507BC3459F09BC3DE27B13FF7642375DE895B6",
		"2DA12C51226E5F7BDEF5FD9087C72957A71B2D9FF0068AB200CDC28CD590C5A4",
		"32EFDD45DF865F5CA8219F78921B930314962DB014AF9FB8B7177339212349C7",
		"33F16A97389F1F91729A140BA2B7A03B30648D3E900674B1A6C09EBC5491F3D5",
		"3495C1DD09808BE80048D1B169855A96509C07C1814FB0CCFB5F7B20314C33A4",
		"356D59F3211F03C15667470A1AC31255C14FC2840C099F5B5250612C7D07F9FE",
		"3733965B49CE9655EAA1AA0AFF7EECE6067448B9BC5C6EA39CDD03CBDF6210E5",
		"391BD7F59BC700D7687F4FCB809D994EB4EAA33C5142115DAAE26623BEDDD801",
		"3920E1FC6362656466D6BB9BF93D754BDFCAAABD0D83EB720D37A35D4483E696",
		"3B471BFCB5D8E839F2DE293CE17F4A9C9C1756D72DD5C897A9FEB25B882DBA6E",
		"3EA40E01E372FF97A42366C64198873EE42D2288F4056AF88A4288E8F6ADF16C",
		"40B62AD6D4A64C2DABDE1649FFB2AC216583749CF71E531C8BFFEEF6A82A9D00",
		"41DCE02268892610CFB0FF0C442C9B701F89A871F8BD1DE0973CF86BC539EC85",
		"42987321CD58E6F174A1FB7703F720CE0B7D78EDD2F2D51FE6EDDD400A4EB881",
		"44A0A75C227C0B5C197071116EFF318DD913798F33CA602F3EE07B6B043ED7DD",
		"469474953796B1E5DE276DA1B41D2EF7D669D5EAAA68660DAF58ACF1CBB66FB8",
		"4755A67A2910562F047CFB3A22A68610B403B49F70BC63479525E92F861AB8C1",
		"47B49FEEE7E5C1C81F953D4445F8999545E640E6F9AF8D0930FA0D205F3238A0",
		"51EDA1F86D2A62121709BCE3BE12DD0E81D8A8CD61F5951CEA740FBDCCB82427",
		"56369450125799237380D7093635342BF99BC2278EABF7CABD8B5DACFB5DAEDC",
		"58F53279C0EB095B012369AF581C981DC79AB517E075204E31F7C4F453AE9159",
		"5968B058110932CA2B0EF181CEAD16B0008BDD2075E84EB2DBA8788C0C391FD6",
		"5AA2917E17476C6A40E67EC7D51B845433A97DB408BDFAB6AF0C2DF628D14CEF",
		"6129474913FDBDDC7C0A40D260DAA581F162033962E617F5B047492917F4CF96",
		"625C4E707AC12244DD657CC0465A280E2B5C64DA37C1F61B70F2DC4269E66760",
		"63E7565BC422ED92693BEB6E4F43254D168199C49F0C154798E31DDEAAAF879A",
		"6AF92FB712157AD24E6369090FBF0EA80ED0096361BBE2167C5C3AD86253E8FB",
		"70506FCE5BA295E5B8DCD2D38885EE1B717A099DF935640D5BA65A3512DE05DA",
		"770B1E9A1152DCED356AC69F421C4374AA81C883690CF1A7DF51E504973EC2B4",
		"77342E2C624EEC3A031EA541498401B60679A46CF793470879EEDE9D95E8B062",
		"77C77406C0EB550F0705838467BBDDA77638BAF2C2BDBE7FA27EAB7421BE32B5",
		"782304A4A0381EA89187D72958E4F12B4CEDEB0E1292538535834FE91EAB8301",
		"7A20AAC7EE7366CB82F46DE3470E7B2E6747AD13A38C8802B4BFC953C9CB58BA",
		"7B150F77ECD44232C242D1C6B0E568DB08E9BD5402701FA2FD9FB6021E687A1E",
		"7E441F57FE6D5297EE7F5E51B5FC7AAE232E015750F73C0B9255CC57D2185888",
		"8096A1316D4BC86CACF76B1F9584A2FAE89660F5B1E5D78CB8716CC2F3C95D33",
		"87766587E59B40ABBA29B3615B96B8A8743C0401B87E3F7B55C1FFBCBA2B70DF",
		"8ABF335AC56A27F028435EC1633474E7CF5CC8549CA418E4B9C69E705E1745B8",
		"8BB5C19CD2CF12BA5FCC505BCC4727155DD5B12200AF81E85176CD3D894417B7",
		"903504658A473094BC517A1435A73EC6331C6FBB9443241E441B29DD3B9C170D",
		"90A4B05A312D048586313CFDA6741B3848E5B7B2B22191DB0ABC1CB9037DE7A2",
		"90BFD95971540E0704CB517D262F0A8C526C0242BCA49A80ACDA2EE4AA06C2A8",
		"96166DE602F8491BB9FBFFDE443FC5A25810CF8B3C76EECB3708D4640F70277F",
		"998E6B367B43FBF729A702C297E0A89D17041610CA8EFD61A1A9D5B802C9D769",
		"9AE32B3D9B7F4325C605B0166B04BF3DB805A54FAEA7180AFF01BAC713960FA2",
		"9D73AA660DDDBCF2F376CAB3B2A45DB9210CB74828C1BA7F5403C4B95546DE34",
		"9E01E3CB16D655CF91FE45477C38265F73C06DEF9684BD91696BEB6635873B0E",
		"A0306B7FCB6A1E626E797E828C79EAD36EB64A1A382BA44444071F15B83F8601",
		"A19B4C3CCA0182EE6508243971D1310E88DA0292C855988B6402B945AF9943A5",
		"A30EA261E2ADC64D115A0FA59A98FF1F060543BFBE44A9DF30FE3E9D7A2097B4",
		"A51A34C375E093E5B4BA8D9F0330FCE0D959A3B1D237F3E40DFFE0A97A65414D",
		"A92CCD142F75EDB112BFBEE7F89AEA3813DFFC6DB54781F19EAD708A5B11557E",
		"AAE34BD6FCE944731AFA09E336B9677F01DAD688D8B16621D951EC874FAB3F52",
		"AFBDBC3D4039EE76A034ED40C5902BA1E887ABC16A3D5A757016AB7DA94FB158",
		"B6730601AF5CFB78E6A040F32BBBCE599F662067D7C3D1D3E825D94CB53FE95A",
		"BB6735BE958AC4C92EF7FD828C39233504E3F4DF2E879AC028924C52A2373FDB",
		"BBF3652682882D05D0B2ACDF8A06ECD1F16CA95877B50AD9EA22A012F0CE22F2",
		"C02BEB8C8A35D3DEF148FD1BEEA1BE74A6E0C1E437CFCE0690342C0D988D7BDB",
		"C3EB617E0DD25AE5AE1ABD5C91CDCDB51435CA515D5FE421462CBA0F37D11EDA",
		"C9AE69BB8982F14CDB7EA9135E0EEA71F9EADD5F9260346302547CC39BAC0ED6",
		"CB0B1034EB9C82D8FB5CEC349FE37C6BE9185EEDB17384294433D360A7A66202",
		"CB390FAD8D646D541046999A2347AABFE25FA73B00382E07C18F62A03C833420",
		"CE27C5A408DD4B30B170A90BB6EC2B04660A35F0ADA107B055812CB668AB6F8D",
		"D2F0BBBE8D17A231A7EF39EF262F368B23C5206096CA1905870F188C0BDFA14B",
		"D4BD757812983E18C839F0C7C071545C53253697C6FF171E479EEDB71D44664C",
		"D4C8299B5A537AE92716125F7706CDF8B7A7E4C8796E2EE6D3235847419957FE",
		"D8039ACA5F751FD777BDDF046064B06962693B6964543FF5EFA65C12A2D76026",
		"D870C04715093BA8180705324F4B5F7BBFAF24D2D9F6FD41825EC3DA0A4848D4",
		"DA80B9FA56213D8F4176D1D81B1FA056EA360CDC103408CF10946B3626E54DC0",
		"E31A6C3943E918C96638F5CDAF558EACE615E43F166D00EC0F97E117F98C46DA",
		"E3BFDD6AAA01B1444DD43DD02F82485C4892FBA10620DD8CB3B8371EE98009E7",
		"E56B1A5410DCA670268BA0EF262BC7C4A8D958E50987EF41A6116B9AE66FFD15",
		"E9EB28E234D15938F8CB13E7312FF7F3DA62CE0D81369FE2CF13B6A9E3B70C60",
		"EE6C4711C360C09B88D399E2000F66EDBC9D88243E977E4DA386575801B6C7BD",
		"F0E8EC18F3C2A264E00BF69139B57B4670735FCAEFD80E4AFDF867A586EEA1C5",
		"F165CD890E63B782E61B497854F6C2E4F12CB1D5BEA22193352586239E513502",
		"F5D933BC96464024C7B176A699C881D8C158D3019673D6E6F4156B1D5D1C2B92",
	}
	removeTransactions(ctx, mgr, toRemoveTxs...)

	// Take the inbound dash into account for the pool
	pool, err := mgr.Keeper().GetPool(ctx, common.DASHAsset)
	if err == nil {
		pool.BalanceAsset = pool.BalanceAsset.Add(cosmos.NewUint(438_32476664))
		err = mgr.Keeper().SetPool(ctx, pool)
		if err != nil {
			ctx.Logger().Error("fail to save pool", "error", err)
		}
	} else {
		ctx.Logger().Error("fail to get pool", "error", err)
	}

	// Sending the amount from reserve to pay out stuck txs
	address := "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz"
	acc, err := cosmos.AccAddressFromBech32(address)
	if err != nil {
		ctx.Logger().Error("fail to parse address: %s", address, "error", err)
	}

	coins := common.NewCoins(common.NewCoin(common.BaseNative, cosmos.NewUint(130_000_0000000000)))
	if err := mgr.Keeper().SendFromModuleToAccount(ctx, ReserveName, acc, coins); err != nil {
		ctx.Logger().Error("fail to send provider reward: %s", address, "error", err)
	}

	// Send node rewards to each of the bond providers
	type providerReward struct {
		Provider string
		Amount   uint64
	}

	// Rewards getting paid out because first few churns they weren't distributed and BPs not able to claim.
	rewards := []providerReward{
		{Provider: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", Amount: 27727_7226000606},
		{Provider: "maya1tndazzezsfka2wgqm52e5neej9q8jqrxv47h7m", Amount: 5298_8905701016},
		{Provider: "maya13yseu9un5f9gwqgzshjqvsqrxew0hhgm3wjh4l", Amount: 318_5469177912},
		{Provider: "maya1rzr9m407svj4jmc6rsxzsg75cx7gm3lsyyttyj", Amount: 1857_4794818020},
		{Provider: "maya1g70v5r9ujxrwewdn3w44pmqcygx49dx7ne82vr", Amount: 2841_3226310891},
		{Provider: "maya1a7gg93dgwlulsrqf6qtage985ujhpu068zllw7", Amount: 348_6034764209},
		{Provider: "maya1zvfwm65cmp9hufk3g800f7d2ejx7slrl4mgh07", Amount: 7803_7630550052},
		{Provider: "maya14udggus78e9hh2my7uxnn0l470dp9yj5u35l00", Amount: 399_8764937105},
		{Provider: "maya1v7gqc98d7d2sugsw5p4pshv0mm24mfmzgmj64n", Amount: 5992_3775801240},
		{Provider: "maya1qsynvzys9l63f0ljgr7vk028n4yk0eyvjakn80", Amount: 3000_6158005970},
		{Provider: "maya1fex4zs3psv8crn6dhx4y7kwwpuag3e6a3e4tc0", Amount: 330_8698303993},
		{Provider: "maya1gekecuwh3njjefpyk96lgjqhyg9mr6ry99nsjh", Amount: 65_5279824954},
		{Provider: "maya1j42xpqgfdyagr57pxkxgmryzdfy2z4l65mjzf9", Amount: 91_4013210909},
		{Provider: "maya1v7adg32vxmhhhmul98j23ut3ryr8r93sat4gkw", Amount: 159_4375777734},
		{Provider: "maya17lz0x3a58ew6qfc23ts68z7axyj7n8ymwqyxxh", Amount: 21_5222490083},
		{Provider: "maya189r94lmqg3hf6flgjdmjkemneruma38hugxqj5", Amount: 120_0202887348},
		{Provider: "maya14alj79vk3vfejtgjrgdjv38e23dd3vmrukqryx", Amount: 1288_7423785536},
		{Provider: "maya1qq30ur49s9fs2srkt6vfxq5hdl5q8f6e652q4y", Amount: 328_5224134319},
		{Provider: "maya109xtpvrzd3gmgjhrjzxjtkqg0veskh2jpg69p8", Amount: 53104058},
		{Provider: "maya1ay4u99j6mv7rtwl4nnv7er7fs67vpyrrangxl9", Amount: 136_4851533017},
		{Provider: "maya1q9v6r2g8lznw7ljp2tyv8wp8q2hrr37ms7trth", Amount: 25_2293461748},
		{Provider: "maya183frtejj5ay6wg5h5z9nll46z57hh35t3q8ssv", Amount: 728_4610283867},
		{Provider: "maya17pxhjm53l3du57wck0pr8jfjm38kx4xmyjw3em", Amount: 151_4680330032},
		{Provider: "maya1m0cza4vpan5sgtkz9yjsncl50e34k244c9wjct", Amount: 102_8997967973},
		{Provider: "maya1s2yw6uqyyaut3da8rrxtkufmy4pvysm93usc4j", Amount: 5_9299403199},
		{Provider: "maya1cpjhj27r04zz36gt5enl2jrhumhkc7eg4aqrk5", Amount: 99_7375922070},
		{Provider: "maya1hh03993slyvggmvdl7q4xperg5n7l86pufhkwr", Amount: 1090_1452789657},
		{Provider: "maya17cyy84n4x94upey4gg2cx0wtc3hf4uzuqsmyhh", Amount: 262_9731534176},
		{Provider: "maya18h4dzee602madsllymve86up6xj0s2n2tlwslm", Amount: 244_5755098764},
		{Provider: "maya192ynka6qjuprdfe040ynnlrzf26nyg38vckr2s", Amount: 73_7255366477},
		{Provider: "maya1guh3n0c84quc7szq9twmlxk9tk9fac3mmpeftt", Amount: 42_0294703033},
		{Provider: "maya1wgwrnw63tn7gxmh5j5eg057ey4ddeemzm4ws8w", Amount: 323_9881527618},
		{Provider: "maya13w6dqa772ndgpfv05sae7l4sue08eqcd8layc8", Amount: 12_1691150829},
		{Provider: "maya1xq3eaf70pdw4wl8esn0kyuxpnunprs05tgppzu", Amount: 44_4895759293},
		{Provider: "maya1y8a0lgl8r6pfwzu7apal07f75cquznvzl5kmea", Amount: 147_3384200708},
		{Provider: "maya1s89srqv03vuz9pacrtsdedqcdxjlkpsnxl8e8g", Amount: 417_2670401686},
		{Provider: "maya1fert275f6afn8hnjypzhq75f9vrwfy3uej2492", Amount: 62_6298239658},
		{Provider: "maya1lghvak02n32tlrgm4xvj9zmjr4s7fwx8wyethm", Amount: 14_5557440690},
		{Provider: "maya15n93tthvzldqykev5cs4l3utqhg8v0m2tn22z7", Amount: 10_0852951238},
		{Provider: "maya1smu8qs5dqrxuvqkyf5v9zrf7pa94gm7e2naq9v", Amount: 2_7552253200},
		{Provider: "maya1u40lr4a2fm9eftwj05wxx3v3nwejw4s7st8ufs", Amount: 417_2670401686},
		{Provider: "maya16f8kzx474xwu9rr9ah4mxrny5rq2nhy0yjkrme", Amount: 19_0164365924},
		{Provider: "maya1f5um8t8d68pk2np2vklpsxcnu799k5h4lj2667", Amount: 20_5120846911},
		{Provider: "maya1mfw8c2agx7tmdxt5ez3qsqfmyslagxny0sl7w8", Amount: 37_2591628642},
		{Provider: "maya1jzpntepl8ukadpejf5m2fccy6vygssn6llw98l", Amount: 36_5232494929},
		{Provider: "maya1gnl3j76rglvw3yfttl5vpgryl2gd6y9y2kmuld", Amount: 5_6707515898},
	}

	for _, reward := range rewards {
		providerAcc, err := cosmos.AccAddressFromBech32(reward.Provider)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", reward.Provider, "error", err)
		}

		if err := mgr.Keeper().SendFromModuleToAccount(ctx, BondName, providerAcc, common.NewCoins(common.NewCoin(common.BaseNative, cosmos.NewUint(reward.Amount)))); err != nil {
			ctx.Logger().Error("fail to send provider reward: %s", reward.Provider, "error", err)
		}
	}
}

func migrateStoreV108(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v108", "error", err)
		}
	}()

	danglingInboundTxIDs := []common.TxID{
		// KUJI queue tx's
		"35EA6C99A16E6334980BA7FDC6FC97A863D13A4B7E01D10EA9DFED5265726819",
		"C2EF8A782F23AFCCF0171C201FE0F70EE0E7A09C3C433EC82658013AC5D526C2",
		"07727D48C95A4F207892A27F723C94DE0643B84DA285C7DBAF47DFBB2F8D1FCF",
		"5C5040FBDCAE945F929A162792DCAA556DC5D76164B9E36F80E05D0A9133C27E",
		"B1568DAE825356B16994DDED2BE3F617F4FE511F48B31774F3D7A35DA2EC1634",
		"B2D81261ED12E2A0714FD049BA8CE1DA43480C49F3623F5CB1CCC32A261088A2",
		"CF998E8483AA9C7121F5300013A0EDE8707B9893456150F8CB2E06CB00454005",
		"0E0474BDB2AD1E9634AF38FAB0D24D379D8F60924E2271E6C2B84B52F2F929EA",
		"8D213C1E3C5ADC149D974905185831D5F66094362A34897894DB2AB8256CA91D",
		"BDD0EEC8BCC6043560EA8D3C85D221BE7C65EE8B4079B7CD0D430075734D8770",
		"11A773C27A30CF4217B137D0342D98CA24C8D5DDD8C60E1A5316D8176878617E",
		"C16BC39E7A63A6EABB202DF0834377FE9842F6CDFD088831E57784E8330F534C",
		"6E281DD6DB9BD448C166627D7AEECF3E72812CB571078E6254888A1E1C57AA5B",
		"426890501A1E5AD99E85B3AF846FB601E707130EAFAEDC2D3B7432E4843E549E",
		"855991ECA50787FB71EF9C2A2476674820CDF1A4A68C200D0345C50A8C9AC3A0",
		"0668A945199574EF3037C228D9AA3366DF9C09402613BC69D30CC0E06E7C1295",
		"5481BE9A943BA530CA0608D5BB8E82D9B5125AC61CE78C7931E6247C61B6C7CD",
		"E1B5BD245EBC6C1A568D08C94F8737FAC3898EEA9DAA80896D58C9D947E070B2",
		"F3C665942C66FDC4A3C29D293FEB5B053294C932F3716445F684F152C73987EB",
		"7916EC08954B2DA559E6A9F51FE8F6CB6C23D5F5E879FCAF3C6580BCCF785EA8",
		"A22C8BCBA695588ACF5E83283326F48CEEDDDF367B132F9DB4D5DB66B14BD038",
		"C79BD8A50DCD2CB0BC7F3EFB035B6D8FC4D75B813EBC36CB52B4C1CC7BF2306B",
		"67F93125F8726902AA17668745CA0F2070E923F28D49E5B9A73136FE2D8995D0",
		"876BEDC2D270C23589F1E8C143EFC5C86E1B0AED8B64A39C30223294A84CED1D",
		"73F2D3DB1BDB8D59ADBD275EA00E0FF6EB997FBA38C66EC291200E25D70318C6",
		"2BB72F4BED36EB707ADEB3BE295B8D69DD2001BD90B0A377587E84A56B6A257E",
		"FFA00E2F26988E51AE7937A723EB29DABD5F36F651BEE0E633A70517118E1402",
		"2F8A29325B2F262C1AEE11692BC5C34F662671377FF5EA8B433710114429D5AF",
		"41B24E943899F44E8DF66E8D12434439D36A548C840727E6A98EA4A258FA792F",
		"3B54CE1184C42991E286CF283942BFA2B190850E9C7F14A81E130580599BE567",
		"A6118444880F187C78283471E9C00B8387021F6C58CBA553B3F60A9A88CF994B",
		"F1983522FA5DA3C6EC9EDB513BB22CBBC469ED9086CA8F7CD1E78CF954C327C7",
		"05375C49190B17D76670A38B55D9DC8BBEC973F1C958B7063BE1439AF0589C35",
		"7533DE41CC57CD64A4337F3A2AA4EE755E47AC7F87748227ECA3D7EA6B60F306",
		"5C8264D030990E1AF18CB1E7C804C34E76D5C36A2892B81D3A1196C1B67EB8C2",
		"E8A7CAF2669D64AA7A53E56E3F9B65F606BA5DEE66B76E6B47DB4779CFBCFF57",
		"E27342795F75B8DCD803E5CE41EA8F431A7DC9C1FB0A1D57D8E31DF433BDC12C",
		"C77643F49ACE542D1E1FC15DA080AB61FCF5923405A0D456BC0C6832D123AE46",
		"C5C8CC605ACFA3CEB4E9D7C2AB7F02D4864E684EBEDFC559B730A579F27B09EB",
		"3812A75DD69D06D9E04A85F28DF32B94EAE1FB3786CA2029540CF5E6EFE749BA",
		"F2EF82E480A1B96EF15B4CB8CBE8C6C5CE2A6E3224BD5DDF27A924F7A7C5162C",
		"484A7B1BABCA27092442F6D45C3FE4E9D86E29857A2FF4FE51EE7FBFEB8D2843",
		"70675963939D19384AEC8BB88BBC3E05E28712691F907BCCDA0BBFA0995058EC",
		"220F7F531D372B182195050C3A409F14891C2872C98C7A10629B07C09C485AD4",
		"ADC8FBB736F45FD7B8E97B4ABB2210611A8E09562DD13A5989DA0E9CAC506490",
		// Refunds consensus failure
		"A07B4C55031F3B9FED8D3BF938740981149A41274CA2B33F54DB83091A75D635",
		"4AB40F14C011C733A0817E65EB5B6395D206C35D56C7912F9EB8207C49AEB20E",
		"E4AA05181989014F462A1BF969D77FCCA2D5E3C7106A895057D80B3E9733FBFA",
		"4C98B47D67ACEB5515BC39A3F053D8BCDD3FE469ABAD7AE042894261868A30E1",
		"FC27513B337E258CC9237638F3F22A370498B659862C39AE9A02E09E9674E8C1",
		"EC1BDC5295F6B98EBD9D472082D15E35319DF9A3FD0088F5CB4E528AA0062A74",
		"0BA546F56B63B131453F1F441D2B200511A6C88CFCD9F196A2759055727EC7A7",
		"5B1F7833601ED07F4DDA87B7A1CD306B981520AE6E934C0D2B60FAE790C22725",
		"A0702FDCB62879B183EC7D565C99BCBFE31653391E3AD181F6D46C84AE35E001",
		"DCC40487D914DB53E7B049D7F3A825743E97C39A6A07687087937AB24475E3EA",
		"C2D445428DCDB2CEEB66770B9E1E462FAC52979AC08E833696E13607E4CD6EAA",
		"8FCA78BD4B74416F00B5505F02EBAF83340F355FFD0729D68BF5E14F08B679E3",
		"3DC8D5E627AF4ACEADBD389EBEB7310CE868F152704B49073DF6C31ECD783E7E",
		// KUJI dropped txs on 08/Dec/2023
		"D7B7AC85DACF35B9FA3F912946A1977398C8F085FAD6C0FC2F13036C726A952E",
		"9F1E084FD5A7EFD9B6A7C304944DB266300C28D91D006ECB0C7E113755D574A0",
		"1A88CDB0AED6FE0047AD1CD27F08002039D584BD02BFE1CF3CB1FF7DACFD4C99",
		"EA33A352BAA5397AFB1F0EC1C30A8FC6B937D49704EFE46436BEF06B9224E99D",
		"99A7F15CF990E8B941A4960CD3E1FE3D8C6F1356891606A4F59DA3C221CA2412",
		"52A5093AF70F6B497E757AB9BE46CF24007C4DD28386EE1BB98AF8B6FD07B4E5",
		"5FB7014A1D750B8D33F22E74528028AD8EBCEC298378793FB99BB8282800BEAE",
		// KUJI dropped txs on 14/Dec/2023
		"2676A985928181376F147C5FEFFA5723B0E001A9C799256B0286AC97FDBED9C4",
		"DD410481BCFEDB5B54298948903EBCB61A55FC610A72DB9AFA92AB4430A43AAE",
		"10F87DEC374B1091C3F9B52735DDA6D0BD262284220952BAAE5FE8F5B007824E",
		"7702783493C209EE0E5879308D93523FBF84B857C39F35BCC05284CECE59ED52",
		"EE95B8AE9CB9D0413348A81C1B30FFA064B8B7CDCF9E9131BDF898707A9F44A1",
		"E395D6A8A748B2FCE7B4251ACADE8B9A4F298F2E93D3560D41C5BA28EE6032BE",
		"E9DD4695E613CE9AAEF2F7BF3BA17647961D2CBADFED08B0DEDA5537404F16C3",
		"0D02DD74026C1EA90FFAC809F811AA349207BEABF2454DA4576FEBD1F5EB7C44",
		"94B17D1B4A9F81E595C6515D037592971F872FFCEB1B525C7EFE1ABBC08EA789",
		"0CAF6DCAF0F9BB2583E639C4DC80B8DD67B0632917EB0DD3B75CA1A27483BE4F",
		"568905978966B1296AB21616479BB9C9C56BE84107F9AA081E04424A0E87DE4F",
		"1B015385AA71EEBE675E3046522E689C9A7F3CE09B1548573CB428A381CD5DCC",
		"46F1A6B6BB222F29DFDAB39BF08BE4357EB0A926CBD07B7252C6C0270A34ED85",
		"60660C31363CD87A6BED81F59D81035FFE274904E457B30CD464C0CF92BC0BC4",
		"C8DB11D3B0A92BF661370947BC23AC2CE2037C4470CD53713D9A0005A35A742A",
		"3892484F2D8F63979148678244622FB17C1F2C8C7324A0D51A95C07FF69B56CB",
		"6219176D387DF508DB80A5DA97F902B396041B936CAA99449F778D991D27E9F1",
		"059F759A98FC1B9AC8D088575EF34C1C718FB98094312FD723523456A2DB27C4",
		"8356489F25D45205C9894F26D4203FDE054E79EFB124272844FF677895A77B8F",
		"AFA4BB32162239BE0BE36C22870E7660C4AF6E6B95A92BE013D4D38BC6C6325C",
		"5EAFA060D7B171DAA2B5D718AFA119313A91C9CE5873CB1A793E18469BD8F591",
	}
	requeueDanglingActions(ctx, mgr, danglingInboundTxIDs)

	// Unbond bond providers from node account
	unbondAddresses := []unbondBondProvider{
		{bondProviderAddress: "maya1f5um8t8d68pk2np2vklpsxcnu799k5h4lj2667", nodeAccountAddress: "maya1v6lt70lqkhxhftlpx26d52ryzc6s3fl4adyzuv"},
		{bondProviderAddress: "maya1gnl3j76rglvw3yfttl5vpgryl2gd6y9y2kmuld", nodeAccountAddress: "maya1v6lt70lqkhxhftlpx26d52ryzc6s3fl4adyzuv"},
	}
	unbondBondProviders(ctx, mgr, unbondAddresses)

	// Send node rewards to each of the bond providers
	type providerReward struct {
		Provider string
		Amount   uint64
	}

	// Rewards getting paid out because first few churns they weren't distributed and BPs not able to claim.
	//
	// Address that have been already paid because of:
	// Store migration v107 or
	// BOND/UNBOND rewards distribution
	//
	// maya10sy79jhw9hw9sqwdgu0k4mw4qawzl7czewzs47: 0
	// maya12amvthg5jqv99j0w4pnmqwqnysgvgljxmazgnq: 0
	// maya1gv85v0jvc0rsjunku3qxempax6kmrg5jqh8vmg: 0
	// maya1q3jj8n8pkvl2kjv3pajdyju4hp92cmxnadknd2: 0
	// maya1vm43yk3jq0evzn2u6a97mh2k9x4xf5mzp62g23: 0
	// maya1gczk5e3slv35y35qyw0jc6jwudm2jg4ztscc5x: 0
	// maya1xfuxhzj2e63yd37z87vmca25v5n8486an9yde2: 0
	// maya1s89srqv03vuz9pacrtsdedqcdxjlkpsnxl8e8g: 0
	// maya189r94lmqg3hf6flgjdmjkemneruma38hugxqj5: 0
	// maya1cujc2sj8avcfnyrxj9grwlcfhyflpxchvq65cg: 0
	// maya14alj79vk3vfejtgjrgdjv38e23dd3vmrukqryx: 0
	// maya1jzpntepl8ukadpejf5m2fccy6vygssn6llw98l: 0
	// maya1gekecuwh3njjefpyk96lgjqhyg9mr6ry99nsjh: 0
	// maya178wqee3z5y9fyqxkdud4rxyldlytq9d6xcs823: 0
	// maya1hh03993slyvggmvdl7q4xperg5n7l86pufhkwr: 0
	// maya1rzr9m407svj4jmc6rsxzsg75cx7gm3lsyyttyj: 0
	// maya1v7gqc98d7d2sugsw5p4pshv0mm24mfmzgmj64n: 0
	// maya1tndazzezsfka2wgqm52e5neej9q8jqrxv47h7m: 0
	// maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz: 0

	rewards := []providerReward{
		{Provider: "maya1zvfwm65cmp9hufk3g800f7d2ejx7slrl4mgh07", Amount: 26725_3597627890},
		{Provider: "maya183frtejj5ay6wg5h5z9nll46z57hh35t3q8ssv", Amount: 19582_1211035418},
		{Provider: "maya1m0cza4vpan5sgtkz9yjsncl50e34k244c9wjct", Amount: 9633_7253828425},
		{Provider: "maya1mfw8c2agx7tmdxt5ez3qsqfmyslagxny0sl7w8", Amount: 9550_5124319582},
		{Provider: "maya1qq30ur49s9fs2srkt6vfxq5hdl5q8f6e652q4y", Amount: 7329_8887338551},
		{Provider: "maya1s7naj6kzxpudy64zka8h5w7uffnzmhzlue4w3p", Amount: 6488_7055306094},
		{Provider: "maya17pxhjm53l3du57wck0pr8jfjm38kx4xmyjw3em", Amount: 5759_6226579937},
		{Provider: "maya1fert275f6afn8hnjypzhq75f9vrwfy3uej2492", Amount: 5635_7251152517},
		{Provider: "maya1guh3n0c84quc7szq9twmlxk9tk9fac3mmpeftt", Amount: 5518_1294840549},
		{Provider: "maya1g70v5r9ujxrwewdn3w44pmqcygx49dx7ne82vr", Amount: 5097_3527408652},
		{Provider: "maya17cyy84n4x94upey4gg2cx0wtc3hf4uzuqsmyhh", Amount: 4397_8618563702},
		{Provider: "maya1cpjhj27r04zz36gt5enl2jrhumhkc7eg4aqrk5", Amount: 4323_0939760062},
		{Provider: "maya18h4dzee602madsllymve86up6xj0s2n2tlwslm", Amount: 3916_3506275542},
		{Provider: "maya14udggus78e9hh2my7uxnn0l470dp9yj5u35l00", Amount: 3883_8904268868},
		{Provider: "maya1qsynvzys9l63f0ljgr7vk028n4yk0eyvjakn80", Amount: 3290_5472396829},
		{Provider: "maya1ay4u99j6mv7rtwl4nnv7er7fs67vpyrrangxl9", Amount: 2326_7311031468},
		{Provider: "maya14dsp7ujkrxqzsv2h2x68ypaeevmg4r5z7500c9", Amount: 2273_5110342539},
		{Provider: "maya14sanmhejtzxxp9qeggxaysnuztx8f5jra7vedl", Amount: 2163_2951513050},
		{Provider: "maya1lghvak02n32tlrgm4xvj9zmjr4s7fwx8wyethm", Amount: 1783_8529169711},
		{Provider: "maya1v7adg32vxmhhhmul98j23ut3ryr8r93sat4gkw", Amount: 1352_3710961279},
		{Provider: "maya15n93tthvzldqykev5cs4l3utqhg8v0m2tn22z7", Amount: 1222_4490527092},
		{Provider: "maya192ynka6qjuprdfe040ynnlrzf26nyg38vckr2s", Amount: 1154_1855063351},
		{Provider: "maya1u40lr4a2fm9eftwj05wxx3v3nwejw4s7st8ufs", Amount: 1091_9836106714},
		{Provider: "maya1q9v6r2g8lznw7ljp2tyv8wp8q2hrr37ms7trth", Amount: 943_1758010293},
		{Provider: "maya1yk4xsaye2m37ytgzulzpr5ajvhvqhg68rpw7ff", Amount: 884_5660977299},
		{Provider: "maya1fex4zs3psv8crn6dhx4y7kwwpuag3e6a3e4tc0", Amount: 861_5494393883},
		{Provider: "maya16f8kzx474xwu9rr9ah4mxrny5rq2nhy0yjkrme", Amount: 835_7844832722},
		{Provider: "maya13yseu9un5f9gwqgzshjqvsqrxew0hhgm3wjh4l", Amount: 780_1025015833},
		{Provider: "maya1j42xpqgfdyagr57pxkxgmryzdfy2z4l65mjzf9", Amount: 727_1695244006},
		{Provider: "maya19jqjqnc7hmvrfez8p5z2tcfmfmq9k5z3wm0rq9", Amount: 629_3112263003},
		{Provider: "maya1a7gg93dgwlulsrqf6qtage985ujhpu068zllw7", Amount: 625_3970832338},
		{Provider: "maya1s2yw6uqyyaut3da8rrxtkufmy4pvysm93usc4j", Amount: 555_1752125982},
		{Provider: "maya1g7c6jt5ynd5ruav2qucje0vuaan0q5xwasswts", Amount: 525_8771743671},
		{Provider: "maya1smu8qs5dqrxuvqkyf5v9zrf7pa94gm7e2naq9v", Amount: 333_9146989413},
		{Provider: "maya1xkdt3ld8xtlfpztdp0k05tmf9g3q622lmahjr2", Amount: 311_7886647361},
		{Provider: "maya1y8a0lgl8r6pfwzu7apal07f75cquznvzl5kmea", Amount: 309_9920691612},
		{Provider: "maya1szmq6kkplsqn7k8lwsm6xajxzgvak0gjvm8c8w", Amount: 260_6258748096},
		{Provider: "maya1ewz79pg6qylpk0p98yzr6jhv23s4wrn0jcnard", Amount: 242_8304775521},
		{Provider: "maya1gnl3j76rglvw3yfttl5vpgryl2gd6y9y2kmuld", Amount: 239_3588509926},
		{Provider: "maya1pf7gg2h9kdq7zuj58r7wk8py99awwj9lwvchdx", Amount: 233_7043832402},
		{Provider: "maya1m7xnnkkrk7e6aa4eq3yndy4nlcre037xnf3zjz", Amount: 224_2073520353},
		{Provider: "maya1f5um8t8d68pk2np2vklpsxcnu799k5h4lj2667", Amount: 217_5743877338},
		{Provider: "maya1h64fpu5uwmzku4xynfc6sevqfpjxp4y36a4t00", Amount: 216_0778531297},
		{Provider: "maya14u40pul8pgpuk42k9502jq5r3wfrpnv9ly8e2j", Amount: 216_0385259062},
		{Provider: "maya1v54s0rwazm5k3ywhaz5rvwnneuccr7rtmqm5yz", Amount: 213_7449170907},
		{Provider: "maya17lz0x3a58ew6qfc23ts68z7axyj7n8ymwqyxxh", Amount: 185_9191116014},
		{Provider: "maya1sclplk79vvlakl8u54r0gr622jfuwar0vfl2l7", Amount: 173_4379838993},
		{Provider: "maya10sdhv0cn0fsfgax6vpzv9pwy8r5872hw3h4tuh", Amount: 169_4603734918},
		{Provider: "maya16k0al0fsslhx8j5cjsjsv4ntmq45sgew8waryj", Amount: 166_4425338415},
		{Provider: "maya1xq3eaf70pdw4wl8esn0kyuxpnunprs05tgppzu", Amount: 162_9076206105},
		{Provider: "maya1c6qrsnstl9l0wtc3fazd6jrfppshs6jk2myeky", Amount: 129_6562049931},
		{Provider: "maya1x64thscxsl39pun3lqzwhge890tvhvgd36c5gs", Amount: 122_3416961124},
		{Provider: "maya1hkqc78uhuc4z8qtt3qjsdn0u7348t2hhlgyzh9", Amount: 119_0618062398},
		{Provider: "maya1vu37n7h7mnk0uxakye2vhh2z2k5cehf6v2lk3r", Amount: 105_5831448383},
		{Provider: "maya18p22jfv43weeyznqg0h9f6dh3adnpj4nwch8hs", Amount: 100_2736412691},
		{Provider: "maya1v7jsyf94rnfdx5v0xjxn5c8vdsyvmym0aegl7k", Amount: 91_0781175605},
		{Provider: "maya1tdp957gs94j7ahgd6cemlunhrwd27l39e52l6l", Amount: 90_9706270720},
		{Provider: "maya1aderqdry6m6vr4qtzkpe5n36xefemfph79pv4a", Amount: 72_2430923098},
		{Provider: "maya18jtxyr7seydqq6q2enhq3c6hx6zc4s8y440swt", Amount: 72_0700051859},
		{Provider: "maya175dn4q74ztt7wzf2n5u0nqkmfvda5sc627vtvd", Amount: 70_2985614795},
		{Provider: "maya1xmn5ecq45fasyt7xqm8nefg8fvpf0w7zqtn2tq", Amount: 66_2528894056},
		{Provider: "maya1nv96km7hgmv76rsjcjj5qmx5ml53alf9r8fy22", Amount: 64_2201125917},
		{Provider: "maya1swcvf06tsytaalk7y6t3urnwyv435gu8fly77g", Amount: 58_0278367469},
		{Provider: "maya1ccf7rs4z6y2spvpmdf7v66v7xy2rd8dye7jhrr", Amount: 55_3287635573},
		{Provider: "maya138cjvf52rr7v5zp6s2gemu0m9wx593juprpgnl", Amount: 38_2125074036},
		{Provider: "maya1f8j08d6p7pqtuhjzcm9297gq4kvhv5lz4p5pma", Amount: 30_5106110499},
		{Provider: "maya19z4xlhxp6hkqe4mlfmqwsnjahrpa3ycjflqczc", Amount: 30_1740875619},
		{Provider: "maya13w6dqa772ndgpfv05sae7l4sue08eqcd8layc8", Amount: 29_8000796134},
		{Provider: "maya1xrn6rw99ncj0qxflwtmvjeuf4kkwuwja4xpwhv", Amount: 27_4001705664},
		{Provider: "maya1wgwrnw63tn7gxmh5j5eg057ey4ddeemzm4ws8w", Amount: 18_5817869273},
		{Provider: "maya1y6lk677q4gdy75qm5x3q4t0sxvx40r8n2kcc4s", Amount: 17_8549597986},
		{Provider: "maya1g286wstwf4vqmegj5324p58gxmy7mnmha80hgz", Amount: 17_1346960497},
		{Provider: "maya1vtzdhyl9sfxh965euupyawn2ql6aa3ee37wz8l", Amount: 9_3962917655},
		{Provider: "maya10n2xw02y4wvv64qulnhmgjdryktzq3nhd53f6x", Amount: 6_5444411103},
		{Provider: "maya1adkthl5cd6h4atrdvxt7tp9xnwu3xpn89c7flu", Amount: 9688364232},
		{Provider: "maya1z4dyge20n7c6g87txma7lv8qmmzvluv2crn8pl", Amount: 8270205310},
		{Provider: "maya109xtpvrzd3gmgjhrjzxjtkqg0veskh2jpg69p8", Amount: 5811384939},
		{Provider: "maya1ha4ypeghxhtdu63dqhhkspqcu4s7375kc3ch4u", Amount: 5735026761},
		{Provider: "maya1nwe0vs65myamknwehgr00r5t2afrlpn26du4vt", Amount: 4822957294},
		{Provider: "maya1kzd9fj58g9exxt44lj8sfzuvc94tsrr2v4gv6g", Amount: 1792051180},
		{Provider: "maya1jttfwrve7mcjfnhnsavpnfzeql4mr5mjns0wpj", Amount: 1250657295},
		{Provider: "maya1ajzlu2p2mnecl6q739fn7hsctlwxyqdulwsslg", Amount: 1210922557},
		{Provider: "maya1ngzyvjtr2xeh4gesxj4wtgl9jxgp2jf3fueah6", Amount: 1094496407},
		{Provider: "maya1kgma45rn0qs22pd45smp2pakng7qz42d2mxmch", Amount: 38386732},
		{Provider: "maya1gmf8lt6ddlq3f0skq77pdskfs4cjz4wcc2s82y", Amount: 14303230},
	}

	for _, reward := range rewards {
		providerAcc, err := cosmos.AccAddressFromBech32(reward.Provider)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", reward.Provider, "error", err)
		}

		if err := mgr.Keeper().SendFromModuleToAccount(ctx, BondName, providerAcc, common.NewCoins(common.NewCoin(common.BaseNative, cosmos.NewUint(reward.Amount)))); err != nil {
			ctx.Logger().Error("fail to send provider reward: %s", reward.Provider, "error", err)
		}
	}

	// Manual Refunds done by Maya team, these are the in_hashes of all manually refunded txs:
	// A62B2036B49D3CA8F8A7FB8A5041BFE21AB4E0CFEA57A6FBAB383FAC84A98911
	// 49B54AD04019CAC6907242D687CA9ADF6BC4C5C69D4EA0C91CE0C9ED76225593
	// 6C0A39FEB57C750FA284D09709E0CCFB0DD30FFDED2067D380BEF8499EE23B51
	// E948D68957DDBEE943723F18DEC8A5A5E66358B0619B894A9C5C160B28EEBB3B
	// 7DC12CA54AC3BFC1EA5177D97C68D32C5C5F8FC344BC92299A151A92E1A6E3D0
	// 17454BAB890705F32C1706A696C16B05464EEAAA38F1BCA915CA7E31BE5FB55B
	// 282A9C93E2351CD22FFFD63C976E2938155F998E2A6BA560FE457A0C010983B4
	// 0AEEEA9FC6E615144C67DE33DEC518631616F58C3397F9C33A82542FB9F7F4CF
	// 548DA2CB06BEA4A7814A389C7E01E30CA7077CEB11B8B95FB1DE21C987C9CDED
	// 140CB4F389366B4A47ECD05EBBD2CB9C051D1AB9766B030F926C4C7A85F8009C
	// 40B0EDF61AA5B22500B2AF2BA80B91ED92633D4797E02E40CEE2D95E99A4CC11
	// E41E7AC9E4A2AE3EEB3B8096CF7B0C4044CF4972029EA68ECC883D50E79E4942
	// 7FA5741EF1D4A653ED3FDB87BC1C167674C17310B9A2B6686365791830F1A788
	// AC2406FDF0C3EB99C3D2D5404587670A86F92C9D9C3F643ED5C503EF5BF46E6E
	// 332E53DA379486730EBF44992AFB8EEBBA93D14F950C0D85AC414A61A692A283
	// F61FBE9EA534D7B5D0B88ED8DA9C926E5DA0DD85F045D8B04B4DE23B9D7D0993
	// FC27513B337E258CC9237638F3F22A370498B659862C39AE9A02E09E9674E8C1
	// 5DBEA8B5591593F03FD9A16D26246EE8DEE581A9798E7F0CA002060604B30AB8
	// 62193BBC0CF9559A8A115269A0903D1B0E9779C5AF9DB6E7FE9A1620BDA65563
	// 5A139A466370218C956C7B9E893E471F235BA6123D51FA8EAFC50BA78D965262
	// 9791E51AF7B54C204B8FE8BE50D98CE0046B4D6BD00A0634A696393A6D95AB27
	// 48856566752CF7082A1ABE3C4452C6496BEEBBC64719CF6222DB3BD5F1D41A5E
	// 49361D535CBDDC2DA954369F23D69147431ECCD3102AB0C0A1FB2651D63A3752
	// 2D3EAC94FB828518B38C7FF122C29B826BC03AF372568ABC3621A4C6A048531A
	// A07B4C55031F3B9FED8D3BF938740981149A41274CA2B33F54DB83091A75D635
	// 7032A84A54D38A4F3FA9A35875848798953472EC5DD24FF232DB5E5DFE906D46
	// 3EA5E2D4E3BA9F5FDF670BA44010AA10D3DD995A6A00EDCAAAA7490901EDA9D6
	// D237147312F23FA85FD207948E24594093AFB883BD6FF21CAC828F1734E7048B
	// B997BBE82DC1654267463CA03C49A195609DDC9D63CFED9608458A44EB808D45
	// 59AFDCDFA99A9194D76CA15957CF1676C11373AC67CEB600AA4FE239EE0CDE75
	// 17DB507BCA325199E4D0F5A238348073743984FE2578A395B0437449BDA2B5E4

	// These are the txid's of all manually refunded txs, or out_hashes
	// https://finder.kujira.network/kaiyo-1/tx/0063EF3F70AD7325AB8756F2CBEA39C4C387FF86EC8B07C00D9F5FDD95D63449
	// https://finder.kujira.network/kaiyo-1/tx/12A46622697E00E67140FEBC124E92A516A29768C844B3B79CFD6E41BE0FC3F9
	// https://finder.kujira.network/kaiyo-1/tx/171FF77979A923898E1648C9FCB0CC30459632AF8EC89175A8E97144D6C5F053
	// https://finder.kujira.network/kaiyo-1/tx/44F321B54C26F49C39D6E465025453238CCAAFCA7857C89BC3740EAD844A8169
	// https://finder.kujira.network/kaiyo-1/tx/4767322ACA41F871FAF4560FB26C4DEA964216ED5460AE789C27071F848C2400
	// https://finder.kujira.network/kaiyo-1/tx/84C62A4902DB6C7AD0BC456FDD4645482EBAB999687150C06967C1CF68F7036A
	// https://finder.kujira.network/kaiyo-1/tx/85E9644E6D6D321CD235D3F15B595E3AF7A28AEAEC81FCCC7F33C31261748765
	// https://finder.kujira.network/kaiyo-1/tx/8B791123AD76BD05F33F130AA197EF020B39BED0233F715D9D8B609E04964272
	// https://finder.kujira.network/kaiyo-1/tx/92DA2232C9D06C46354CFF616CB05CA5D27BA39214AF11A5111BA21C98828F17
	// https://finder.kujira.network/kaiyo-1/tx/C458CC31B6908D86271E3D1090B9590073BD57A901361DD8A1EEEBF11B0AFC0C
	// https://finder.kujira.network/kaiyo-1/tx/E6D61FEDF8925099B4D489BFC92CE52FE64C275208311D7367A774F047851754
	// https://finder.kujira.network/kaiyo-1/tx/E6D61FEDF8925099B4D489BFC92CE52FE64C275208311D7367A774F047851754
	// https://finder.kujira.network/kaiyo-1/tx/E6D61FEDF8925099B4D489BFC92CE52FE64C275208311D7367A774F047851754
	// https://finder.kujira.network/kaiyo-1/tx/E6D61FEDF8925099B4D489BFC92CE52FE64C275208311D7367A774F047851754
	// https://finder.kujira.network/kaiyo-1/tx/EA2294FCC7E1C8BFF8E2692011293EA02BD79FE73CE3617E79CF96049101A60B
	// https://finder.kujira.network/kaiyo-1/tx/F7EE8C3FF401C5F70E536F4255B815A2965F826975312F3DBC241587C630B137
	// https://finder.kujira.network/kaiyo-1/tx/F964F5C87EF1C97EBC0B9663A745BB1E8BB2005CC22F0028849E77EA6D5E812E
	// https://runescan.io/tx/6128A5AA0882F994E101CB7792EA6779AADA6C15AF6DE734670DA8144859176F
	// https://viewblock.io/thorchain/tx/136B9EB5B7FF520A62DAFCBEBEFAA262B42BF5520967970A580057418E55A84B
	// https://viewblock.io/thorchain/tx/136B9EB5B7FF520A62DAFCBEBEFAA262B42BF5520967970A580057418E55A84B
	// https://viewblock.io/thorchain/tx/136B9EB5B7FF520A62DAFCBEBEFAA262B42BF5520967970A580057418E55A84B

	manualRefunds := []RefundTxCACAO{
		// Total refunds:
		// 38,500.00	CACAO
		// 13,360.00	KUJI
		//  5,231.00	RUNE
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(151447_9545000000)},
		// Kayaba ImmuneFi bug bounty for TSS vulnerability: $38,500
		// https://etherscan.io/tx/0xdae2f4f89726f8d71a2802958dc80e30edf63aafcb725c8396c26f64d756793d
		// https://etherscan.io/tx/0x578024d6aac35ba5985989ee2b80ecde9484a80b189cc2649f52ed9c6d2a1633
		// https://etherscan.io/tx/0xf596f21b2717cc1b1730f3f52fa95af3aac1c9b607eaf2b63e17de484705476f
		// https://etherscan.io/tx/0x8980713f6dbe62b4a1c971ffd26806698603e7733bb1bc7f4c0373f5ee6a7df7
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(39772_7272700000)},
		// From Reserve
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(230000_0000000000)},
	}
	refundTxsCACAO(ctx, mgr, manualRefunds)

	// Send CACAO due to invalidad memo (not specifying kuji address or using maya address instead for non synth asset)
	swapKujiFail := []RefundTxCACAO{
		// Tx Hash: 2E9E6A8BC525635892E1809DC5BBAC9FE7906A41AA821B174DCFEDF130042C32 memo: =:kuji.kuji:
		{sendAddress: "maya1xv2tqx22t4666awa2eyu4uayucxwyy6jk05yxn", amount: cosmos.NewUint(60000000000)},
		// Tx Hash: D754704FEEAA474F27665AB7D19351543E7BC4AFC5AA1F07453B6ED36717535A memo: swap:kuji.kuji
		{sendAddress: "maya1a7gg93dgwlulsrqf6qtage985ujhpu068zllw7", amount: cosmos.NewUint(27700000000000)},
		// Tx Hash: E7E298C775F1480027907ACAD29F6A7E872D633F4EFBB1E28FE16577BA8D5257 memo: swap:kuji.kuji
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(46560000000000)},
		// Tx Hash: 2C64F982BB6FE20EEC54C3F05E277EE7EFC1762CB43E50FA3C047B80B34481B3 memo: swap:kuji.kuji
		{sendAddress: "maya1m3t5wwrpfylss8e9g3jvq5chsv2xl3uchjja6v", amount: cosmos.NewUint(79310000000000)},
		// Tx Hash: 59399A6F94707513314DA42B8D0490A2C9C7F4395FB60385F3C87B827A541484 memo: =:KUJI.KUJI:maya1kh5lvr8msnvwpgl4sdk8r572d29dtvw3wq5j69
		// Using WSTETH = 2577 and CACAO = 0.74
		// (0.20443 * 2577) / 0.74 = 711.9136621622
		{sendAddress: "maya1kh5lvr8msnvwpgl4sdk8r572d29dtvw3wq5j69", amount: cosmos.NewUint(711_9136621622)},
		// Tx Hash: 99339C24ACF9C1EE3239F3F0EA11A050CC1F8A39923407DAA1ECE0F3F55E369E memo: =:KUJI.KUJI:maya1sqj6zr2wy762rg9ugn7l3s599xczl4uvetv4z6
		// Using ETH = 2233 and CACAO = 0.74
		// (0.15 * 2233) / 0.74 = 452.6351351351
		{sendAddress: "maya1sqj6zr2wy762rg9ugn7l3s599xczl4uvetv4z6", amount: cosmos.NewUint(452_6351351351)},
	}

	refundTxsCACAO(ctx, mgr, swapKujiFail)
}

func migrateStoreV109(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v109", "error", err)
		}
	}()

	// Send node rewards to each of the bond providers
	type providerReward struct {
		Provider string
		Amount   uint64
	}

	// unpaid BPs from 2023-12-08 to 2024-01-25
	rewards := []providerReward{
		{Provider: "maya13yseu9un5f9gwqgzshjqvsqrxew0hhgm3wjh4l", Amount: 586_7249227958},
		{Provider: "maya1qq30ur49s9fs2srkt6vfxq5hdl5q8f6e652q4y", Amount: 1186539066},
		{Provider: "maya109xtpvrzd3gmgjhrjzxjtkqg0veskh2jpg69p8", Amount: 1162261524},
		{Provider: "maya1qsynvzys9l63f0ljgr7vk028n4yk0eyvjakn80", Amount: 2105_2928566590},
		{Provider: "maya1fex4zs3psv8crn6dhx4y7kwwpuag3e6a3e4tc0", Amount: 310_3463000064},
		{Provider: "maya1adkthl5cd6h4atrdvxt7tp9xnwu3xpn89c7flu", Amount: 1451832905},
		{Provider: "maya1ngzyvjtr2xeh4gesxj4wtgl9jxgp2jf3fueah6", Amount: 150846361},
		{Provider: "maya1z4dyge20n7c6g87txma7lv8qmmzvluv2crn8pl", Amount: 1287912860},
		{Provider: "maya1gekecuwh3njjefpyk96lgjqhyg9mr6ry99nsjh", Amount: 110_0299240344},
		{Provider: "maya1j42xpqgfdyagr57pxkxgmryzdfy2z4l65mjzf9", Amount: 362_2294826144},
		{Provider: "maya1v7adg32vxmhhhmul98j23ut3ryr8r93sat4gkw", Amount: 1019_6008344068},
		{Provider: "maya17lz0x3a58ew6qfc23ts68z7axyj7n8ymwqyxxh", Amount: 79_7492876313},
		{Provider: "maya1ay4u99j6mv7rtwl4nnv7er7fs67vpyrrangxl9", Amount: 1348_8535910543},
		{Provider: "maya1szmq6kkplsqn7k8lwsm6xajxzgvak0gjvm8c8w", Amount: 554_5526331001},
		{Provider: "maya16k0al0fsslhx8j5cjsjsv4ntmq45sgew8waryj", Amount: 383_0590203067},
		{Provider: "maya1q9v6r2g8lznw7ljp2tyv8wp8q2hrr37ms7trth", Amount: 127_0409512208},
		{Provider: "maya183frtejj5ay6wg5h5z9nll46z57hh35t3q8ssv", Amount: 2584_1288629325},
		{Provider: "maya17pxhjm53l3du57wck0pr8jfjm38kx4xmyjw3em", Amount: 779_1406136992},
		{Provider: "maya1zvfwm65cmp9hufk3g800f7d2ejx7slrl4mgh07", Amount: 2610_1017855368},
		{Provider: "maya1m0cza4vpan5sgtkz9yjsncl50e34k244c9wjct", Amount: 1173_7549946213},
		{Provider: "maya1s2yw6uqyyaut3da8rrxtkufmy4pvysm93usc4j", Amount: 67_6415045011},
		{Provider: "maya1guh3n0c84quc7szq9twmlxk9tk9fac3mmpeftt", Amount: 895_2277036614},
		{Provider: "maya1ha4ypeghxhtdu63dqhhkspqcu4s7375kc3ch4u", Amount: 1076643315},
		{Provider: "maya1nwe0vs65myamknwehgr00r5t2afrlpn26du4vt", Amount: 919619312},
		{Provider: "maya1y6lk677q4gdy75qm5x3q4t0sxvx40r8n2kcc4s", Amount: 193_9880334042},
		{Provider: "maya1swcvf06tsytaalk7y6t3urnwyv435gu8fly77g", Amount: 196_9506511639},
		{Provider: "maya1xmn5ecq45fasyt7xqm8nefg8fvpf0w7zqtn2tq", Amount: 557_1829671102},
		{Provider: "maya10n2xw02y4wvv64qulnhmgjdryktzq3nhd53f6x", Amount: 74_8255850119},
		{Provider: "maya1fert275f6afn8hnjypzhq75f9vrwfy3uej2492", Amount: 1468_3170205498},
		{Provider: "maya1lghvak02n32tlrgm4xvj9zmjr4s7fwx8wyethm", Amount: 218_6476571398},
		{Provider: "maya15n93tthvzldqykev5cs4l3utqhg8v0m2tn22z7", Amount: 153_1323322462},
		{Provider: "maya1smu8qs5dqrxuvqkyf5v9zrf7pa94gm7e2naq9v", Amount: 41_8405069148},
		{Provider: "maya1ccf7rs4z6y2spvpmdf7v66v7xy2rd8dye7jhrr", Amount: 130_4120103595},
		{Provider: "maya1x64thscxsl39pun3lqzwhge890tvhvgd36c5gs", Amount: 221_6570161033},
		{Provider: "maya1nv96km7hgmv76rsjcjj5qmx5ml53alf9r8fy22", Amount: 305_9885056479},
		{Provider: "maya1ewz79pg6qylpk0p98yzr6jhv23s4wrn0jcnard", Amount: 429_9454195178},
		{Provider: "maya1kzd9fj58g9exxt44lj8sfzuvc94tsrr2v4gv6g", Amount: 312366982},
		{Provider: "maya1ajzlu2p2mnecl6q739fn7hsctlwxyqdulwsslg", Amount: 190156170},
		{Provider: "maya1jttfwrve7mcjfnhnsavpnfzeql4mr5mjns0wpj", Amount: 184190874},
		{Provider: "maya14udggus78e9hh2my7uxnn0l470dp9yj5u35l00", Amount: 651_7631226817},
		{Provider: "maya19jqjqnc7hmvrfez8p5z2tcfmfmq9k5z3wm0rq9", Amount: 886_8425791225},
		{Provider: "maya1c6qrsnstl9l0wtc3fazd6jrfppshs6jk2myeky", Amount: 246_2729280723},
		{Provider: "maya1xkdt3ld8xtlfpztdp0k05tmf9g3q622lmahjr2", Amount: 439_3811075027},
		{Provider: "maya1sclplk79vvlakl8u54r0gr622jfuwar0vfl2l7", Amount: 244_4135469557},
		{Provider: "maya1v7jsyf94rnfdx5v0xjxn5c8vdsyvmym0aegl7k", Amount: 337_5376073612},
		{Provider: "maya1g286wstwf4vqmegj5324p58gxmy7mnmha80hgz", Amount: 179_4228504620},
		{Provider: "maya1xrn6rw99ncj0qxflwtmvjeuf4kkwuwja4xpwhv", Amount: 365_3826916897},
		{Provider: "maya16f8kzx474xwu9rr9ah4mxrny5rq2nhy0yjkrme", Amount: 237_7645982361},
		{Provider: "maya1mfw8c2agx7tmdxt5ez3qsqfmyslagxny0sl7w8", Amount: 969_6826822177},
		{Provider: "maya1s7naj6kzxpudy64zka8h5w7uffnzmhzlue4w3p", Amount: 746_0734312238},
		{Provider: "maya14sanmhejtzxxp9qeggxaysnuztx8f5jra7vedl", Amount: 263_1238239972},
		{Provider: "maya1g7c6jt5ynd5ruav2qucje0vuaan0q5xwasswts", Amount: 261_7814317799},
		{Provider: "maya1k3r9mtedeurcnjzfhgxzkqrum9f3yy2kkpgt34", Amount: 231_9382825684},
		{Provider: "maya175dn4q74ztt7wzf2n5u0nqkmfvda5sc627vtvd", Amount: 673_6157260911},
		{Provider: "maya1hkqc78uhuc4z8qtt3qjsdn0u7348t2hhlgyzh9", Amount: 1140_8754798383},
		{Provider: "maya1kpm9vz8cc2w984vghz40z0ekqef4xlglyx2yeg", Amount: 496_4215432024},
		{Provider: "maya1pf7gg2h9kdq7zuj58r7wk8py99awwj9lwvchdx", Amount: 670_4760996117},
		{Provider: "maya1tdp957gs94j7ahgd6cemlunhrwd27l39e52l6l", Amount: 186_6576821491},
		{Provider: "maya10sdhv0cn0fsfgax6vpzv9pwy8r5872hw3h4tuh", Amount: 548_2678887003},
		{Provider: "maya1vu37n7h7mnk0uxakye2vhh2z2k5cehf6v2lk3r", Amount: 203_4014261306},
		{Provider: "maya18p22jfv43weeyznqg0h9f6dh3adnpj4nwch8hs", Amount: 205_7460310253},
		{Provider: "maya19z4xlhxp6hkqe4mlfmqwsnjahrpa3ycjflqczc", Amount: 61_9125692167},
		{Provider: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", Amount: 39891_0000000000},
	}

	for _, reward := range rewards {
		providerAcc, err := cosmos.AccAddressFromBech32(reward.Provider)
		if err != nil {
			ctx.Logger().Error("fail to parse address: %s", reward.Provider, "error", err)
		}

		if err := mgr.Keeper().SendFromModuleToAccount(ctx, BondName, providerAcc, common.NewCoins(common.NewCoin(common.BaseNative, cosmos.NewUint(reward.Amount)))); err != nil {
			ctx.Logger().Error("fail to send provider reward: %s", reward.Provider, "error", err)
		}
	}

	manualRefunds := []RefundTxCACAO{
		// Manual refund paid out by team https://www.mayascan.org/tx/A40EA9E982A794CE0ED9B813F553C218CAE975975A359326FEF9F9AABF749643
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(28000_0000000000)},
		// Manual refund by team (sent to old vault by TW due to mayanode public API endpoint downtime) CACAO (5,000)
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(5000_0000000000)},
		// Manual refund by team https://etherscan.io/tx/0xae68c0c59977087b340c5226f37c6e3c96b510caf151a13d06822a9061942138
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(3735_1430430000)},
		// Manual refund by team https://etherscan.io/tx/0xfe113a3d316cd4cccd2a79077628bffdc740bc1eecef5fa9408c76aed1f8dd44
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(1223_4625390000)},
		// Manual refund by team https://runescan.io/tx/210512AF36E7925AD80344CA86CFC3732359EA3708CE54E82A1BB33FC69F01AA
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(175_6608684000)},
		// Manual refund by team https://runescan.io/tx/9DC1F05BCCB1CEA628266FA3A02483074C20EEB40B3F22089B6FD1E9470ACA28
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(702_6434735000)},
		// Manual refund by team https://www.mayascan.org/tx/9216104566E91EEC694E90EDF7686B08164F26CF73E3B9E3EA959A4C1D248580
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(3366_0000000000)},

		// Refund Dropped ETH Outbounds during a stuck queue in January 2024
		// https://www.mayascan.org/tx/47CA1E330362982D79680583B11FA1AB7E5F64452BA21B39D6B244D09054925C
		{sendAddress: "maya1xq3eaf70pdw4wl8esn0kyuxpnunprs05tgppzu", amount: cosmos.NewUint(152_2000000000)},
		// https://www.mayascan.org/tx/A8745F4940662F648CA601B8D63A7EB25393072E87D3DAAE2BEA3EA51434ACE6 * 1.20
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(6956_4662380000)},
		// https://www.mayascan.org/tx/9B8361F471C4CB8D803F92DF086DD4E30806EB2B9273C9A579A046AF218B7C0D * 1.20
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(13000_0000000000)},

		// Refund 19.98 RUNE per tx (in CACAO) to the below addresses, bifröst was setting gas_rate at 20 RUNE instead of 0.02 RUNE, overcharging customers
		// https://www.mayascan.org/tx/7C5935C87A1F383E5F075249D511C1C19094323E6D36181DED6A2B736189DCC8
		{sendAddress: "maya1ym3vk67ldc2jwlwmgzpenq78kkln7naxcqy005", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1F50570FAE74B4DE0FB57601FEA7657A750EFF1379FD9EF515984D14612ECDE9
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3C44C9CC227C68AA5EA79511BDBFCB1A294A30576FD1099E66C7E8932F7ACC1B
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/8C0F22BD6B8C1D79A801F5C95F356EB73A4DB9BD8D52206A6B5101D496A36E25
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9758A4E3E8F414A8C6B7670CC114819A42E225DC866E75E0B7EA1154EC670279
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E57E2599F72E229CA77B2BCEF0EA632B2F77F3FEF9C04EEF8161A4214804C1FF
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/4B7571EF61561C8AE287BC83A4403786626F5F2E87A81413E04827BE833745F6
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/DEF1492B6322CFA6A46F521EB8864F894427512D35F601553F5C3E172F460514
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9CE8B91AE846E558DD8F2294C17BB7932358EAD87F85B826E5C47D44BCE4695F
		{sendAddress: "maya1gsvhkjqe42z7h0hq589kfgd7xywwfud04awmwd", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/F23E27C695D88F6F91FB901C25DE6635025C48B7419C23B65C5715419EB24110
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/2476598E54F867E9498CCACAEED1E6E4FC0A8E0FA3B0BE028BAC30A5CAC92283
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/C760B881C06A46271DA845C5B74C231AEAE5FAE2C708F5DF9C2938148CA13D20
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D2D2210402370E0E12BF828051CBFBBD6D840BDA61662403912198581A5B9696
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/EEBCB3D88A9C53820DF3FC1AD6498A374B04859CF01BE3880DCAE1F34C36E35F
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D62641CB9FEF4E67EDADF0EA86A8185CD2EFA8D25BA778D1D3945F7A7E70D34B
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CF340B8DD282D080BA515AF7D1B732462073BC04A3E47DD5D4D596D4E37185D2
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0EC3288D9D7C43215E7058E0A3A997D42E07FDC4D067D108B66FDF099CFF6806
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/52B5931FDA4578A8B6FBA5D6D373EF3E923C415BE3EEBF0FD346835A2940FDD9
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/58C49A458A3E455260E5AAF9B322E76C8FAB5F06D7A510353C1014FD245C18BB
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0CFD1055C2D0D7ED8C017DD24BA9023F6163FCB405D8A477EDD3EAF3100EB9E5
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E52340F2D33CDC42EE38956848E9AC1E88F27D57A4173E9452006BAC7A952CDB
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E91651E8B622190E15ED2CC7071A4A063B44DDB8CEF66A7D9A506578541F4466
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1B5D64044B41D51FF8D2B0C41B01C41248A65848504E27388528148808548A3E
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/037B6F40CD59D3E548D1B2FD320F35CDBE907431770CAB1F186DF60F0588B50E
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/7365DF38557B0CC36D86C78D37475E6AFFB5AA63B5EE0AEB953E3DBF990BA9B4
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/8D526E6E39B00119078B7EE64A4656C1159941DB8456F44054D38F6CFE03D029
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/7EAD0D66085B6A278AB59CF3D92AA97CAB977AE8E8046CA8CC758033E8052E5E
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B90F321681937CC2154C27E8E30723A180C1CFA8AC8EE428C144AC2DB2016455
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9D8600866E0C6A0E1A1586A425CE723C00943AC618C6561BBE4116C1E686EEEA
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9D793B7842FC6F57BFBE92E107FB24165C95E0B6FFEC43D6CF7128358F3F6DFB
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D8A3AD9D6057AC27972378BFD3AD699AC330FDDEEC17BAE6B105B048D643E29A
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/5E60DD56A79749D2CF618EC54F466E473A97DB2659C5B8A6F65204524C2690A7
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A96BF6FC8DF39DFBAEDF7CDAC917B10FD7D591F0E3D08AD90C534EEE081CEB27
		{sendAddress: "maya1ymvh7rg4a6thqsgmf9z3fa5u4alswkftml80sf", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/BCAD48DDD5645BEDCB8C9EEBA8D3FA4DF0277E07E955FC61F2B4E47543FA75F6
		{sendAddress: "maya16adcj4245hw5jdn023cg4nnt4f089l6vhgcpyt", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/DEC27A703D7F7F1E81D2AF4EE2A88959BB3C9C6A07CA8B7E31FA8DC94CD78011
		{sendAddress: "maya1j9hg086cp4kc79jhauqkqnmp6j6u402zsax94c", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E007E360FCCF46A1D5F6780ECEFBB7EFF82D85B4038AF3EF0698FAA5BCF3ED15
		{sendAddress: "maya16adcj4245hw5jdn023cg4nnt4f089l6vhgcpyt", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/21E996872215D4961BDD8A16E20F886659724415EA3BAB670FAE8C42D01FB511
		{sendAddress: "maya16adcj4245hw5jdn023cg4nnt4f089l6vhgcpyt", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/21A7A0E345D144225938257A47096EC23A462D37FE5F2836319E78653B1FFE85
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/650B41AA0D096E5BD3DA801B06A26F7BB6B92E28B5A9D467098423AD60FE45DD
		{sendAddress: "maya187atpn8wgf45vah47hawkfhcrwyf4r7c8ersxs", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/F3DA3A3758A2F81749FF5ADFD8C4A81B58E03B6E5914BC115FE0593C879DE417
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/201F0EB6941A844C85C731F722D47255F1A5D3F608824510B4EE66A06A44BDA0
		{sendAddress: "maya19t4cjgcm6s47qshu2lt4d2yj4wy0rd8xs82vea", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/29A5009402A27423543E612988429D089C8FE879D56ED56F12B8022C0B21D7DC
		{sendAddress: "maya14sanmhejtzxxp9qeggxaysnuztx8f5jra7vedl", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3A989A22D4A51791271D56535D079842FCF134D8F9602692C51B29C5183A86A2
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B7B96D37D74A033612F11347B015EC21882BDFC104400CDDE2964406925684A6
		{sendAddress: "maya1qsynvzys9l63f0ljgr7vk028n4yk0eyvjakn80", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/C45B44E24379BEB601AF067E6F6A787D8DF5B6FCD7AE76B4F099789F16F7FB70
		{sendAddress: "maya1ahzf4jvc93j74m37ex3nfuvhupuevyanx5lfjc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/5BA778806FF1D10BC96B619E1339D2C85C192C85DA3B705FC0E587AD6CC813DB
		{sendAddress: "maya1ym3vk67ldc2jwlwmgzpenq78kkln7naxcqy005", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0C8D2B62D9C3FA906A947AF5457583DC81D96877A40D01D43C69879FA5AA2A88
		{sendAddress: "maya12qj3aw9e2ec45n4fv9jx56fl7lcrccupc7qj6p", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/523FDF05B87DC9EC2211F8AE3375A5EC054F4C41323F2D7D4CE0A175F781CFF4
		{sendAddress: "maya1mxes04w6mu9fy32w7zml20anux6smtajqprvfq", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CA5FD4525A1C39BD3D66300EDBCD47332EEF516C30A57D61F66A55FE2B432455
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/DF471498F03AC94F8C2051174C4962DAC69A5F5FD111BA92B55134535A34C15D
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3A91288B57B557FCC774CB2C9800240B72B891B8803D76A7943EFA73AF100A52
		{sendAddress: "maya1nxvlefnx27pqmj2q5trf4pzpgquhedxhxm0n4k", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0A91AD07C15334F4AE78BA61FEC2C1D47407793DA8BE7D040AF2F7D9988261B3
		{sendAddress: "maya1ayzv7vga724m95g09unzcfkuxp6838md5tsy9k", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D2D81ADC431C467E02D8EBFDEF90A442238680C26452797A980CEE4C5695C079
		{sendAddress: "maya1j0t0gwvz7untqwk8jv535hlxqrtzp8kyefp7n3", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/11C43A11A2F99375D8647D425D6E04722A970C5301E71043E4BA12DBB89C2FD9
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3EBCE1628995DD0B8212CB11F76D6D37C51B53EE32A59B81E47D8417F77BDD4F
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D7F91E02A7D6765B423F93B03565A078832CD267CC0C61608ED4C310687984BE
		{sendAddress: "maya12aetsuee7cyzxahn8u0xzpqftkdmyr3s3t3h88", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/004D4C4F2935C583F488EDBC5B93CD70DBC88060DCEFC0AB807A20F320F5071B
		{sendAddress: "maya13su7x39jun8mmup83c53lv5dqj58mg0nu24efp", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A6467A6ABFC0B51591FECD7FE1B1AF272645D4EF128AAF1D3B91578E47B586EF
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/70694723EBAD82D6E4000A20B3EE9EA849C6E6C0A6CC2E8097D1F082BA613A06
		{sendAddress: "maya13su7x39jun8mmup83c53lv5dqj58mg0nu24efp", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/F8E958984DCEDFEC9EC5793012DEC0FE57EA04CEA149D7E1E2ABC39D170ADB91
		{sendAddress: "maya1tm0qxkqylv0xly82qykafvmg4ul8vj4yxma7h7", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/DA060D5149FEBE2242C968E977192153E7846B2F7A8F57AF045D58B38EFAF652
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A718298A90C08695D88CEBE1B20E3B6D1645D11DA6F9405C85DB9DEAC584039A
		{sendAddress: "maya1jmk9wak5xhjl7kljv636ayym9as8hvn7z667m4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3A8068B1D8B43E5118B0F3B6A235D2F341D23A1676671AA7559A3D9342E1448B
		{sendAddress: "maya1hcaqh4ez3p24pr79gu226cvfw9473crxnyff9a", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A898FBF21E878F69E4C5B74534CBDB50519083E2A71958DF6BE1C01C4618B911
		{sendAddress: "maya1vx2s0dqy0unedgmruc52sg0u2emq7rq0nvcm3u", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/736E667D6F0BD397C5D71500AB298AC945AFC2F6EF1A2EFBB81F65486901C5E5
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/8F1D44840B4D2CF6C60E0E8F6DD3BADB79783053099CB5522D16334547E60234
		{sendAddress: "maya1qefmyzkgkvu2kz57yv5x5r6k9tvr65v3u2dgvh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/2FC67EE02E2679687816A3A53438BF815AA0CC488BB58CF74034F8E381E6E8AC
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/65651525C93198E5114610049FFDDFABEA86FA2BC871A549298159E86A8B6327
		{sendAddress: "maya1d059pckjsfzv7wzn4wa5zfjm6snzen7mysmth2", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E2BE7F1B12538B1C1933862584D0EFE422FB7B3F78A66C3330BF173D76DDF6C2
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9412C099A738001BD2E22A2E9D595C49FC67CE4F7C12820A8E5C55D95995415F
		{sendAddress: "maya1605alvdhp7y990a9grnf2kvql8d66lxvwy9g0m", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9659BCF686FF915188351BE4F39B281E07344F61BA93595E5EA18E4FDCB60C6A
		{sendAddress: "maya15n0nanaymr0aupvvwrxx3yde04nff9pxta682f", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/6072B245D02CCB0F47B96E8293033A3414B59B9C5244ECAE631927DEC5E3AAE4
		{sendAddress: "maya1sdnxgkh2te46ayhkj0au63k8klm6wj7tzzztnk", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CBCDDA3779DBBBAA3B24C67DFE3F058AEB96D28D62C07397928D03E61CD5E271
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/499097246E2DE9EE4C718FDAD7EC83E7783D6F52A22349DB62FAF6792541CA6C
		{sendAddress: "maya1n0tp2rgc9pnl5r7nq705adpmk6lfvss6nhkh6y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/AD0756F721D9442B6DECBF7353A77A8C6E3AE3895CBFA9B084DC535FC100E00D
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/391A614C22F391A9F8C7A75E43527D26B7B681C3E04579776B0D7D61BDC19761
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/19CC3BB9019BB72B6EE2BD53439F89A090ACDF3C233BE29FDBA8FA0019DEB146
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B2F1CD0AFA0FA306296FD68141CA0174D2080DD4A68E848720A954C2A0DA9FF2
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/4CA9E5BEC621135F76562D21B986C18D70F50AF5EDA7FFD5422EFF895249E800
		{sendAddress: "maya1n0tp2rgc9pnl5r7nq705adpmk6lfvss6nhkh6y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/F959D7784EDD1048E60A1B4E38DA5DC3F135F111D5AF2D21DB20DB02C8EADCD7
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1D00A91DFD3411D19495A678FBBB22BC23671FB7B2FF764B747107EB6A5C68C6
		{sendAddress: "maya1s7mga7ztwx8sf4ptuyspd4zyj8hac7ltulnegh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/510EA13B676027720F10FC1C8A33C7993DF5838339E2CC4C288A2DD747A9317C
		{sendAddress: "maya1n0tp2rgc9pnl5r7nq705adpmk6lfvss6nhkh6y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1E751FA1E6887647A061D3F76EE100F34F5C74A9521A3580A786B17139ED8D8F
		{sendAddress: "maya12fjzarhhc0f439yy5372xz4jwzgxg60un9gg9g", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/C52F328B736319D2DA5ACC20D486AF6F0B17256D615E86F5512EF6F18EF1B98C
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/49294FF51BC204FE461CC2EA7A9DD970BA256F01E03B91792CF2FEAADE8B46D5
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/535E56823D5DF0C65C7F22FAD56025DD1EDFFE4223C8A652A46073232401FE16
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A90CD0FFD6B1500D854C29A2D37D55FC5F96FC7083B6A3D3C921488E545FECA4
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/10931011F0BD898D1EBB2C957332A74ADC782E293A0E6DB792DB351A8CAE7221
		{sendAddress: "maya1quefasy83stwlghdh79tpyr3hm9lapvc27ud5y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/970FD9C8723A05BEC3F3FFA1B10B60C6C2E3F40C466240818BE4BE3576A9FD75
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/BFBEBC677BF333A624E23FE764B140616AA7BCFAC687F4406EB59EDC07472811
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E9C3D660F88FD29982B3F1FF7575CE4A5957341AF75D5FBCD06D14D93E3E3282
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/12F9ABBBBE1FFCC15D5DD3DA4B3F6C9248A1FE15D205936E8FEFFAE7B2F0DB6F
		{sendAddress: "maya1quefasy83stwlghdh79tpyr3hm9lapvc27ud5y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3585AC665644FAFCBE104EB8679B2C00FE8B797BBA72D040E5D63BA34973CDC6
		{sendAddress: "maya1l42u3x2xr0srtvllupu4kgtnxkgvwvvmmpdh6u", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9ED4FB15585F2CD6A0A4113673E4EDC067977282DD7866A76E6F295606B3CFDD
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/6F319008C963EA115FB7631D86F46CB060A8B40ED684C032B453FE852819099E
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/98C5A92F45A9D1BB9A30BF39E54FF0E89739EE3FA5C415526FF8D8B923741C72
		{sendAddress: "maya1n0tp2rgc9pnl5r7nq705adpmk6lfvss6nhkh6y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9BC6DE77E4A1447BB8DB1AB374B197DC996143AA80E9FD9BF1C68D4889EC4AC2
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1898ECD5449610F17998C5F2DE24286DDA5F6155DF28B8C19A1E74ECBE8C153E
		{sendAddress: "maya1ym3vk67ldc2jwlwmgzpenq78kkln7naxcqy005", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/AB31AC78F5BAFA0B91BADDEC1CE658F0D5F4C19C360B6C4FB1F815F42A4959FF
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/142712796C854CC16787FA33C59A8B997F02AE911A82AFC77DD838B3DE55BF1B
		{sendAddress: "maya1ym3vk67ldc2jwlwmgzpenq78kkln7naxcqy005", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3A70C367BEB973ACA2018A304857527998973D45F091D09C6C519AC550ACE343
		{sendAddress: "maya16adcj4245hw5jdn023cg4nnt4f089l6vhgcpyt", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/638EBEE98D7185736BCA369CC6277EAFB38C7DB02875367EBFA2BE405171AFCC
		{sendAddress: "maya1wx5av89rghsmgh2vh40aknx7csvs7xj2c5tjrr", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/5284F153C7D65D78188D1F477B6CDE7EBAF22BE52958028778150501548373A5
		{sendAddress: "maya12aetsuee7cyzxahn8u0xzpqftkdmyr3s3t3h88", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D0B3263CC81D0C083525697AB4DDC8D5A68DA43CD610157B6FD6365488000466
		{sendAddress: "maya1ulzsr523wmx0gkndw029flkh3j9q5dljg8veen", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D08B84BDADC7E56D170EAF2D4D027330F902CCBE39C5820763CA791A811F446B
		{sendAddress: "maya1mxes04w6mu9fy32w7zml20anux6smtajqprvfq", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/7EB3982BC22519C99E69A6D2E0F2097713A4BD531E1665C05875DDF718C317C4
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E58AA905E06D33DCDED4D9EB073A9A535B5DEEA32DBBB6FFDADD2C9A0B1D2EC2
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/1FCC63BDD9BBDF889EDA14D93A4B4653CF0CC5870D33FCB75DAB6FC07152CE60
		{sendAddress: "maya17cyy84n4x94upey4gg2cx0wtc3hf4uzuqsmyhh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/66A0A42BF830A5564546455FC8D819C6DC72F36E1A60A0F9FE5527592D9F03D0
		{sendAddress: "maya1vnalntj68qzv9mr0sftxgvucqlw2ht6yjqkr8d", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0BDCD6A396F9074D7DD6686088D104FD19CC05E3730837AA3C116A8AF8BAD969
		{sendAddress: "maya17cyy84n4x94upey4gg2cx0wtc3hf4uzuqsmyhh", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/41E3F70245CC22AD4A237F4F68447E51457441DEE081973D27EB2ACE6A56B513
		{sendAddress: "maya1mxes04w6mu9fy32w7zml20anux6smtajqprvfq", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/F8BA206729778CC8EB13616193121743441331424163D755EE6295C76E7A6AB7
		{sendAddress: "maya18jxee48ah3vlkfrndlm32y6urwhgyjgpdfxw7e", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CE8FE08CE5E7C31AC03C39EB35EC8AB5219760CA3B10C2644D3E7C319E53889D
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/2C9B224C4E49D99EA13F3572BE31E3416D4484E4785BF8814D39DEFBCFFE302A
		{sendAddress: "maya12fjzarhhc0f439yy5372xz4jwzgxg60un9gg9g", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CEE478CC58D8E28393DC1C7F7C3B3C616894C62A4CBB5ABC7A033D24786BF3F9
		{sendAddress: "maya1hklqwgqfe9dk43xg0lfz6zf3rsers3hslllms5", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/CFBB4EB812B706692830758CEEB92893955C45BEB0433B0D2ACEDF7D58EB1863
		{sendAddress: "maya12fjzarhhc0f439yy5372xz4jwzgxg60un9gg9g", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/03A4C470D4F360D6A4AF7D71131D87D60478D6E7CA9AB10E9EFA52974DA44D33
		{sendAddress: "maya17x5sz7lmayh4vw67nj2wynquf6vrt50f0u8lxx", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/87B910C7805F11902ABCE796141389101FEC3C217396495CC8D943A8C3914CA2
		{sendAddress: "maya17x5sz7lmayh4vw67nj2wynquf6vrt50f0u8lxx", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/486639334B1EA058DE2ACA87A8B8180DC8518C952C801F99661F61916FB5F413
		{sendAddress: "maya1wx5av89rghsmgh2vh40aknx7csvs7xj2c5tjrr", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/059024A6626F3AA6AA9B33CEFD97B48FE2B9E9B9715EC8D0B48037BC50277041
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B195D9EC30FCA896F93490645C1D167E05CFD2E150921AFABAE82697B5A71A0B
		{sendAddress: "maya1n0tp2rgc9pnl5r7nq705adpmk6lfvss6nhkh6y", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/C94DF785FFF232CDB43A9A6C8470BDCAF254C1301EC7725E30DAFA998016584C
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D2E6B758E80C765D4F57403D5546C6503D86E36F69572AA976B0CC91F641523D
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B60AB62E3FC90E2FB3FCDE3B1F6A762B63432AFEFD6D587D0FC8A4621D5F202B
		{sendAddress: "maya1wx5av89rghsmgh2vh40aknx7csvs7xj2c5tjrr", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/61DF0A907F7C9AC310CAAEECE7179EAFB07E8130BB6F7548C9BD58EBF8252DF6
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/6CE11F3A150E65FED8F7A61159CEF1EB683A9203A44E7753F77DC15CCCA9A4EF
		{sendAddress: "maya1mxes04w6mu9fy32w7zml20anux6smtajqprvfq", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/E5AFBDEA0BF36EB7369770765C9056E435B61E14A0BA05FF3D2197DBD398B1E3
		{sendAddress: "maya1hqkucmh30dqq9lx4gtk4dx5m937s94cc5cxjh4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/134F5C10ACEEA2B427AD53E3DD94D5A01065BDDB8201F0C330FE5331EF3AAC8E
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D38A09B04E117DBA4326C7020A6E3ED43713498F96DC13722EFD1E6A0C36A76C
		{sendAddress: "maya1x0uv04mglg439a8vw8q4mn83jm5s3uuj0vuuj4", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D70FCFF3CB17FAA67F760FDC7EE2DCC2AC2977F07694C3351348483CC0B7509A
		{sendAddress: "maya1wx5av89rghsmgh2vh40aknx7csvs7xj2c5tjrr", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/0342DE3969CD3FDCDDB0203D46875A8EC244E17E962F3949AD936A1EDD3431D2
		{sendAddress: "maya1gyap83aenguyhce3a0y3gprap32ypuc99vtzlc", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/A3ADCEA068C52F73BA405BAD527A9FFE95C4E3B1F6C8E12DF1E7A20563E0EE50
		{sendAddress: "maya17hwqt302e5f2xm4h95ma8wuggqkvfzgvsyfc54", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/182F09F0E2DC9FF9AEAC8C6BFB974C79EC4F18D68D53C8399F50EEE3DE7AC8D1
		{sendAddress: "maya1tjw4nxtezank4kgwvupxfyc3r4shw4j86cjxpv", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/C7C09BEB97418AFA066A796F4DE5CE9D40D783411A011FDEB2C048286732FE8A
		{sendAddress: "maya1mxes04w6mu9fy32w7zml20anux6smtajqprvfq", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/67F34BFF74EAE24676A7CB573C9DACD9C271FC6B568FE9A332FB673A595F41AC
		{sendAddress: "maya1vx2s0dqy0unedgmruc52sg0u2emq7rq0nvcm3u", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/46932FDE36416702969BE329BF860AF6DA9A93A8906484A38D6455B7343BEAC8
		{sendAddress: "maya1dzawlfg28lp0lysxx3p6ucnwfy948xh7u0yjks", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/D69114D6B5989E11543625F1D52BC6BF457C9ED85474DB0D2DB6B0D24020054A
		{sendAddress: "maya18pd64frgh5eg4z374dsvsap3lazqhfxgcfe8d6", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/9D342F36518DAA0114D23241DC457598DFFA56E7988DEC2628CAB30C52C36A86
		{sendAddress: "maya12aetsuee7cyzxahn8u0xzpqftkdmyr3s3t3h88", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/287196CB09AE12FC0A67C5D997809CFCBEC4F0484DB779F4EB1FF7FF645C53D4
		{sendAddress: "maya1ehsjen3p9kfw93fasul299weujgl3pvp3hg7p2", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/B7CC19A8456339711473F35716D434D9D4B9D1A3679BF97DB3F72DFEC197EE88
		{sendAddress: "maya1rh4s4weg8ewvagnm27y3l99v342hxnda8z774j", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/BED0D529DD716A16ADF7C1C7FE33755818174E416BC7B3EBB202591F7CA4CA47
		{sendAddress: "maya13ucfs2scdcm5kusxt7khs6g4e9s2ycekk32q4u", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/5BC380417C867E9923C91786CE3AAFAF60778E948267FC429100056E84DAF4CB
		{sendAddress: "maya1lzwwgvdw6amvmt3en66tyhe64krkc6eul7d9m6", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/14029795AB815B00BFE18A038E825773D67A7443D442677F1DB9CACA80A46978
		{sendAddress: "maya1kxqckmd770ntr52qq8c4ryzhcrehsdqe38ydp2", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/3DE6C788FBC563801A8C9FECAD9EA9A5806AE20D78893BBCA2E617222A98927C
		{sendAddress: "maya13euhshrsncukjx6x0rpatrja4xvyjcmpwg08fj", amount: cosmos.NewUint(140_5286947000)},
		// https://www.mayascan.org/tx/14EDD9FDDA114FB1CF219CC1D668ADCD54E2ECF4F73FAECDF3F8BAB074A8F0DC
		{sendAddress: "maya1hwvf5pwa3379t2c9w3j6ys6h9djw7v2ldv4nlq", amount: cosmos.NewUint(140_5286947000)},

		// Ticket 1773 https://www.mayascan.org/tx/BBB7474251BED8D6CDFF6D8888678CA205A02F1912412D46B8D44AB63EDCD3B2
		{sendAddress: "maya16fpyktrqntc2q2hr98p366k6peme7mlp8eyu77", amount: cosmos.NewUint(141_7230769000)},
	}
	refundTxsCACAO(ctx, mgr, manualRefunds)

	// Move node op addr
	// https://www.mayascan.org/address/maya1lx06jmugq3s9n6rz6y5up2d0wh0vsj6gy5rh3e
	nodeOpAssets := common.NewCoin(common.BaseAsset(), cosmos.NewUint(347_7600000000))
	nodeOpMaya := common.NewCoin(common.MayaNative, cosmos.NewUint(1247_8600))
	nodeOpOriginAddr, err := cosmos.AccAddressFromBech32("maya1lx06jmugq3s9n6rz6y5up2d0wh0vsj6gy5rh3e")
	if err != nil {
		ctx.Logger().Error("failed to parse origin address", "error", err)
		return
	}

	nodeOpDestinationAddr, err := cosmos.AccAddressFromBech32("maya1jzpntepl8ukadpejf5m2fccy6vygssn6llw98l")
	if err != nil {
		ctx.Logger().Error("failed to parse destination address", "error", err)
		return
	}

	nodeOpAssetsCoin := cosmos.NewCoin(nodeOpAssets.Asset.Native(), cosmos.NewIntFromBigInt(nodeOpAssets.Amount.BigInt()))
	if err := mgr.Keeper().SendCoins(ctx, nodeOpOriginAddr, nodeOpDestinationAddr, cosmos.NewCoins(nodeOpAssetsCoin)); err != nil {
		ctx.Logger().Error("fail to send node op asset", "error", err)
		return
	}
	nodeOpMayaCoin := cosmos.NewCoin(nodeOpMaya.Asset.Native(), cosmos.NewIntFromBigInt(nodeOpMaya.Amount.BigInt()))
	if err := mgr.Keeper().SendCoins(ctx, nodeOpOriginAddr, nodeOpDestinationAddr, cosmos.NewCoins(nodeOpMayaCoin)); err != nil {
		ctx.Logger().Error("fail to send node op maya", "error", err)
		return
	}
}

func migrateStoreV110(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v110", "error", err)
		}
	}()

	manualRefunds := []RefundTxCACAO{
		// Refund team for other stuck txs
		// https://www.mayascan.org/tx/3630A7C5506D49C27F580B59E3C9E2308851E9A4A68DB62E4CCAFE5D8883D5B0
		// https://mayascan.org/tx/660A3EF28891B31963E0D7768F20058E80257C519C183EC46D45C4ABCF05A23D
		// https://runescan.io/tx/72EC71610AFF80028A71F91A74F481A46D471BE05290B71D8CBB5ABABC228255
		// https://www.mayascan.org/tx/D44DE3211E84DDD5F384700354F35A58DC07C7358C049E87AD4872F99A0F7F15

		// Refund team for Thorswap Grant of 35,294 CACAO

		// Refund for limbo RUNE saver’s add (7,284 CACAO worth of RUNE) to be store migrated to pool balance https://runescan.io/tx/C3D527932C11A8B59E7FE0E0586B660DDB39EBB289011EF2E8B0491D3371BA47

		// Refund team for manual Savers withdrawals
		//  XmwJn2ZAs59ZgCcH3epfsL2mTVpzps4ViE
		// thor12xdyyc9dep9ye7fzs8lnyaw7pk88w0f6hnvwus
		// thor1r08ls3k4k4fj84n27cryfmle9sdmp9jkpes3rj
		// Thor1xmwrh6626q7qvwnpyrw3fjh8q6f59r9mvkgcam
		// thor1a7gg93dgwlulsrqf6qtage985ujhpu0684pncw
		// https://viewblock.io/thorchain/tx/43232BDBFAF977A5DFE0D8C619D79055974D8D7AAD7E3D87B8DC2C957F058904
		// https://www.mayascan.org/tx/614D9F7AFBB0FD605E574F8428A57E2DBBA200AD42FA37B8C5517E845846D6AE
		// https://viewblock.io/thorchain/tx/BC6FF661171A1915BC98CE1358B11779770CE8A7BFD3975AA132A8734F38E65C
		// https://www.mayascan.org/tx/91826c0d44c41defffc7fcc179825f2b3a8189aadb91551f358300a11c73bdfc

		// Refunds from stuck RUNE when Thorchain halted its chain
		// 2508.40747108 A6E2C37C9BF658890657030393D7A3CCC1D2859169A0C37AA0FEEDAD4B377A52 done
		// 	53.31317576 B3AB2A81FB9624EFE2F64A3CF83EA87DE0F082DB79D02BFEB90BC3763674AE1C done
		// 3026.83465896 1A50CC9165B6F161632E185C34E198162D0C123AB5B8DFBF4015226E2B96595E double
		//  180.84873985 E2419D67A47818EBC45BB9EA6AF9D10D202D10973F25C2C133BD603EF8FA6803 double
		//  131.64379267 2ACC77FCB56FDA812BB1364721CD42A9B1B1F8DAC9C61C2C92AB7307D615F4D3 double
		// 	 1.78879370 58FDDB3DCBFC925BEC235D13B9F9B5B50904BDDA6781C8A3A51B25A9C84BA8EF double
		//  120.38799809 EF5E3B1CD8D54EAED01B3EE4F42E8E3304623D7C36E60BB77AC86620940C5770 double
		//  172.25939246 DD8D8FDE30D3ABBCC5D16DF627C8C50A7D435ECA858AB71870C8B0F5BFDB281E done
		// 		. 3175966 2B7A32B9C86E4E17F78D708F1FBEBDBBBFFFD4D437E4BE5B9F34799A6FE9854A done
		//  105.13197943 1927FB09E9382594CD794DEF489D2F95FFA4276D2709F8E87F517ACB9BB22F3A done
		// 	54.32749518 BF1CA13F217094EEECF19F04AAB86999469FA98BDB19A64249E56D9C3FC567AF done
		// 	 0.10919544 7B3A287932F6F119D93B8865FCEADD398D2F93B696879261D61F58005A1635F3 done
		// 1266.25643944 5B5171536EED1AEC6DB8A2BECBFC9C463490F5A5C9AF150D2E7D8A5952B49A84 done
		// 	77.95788519 AC189F12FCF832F6C1C50014C3607E77CC0D765B10FCCE70D7F05EA3C6DB11EE done
		//  113.81770482 FAF927B60DCB02220547B255D10C00B6E4C8D49853D64F79C1B0424CEA6EABA8 done
		//  104.97360247 695A0A2749B21A6043BDCC4C8AEBEF7ABA97A99FFEC6579ED124A35E02BCE3AB done
		// 	75.34835976 2907FDE0D445F142DA337DFD1626CD9710AA4E3BCC4621507F68C0A575F0FC3A done
		{sendAddress: "maya18z343fsdlav47chtkyp0aawqt6sgxsh3vjy2vz", amount: cosmos.NewUint(112915_1100000000)},
	}
	refundTxsCACAO(ctx, mgr, manualRefunds)
}

func migrateStoreV111(ctx cosmos.Context, mgr *Mgrs) {
	defer func() {
		if err := recover(); err != nil {
			ctx.Logger().Error("fail to migrate store to v111", "error", err)
		}
	}()

	// For any in-progress streaming swaps to non-RUNE Native coins,
	// mint the current Out amount to the Pool Module.
	var coinsToMint common.Coins

	iterator := mgr.Keeper().GetSwapQueueIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var msg MsgSwap
		if err := mgr.Keeper().Cdc().Unmarshal(iterator.Value(), &msg); err != nil {
			ctx.Logger().Error("fail to fetch swap msg from queue", "error", err)
			continue
		}

		if !msg.IsStreaming() || !msg.TargetAsset.IsNative() || msg.TargetAsset.IsBase() {
			continue
		}

		swp, err := mgr.Keeper().GetStreamingSwap(ctx, msg.Tx.ID)
		if err != nil {
			ctx.Logger().Error("fail to fetch streaming swap", "error", err)
			continue
		}

		if !swp.Out.IsZero() {
			mintCoin := common.NewCoin(msg.TargetAsset, swp.Out)
			coinsToMint = coinsToMint.Add(mintCoin)
		}
	}

	// The minted coins are for in-progress swaps, so keeping the "swap" in the event field and logs.
	var coinsToTransfer common.Coins
	for _, mintCoin := range coinsToMint {
		if err := mgr.Keeper().MintToModule(ctx, ModuleName, mintCoin); err != nil {
			ctx.Logger().Error("fail to mint coins during swap", "error", err)
		} else {
			// MintBurn event is not currently implemented, will ignore

			// mintEvt := NewEventMintBurn(MintSupplyType, mintCoin.Asset.Native(), mintCoin.Amount, "swap")
			// if err := mgr.EventMgr().EmitEvent(ctx, mintEvt); err != nil {
			// 	ctx.Logger().Error("fail to emit mint event", "error", err)
			// }
			coinsToTransfer = coinsToTransfer.Add(mintCoin)
		}
	}

	if err := mgr.Keeper().SendFromModuleToModule(ctx, ModuleName, AsgardName, coinsToTransfer); err != nil {
		ctx.Logger().Error("fail to move coins during swap", "error", err)
	}
}
