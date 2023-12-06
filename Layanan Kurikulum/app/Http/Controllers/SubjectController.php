<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Subject;
use Illuminate\Validation\ValidationException;

class SubjectController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $subjectData = Subject::all();

        if ($subjectData->isEmpty()) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $subjectData
            ], 200);
        }
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        try {
            // Validasi input dari user
            $validatedData = $request->validate([
                'name' => 'required|string',
                'curriculum' => 'required|string',
                'description' => 'required|string|max:255',
                'minimum_avarage_value' => 'required|numeric|min:0|max:100'
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        try {
            // Mencoba membuat objek Subject dengan data yang sudah divalidasi
            $subject = Subject::create($validatedData);

            // Jika berhasil disimpan
            return response()->json([
                'message' => 'Data stored successfully',
                'data' => $subject
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
        // Mencari data Subject berdasarkan ID
        $subject = Subject::find($id);

        // Validasi apakah data ditemukan
        if (!$subject) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $subject
            ], 200);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        try {
            // Validasi input dari user
            $validatedData = $request->validate([
                'name' => 'required|string',
                'curriculum' => 'required|string',
                'description' => 'required|string|max:255',
                'minimum_avarage_value' => 'required|numeric|min:0|max:100'
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        // Mencari data Focus berdasarkan ID
        $subject = Subject::find($id);

        // Validasi apakah data ditemukan
        if (!$subject) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba melakukan update data
        try {
            $subject->update($validatedData);
            return response()->json([
                'message' => 'Data record updated succesfuly',
                'data' => $subject
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
        // Mencari data Focus berdasarkan ID
        $subject = Subject::find($id);

        // Validasi apakah data ditemukan
        if (!$subject) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba menghapus data
        try {
            $subject->delete();

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
