//idが'price'のものを全部取得する
//この時点での型は、[object HTMLCollection]
var get_price = document.getElementsByClassName("price");


//get_priceを配列にする
//この時点での”要素の”型は、[object HTMLParagraphElement]
var price_list = Array.from(get_price);


var total = 0;
//[object HTMLParagraphElement]のタグを取り除いた中身だけ欲しいので、.innerHTMLをつける
for (var i=0; i<price_list.length; i++) {
    total += parseInt(price_list[i].innerHTML, 10);
}

//idが'total'のタグの中に指定したものを入れる
document.getElementById('total').innerHTML = total;