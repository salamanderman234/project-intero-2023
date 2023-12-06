<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Grade;
use Illuminate\Validation\ValidationException;

class GradeController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $gradeData = Grade::all();

        if ($gradeData->isEmpty()) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $gradeData
            ], 200);
        }
    }


    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        // Validasi input dari user
        try {
            $validatedData = $request->validate([
                'grade' => 'required|string',
                'description' => 'required|string|max:255',
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        // Mencoba membuat objek grade dengan data yang sudah divalidasi
        try {
            $grade = Grade::create($validatedData);
            // Jika berhasil disimpan
            return response()->json([
                'message' => 'Data stored successfully',
                'data' => $grade
            ], 201);
        } catch (\Exception $e) {
            // Jika gagal disimpan
            return response()->json([
                'message' => 'Data failed to stored'
            ], 500);
        }
    }

    /**
     * Display the specified resource.
     */
    public function show(string $id)
    {
        // Mencari data grade berdasarkan ID
        $grade = Grade::find($id);

        // Validasi apakah data ditemukan
        if (!$grade) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $grade
            ], 200);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        // Validasi input dari user
        try {
            $validatedData = $request->validate([
                'grade' => 'required|string',
                'description' => 'required|string|max:255',
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        // Mencari data grade berdasarkan ID
        $grade = Grade::find($id);

        // Validasi apakah data ditemukan
        if (!$grade) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba melakukan update data
        try {
            $grade->update($validatedData);
            // Jika berhasil diupdate
            return response()->json([
                'message' => 'Data record updated succesfuly',
                'data' => $grade
            ], 200);
        } catch (\Exception $e) {
            // Jika gagal diupdate
            return response()->json([
                'message' => 'Failed to update data'
            ], 500);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        // Mencari data Grade berdasarkan ID
        $grade = Grade::find($id);

        // Validasi apakah data ditemukan
        if (!$grade) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba menghapus data
        try {
            $grade->delete();

            // Jika berhasil dihapus
            return response()->json([
                'message' => 'Data record succesfuly deleted'
            ], 200);
        } catch (\Exception $e) {
            // Jika gagal dihapus
            return response()->json([
                'message' => 'Failed to delete data'
            ], 500);
        }
    }
}

