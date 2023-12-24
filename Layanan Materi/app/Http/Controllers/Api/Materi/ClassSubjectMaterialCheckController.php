<?php

namespace App\Http\Controllers\Api\Materi;

use App\Http\Controllers\Controller;
use App\Models\ClassSubjectMaterialCheck;
use App\Http\Requests\ClassSubjectMaterialCheckRequest;
use App\Traits\ResponseTrait;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Http\Response;

class ClassSubjectMaterialCheckController extends Controller
{
    use ResponseTrait;
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        //
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(ClassSubjectMaterialCheckRequest $request)
    {
        try {
            // Create data
            $data = ClassSubjectMaterialCheck::create($request->all());
            return $this->responseSuccess($data, 'New Class Subject Material Check Created Successfully!');
        } catch (\Exception $exception) {
            return $this->responseError(null, $exception->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Display the specified resource.
     */
    public function show($id)
    {
        //
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(ClassSubjectMaterialCheckRequest $request, $id)
    {
        //
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy($id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectMaterialCheck::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Material Check Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Save Data
            $data->deleted_at = date('Y-m-d H:i:s');
            $data->save();
            if (!$data) {
                return $this->responseError(null, 'Failed to delete the Class Subject Material Check.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            return $this->responseSuccess($data, 'Class Subject Material Check Deleted Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }
}
