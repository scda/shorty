package main

import (
	"strings"
	"testing"
)

var amazonLinksRaw = `amazon.de/dp/B01NCEC2NT
www.amazon.de/dp/B01NCEC2NT
www.smile.amazon.de/dp/B01NCEC2NT
https;//www.smile.amazon.de/dp/B01NCEC2NT
https;//smile.amazon.de/dp/B01NCEC2NT
amazon.de/BRITA-Filterkartuschen-MAXTRA-Pack-geschmacksstörenden/dp/B01NCEC2NT
https://www.amazon.de/BRITA-Filterkartuschen-MAXTRA-Pack-geschmacksst%C3%B6renden/dp/B01NCEC2NT/ref=sr_1_5?m=A8KICS1PHF7ZO&pf_rd_i=3581963031&pf_rd_m=A3JWKAKR8XB7XF&pf_rd_p=1b438f99-9a75-4c4c-9990-2ed6e82fd67e&pf_rd_r=WW77G09SV5CC1GDBGW66&pf_rd_s=merchandised-search-5&pf_rd_t=101&qid=1656078690&s=warehouse-deals&sr=1-5
https://www.amazon.de/dp/B07FQ4DJ7X/ref=gw_de_desk_mso_eink_jg_bau_xcat1?pf_rd_r=43XXHB6JXJS7YDHPP6DK&pf_rd_p=3eeebfd9-ba47-4ef4-aebc-2158845c16da&pd_rd_r=40ed4323-24d1-4cef-88d0-8b4df55d8766&pd_rd_w=B3u3s&pd_rd_wg=E1v8b&ref_=pd_gw_unk
https://www.amazon.de/dp/B08LR3G17D/ref=gw_de_desk_mso_vicc_grc_bau_xcat1?pf_rd_r=43XXHB6JXJS7YDHPP6DK&pf_rd_p=3eeebfd9-ba47-4ef4-aebc-2158845c16da&pd_rd_r=40ed4323-24d1-4cef-88d0-8b4df55d8766&pd_rd_w=B3u3s&pd_rd_wg=E1v8b&ref_=pd_gw_unk
https://www.amazon.de/HOMEFAVOR-Herausnehmbarer-Galvanisierung-Regenbogenfarbe-Brotzeitdose/dp/B08GKFRPQS/ref=sr_1_19?__mk_de_DE=%C3%85M%C3%85%C5%BD%C3%95%C3%91&crid=2H95X6OUCGRXZ&keywords=brotdose&qid=1656078648&sprefix=brotdos%2Caps%2C103&sr=8-19&th=1
https://www.amazon.de/Morbius-Blu-ray-Jared-Leto/dp/B09WRTL2TK/?_encoding=UTF8&pd_rd_w=bQB86&content-id=amzn1.sym.334f809e-3300-40ce-8741-870dea477993&pf_rd_p=334f809e-3300-40ce-8741-870dea477993&pf_rd_r=43XXHB6JXJS7YDHPP6DK&pd_rd_wg=H7hu2&pd_rd_r=b82a7ecb-1529-441d-b897-3572864b6b49&ref_=pd_gw_crs_zg_bs_284266
https://www.amazon.de/SanDisk-256GB-iXpand-Flash-Laufwerk-iPhone/dp/B07VQPDM56?ref_=ast_sto_dp
https://www.amazon.de/verstellbarer-stander-fur-echo-show-5-2-gen/dp/B08MGX9X3K/?pf_rd_r=43XXHB6JXJS7YDHPP6DK&pf_rd_p=73bd4b07-8bcc-458b-8951-0fd9aa76b2f7&pd_rd_r=40ed4323-24d1-4cef-88d0-8b4df55d8766&pd_rd_w=Qwwvd&pd_rd_wg=E1v8b&ref_=pd_gw_unk
https://www.amazon.de/gp/product/B079V367MM/ref=ox_sc_act_title_1?smid=A3JWKAKR8XB7XF&psc=1
https://www.amazon.de/gp/product/B079QFKJ13/ref=ox_sc_act_title_2?smid=A3JWKAKR8XB7XF&psc=1
https://www.amazon.de/gp/product/B07W4DGC27/ref=ox_sc_act_title_3?smid=A3JWKAKR8XB7XF&psc=1
https://www.amazon.de/gp/product/B07W5JHLCB/ref=ox_sc_act_title_4?smid=A3JWKAKR8XB7XF&psc=1
https://www.amazon.de/-/en/Entfernungsmesser-LM50-Entfernungsmesse-Hintergrundbeleuchtung-Pythagoras/dp/B08G892B9W?dib=eyJ2IjoiMSJ9.Gheo0TVDwnsRkjyKhP4dOLJ2RZ_7f3vvn9SfWCuFMo_3oGEsaTNQDATwOw6ZTb5p0RnOjBa6mT9qdW3LeAnxV0H6uc1FQf95saWD0Y_Gd034omjLtW0jzz9NmdKkVLKzGO6pZOJJFhC2oTmzkWgrNPLRhdS3aQAe-37q2ouKUFa3IyjbrxPkW0BQTH1Jzl9gQ7gIOmxKDqxlqY0oCqdHngrBGl8aVv_hyPjxy5pEtwY.uNtyGOfrhdGhHL7hSRLjOUJXUeJROomoqVytDWuNynU&dib_tag=se&keywords=laser%2Bmessger%C3%A4t&qid=1750184288&sr=8-12&th=1
https://www.amazon.de/-/en/Rangefinder-RockSeed-Measuring-Portable-Pythagorean/dp/B0863RK1KX?dib=eyJ2IjoiMSJ9.Gheo0TVDwnsRkjyKhP4dOLJ2RZ_7f3vvn9SfWCuFMo_3oGEsaTNQDATwOw6ZTb5p0RnOjBa6mT9qdW3LeAnxV0H6uc1FQf95saWD0Y_Gd034omjLtW0jzz9NmdKkVLKzGO6pZOJJFhC2oTmzkWgrNPLRhdS3aQAe-37q2ouKUFa3IyjbrxPkW0BQTH1Jzl9gQ7gIOmxKDqxlqY0oCqdHngrBGl8aVv_hyPjxy5pEtwY.uNtyGOfrhdGhHL7hSRLjOUJXUeJROomoqVytDWuNynU&dib_tag=se&keywords=laser%2Bmessger%C3%A4t&qid=1750184288&sr=8-1&th=1
https://www.amazon.de/-/en/Professional-measure-memory-function-measuring/dp/B00R0Z7TFM?dib=eyJ2IjoiMSJ9.Gheo0TVDwnsRkjyKhP4dOLJ2RZ_7f3vvn9SfWCuFMo_3oGEsaTNQDATwOw6ZTb5p0RnOjBa6mT9qdW3LeAnxV0H6uc1FQf95saWD0Y_Gd034omjLtW0jzz9NmdKkVLKzGO6pZOJJFhC2oTmzkWgrNPLRhdS3aQAe-37q2ouKUFa3IyjbrxPkW0BQTH1Jzl9gQ7gIOmxKDqxlqY0oCqdHngrBGl8aVv_hyPjxy5pEtwY.uNtyGOfrhdGhHL7hSRLjOUJXUeJROomoqVytDWuNynU&dib_tag=se&keywords=laser%2Bmessger%C3%A4t&qid=1750184021&sr=8-4&th=1
`

var etsyLinksRaw = `
https://www.etsy.com/de/listing/4311408001/1516-zoll-eleganter-moderner-laptop?ls=s&ga_order=most_relevant&ga_search_type=all&ga_view_type=gallery&ga_search_query=laptop+st%C3%A4nder&ref=sr_gallery-1-24&local_signal_search=1&content_source=616660b9-a24a-4be6-981d-cb20aa90650c%253ALTfde0873dc95278fcadd0e34ecdb9ee1d4de94ca8&organic_search_click=1&logging_key=616660b9-a24a-4be6-981d-cb20aa90650c%3ALTfde0873dc95278fcadd0e34ecdb9ee1d4de94ca8
https://www.etsy.com/de/listing/1884499559/vertikaler-laptopstander-perfekt-fur?ls=s&ga_order=most_relevant&ga_search_type=all&ga_view_type=gallery&ga_search_query=laptop+st%C3%A4nder&ref=sr_gallery-1-6&sts=1&local_signal_search=1&content_source=616660b9-a24a-4be6-981d-cb20aa90650c%253ALT0c995e414a7a31f4772fa2bea70f13e6a4f314b2&organic_search_click=1&logging_key=616660b9-a24a-4be6-981d-cb20aa90650c%3ALT0c995e414a7a31f4772fa2bea70f13e6a4f314b2
https://www.etsy.com/de/listing/4396295569/macbook-pro-macbook-air-stander-voronoy?ls=s&ga_order=most_relevant&ga_search_type=all&ga_view_type=gallery&ga_search_query=laptop+st%C3%A4nder&ref=sr_gallery-1-30&pro=1&local_signal_search=1&content_source=616660b9-a24a-4be6-981d-cb20aa90650c%253ALTcf12362f516f7e0c7f4b8d4116965c6ef26440c4&organic_search_click=1&logging_key=616660b9-a24a-4be6-981d-cb20aa90650c%3ALTcf12362f516f7e0c7f4b8d4116965c6ef26440c4
`

var amazonLinksReduced = `amazon.de/dp/B01NCEC2NT
amazon.de/dp/B01NCEC2NT
amazon.de/dp/B01NCEC2NT
amazon.de/dp/B01NCEC2NT
amazon.de/dp/B01NCEC2NT
amazon.de/dp/B01NCEC2NT
https://www.amazon.de/dp/B01NCEC2NT
https://www.amazon.de/dp/B07FQ4DJ7X
https://www.amazon.de/dp/B08LR3G17D
https://www.amazon.de/dp/B08GKFRPQS
https://www.amazon.de/dp/B09WRTL2TK
https://www.amazon.de/dp/B07VQPDM56
https://www.amazon.de/dp/B08MGX9X3K
https://www.amazon.de/gp/product/B079V367MM
https://www.amazon.de/gp/product/B079QFKJ13
https://www.amazon.de/gp/product/B07W4DGC27
https://www.amazon.de/gp/product/B07W5JHLCB
https://www.amazon.de/dp/B08G892B9W
https://www.amazon.de/dp/B0863RK1KX
https://www.amazon.de/dp/B00R0Z7TFM`

var etsyLinksReduced = `https://www.etsy.com/de/listing/4311408001
https://www.etsy.com/de/listing/1884499559
https://www.etsy.com/de/listing/4396295569`

var randomOtherLinks = `https://go.dev/doc/tutorial/add-a-test
https://go.dev/doc/diagnostics`

var nonsenseString = `abcdefgN0ns3n$3`

var mixedLinks = strings.Join([]string{amazonLinksRaw, randomOtherLinks}, "\n")

func Test_Replace_Should_Reduce_Content_With_Amazon_Links(t *testing.T) {
	// arrange
	input := amazonLinksRaw
	expectedResult := amazonLinksReduced

	// act
	result, err := ReduceUrls(input)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReplaceClipboardContent() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_Replace_Should_Reduce_Content_With_Etsy_Links(t *testing.T) {
	// arrange
	input := etsyLinksRaw
	expectedResult := etsyLinksReduced

	// act
	result, err := ReduceUrls(input)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReduceUrls() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_Replace_Should_Reduce_Only_Amazon_Links_On_Mixed_Content(t *testing.T) {
	// arrange
	input := mixedLinks
	expectedResult := strings.Join([]string{amazonLinksReduced, randomOtherLinks}, "\n")

	// act
	result, err := ReduceUrls(input)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReplaceClipboardContent() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_Replace_Should_Return_Input_On_Content_Without_Amazon_Links(t *testing.T) {
	// arrange
	input := randomOtherLinks
	expectedResult := randomOtherLinks

	// act
	result, err := ReduceUrls(input)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReplaceClipboardContent() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_Replace_Should_Reduce_Mixed_Amazon_And_Etsy_Links(t *testing.T) {
	// arrange
	// Take 2 Amazon links and 2 Etsy links and mix them
	amazonLines := strings.Split(strings.TrimSpace(amazonLinksRaw), "\n")
	etsyLines := strings.Split(strings.TrimSpace(etsyLinksRaw), "\n")

	// Create mixed input with Amazon[0], Etsy[0], Amazon[1], Etsy[1]
	mixedInput := strings.Join([]string{
		amazonLines[0],
		etsyLines[0],
		amazonLines[1],
		etsyLines[1],
	}, "\n")

	// Create expected output in the same order
	amazonReducedLines := strings.Split(strings.TrimSpace(amazonLinksReduced), "\n")
	etsyReducedLines := strings.Split(strings.TrimSpace(etsyLinksReduced), "\n")

	expectedResult := strings.Join([]string{
		amazonReducedLines[0],
		etsyReducedLines[0],
		amazonReducedLines[1],
		etsyReducedLines[1],
	}, "\n")

	// act
	result, err := ReduceUrls(mixedInput)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReduceUrls() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_Replace_Should_Reduce_Mixed_Content_With_Amazon_Etsy_And_Other_Links(t *testing.T) {
	// arrange
	// Take 2 Amazon, 2 Etsy, and 2 other links and mix them
	amazonLines := strings.Split(strings.TrimSpace(amazonLinksRaw), "\n")
	etsyLines := strings.Split(strings.TrimSpace(etsyLinksRaw), "\n")
	otherLines := strings.Split(strings.TrimSpace(randomOtherLinks), "\n")

	// Create mixed input
	mixedInput := strings.Join([]string{
		amazonLines[0],
		otherLines[0],
		etsyLines[0],
		amazonLines[5],
		etsyLines[1],
		otherLines[1],
	}, "\n")

	// Create expected output in the same order
	amazonReducedLines := strings.Split(strings.TrimSpace(amazonLinksReduced), "\n")
	etsyReducedLines := strings.Split(strings.TrimSpace(etsyLinksReduced), "\n")

	expectedResult := strings.Join([]string{
		amazonReducedLines[0],
		otherLines[0],
		etsyReducedLines[0],
		amazonReducedLines[5],
		etsyReducedLines[1],
		otherLines[1],
	}, "\n")

	// act
	result, err := ReduceUrls(mixedInput)

	// assert
	if expectedResult != result || err != nil {
		t.Fatalf(`ReduceUrls() = %q, %v. Expected result %#q, nil`, result, err, expectedResult)
	}
}

func Test_WriteClipboard_Should_Update_Clipboard_Contents(t *testing.T) {
	// arrange
	input := nonsenseString
	WriteClipboard(input)

	// act
	main()

	// assert
	result := ReadClipboard()
	if result != input {
		t.Fatalf(`WriteClipboard() should have updated clipboard contents. Found %q. Expected result %#q`, result, input)
	}

}

func Test_ReadClipboard_Should_Return_Clipboard_Contents(t *testing.T) {
	// arrange
	input := nonsenseString
	WriteClipboard(input)

	// act
	main()

	// assert
	result := ReadClipboard()
	if result != input {
		t.Fatalf(`ReadClipboard() should have returned clipboard contents. Found %q. Expected result %#q`, result, input)
	}
}

func Test_Main_Should_Not_Update_Clipboard_Without_Amazon_Links(t *testing.T) {
	// arrange
	input := randomOtherLinks
	WriteClipboard(input)

	// act
	main()

	// assert
	clipboardContent := ReadClipboard()
	if clipboardContent != input {
		t.Fatalf(`main() should not have altered clipboard content. Result = %q. Expected result %#q`, clipboardContent, input)
	}
}

func Test_Main_Should_Update_Clipboard_On_Amazon_Links(t *testing.T) {
	// arrange
	input := amazonLinksRaw
	WriteClipboard(input)

	// act
	main()

	// assert
	clipboardContent := ReadClipboard()
	if clipboardContent != amazonLinksReduced {
		t.Fatalf(`main() should have altered clipboard content. Result = %q. Expected result %#q`, clipboardContent, amazonLinksReduced)
	}
}

func Test_Main_Should_Update_Clipboard_On_Mixed_Content(t *testing.T) {
	// arrange
	input := mixedLinks
	expected := strings.Join([]string{amazonLinksReduced, randomOtherLinks}, "\n")
	WriteClipboard(input)

	// act
	main()

	// assert
	clipboardContent := ReadClipboard()
	if clipboardContent != expected {
		t.Fatalf(`main() should have altered clipboard content. Result = %q. Expected result %#q`, clipboardContent, amazonLinksReduced)
	}
}
