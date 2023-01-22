export const Practice1 = () => {

    /* 数値を受け取り、税率計算を行う */
    const taxRateCalculation = (num: number) => {
        const total: number = num * 1.1;
        console.log(total);
    };

    const onClickPractice = () => taxRateCalculation(1000);

    return (
        <div>
            <p>test1: 引数の指定</p>
            <button onClick={onClickPractice}>
                test1を実行
            </button>
        </div>
    )
}