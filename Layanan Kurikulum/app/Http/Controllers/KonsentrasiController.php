<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Focus;
use Illuminate\Validation\ValidationException;

class KonsentrasiController extends Controller
{
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        $focusData = Focus::all();

        if ($focusData->isEmpty()) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $focusData
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
                'focus' => 'required|string',
                'description' => 'required|string|max:255',
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        // Mencoba membuat objek Focus dengan data yang sudah divalidasi
        try {
            $focus = Focus::create($validatedData);

            // Jika berhasil disimpan
            return response()->json([
                'message' => 'Data stored successfully',
                'data' => $focus
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
        // Mencari data Focus berdasarkan ID
        $focus = Focus::find($id);

        // Validasi apakah data ditemukan
        if (!$focus) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        } else {
            return response()->json([
                'message' => 'Data retrieved successfully',
                'data' => $focus
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
                'focus' => 'string',
                'description' => 'string|max:255',
            ]);
        } catch (ValidationException $e) {
            // Jika validasi gagal, kembalikan respons JSON dengan pesan kesalahan validasi
            return response()->json([
                'message' => $e->errors()
            ], 422);
        }

        // Mencari data Focus berdasarkan ID
        $focus = Focus::find($id);

        // Validasi apakah data ditemukan
        if (!$focus) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba melakukan update data
        try {
            $focus->update($validatedData);
            return response()->json([
                'message' => 'Data record updated succesfuly',
                'data' => $focus
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
        $focus = Focus::find($id);

        // Validasi apakah data ditemukan
        if (!$focus) {
            return response()->json([
                'message' => 'Data record not found'
            ], 404);
        }

        // Mencoba menghapus data
        try {
            $focus->delete();

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
