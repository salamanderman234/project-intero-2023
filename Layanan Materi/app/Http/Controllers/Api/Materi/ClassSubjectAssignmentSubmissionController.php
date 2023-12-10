<?php

namespace App\Http\Controllers\Api\Materi;

use App\Http\Controllers\Controller;
use App\Models\ClassSubjectAssignment;
use App\Models\ClassSubjectAssignmentSubmission;
use App\Http\Requests\ClassSubjectAssignmentSubmissionRequest;
use App\Traits\ResponseTrait;
use App\Helpers\UploadHelper;
use Illuminate\Http\Request;
use Illuminate\Support\Facades\Auth;
use Illuminate\Support\Facades\Validator;
use Illuminate\Http\Response;
use Illuminate\Support\Str;

class ClassSubjectAssignmentSubmissionController extends Controller
{
    use ResponseTrait;
    /**
     * Display a listing of the resource.
     */
    public function index()
    {
        try {
            // Get data
            $data = ClassSubjectAssignmentSubmission::all();
            return $this->responseSuccess($data, 'Class Subject Assignment Submission List Fetch Successfully');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Store a newly created resource in storage.
     */
    public function store(Request $request)
    {
        try {
            $fileType = ClassSubjectAssignment::find($request->class_subject_assignment_id);
            $validator = Validator::make($request->all(), 
                [
                    'attachment' => 'nullable|mimes:'.$fileType->submission_content_type.'|max:2048',
                ],
                [
                    'attachment' => 'Kolom :attribute harus berupa file '.$fileType->submission_content_type.' dengan ukuran maksimal :max kilobyte.',
                ]
            );
            if ($validator->fails()) {
                return $this->responseError($validator->errors(), 'Data tidak valid.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            $validatedData = $request->all();
            // Upload image
            if (!empty($validatedData['attachment'])) {
                $titleShort      = Str::slug(substr($fileType->description, 0, 20));
                $validatedData['attachment'] = UploadHelper::upload('attachment', $validatedData['attachment'], $titleShort . '-' . time(), 'attachment/submission');
            }
            // Create data
            $data = ClassSubjectAssignmentSubmission::create($validatedData);
            return $this->responseSuccess($data, 'New Class Subject Assignment Submission Created Successfully!');
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
            $data = ClassSubjectAssignmentSubmission::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Assignment Submission Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Get Data
            return $this->responseSuccess($data, 'Class Subject Assignment Submission Details Fetch Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }

    /**
     * Update the specified resource in storage.
     */
    public function update(ClassSubjectAssignmentSubmissionRequest $request, $id)
    {
        try {
            // Find data by ID
            $data = ClassSubjectAssignmentSubmission::find($id);
            if (is_null($data)){
                return $this->responseError(null, 'Class Subject Assignment Submission Not Found.', Response::HTTP_NOT_FOUND);
            }
            $fileType = ClassSubjectAssignment::find($request->class_subject_assignment_id);
            $validator = Validator::make($request->all(), 
                [
                    'attachment' => 'nullable|mimes:'.$fileType->submission_content_type.'|max:2048',
                ],
                [
                    'attachment' => 'Kolom :attribute harus berupa file '.$fileType->submission_content_type.' dengan ukuran maksimal :max kilobyte.',
                ]
            );
            if ($validator->fails()) {
                return $this->responseError($validator->errors(), 'Data tidak valid.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            // Update Data
            $data->update($request->all());
            return $this->responseSuccess($data, 'Class Subject Assignment Submission Updated Successfully!');
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
            $data = ClassSubjectAssignmentSubmission::find($id);
            if (is_null($data)) {
                return $this->responseError(null, 'Class Subject Assignment Submission Not Found.', Response::HTTP_NOT_FOUND);
            }
            // Save Data
            $data->deleted_at = date('Y-m-d H:i:s');
            $data->save();
            if (!$data) {
                return $this->responseError(null, 'Failed to delete the Class Subject Assignment Submission.', Response::HTTP_INTERNAL_SERVER_ERROR);
            }
            return $this->responseSuccess($data, 'Class Subject Assignment Submission Deleted Successfully!');
        } catch (\Exception $e) {
            return $this->responseError(null, $e->getMessage(), Response::HTTP_INTERNAL_SERVER_ERROR);
        }
    }
}
