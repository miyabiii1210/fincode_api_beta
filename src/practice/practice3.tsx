export const Practice3 = () => {

    /* 受け取った数値で税率計算を行い、返却する(戻り値にも型指定が可能) */
    const getTaxTotalRateCalculation = (num: number): number => {
        const total: number = num * 1.1;
        return total;
    };

    const onClickPractice = () => {
        // 変数に型を指定する
        let total: number = 0;
        total = getTaxTotalRateCalculation(1000);
        console.log(total);
    };

    return (
        <div>
            <p>test2: 変数の型指定</p>
            <button onClick={onClickPractice}>
                test3を実行
            </button>
        </div>
    )
}