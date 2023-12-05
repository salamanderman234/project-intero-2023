<?php

use App\Http\Controllers\ProfileController;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;

/*
|--------------------------------------------------------------------------
| API Routes
|--------------------------------------------------------------------------
|
| Here is where you can register API routes for your application. These
| routes are loaded by the RouteServiceProvider and all of them will
| be assigned to the "api" middleware group. Make something great!
|
*/

Route::get('/siswa', [ProfileController::class, 'getSiswaInfo']);
Route::get('/teacher', [ProfileController::class, 'getTeacherInfo']);

Route::put('/siswa', [ProfileController::class, 'updateSiswa']);
Route::put('/teacher', [ProfileController::class, 'updateTeacher']);

Route::post('/siswa', [ProfileController::class, 'createSiswa']);
Route::post('/teacher', [ProfileController::class, 'createGuru']);