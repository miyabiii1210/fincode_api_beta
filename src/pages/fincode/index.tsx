import Head from "next/head";
export default function GeneratePaymentToken() {
    const onClickGenerateToken = () => {
        let fincode = Fincode('p_test_YWYzYzFhMjktNjhhZS00MDU0LTljMTMtYjcxY2ZiMDdkNDY5YjNlNDZkOTYtZTVhOC00MTc1LTgxYmUtZjk3OGZjMGI4N2U1c18yMzAxMDQzNzAwNw')
        interface Token {
            card_no: string;
            expire: string;
            holder_name: string;
            security_code: string;
            number: string;
        }

        const card: Token = {
            card_no: "4111111111111111",  // VISA: 10-16桁(スペースは含まない)
            // card_no: "5111111111111111",  // MASTER:
            // card_no: "3531111111111111",  // JCB:
            expire: "2212",               // YYMM形式
            holder_name: "ma sand",     // カード名義人
            security_code: "123",         // セキュリティコード
            number: "1"                   // トークン発行数
        }

        // customer_id: c_ZrAH4kSeRbuX9YxXbVpgSQ
        // card_id    : cs_ErBt93JARcCi_M0AZ7Tkeg

        // オブジェクトに配列いれて更にオブジェクト入れた
        // var s: {key:{key:string}[]} = {key:[{key:'123'},{key:'456'}]}
        // var s: {list:{key1:string, key2:string,}[]} = {list:[{key1:'123', key2:'123'}]}
        // const fincodeCardList: {list:{
        //     customerID:string, 
        //     cardId:string,
        //     cardNumber:string,
        //     expire:string,
        //     holderName:string,
        //     cardNumberHash:string,
        //     cardType:string,
        //     cardBrand:string
        //   }[]} = {list:[
        //     {customerID:"c_ZrAH4kSeRbuX9YxXbVpgSQ", 
        //     cardId:"cs_ErBt93JARcCi_M0AZ7Tkeg",
        //     cardNumber:"************1111",
        //     expire:"2212",
        //     holderName:"Masato Takeuchi",
        //     cardNumberHash:"9bbef19476623ca56c17da75fd57734dbf82530686043a6e491c6d71befe8f6e",
        //     cardType:"0",
        //     cardBrand:"VISA"
        //   }]}
        // console.log(fincodeCardList);

        fincode.tokens(card,
            function (status: number, response: any) {
                if (200 === status) {
                    console.log("Token successfully issued.", response.list);
                    console.log("===== [ DEBUG ] =====")
                    console.log(response)
                    console.log(status)
                    const debug = typeof(response) // object
                    console.log(debug)
                } else {
                    console.error(response);
                    console.error("Failed to issue token, Status: ", status);
                }
            },
            function () {
                console.error("Request Error");
            }
        );
    }
    return (
        <>
            <Head>
                <script type="text/javascript" src="https://js.test.fincode.jp/v1/fincode.js"></script>
            </Head>
            <div className="p-3 bg-black-bg-img min-h-screen">
                <div className="mb-4 flex flex-col bg-gray-200 m-5">
                    <button
                        type='button'
                        className='border p-1 bg-gray-500 text-white'
                        onClick={() => onClickGenerateToken()}
                    >
                        トークンを発行する
                    </button>
                </div>
            </div>
        </>
    )
}