<?php

use Illuminate\Http\Request;
use Illuminate\Support\Facades\Route;
use App\Http\Controllers\Api\Materi\ClassSubjectMaterialController;
use App\Http\Controllers\Api\Materi\ClassSubjectMaterialCheckController;
use App\Http\Controllers\Api\Materi\ClassSubjectAssignmentController;
use App\Http\Controllers\Api\Materi\ClassSubjectAssignmentSubmissionController;

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

Route::middleware('auth:sanctum')->get('/user', function (Request $request) {
    return $request->user();
});

Route::apiResource('class-subject-material', ClassSubjectMaterialController::class);
Route::apiResource('class-subject-material-check', ClassSubjectMaterialCheckController::class);
Route::apiResource('class-subject-assignment', ClassSubjectAssignmentController::class);
Route::apiResource('assignment-submission', ClassSubjectAssignmentSubmissionController::class);

