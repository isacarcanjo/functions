
class firstDegreeEquation {
  private a: number;
  private b: number;
  constructor(a: number, b: number) {
    this.a = a;
    this.b = b;
  }
  calculate() {
    return -this.b / this.a;
  }
}