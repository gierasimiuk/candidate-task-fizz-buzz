const fizzBuzz = require('./fizz-buzz')

describe('FizzBuzz function', () => {
  it('doesn\'t throw', () => {
    expect(() => fizzBuzz(1, 15)).not.toThrow()
  })

  it('calculates FizzBuzz', () => {
    expect(fizzBuzz(1, 15)).toStrictEqual([
      { input: 1, output: '' },
      { input: 2, output: '' },
      { input: 3, output: 'Fizz' },
      { input: 4, output: '' },
      { input: 5, output: 'Buzz' },
      { input: 6, output: 'Fizz' },
      { input: 7, output: '' },
      { input: 8, output: '' },
      { input: 9, output: 'Fizz' },
      { input: 10, output: 'Buzz' },
      { input: 11, output: '' },
      { input: 12, output: 'Fizz' },
      { input: 13, output: '' },
      { input: 14, output: '' },
      { input: 15, output: 'FizzBuzz' },
    ])
  })
})