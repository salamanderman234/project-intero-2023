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
            return response()->json(['message' => 'Tidak ada data Materi yang tersedia'], 404);
        }

        return $subjectData;
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
        try {
            // Validasi input dari user
            $validatedData = $request->validate([
                'name' => 'required|string',
                'curriculum' => 'required|string',
                'description' => 'required|string|max:255',
                'minimum_avarage_value' => 'required|numeric|min:0|max:100'
            ]);

            // Mencoba membuat objek Subject dengan data yang sudah divalidasi
            $subject = Subject::create($validatedData);

            // Jika berhasil disimpan
            return response()->json(['message' => 'Data berhasil disimpan'], 201);
        } catch (ValidationException $e) {
            // Jika validasi tidak sesuai, kembalikan pesan validasi yang tidak sesuai
            return response()->json(['message' => 'Gagal menyimpan data', 'errors' => $e->validator->errors()], 422);
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
        // Mencari data Subject berdasarkan ID
        $subject = Subject::find($id);

        // Validasi apakah data ditemukan
        if (!$subject) {
            return response()->json(['message' => 'Data tidak ditemukan'], 404);
        }

        // Jika data ditemukan, kembalikan sebagai respons
        return $subject;
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
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy(string $id)
    {
        //
    }
}
