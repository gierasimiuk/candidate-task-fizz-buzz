<?php


use App\FizzBuzz;
use PHPUnit\Framework\TestCase;

class FizzBuzzTest extends TestCase
{
    private FizzBuzz $fizzBuzz;

    public function setUp(): void
    {
        $this->fizzBuzz = new FizzBuzz();
    }

    public function testCalculatesFizzBuzz(): void
    {
        $this->assertEquals([
            [ 'input' => 1, 'output' => '' ],
            [ 'input' => 2, 'output' => '' ],
            [ 'input' => 3, 'output' => 'Fizz' ],
            [ 'input' => 4, 'output' => '' ],
            [ 'input' => 5, 'output' => 'Buzz' ],
            [ 'input' => 6, 'output' => 'Fizz' ],
            [ 'input' => 7, 'output' => '' ],
            [ 'input' => 8, 'output' => '' ],
            [ 'input' => 9, 'output' => 'Fizz' ],
            [ 'input' => 10, 'output' => 'Buzz' ],
            [ 'input' => 11, 'output' => '' ],
            [ 'input' => 12, 'output' => 'Fizz' ],
            [ 'input' => 13, 'output' => '' ],
            [ 'input' => 14, 'output' => '' ],
            [ 'input' => 15, 'output' => 'FizzBuzz' ],
        ], $this->fizzBuzz->calculate(1, 15));
    }
}