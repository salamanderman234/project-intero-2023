<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;
use App\Models\Focus;
use Illuminate\Validation\Rule;
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
            return response()->json(['message' => 'Tidak ada data Focus yang tersedia'], 404);
        }

        return $focusData;
    }

    /**
     * Show the form for creating a new resource.
     */
    public function create()
    {
        //
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
            return response()->json(['message' => $e->errors()], 422);
        }

        // Mencoba membuat objek Focus dengan data yang sudah divalidasi
        try {
            $focus = Focus::create($validatedData);

            // Jika berhasil disimpan
            return response()->json(['message' => 'Data berhasil disimpan'], 201);
        } catch (\Exception $e) {
            // Jika gagal disimpan
            return response()->json(['message' => 'Gagal menyimpan data'], 500);
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
            return response()->json(['message' => 'Data tidak ditemukan'], 404);
        }

        // Jika data ditemukan, kembalikan sebagai respons
        return $focus;
    }

    /**
     * Show the form for editing the specified resource.
     */
    public function edit(string $id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(Request $request, string $id)
    {
        // Validasi input dari user
        $validatedData = $request->validate([
            'focus' => 'string',
            'description' => 'string|max:255',
        ]);

        // Mencari data Focus berdasarkan ID
        $focus = Focus::find($id);

        // Validasi apakah data ditemukan
        if (!$focus) {
            return response()->json(['message' => 'Data tidak ditemukan'], 404);
        }

        // Mencoba melakukan update data
        try {
            $focus->update($validatedData);

            // Jika berhasil diupdate
            return response()->json(['message' => 'Data berhasil diupdate'], 200);
        } catch (\Exception $e) {
            // Jika gagal diupdate
            return response()->json(['message' => 'Gagal mengupdate data'], 500);
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
            return response()->json(['message' => 'Data tidak ditemukan'], 404);
        }

        // Mencoba menghapus data
        try {
            $focus->delete();

            // Jika berhasil dihapus
            return response()->json(['message' => 'Data berhasil dihapus'], 200);
        } catch (\Exception $e) {
            // Jika gagal dihapus
            return response()->json(['message' => 'Gagal menghapus data'], 500);
        }
    }
}
