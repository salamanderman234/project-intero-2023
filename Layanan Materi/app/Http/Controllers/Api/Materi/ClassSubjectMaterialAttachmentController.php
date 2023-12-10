<?php

namespace App\Http\Controllers\Api\Materi;

use App\Http\Controllers\Controller;
use App\Models\ClassSubjectMaterialAttachment;
use App\Http\Requests\ClassSubjectMaterialAttachmentRequest;
use App\Traits\ResponseTrait;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Http\Response;

class ClassSubjectMaterialAttachmentController extends Controller
{
    use ResponseTrait;
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        try {
            // Get data
            $data = ClassSubjectMaterialAttachment::all();
            return $this->responseSuccess($data, 'Class Subject Material Attachment List Fetch Successfully');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(ClassSubjectMaterialAttachmentRequest $request)
    {
        try {
            // Create data
            $data = ClassSubjectMaterialAttachment::create($request->all());
            return $this->responseSuccess($data, 'New Class Subject Material Attachment Created Successfully!');
        } catch (\Exception $exception) {
            return $this->responseError(null, $exception->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Display the specified resource.
     */
    public function show($id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectMaterialAttachment::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Material Attachment Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Get Data
            return $this->responseSuccess($data, 'Class Subject Material Attachment Details Fetch Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(ClassSubjectMaterialAttachmentRequest $request, $id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectMaterialAttachment::find($id);
            if (is_null($data)){
                return $this->responseError(null, 'Class Subject Material Attachment Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Update Data
            $data->update($request->all());
            return $this->responseSuccess($data, 'Class Subject Material Attachment Updated Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Remove the specified resource from storage.
     */
    public function destroy($id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectMaterialAttachment::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Material Attachment Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Save Data
            $data->deleted_at = date('Y-m-d H:i:s');
            $data->save();
            if (!$data) {
                return $this->responseError(null, 'Failed to delete the Class Subject Material Attachment.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            return $this->responseSuccess($data, 'Class Subject Material Attachment Deleted Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }
}
