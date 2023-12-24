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
Route::get("/profile/{user}",[ProfileController::class, "getProfile"]);
Route::get("/profileSiswa/{student}",[ProfileController::class, "profileSiswa"]);
Route::get("/profileGuru/{teacher}",[ProfileController::class, "profileGuru"]);
Route::get("/siswas",[ProfileController::class, "getSiswas"]);
Route::get("/teachers",[ProfileController::class, "getGurus"]);

Route::get('/siswa', [ProfileController::class, 'getSiswaInfo']);
Route::get('/teacher', [ProfileController::class, 'getTeacherInfo']);

Route::put("/student/{student}",[ProfileController::class, 'updateSiswaProfile']);
Route::put('/siswa', [ProfileController::class, 'updateSiswa']);
Route::put('/teacher', [ProfileController::class, 'updateTeacher']);

Route::post('/siswa', [ProfileController::class, 'createSiswa']);
Route::post('/teacher', [ProfileController::class, 'createGuru']);