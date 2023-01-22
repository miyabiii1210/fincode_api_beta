/*eslint-disable @typescript-eslint/no-unused-vars */
/** TypeScript 基本の型 */

// boolean
let bool: boolean = true;

// number
let num: number = 0;

// string
let str: string = "A"

// Array
let arr1: Array<number> = [0, 1, 2];
let arr2: number[] = [3 ,4, 5];

// tuple
let tuple: [number, string] = [0, "A"];

// any (全ての型に対応できる, tsを使用する意味がなくなるため多様禁物)
let any1: any = false;

// void (何も返さない、引数未指定とした場合はデフォルトでこの形式となる)
const funcA = (): void => {
    const test = "TEST";
}

// null
let null1: null = null;

// undefined (何も設定されていない)
let undefined1: undefined = undefined;

// object
let obj1: object = {};
let obj2: {id: number, name: string} = {id: 0, name: "A"};

export {}