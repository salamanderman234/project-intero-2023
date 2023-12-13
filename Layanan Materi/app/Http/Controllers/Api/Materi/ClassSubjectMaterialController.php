<?php

namespace App\Http\Controllers\Api\Materi;

use App\Http\Controllers\Controller;
use App\Models\ClassSubjectMaterial;
use App\Http\Requests\ClassSubjectMaterialRequest;
use App\Traits\ResponseTrait;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Http\Response;

class ClassSubjectMaterialController extends Controller
{
    use ResponseTrait;
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        try {
            $class_subject_id = request('class_subject_id');
            $student_id = request('student_id');

            $data = ClassSubjectMaterial::with(['progres_checks' => function ($query) use ($student_id) {
                if ($student_id) {
                    $query->where('student_id', $student_id);
                }
            }])
            ->when($class_subject_id, function ($query) use ($class_subject_id) {
                $query->where('class_subject_id', $class_subject_id);
            })
            ->get();
            
            return $this->responseSuccess($data, 'Class Subject Material List Fetch Successfully');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(ClassSubjectMaterialRequest $request)
    {
        try {
            $data = ClassSubjectMaterial::create($request->all());
            return $this->responseSuccess($data, 'New Class Subject Material Created Successfully!');
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
            $data = ClassSubjectMaterial::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Material Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Get Data
            return $this->responseSuccess($data, 'Class Subject Material Details Fetch Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(ClassSubjectMaterialRequest $request, $id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectMaterial::find($id);
            if (is_null($data)){
                return $this->responseError(null, 'Class Subject Material Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Update Data
            $data->update($request->all());
            return $this->responseSuccess($data, 'Class Subject Material Updated Successfully!');
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
            $data = ClassSubjectMaterial::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Material Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Save Data
            $data->deleted_at = date('Y-m-d H:i:s');
            $data->save();
            if (!$data) {
                return $this->responseError(null, 'Failed to delete the Class Subject Material.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            return $this->responseSuccess($data, 'Class Subject Material Deleted Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }
}
