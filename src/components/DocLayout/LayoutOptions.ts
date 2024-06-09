export interface LayoutOptions {
  width: number
  height: number
  padding: number
  blockGap?: number
  fontSize: number
  lineHeight: number
  fontFamily: string
}

const PageConfig = {
  A4: {
    width: 595,
    height: 842,
    padding: 50,
    fontSize: 14,
    lineHeight: 1.5,
    fontFamily: 'Arial'
  },
  A3: {
    width: 842,
    height: 1191,
    padding: 50,
    fontSize: 16,
    lineHeight: 1.5,
    fontFamily: 'Arial'
  },
  Letter: {
    width: 612,
    height: 792,
    padding: 50,
    fontSize: 16,
    lineHeight: 1.5,
    fontFamily: 'Arial'
  }
}


export declare type PageOption = keyof typeof PageConfig

export function getLayoutOptions(pageType: keyof typeof PageConfig): Readonly<LayoutOptions> {
  return PageConfig[pageType]
}
