<?php


namespace App\Controllers;


use App\FizzBuzz;
use Symfony\Component\HttpFoundation\JsonResponse;

class ApiController
{
    protected FizzBuzz $fizzBuzz;

    public function __construct()
    {
        $this->fizzBuzz = new FizzBuzz();
    }

    public function handle(): JsonResponse
    {
        // TODO
    }
}