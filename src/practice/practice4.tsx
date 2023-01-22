export const Practice4 = () => {

    /* 数値を受け取り、税率計算を行う */
    const taxRateCalculation = (num: number) => {
        const total: number = num * 1.1;
        console.log(total);
    };

    const onClickPractice = () => taxRateCalculation(1000);

    return (
        <div>
            <p>test4: 設定ファイルの編集</p>
            <button onClick={onClickPractice}>
                test4を実行
            </button>
        </div>
    )
}