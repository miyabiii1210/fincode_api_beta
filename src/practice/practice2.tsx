export const Practice2 = () => {

    /* 受け取った数値で税率計算を行い、返却する(戻り値にも型指定が可能) */
    const getTaxTotalRateCalculation = (num: number): number => {
        const total: number = num * 1.1;
        return total;
    };

    const onClickPractice = () => {
        console.log(getTaxTotalRateCalculation(1000));
    };

    return (
        <div>
            <p>test2: 戻り値の型指定</p>
            <button onClick={onClickPractice}>
                test2を実行
            </button>
        </div>
    )
}